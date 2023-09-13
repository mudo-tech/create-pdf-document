package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mudo-tech/create-pdf-document/dto"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/primitives"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	pdfcpu.LoadConfiguration()

	err := pdfcpu.InstallFonts([]string{"./resources/CONSOLA.ttf"})
	if err != nil {
		log.Println(err)
		return
	}

	content, err := populatePDFData(dto.NotaBody{
		TransactionDetail: dto.NotaTransactionDetail{
			ReferenceNumber: "abasdfalk234kdf234",
			NotaWa:          "+6282186266734",
			Name:            "Ridho Muhammad",
			Phone:           "+6282186266734",
			StartedAt:       "Senin, 27 Juli 2023 14:00 WIB",
			FinsihedAt:      "Selasa, 28 Juli 2023 14:00 WIB",
		},
		NotaBranchDetail: dto.NotaBranchDetail{
			ImageUrl: "https://res.cloudinary.com/dukuh51km/image/upload/v1691834622/staging/mobxL-1691834622.png",
			Name:     "Kassir Bersih",
			Address:  "6, Jl. H. Shibi No.14, RT.6/RW.1, Srengseng Sawah, Kec. Jagakarsa, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 10550",
			Phone:    "+6282186266734",
			Divider:  "-------------------------------------------------------",
		},
		ServiceDetail: dto.NotaServiceDetail{
			TotalPrice: 1000,
			Services: []dto.NotaService{
				{
					Name:       "Baji",
					Quantity:   10,
					UnitAmount: "1",
					Units:      "1",
					Price:      1000,
				},
			},
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	payloadBytes, err := json.Marshal(content)
	op := "utils.http.ReadRequest:"
	if err != nil {
		log.Println(op, err)
		return
	}
	f0 := bytes.NewBuffer(payloadBytes)

	var f2 *os.File
	outFilePDF := "./example-pdf.pdf"
	f2, err = os.Create(outFilePDF)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if err = f2.Close(); err != nil {
			log.Println(err)
			return
		}
		return
	}()

	err = pdfcpu.Create(nil, f0, f2, pdfcpu.LoadConfiguration())
	if err != nil {
		log.Println(err)
		return
	}
}

type CreatePDF struct {
	component *primitives.PDF
	config    *ConfigPDF
}

type ConfigPDF struct {
	PaperWidth  int
	PaperHeight int
	CurrentY    int
	CurrentX    int
	Padding     int
	DefaultFont DefaultFont
}

type DefaultFont struct {
	Size int
	Name string
}

func NewCreatePDf() CreatePDF {
	return CreatePDF{
		config: &ConfigPDF{
			PaperWidth:  125,
			PaperHeight: 184,
			CurrentY:    0,
			CurrentX:    0,
			Padding:     3,
			DefaultFont: DefaultFont{
				Size: 4,
				Name: "Consolas",
			},
		},
		component: &primitives.PDF{
			Paper: "A8",
			// Paper size is W X H = 125 X 184
			Crop:            "10",
			Origin:          "upperLeft",
			ContentBox:      false,
			Debug:           false,
			Guides:          false,
			TimestampFormat: time.Now().Format("Monday, 2.Jan 2006 15:04:05"),
			Pages: map[string]*primitives.PDFPage{
				"1": {
					Content: &primitives.Content{},
				},
			},
		},
	}
}

func (cp *CreatePDF) ApplyStyling(pType reflect.Type, pVal reflect.Value) error {
	for i := 0; i < pType.NumField(); i++ {
		tag := pType.Field(i).Tag.Get("pdfField")
		if pType.Field(i).Type.Kind() == reflect.Struct &&
			tag == "" {
			err := cp.ApplyStyling(pType.Field(i).Type, pVal.Field(i))
			if err != nil {
				return err
			}
			continue
		}
		var keyValMap = map[string]string{}
		for _, s := range strings.Split(tag, ";") {
			if len(s) < 1 {
				continue
			}
			keyVal := strings.Split(s, ":")
			if len(keyVal) < 2 {
				return fmt.Errorf("invalid on key val")
			}
			keyValMap[keyVal[0]] = keyVal[1]
		}
		err := cp.CreateComponent(keyValMap, pType.Field(i), pVal.Field(i))
		if err != nil {
			return err
		}

	}

	return nil
}

type KeyValRules struct {
	Type           string
	YPos           int64
	XPos           int64
	FontSize       int64
	ImageWidth     int64
	ImageHeight    int64
	TableRowHeight int64
	ColAnchor      string
	Anchor         string
}

func (rl *KeyValRules) extractFromMap(keyvals map[string]string) (err error) {
	rl.Type = keyvals["type"]

	if keyvals["dx"] != "" {
		rl.XPos, err = strconv.ParseInt(keyvals["dx"], 10, 64)
		if err != nil {
			return fmt.Errorf("dy should be number")
		}
	}

	if keyvals["dy"] != "" {
		rl.YPos, err = strconv.ParseInt(keyvals["dy"], 10, 64)
		if err != nil {
			return fmt.Errorf("dx should be number")
		}
	}

	if keyvals["fontSize"] != "" {
		rl.FontSize, err = strconv.ParseInt(keyvals["fontSize"], 10, 64)
		if err != nil {
			return fmt.Errorf("fontSize should be number")
		}
	}

	if keyvals["imageWidth"] != "" {
		rl.ImageWidth, err = strconv.ParseInt(keyvals["imageWidth"], 10, 64)
		if err != nil {
			return fmt.Errorf("imageWidth should be number")
		}

	}

	if keyvals["imageHeight"] != "" {
		rl.ImageHeight, err = strconv.ParseInt(keyvals["imageHeight"], 10, 64)
		if err != nil {
			return fmt.Errorf("imageHeight should be number")
		}
	}

	if keyvals["tableRowHeight"] != "" {
		rl.TableRowHeight, err = strconv.ParseInt(keyvals["tableRowHeight"], 10, 64)
		if err != nil {
			return fmt.Errorf("tableRowHeight should be number")
		}

	}

	if keyvals["usingColon"] != "" {
		rl.UsingColon, err = strconv.ParseBool(keyvals["usingColon"])
		if err != nil {
			return fmt.Errorf("usingColon should be boolean")
		}
	}

	if keyvals["colWidths"] != "" {
		cs := strings.Split(keyvals["colWidths"], ",")
		var cwd = make([]int, len(cs))
		for i, c := range cs {
			num, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return fmt.Errorf("kesalahan format pada colWidths")
			}
			cwd[i] = int(num)
		}
		rl.ColWidths = cwd
	}

	rl.ColAnchor = keyvals["colAnchor"]
	rl.Anchor = keyvals["anchor"]

	return nil
}

