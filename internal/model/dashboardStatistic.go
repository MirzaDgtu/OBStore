package model

type Statistic struct {
	UserID               int     `json:"user_id"`
	UserName             string  `json:"user_name"`
	OrdersCollectedCount int     `json:"orders_collectedcount"`
	PackageCount         float32 `json:"package_count"`
	TotalWeight          float32 `json:"total_weight"`
	TotalSum             float32 `json:"total_sum"`
	ProductCount         float32 `json:"product_count"`
}
