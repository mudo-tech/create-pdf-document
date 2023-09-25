package dto

type NotaBody struct {
	TransactionID     uint                  `json:"transactionId"`
	ActionType        string                `json:"actionType"`
	NotaBranchDetail  NotaBranchDetail      `json:"notaBranchDetail"`
	TransactionDetail NotaTransactionDetail `pdfField:"type:tablePivot;tableRowHeight:6;usingColon:true;colWidths:35,5,60" json:"transactionDetail"`
	Divider           string                `pdfField:"type:text" json:"-"`
	CustomerName      string                `pdfField:"type:text;fontSize:8" json:"customerName"`
	ServiceDetail     NotaServiceDetail     `json:"service"`
	Divider1          string                `pdfField:"type:text" json:"-"`
	Payment           NotaPayment           `pdfField:"type:tablePivot;tableRowHeight:6;dx:39;colWidths:30,70" json:"payment"`
	Divider2          string                `pdfField:"type:text" json:"-"`
	FootNote          string                `pdfField:"type:text" json:"footNote"`
	Barcode           string                `pdfField:"type:image;imageHeight:35" json:"barcode"`
}

type NotaBranchDetail struct {
	ID       uint   `json:"-"`
	ImageUrl string `pdfField:"type:image;imageHeight:25" json:"ImageUrl"`
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
	PriceFormatted    string `pdfField:"colName:Price" json:"-"`
	Price             uint   `json:"price"`
}

type NotaPayment struct {
	ProcessedBy         string `json:"processedBy"`
	MethodFormatted     string `json:"methodFormatted"`
	PaidAt              string `json:"paidAt"`
	Status              string `json:"status"`
	Method              string `json:"method"`
	TotalPriceFormatted string `pdfField:"colName:Total Harga" json:"-"`
	TotalPrice          uint   `json:"totalPrice"`
	PaidFormatted       string `pdfField:"colName:Terbayar" json:"-"`
	Paid                uint   `json:"paid"`
	RemainFormatted     string `pdfField:"colName:Sisa" json:"-"`
	Remain              uint   `json:"remain"`
}

type NotaTransactionDetail struct {
	CreatedBy       string `json:"createdBy"`
	Parfum          string `json:"parfum"`
	Rack            string `json:"rack"`
	Name            string `pdfField:"colName:Nama Pelanggan" json:"name"`
	Phone           string `pdfField:"colName:Nomor Wa" json:"phone"`
	StartedAt       string `pdfField:"colName:Laundri Masuk" json:"startedAt"`
	FinsihedAt      string `pdfField:"colName:Estimasi" json:"finishedAt"`
	ReferenceNumber string `pdfField:"colName:Nomor Nota" json:"referenceNumber"`
	NotaPdf         string `json:"notaFile"`
	NotaWa          string `json:"whatsappCTA"`
	PaidAt          string `pdfField:"colName:Dibayar Pada" json:"paidAt"`
}
