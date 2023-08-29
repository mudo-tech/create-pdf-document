package main

import (
	"bytes"
	"encoding/json"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/primitives"
	"log"
	"os"
	"time"
)

func main() {

	pdfcpu.LoadConfiguration()

	err := pdfcpu.InstallFonts([]string{"./resources/CONSOLA.ttf"})
	if err != nil {
		log.Println("1", err)
		return
	}

	fontConsola := &primitives.FormFont{
		Name: "Consolas",
		Size: 4,
	}

	tableNoBorder := &primitives.Border{
		Width: 1,
		Color: "#FFFFFF",
	}

	content := primitives.PDF{
		Paper:           "A8",
		Crop:            "10",
		Origin:          "upperLeft",
		ContentBox:      false,
		Debug:           false,
		Guides:          false,
		TimestampFormat: time.Now().Format("Monday, 2.Jan 2006 15:04:05"),
		Pages: map[string]*primitives.PDFPage{
			"1": {
				Content: &primitives.Content{
					ImageBoxes: []*primitives.ImageBox{
						{
							Src:    "https://res.cloudinary.com/dukuh51km/image/upload/v1691834622/staging/mobxL-1691834622.png",
							Height: 30,
							Anchor: "tc",
							Dy:     2,
						},
					},
					TextBoxes: []*primitives.TextBox{
						{
							Font:      fontConsola,
							Alignment: "center",
							Value:     "Jalan kedoya 2 Jakarta Selatan, DKI Jakarta Indonesia\nEmail ke kassirbersih@gmail.com, Wa Ke +082186266734.",
							Anchor:    "tc",
							Dy:        34,
						},
						{
							Font:      fontConsola,
							Alignment: "center",
							Value:     "--------------------------------------------------------",
							Anchor:    "tc",
							Dy:        43,
						},
						{
							Font:      fontConsola,
							Alignment: "center",
							Value:     "Tagihan Untuk Layanan:",
							Anchor:    "tc",
							Dy:        48,
							Dx:        3,
						},
					},
					Tables: []*primitives.Table{
						{
							Hide: false,
							Font: fontConsola,
							Values: [][]string{
								{"", "ID Referensi", ":", "abasdfalk234kdf234"},
								{"", "Nama Pelanggan", ":", "Ridho Muhammad"},
								{"", "Nomor Wa", ":", "+6282186266734"},
								{"", "Dimulai dari", ":", "Senin, 27 Juli 2023 14:00 WIB"},
								{"", "Estimasi", ":", "Selasa, 28 Juli 2023 14:00 WIB"},
							},
							Border:     tableNoBorder,
							ColWidths:  []int{5, 30, 5, 60},
							Width:      125,
							Grid:       false,
							Rows:       5,
							Cols:       4,
							ColAnchors: []string{"Center", "Left", "Center", "Left"},
							LineHeight: 6,
							Anchor:     "tc",
							Dy:         -27,
						},
						{
							Hide: false,
							Font: fontConsola,
							Header: &primitives.TableHeader{
								BackgroundColor: "#D8D8D8",
								Values:          []string{"Nama Layanan", "Harga Satuan", "Jumlah", "Harga Kolektif"},
								ColAnchors:      []string{"Left", "Left", "Left", "Left"},
							},
							Border: tableNoBorder,
							Values: [][]string{
								{"Cuci Reguler", "Rp. 6.000", "4 Kg", "Rp. 24.000"},
								{"Cuci Kilat", "Rp. 12.000", "3 Kg", "Rp. 36.000"},
								{"Karpet", "Rp. 10.000", "12 X 300 Meter", "Rp. 24.000"},
								{"Sepatu", "Rp. 12.000", "1 Pasang", "Rp. 12.000"},
								{"Tuxedo", "Rp. 12.000", "1 Item", "Rp. 12.000"},
							},
							Rows:       5,
							Cols:       4,
							Width:      125,
							ColWidths:  []int{25, 25, 25, 25},
							ColAnchors: []string{"Left", "Left", "Left", "Left"},
							LineHeight: 6,
							Anchor:     "tc",
							Dy:         -46,
						},
						{
							Hide:   false,
							Font:   fontConsola,
							Border: tableNoBorder,
							Values: [][]string{
								{"", "", "Total:", "Rp. 108.000"},
							},
							Rows:       1,
							Cols:       4,
							Width:      125,
							ColWidths:  []int{25, 25, 25, 25},
							LineHeight: 6,
							Anchor:     "tc",
							Dy:         -65,
						},
					},
				},
			},
		},
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