func (cp *CreatePDF) CreateComponent(keyVals map[string]string, field reflect.StructField, val reflect.Value) error {
	rule := &KeyValRules{}
	err := rule.extractFromMap(keyVals)
	if err != nil {
		return err
	}

	pageLen := len(cp.component.Pages)
	if cp.config.CurrentY > cp.config.PaperHeight-30 {
		pageLen += 1
		cp.component.Pages[strconv.FormatInt(int64(pageLen), 10)] =
			&primitives.PDFPage{
				Content: &primitives.Content{},
			}
		cp.config.CurrentY = 0
	}

	if rule.FontSize == 0 {
		rule.FontSize = int64(cp.config.DefaultFont.Size)
	}

	comp := cp.component.Pages[strconv.FormatInt(int64(pageLen), 10)].Content
	switch rule.Type {
	case "text":
		text, indent := cp.config.ParseText(fmt.Sprintf("%v", val.Interface()))
		comp.TextBoxes = append(comp.TextBoxes, &primitives.TextBox{
			Value: text,
			Font: &primitives.FormFont{
				Name: cp.config.DefaultFont.Name,
				Size: cp.config.DefaultFont.Size,
			},
			Anchor: "tc",
			Dy:     float64(cp.config.CurrentY + int(rule.YPos)),
			Dx:     float64(cp.config.CurrentX + int(rule.XPos)),
		})

		if rule.FontSize == 0 {
			rule.FontSize = int64(cp.config.DefaultFont.Size)
		}

		cp.config.CurrentY += (int(rule.FontSize+1) * indent) + int(rule.YPos)
		cp.config.CurrentX += int(rule.XPos)
	case "image":
		comp.ImageBoxes = append(comp.ImageBoxes, &primitives.ImageBox{
			Src:    val.String(),
			Anchor: "tc",
			Width:  float64(rule.ImageWidth),
			Height: float64(rule.ImageHeight),
			Dy:     float64(cp.config.CurrentY + int(rule.YPos)),
			Dx:     float64(cp.config.CurrentX + int(rule.XPos)),
		})
		cp.config.CurrentY += int(rule.ImageHeight) + int(rule.YPos)
		cp.config.CurrentX += int(rule.XPos)
	case "tablePivot":
		pType := field.Type
		if pType.Kind() != reflect.Struct &&
			field.Name != reflect.TypeOf(time.Time{}).Name() {
			return fmt.Errorf("tablePivot should be a struct")
		}
		var rows = make([][]string, 0)
		for i := 0; i < pType.NumField(); i++ {
			fieldName := cp.getTableColName(pType, i)
			if fieldName == "" {
				continue
			}
			rows = append(rows, []string{" ", fieldName, ":", fmt.Sprintf("%v", val.Field(i).Interface())})
		}

		comp.Tables = append(comp.Tables, &primitives.Table{
			Hide:   false,
			Values: rows,
			Font: &primitives.FormFont{
				Name: cp.config.DefaultFont.Name,
				Size: cp.config.DefaultFont.Size,
			},
			Anchor:     "tc",
			Width:      float64(cp.config.PaperWidth - 3),
			LineHeight: int(rule.TableRowHeight),
			Rows:       len(rows),
			Cols:       len(rows[0]),
			ColWidths:  []int{5, 30, 5, 60},
			Border: &primitives.Border{
				Width: 1,
				Color: "#FFFFFF",
			},
			ColAnchors: []string{"Center", "Left", "Center", "Left"},
			Dy:         float64(-(cp.config.CurrentY + int(rule.YPos))),
			Dx:         float64(cp.config.CurrentX + int(rule.XPos)),
		})
		cp.config.CurrentY += int(rule.TableRowHeight+1)*len(rows) + int(rule.YPos)
		cp.config.CurrentX += int(rule.XPos)
	case "table":
		pType := field.Type
		if pType.Kind() != reflect.Slice {
			if val.Index(0).Kind() != reflect.Struct {
				return fmt.Errorf("invalid type for field with type table")
			}
			return fmt.Errorf("invalid type for field with type table")
		}

		var colName = make([]string, 0)
		var colNameMap = map[string]string{}
		for i := 0; i < val.Index(0).NumField(); i++ {
			fieldName := cp.getTableColName(val.Index(0).Type(), i)
			if fieldName == "" {
				continue
			}
			colNameMap[val.Index(0).Type().Field(i).Name] = fieldName
			colName = append(colName, fieldName)
		}

		var rows = make([][]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			rowVals := val.Index(i)
			var row []string
			for j := 0; j < rowVals.NumField(); j++ {
				_, ok := colNameMap[rowVals.Type().Field(j).Name]
				if ok {
					row = append(row, fmt.Sprintf("%v", rowVals.Field(j).Interface()))
				}
			}
			rows[i] = row
		}

		comp.Tables = append(comp.Tables, &primitives.Table{
			Hide:   false,
			Values: rows,
			Header: &primitives.TableHeader{
				Values:          colName,
				BackgroundColor: "#D8D8D8",
				ColAnchors:      []string{"Left", "Left", "Left", "Left"},
			},
			Font: &primitives.FormFont{
				Name: cp.config.DefaultFont.Name,
				Size: cp.config.DefaultFont.Size,
			},
			Anchor:     "tc",
			Width:      float64(cp.config.PaperWidth - 3),
			LineHeight: int(rule.TableRowHeight),
			Rows:       len(rows),
			Cols:       len(rows[0]),
			ColWidths:  []int{25, 25, 25, 25},
			Border: &primitives.Border{
				Width: 1,
				Color: "#FFFFFF",
			},
			ColAnchors: []string{"Left", "Left", "Left", "Left"},
			Dy:         float64(-(cp.config.CurrentY + int(rule.YPos))),
			Dx:         float64(cp.config.CurrentX + int(rule.XPos)),
		})
		cp.config.CurrentY += int(rule.TableRowHeight+1)*len(rows) + int(rule.YPos)
		cp.config.CurrentX += int(rule.XPos)

	}

	cp.component.Pages[strconv.FormatInt(int64(pageLen), 10)].Content = comp
	return nil
}

