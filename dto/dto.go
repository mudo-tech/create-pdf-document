package dto

type NotaBody struct {
	TransactionID     uint                  `json:"transactionId"`
	NotaBranchDetail  NotaBranchDetail      `json:"notaBranchDetail"`
	TransactionDetail NotaTransactionDetail `pdfField:"type:tablePivot;tableRowHeight:6;usingColon:true;colWidths:35,5,60" json:"transactionDetail"`
	Divider           string                `pdfField:"type:text" json:"-"`
	CustomerName      string                `pdfField:"type:text;fontSize:8" json:"customerName"`
	ServiceDetail     NotaServiceDetail     `json:"service"`
	Divider1          string                `pdfField:"type:text" json:"-"`

	Payment  NotaPayment `pdfField:"type:tablePivot;tableRowHeight:6;dx:39;colWidths:30,70" json:"payment"`
	Divider2 string      `pdfField:"type:text" json:"-"`
	FootNote string      `pdfField:"type:text" json:"footNote"`
}

type NotaBranchDetail struct {
	ImageUrl string `pdfField:"type:image;imageHeight:15" json:"ImageUrl"`
	Name     string `pdfField:"type:text" json:"name"`
	Address  string `pdfField:"type:text" json:"address"`
	Phone    string `pdfField:"type:text" json:"phone"`
	Divider  string `pdfField:"type:text" json:"-"`
}

type NotaServiceDetail struct {
	Services   []NotaService `pdfField:"type:table;tableRowHeight:6" json:"services"`
	TotalPrice uint          `json:"totalPrice"`
}

type NotaService struct {
	Name              string `pdfField:"colName:Name" json:"name"`
	QuantityFormatted string `pdfField:"colName:Quantity" json:"quantityFormatted"`
	Quantity          uint   `json:"quantity"`
	UnitAmount        string `json:"unitAmount"`
	Units             string `json:"units"`
	PriceFormatted    string `pdfField:"colName:Price" json:"priceFormatted"`
	Price             uint   `json:"price"`
}

type NotaPayment struct {
	Status              string `json:"status"`
	Method              string `json:"method"`
	TotalPriceFormatted string `pdfField:"colName:Total Harga" json:"totalPriceFormatted"`
	TotalPrice          uint   `json:"totalPrice"`
	PaidFormatted       string `pdfField:"colName:Terbayar" json:"paidFormatted"`
	Paid                uint   `json:"paid"`
	RemainFormatted     string `pdfField:"colName:Sisa" json:"remainFormatted"`
	Remain              uint   `json:"remain"`
}

type NotaTransactionDetail struct {
	ReferenceNumber string `pdfField:"colName:Nomor Nota" json:"referenceNumber"`
	NotaPdf         string `json:"notaFile"`
	NotaWa          string `json:"whatsappCTA"`
	Phone           string `pdfField:"colName:Nomor Wa" json:"phone"`
	StartedAt       string `pdfField:"colName:Laundri Masuk" json:"startedAt"`
	FinsihedAt      string `pdfField:"colName:Estimasi" json:"finishedAt"`
	PaidAt          string `pdfField:"colName:Dibayar Pada" json:"paidAt"`
}
