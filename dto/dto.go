package dto

type NotaBody struct {
	TransactionID     uint                  `json:"transactionId"`
	NotaBranchDetail  NotaBranchDetail      `json:"notaBranchDetail"`
	ServiceDetail     NotaServiceDetail     `json:"service"`
	TransactionDetail NotaTransactionDetail `json:"transactionDetail"`
	Payment           NotaPayment           `json:"payment"`
}

type NotaBranchDetail struct {
	ImageUrl string `pdfField:"type:Textfield;page:1;name:ImageUrl" json:"ImageUrl"`
	Name     string `pdfField:"type:Textfield;page:1;name:name" json:"name"`
	Address  string `pdfField:"type:Textfield;page:1;name:address" json:"address"`
	Phone    string `pdfField:"type:Textfield;page:1;name:phone" json:"phone"`
}

type NotaServiceDetail struct {
	TotalPrice uint          `pdfField:"type:Textfield;page:1;name:totalPrice"  json:"totalPrice"`
	Services   []NotaService `json:"services"`
}

type NotaService struct {
	Name       string `pdfField:"type:Textfield;page:1;name:ServiceName" json:"name"`
	Quantity   uint   `pdfField:"type:Textfield;page:1;name:ServiceQuantity" json:"quantity"`
	UnitAmount string `json:"unitAmount"`
	Units      string `pdfField:"type:Textfield;page:1;name:ServiceUnit" json:"units"`
	Price      uint   `pdfField:"type:Textfield;page:1;name:ServicePrice" json:"price"`
}

type NotaPayment struct {
	Status string `pdfField:"type:Textfield;page:1;name:paymentStatus" json:"status"`
	Method string `pdfField:"type:Textfield;page:1;name:paymentMethod" json:"method"`
	PaidAt string `pdfField:"type:Textfield;page:1;name:paymentPaidAt" json:"paidAt"`
	Paid   uint   `pdfField:"type:Textfield;page:1;name:paymentPaidAmount" json:"paid"`
	Remain uint   `pdfField:"type:Textfield;page:1;name:paymentRemain" json:"remain"`
}

type NotaTransactionDetail struct {
	ReferenceNumber string `pdfField:"type:Textfield;page:1;name:referenceNumber" json:"referenceNumber"`
	NotaPdf         string `json:"notaFile"`
	NotaWa          string `json:"whatsappCTA"`
	Name            string `pdfField:"type:Textfield;page:1;name:custName"  json:"name"`
	Phone           string `pdfField:"type:Textfield;page:1;name:custPhone"  json:"phone"`
	StartedAt       string `pdfField:"type:Textfield;page:1;name:startedAt" json:"startedAt"`
	FinsihedAt      string `pdfField:"type:Textfield;page:1;name:finishedAt" json:"finishedAt"`
}
