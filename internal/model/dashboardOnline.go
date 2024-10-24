package model

type Online struct {
	OnlineUsers            int `json:"online_users"`
	OrdersCount            int `json:"orders_count"`
	OrdersCollectedCount   int `json:"orders_collectedcount"`
	OrdersUncollectedCount int `json:"orders_uncollectedcount"`
}