func (cp *CreatePDF) getTableColName(pType reflect.Type, i int) string {
	var fieldName string
	for _, s := range strings.Split(pType.Field(i).Tag.Get("pdfField"), ";") {
		if s == "" {
			return ""
		}
		re := regexp.MustCompile(`colName:(.+?);`)
		match := re.FindStringSubmatch(s + ";")
		if len(match) > 0 {
			fieldName = match[1]
			break
		}
	}
	return fieldName
}

func populatePDFData(body dto.NotaBody) (primitives.PDF, error) {
	//fontConsola := &primitives.FormFont{
	//	Name: "Consolas",
	//	Size: 4,
	//}
	//
	//logoHeight := 30
	//padding := 4
	//fontSize := 4

	tmp := NewCreatePDf()
	err := tmp.ApplyStyling(reflect.TypeOf(body), reflect.ValueOf(body))
	if err != nil {
		return primitives.PDF{}, err
	}
	return *tmp.component, nil
}

func (cf ConfigPDF) ParseText(tx string) (string, int) {
	r := strings.NewReplacer("\n", "", "  ", "")
	tx = r.Replace(tx)
	segments := len(tx) / 56
	if segments*56 < len(tx) {
		segments += 1
	}
	start := 0
	end := 56
	var strs []string
	for i := 0; i < segments && end < len(tx); i++ {
		for i := end; i > 0; i-- {
			if string(tx[i]) == " " {
				end = i
				break
			}
		}
		strs = append(strs, tx[start:end])
		start = end
		end += 56
	}
	strs = append(strs, tx[start:])

	return strings.Join(strs, "\n"), segments
}

