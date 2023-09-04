package dto

type NotaBody struct {
	TransactionID     uint                  `json:"transactionId"`
	NotaBranchDetail  NotaBranchDetail      `json:"notaBranchDetail"`
	TransactionDetail NotaTransactionDetail `pdfField:"type:tablePivot;tableRowHeight:6;dy:-20" json:"transactionDetail"`
	ServiceDetail     NotaServiceDetail     `json:"service"`
	Payment           NotaPayment           `json:"payment"`
}

type NotaBranchDetail struct {
	ImageUrl string `pdfField:"type:image;imageHeight:15" json:"ImageUrl"`
	Name     string `pdfField:"type:text" json:"name"`
	Address  string `pdfField:"type:text" json:"address"`
	Phone    string `pdfField:"type:text" json:"phone"`
	Divider  string `pdfField:"type:text" json:"-"`
}

type NotaServiceDetail struct {
	Services   []NotaService `pdfField:"type:table;tableRowHeight:6;dy:-18" json:"services"`
	TotalPrice uint          `pdfField:"type:text;name:totalPrice;dy:55;dx:35"  json:"totalPrice"`
}

type NotaService struct {
	Name       string `pdfField:"colName:Name" json:"name"`
	Quantity   uint   `pdfField:"colName:Quantity" json:"quantity"`
	UnitAmount string `pdfField:"colName:Unit Amount" json:"unitAmount"`
	Units      string `json:"units"`
	Price      uint   `pdfField:"colName:Price" json:"price"`
}

type NotaPayment struct {
	Status string `json:"status"`
	Method string `json:"method"`
	PaidAt string `json:"paidAt"`
	Paid   uint   `json:"paid"`
	Remain uint   `json:"remain"`
}

type NotaTransactionDetail struct {
	ReferenceNumber string `pdfField:"colName:ID Referensi" json:"referenceNumber"`
	NotaPdf         string `json:"notaFile"`
	NotaWa          string `json:"whatsappCTA"`
	Name            string `pdfField:"colName:Nama Pelanggan" json:"name"`
	Phone           string `pdfField:"colName:Nomor Wa" json:"phone"`
	StartedAt       string `pdfField:"colName:Dimulai dari" json:"startedAt"`
	FinsihedAt      string `pdfField:"colName:Estimasi" json:"finishedAt"`
}
