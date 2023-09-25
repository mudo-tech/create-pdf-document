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
	start := time.Now()
	pdfcpu.LoadConfiguration()

	err := pdfcpu.InstallFonts([]string{"./resources/CONSOLA.ttf"})
	if err != nil {
		log.Println(err)
		return
	}

	content, err := populatePDFData(dto.NotaBody{
		TransactionDetail: dto.NotaTransactionDetail{
			ReferenceNumber: "b576c0147cbea97f44665f29294d87cb",
			NotaWa:          "+6282186266734",
			Phone:           "+6282186266734",
			StartedAt:       "27 Juli 2023 14:00 WIB",
			FinsihedAt:      "28 Juli 2023 14:00 WIB",
			PaidAt:          "28 Juli 2023 14:00 WIB",
		},
		CustomerName: "Ridho Muhammad",
		Barcode:      "https://www.bdc.ca/globalassets/digizuite/40415-bdc-qr-code.jpg?v=498d76",
		//Barcode:      "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAALEAAACxCAIAAAAES8uSAAAABmJLR0QA/wD/AP+gvaeTAAAEI0lEQVR4nO3d227dIBAF0J6q///L6VvkbIkGwgz2qdZ6jHw70RbGGMavj4+PX3Dx++4L4HFkgiQTJJkgyQRJJkgyQZIJkkyQZIIkEySZIMkESSZIMkGSCZJMkGSCJBMkmSDJBOnP/iFer9f+QcJ1NvnM8Ufbd89KH51rdM2rv2tVye/VTpBkgiQTpIL+xNXO/azq/rp6DVXn3enHPOH/9kk7QZIJkkyQivsTV6vjCjNGz/cz4wTdYwNXM+MTq/uuHufHtBMkmSDJBKmxP9Gh+/3C6D69OvZw8p1LOe0ESSZIMkF6g/7EzP14ph8w2n40zrFz/LemnSDJBEkmSI39iarn8pln/dX7+s77kevfO8Yebh/P0E6QZIIkE6Ti/sTJ5/WqPsFdfx/9lttpJ0gyQZIJUkF/ovt5unst6Mwci9W/r573UbQTJJkgyQSpsf7EzrP4aN/V+ZI7/YCTYwlV61DUn6CFTJBkgnSo/sTO+snV41SNYczMmZjpr+xcwy1jGNoJkkyQZIL06rtjdcwzmDnXqpN1LdWf4C3JBEkmSAX9iZ2x+o66kFX1rHbqTOzMwbyrdvgn7QRJJkgyQTo0H7Oq3sPq8We2v5rZd/Udx0yfoKMv8mPaCZJMkGSC1Dg+MbPNTr3L1WsbHad7nWf3WtNy2gmSTJBkglQ8f6KjpnX3vbbj3cfqvjO6j/9JO0GSCZJMkIrXi+4831fNedypcbnzvmPnemYYn+A2MkGSCVJjfcyTfYud+hM7cxq630fc8u0x7QRJJkgyQWpcL/rlNBvzDGa2H52r+9c9oVZ3Oe0ESSZIMkEqno858r/Wn+ioAdo9R/Vb2gmSTJBkgnRzfczVfavu/d19oNX/w+01rK60EySZIMkE6dB60dH2IzvvMqq+J9I9d7JqHML8CdrJBEkmSDf3J072P3b27V4vevs7jivtBEkmSDJBetD3yju+27nz3uEJ4xkzdUXNn6CdTJBkgnToe2BXT3jW735XsrPvsXkSI9oJkkyQZIJUXM+qat/uut2r19DxG6uux/gE7WSCJBOkQ/UndnTMT+yYr1C1Dvb2upnaCZJMkGSCdPP4xMjO2suqbbrXU+zU6GylnSDJBEkmSIfqT8zYmSM5uoadORPd70FW/1fH5mZqJ0gyQZIJUnF/4qqjRlNVbcoZHWMDVd8Sax230E6QZIIkE6TG+hNVNShH5xodZ2b7mWuYOf7Iydqd5bQTJJkgyQSpcXziCbrHMKrWboyYP8EjyARJJkiN/YmOOlcn11lUjbWMtrnyfVEeTSZIMkEq7k8ce4b+x7lm7usda0o7xjw66pd/SztBkgmSTJDeoP4Eh2knSDJBkgmSTJBkgiQTJJkgyQRJJkgyQZIJkkyQZIIkEySZIMkESSZIMkGSCZJMkGSCJBOkv6w975dWg7c6AAAAAElFTkSuQmCC",
		Divider: "-------------------------------------------------------",
		NotaBranchDetail: dto.NotaBranchDetail{
			ImageUrl: "https://res.cloudinary.com/dukuh51km/image/upload/v1691834622/staging/mobxL-1691834622.png",
			Name:     "Kassir Bersih",
			Address:  "6, Jl. H. Shibi No.14, RT.6/RW.1, Srengseng Sawah, Kec. Jagakarsa, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 10550",
			Phone:    "+6282186266734",
			Divider:  "-------------------------------------------------------",
		},
		Payment: dto.NotaPayment{
			Status:              "Lunas",
			Method:              "Tunai",
			TotalPriceFormatted: "Rp. 100.000",
			TotalPrice:          0,
			PaidFormatted:       "Rp. 100.000",
			Paid:                10000,
			RemainFormatted:     "Rp. 0",
			Remain:              0,
		},
		FootNote: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book",
		ServiceDetail: dto.NotaServiceDetail{
			TotalPrice: 1000,
			Services: []dto.NotaService{
				{
					Name:              "Reguler Kiloan",
					QuantityFormatted: "1 * 1.00 Kg",
					Quantity:          10,
					UnitAmount:        "",
					Units:             "1",
					PriceFormatted:    "Rp. 10.000",
					Price:             1000,
				},
				{
					Name:              "Karpet",
					QuantityFormatted: "2 * 10 Meter Kubik",
					Quantity:          1,
					UnitAmount:        "1",
					Units:             "1",
					PriceFormatted:    "Rp. 40.000",
					Price:             1000,
				}, {
					Name:              "Baju",
					QuantityFormatted: "1 * Pasang",
					Quantity:          10,
					UnitAmount:        "",
					Units:             "1",
					PriceFormatted:    "Rp. 10.000",
					Price:             1000,
				},
				{
					Name:              "Karpet",
					QuantityFormatted: "2 * 10 Meter Kubik",
					Quantity:          1,
					UnitAmount:        "1",
					Units:             "1",
					PriceFormatted:    "Rp. 40.000",
					Price:             1000,
				},
			},
		},
		Divider1: "-------------------------------------------------------",
		Divider2: "-------------------------------------------------------",
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

	log.Println(time.Since(start).Milliseconds())
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
	MaxY        int
	MaxX        int
	PageStack   map[string]bool
	CurrentPage int
	DefaultFont DefaultFont
}

type DefaultFont struct {
	Size int
	Name string
}

func NewCreatePDf() CreatePDF {
	return CreatePDF{
		config: &ConfigPDF{
			PaperWidth: 125,
			MaxY:       180,
			PageStack: map[string]bool{
				"1": true,
			},
			CurrentPage: 1,
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
				"2": {
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
	UsingColon     bool
	ColWidths      []int
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

	if rule.FontSize == 0 {
		rule.FontSize = int64(cp.config.DefaultFont.Size)
	}

	comp := cp.component.Pages[strconv.FormatInt(int64(cp.config.CurrentPage), 10)].Content
	switch rule.Type {
	case "text":
		text, indent := cp.config.ParseText(fmt.Sprintf("%v", val.Interface()))
		textBox := &primitives.TextBox{
			Value: text,
			Font: &primitives.FormFont{
				Name: cp.config.DefaultFont.Name,
				Size: int(rule.FontSize),
			},
			Anchor: "tc",
			Dy:     float64(cp.config.CurrentY + int(rule.YPos)),
			Dx:     float64(cp.config.CurrentX + int(rule.XPos)),
		}
		cp.config.CurrentY += (int(rule.FontSize+1) * indent) + int(rule.YPos)
		cy := cp.config.CurrentY
		if cp.config.changingPage() {
			overIndent := (cy - cp.config.MaxY) / int(rule.FontSize+1)
			indent -= overIndent
			txts := cp.config.CutStringOnNewLine(text, indent)
			var nextTextBox primitives.TextBox
			if len(txts) > 1 {
				textBox.Value = txts[0]
				cp.component.Pages[strconv.Itoa(cp.config.CurrentPage-1)].Content.TextBoxes = append(
					cp.component.Pages[strconv.Itoa(cp.config.CurrentPage-1)].Content.TextBoxes,
					textBox,
				)
				nextTextBox = *textBox
				nextTextBox.Value = txts[1]
			}

			nextTextBox.Dy = float64(cp.config.CurrentY)
			comp = &primitives.Content{
				TextBoxes: []*primitives.TextBox{&nextTextBox},
			}
			cp.config.CurrentY += (int(rule.FontSize+1) * indent) + int(rule.YPos)
			break
		}
		comp.TextBoxes = append(comp.TextBoxes, textBox)
	case "image":
		imageBox := &primitives.ImageBox{
			Src:    val.String(),
			Anchor: "tc",
			Width:  float64(rule.ImageWidth),
			Height: float64(rule.ImageHeight),
			Dy:     float64(cp.config.CurrentY + int(rule.YPos)),
			Dx:     float64(cp.config.CurrentX + int(rule.XPos)),
		}
		cp.config.CurrentY += int(rule.ImageHeight) + int(rule.YPos)
		if cp.config.changingPage() {
			imageBox.Dy = float64(cp.config.CurrentY)
			comp = &primitives.Content{
				ImageBoxes: []*primitives.ImageBox{imageBox},
			}
			cp.config.CurrentY += int(rule.ImageHeight) + int(rule.YPos)
			break
		}
		comp.ImageBoxes = append(comp.ImageBoxes, imageBox)
	case "tablePivot":
		pType := field.Type
		if pType.Kind() != reflect.Struct &&
			field.Name != reflect.TypeOf(time.Time{}).Name() {
			return fmt.Errorf("tablePivot should be a struct")
		}
		var rows = make([][]string, 0)
		var col []string
		for i := 0; i < pType.NumField(); i++ {
			fieldName := cp.getTableColName(pType, i)
			if fieldName == "" {
				continue
			}

			col = append(col, fieldName)
			if rule.UsingColon {
				col = append(col, ":")
			}
			col = append(col, fmt.Sprintf("%v", val.Field(i).Interface()))

			rows = append(rows, col)
			col = []string{}
		}

		var colAnchors = make([]string, len(rows[0]))
		for i := range colAnchors {
			colAnchors[i] = "Left"
		}

		tables := &primitives.Table{
			Hide:   false,
			Values: rows,
			Font: &primitives.FormFont{
				Name: cp.config.DefaultFont.Name,
				Size: int(rule.FontSize),
			},
			Anchor:     "tc",
			Width:      float64(cp.config.PaperWidth - 3),
			LineHeight: int(rule.TableRowHeight),
			Rows:       len(rows),
			Cols:       len(rows[0]),
			ColWidths:  rule.ColWidths,
			Border: &primitives.Border{
				Width: 1,
				Color: "#FFFFFF",
			},
			ColAnchors: colAnchors,
			Dy:         float64(-(cp.config.CurrentY + int(rule.YPos)) / 2),
			Dx:         float64(cp.config.CurrentX + int(rule.XPos)),
		}
		cp.config.CurrentY += (int(rule.TableRowHeight) * (len(rows))) + int(rule.YPos)
		if cp.config.changingPage() {
			tables.Dy = float64(cp.config.CurrentY)
			comp = &primitives.Content{
				Tables: []*primitives.Table{tables},
			}
			cp.config.CurrentY += (int(rule.TableRowHeight) * (len(rows))) + int(rule.YPos)
			break
		}
		comp.Tables = append(comp.Tables, tables)
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

		tables := &primitives.Table{
			Hide:   false,
			Values: rows,
			Header: &primitives.TableHeader{
				Values:          colName,
				BackgroundColor: "#D8D8D8",
				ColAnchors:      []string{"Left", "Left", "Left"},
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
			ColWidths:  []int{30, 40, 30},
			Border: &primitives.Border{
				Width: 1,
				Color: "#FFFFFF",
			},
			ColAnchors: []string{"Left", "Left", "Left"},
			//Dy:         -5,
			Dy: float64(-(cp.config.CurrentY + int(rule.YPos)) / 2),
			Dx: float64(cp.config.CurrentX + int(rule.XPos)),
		}
		cp.config.CurrentY += (int(rule.TableRowHeight+1) * (len(rows) + 1)) + int(rule.YPos)
		if cp.config.changingPage() {
			tables.Dy = float64(cp.config.CurrentY)
			comp = &primitives.Content{
				Tables: []*primitives.Table{tables},
			}
			cp.config.CurrentY += (int(rule.TableRowHeight+1) * (len(rows) + 1)) + int(rule.YPos)
			break
		}
		comp.Tables = append(comp.Tables, tables)
	}

	cp.component.Pages[strconv.FormatInt(int64(cp.config.CurrentPage), 10)].Content = comp
	return nil
}

func (cf *ConfigPDF) changingPage() bool {
	if !cf.PageStack[strconv.Itoa(cf.CurrentPage+1)] && cf.CurrentY > cf.MaxY {
		cf.CurrentPage += 1
		cf.PageStack[strconv.Itoa(cf.CurrentPage)] = true
		cf.CurrentY = 0
		return true
	}
	return false
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
	tmp := NewCreatePDf()
	err := tmp.ApplyStyling(reflect.TypeOf(body), reflect.ValueOf(body))
	if err != nil {
		return primitives.PDF{}, err
	}
	return *tmp.component, nil
}

func (cf ConfigPDF) CutStringOnNewLine(tx string, count int) []string {
	var (
		step int
		strs = make([]string, 0)
	)
	for i := range tx {
		if string(tx[i]) == "\n" {
			step += 1
			if count == step {
				strs = append(strs, tx[:i])
				strs = append(strs, tx[i:])
			}
		}
	}

	return strs
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