//
//const template = primitives.PDF{
//	Paper: "A8",
//	// Paper size is W X H = 125 X 184
//	Crop:            "10",
//	Origin:          "upperLeft",
//	ContentBox:      false,
//	Debug:           false,
//	Guides:          false,
//	TimestampFormat: time.Now().Format("Monday, 2.Jan 2006 15:04:05"),
//	Pages: map[string]*primitives.PDFPage{
//		"1": {
//			Content: &primitives.Content{
//				ImageBoxes: []*primitives.ImageBox{
//					{
//						Src:    body.NotaBranchDetail.ImageUrl,
//						Height: float64(logoHeight),
//						Anchor: "tc",
//						Dy:     2,
//					},
//				},
//				TextBoxes: []*primitives.TextBox{
//					{
//						Font:      fontConsola,
//						Alignment: "center",
//						Value:     parseText(body.NotaBranchDetail.Address),
//						Anchor:    "tc",
//						Dy:        float64(logoHeight + padding),
//					},
//					{
//						Font:      fontConsola,
//						Alignment: "center",
//						Value:     "--------------------------------------------------------",
//						Anchor:    "tc",
//						Dy:        float64(logoHeight + (fontSize * 2) + padding),
//					},
//					{
//						Font:      fontConsola,
//						Alignment: "left",
//						Value:     "Tagihan Untuk Layanan:",
//						Anchor:    "tl",
//						Dy:        float64(logoHeight + (fontSize * 3) + (padding * 2)),
//						Dx:        3,
//					},
//				},
//				Tables: []*primitives.Table{
//					{
//						Hide: false,
//						Font: fontConsola,
//						Values: [][]string{
//							{"", "ID Referensi", ":", "abasdfalk234kdf234"},
//							{"", "Nama Pelanggan", ":", "Ridho Muhammad"},
//							{"", "Nomor Wa", ":", "+6282186266734"},
//							{"", "Dimulai dari", ":", "Senin, 27 Juli 2023 14:00 WIB"},
//							{"", "Estimasi", ":", "Selasa, 28 Juli 2023 14:00 WIB"},
//						},
//						Border:     tableNoBorder,
//						ColWidths:  []int{5, 30, 5, 60},
//						Width:      125,
//						Grid:       false,
//						Rows:       5,
//						Cols:       4,
//						ColAnchors: []string{"Center", "Left", "Center", "Left"},
//						LineHeight: 6,
//						Anchor:     "tc",
//						Dy:         -27,
//					},
//					{
//						Hide: false,
//						Font: fontConsola,
//						Header: &primitives.TableHeader{
//							BackgroundColor: "#D8D8D8",
//							Values:          []string{"Nama Layanan", "Harga Satuan", "Jumlah", "Harga Kolektif"},
//							ColAnchors:      []string{"Left", "Left", "Left", "Left"},
//						},
//						Border: tableNoBorder,
//						Values: [][]string{
//							{"Cuci Reguler", "Rp. 6.000", "4 Kg", "Rp. 24.000"},
//							{"Cuci Kilat", "Rp. 12.000", "3 Kg", "Rp. 36.000"},
//							{"Karpet", "Rp. 10.000", "12 X 300 Meter", "Rp. 24.000"},
//							{"Sepatu", "Rp. 12.000", "1 Pasang", "Rp. 12.000"},
//							{"Tuxedo", "Rp. 12.000", "1 Item", "Rp. 12.000"},
//						},
//						Rows:       5,
//						Cols:       4,
//						Width:      125,
//						ColWidths:  []int{25, 25, 25, 25},
//						ColAnchors: []string{"Left", "Left", "Left", "Left"},
//						LineHeight: 6,
//						Anchor:     "tc",
//						Dy:         -46,
//					},
//					{
//						Hide:   false,
//						Font:   fontConsola,
//						Border: tableNoBorder,
//						Values: [][]string{
//							{"", "", "Total:", "Rp. 108.000"},
//						},
//						Rows:       1,
//						Cols:       4,
//						Width:      125,
//						ColWidths:  []int{25, 25, 25, 25},
//						LineHeight: 6,
//						Anchor:     "tc",
//						Dy:         -65,
//					},
//				},
//			},
//		},
//	},
//}
