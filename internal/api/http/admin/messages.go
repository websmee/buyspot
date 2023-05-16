package admin

type User struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	BinanceAPIKey    string `json:"binanceAPIKey"`
	BinanceSecretKey string `json:"binanceSecretKey"`
	NotificationsKey string `json:"notificationsKey"`
	IsDemo           bool   `json:"isDemo"`
}

type Balance struct {
	UserID   string  `json:"userID"`
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount"`
	IsActive bool    `json:"isActive"`
}
