package domain

type OrderRequest struct {
	OrderType     string `json:"orderType" binding:"required"`
	TransactionID string `json:"transactionId" binding:"required"`
	UserId        string `json:"userId" binding:"required"`
	PackageId     string `json:"packageId" binding:"required"`
}
