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
	TotalPrice uint          `pdfField:"type:text;name:totalPrice"  json:"totalPrice"`
	Services   []NotaService `json:"services"`
}

type NotaService struct {
	Name       string `json:"name"`
	Quantity   uint   `json:"quantity"`
	UnitAmount string `json:"unitAmount"`
	Units      string `json:"units"`
	Price      uint   `json:"price"`
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
