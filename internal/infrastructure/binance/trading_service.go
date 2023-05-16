package binance

import (
	"context"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"

	"websmee/buyspot/internal/domain"
)

type TradingService struct{}

func NewTradingService() *TradingService {
	return &TradingService{}
}

func (s *TradingService) Buy(
	ctx context.Context,
	user *domain.User,
	balanceSymbol string,
	balanceAmount float64,
	tradeSymbol string,
) (string, error) {
	order, err := binance.NewClient(user.BinanceAPIKey, user.BinanceSecretKey).NewCreateOrderService().
		Symbol(tradeSymbol + balanceSymbol).
		Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).
		QuoteOrderQty(fmt.Sprintf("%f", balanceAmount)).
		Do(ctx)
	if err != nil {
		return "0", fmt.Errorf(
			"could not buy %s%s on binance, err: %w",
			tradeSymbol,
			balanceSymbol,
			err,
		)
	}

	if order.Status != binance.OrderStatusTypeFilled {
		return "0", fmt.Errorf(
			"could not buy %s%s on binance, order status: %s",
			tradeSymbol,
			balanceSymbol,
			order.Status,
		)
	}

	return order.ExecutedQuantity, nil
}

func (s *TradingService) Sell(
	ctx context.Context,
	user *domain.User,
	tradeSymbol string,
	tradeAmount string,
	balanceSymbol string,
) (float64, error) {
	if user.BinanceAPIKey == "" || user.BinanceSecretKey == "" {
		return 0, fmt.Errorf("user '%s' does not have binance api access", user.ID.Hex())
	}

	order, err := binance.NewClient(user.BinanceAPIKey, user.BinanceSecretKey).NewCreateOrderService().
		Symbol(tradeSymbol + balanceSymbol).
		Side(binance.SideTypeSell).
		Type(binance.OrderTypeMarket).
		Quantity(tradeAmount).
		Do(ctx)
	if err != nil {
		return 0, fmt.Errorf(
			"could not sell %s%s on binance, err: %w",
			tradeSymbol,
			balanceSymbol,
			err,
		)
	}

	if order.Status != binance.OrderStatusTypeFilled {
		return 0, fmt.Errorf(
			"could not sell %s%s on binance, order status: %s",
			tradeSymbol,
			balanceSymbol,
			order.Status,
		)
	}

	balanceAmount := 0.0
	for i := range order.Fills {
		q, _ := strconv.ParseFloat(order.Fills[i].Quantity, 64)
		p, _ := strconv.ParseFloat(order.Fills[i].Price, 64)
		balanceAmount += q * p
	}

	return balanceAmount, nil
}
