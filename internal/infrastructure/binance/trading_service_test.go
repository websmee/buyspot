package binance

import (
	"context"
	"fmt"
	"testing"

	"websmee/buyspot/internal/domain"
)

func TestBuySell(t *testing.T) {
	ctx := context.Background()
	tradingService := NewTradingService()
	user := &domain.User{
		BinanceAPIKey:    "EyXXD6gh4EpfFOLilEb45IeFFFJHyAk6IfmVV926d5oHqoRjcqzkuQlHYvObF6S1",
		BinanceSecretKey: "3DVkRLaoG9AxvVwbgYwDVoQs7rcmZLXbGdTVrTQ19fPNM2THJfR7dvFXw8MEBEwB",
	}

	//fmt.Println(tradingService.Buy(ctx, user, "USDT", 98, "BTC"))
	fmt.Println(tradingService.Sell(ctx, user, "BTC", 0.0039, "USDT"))
}
