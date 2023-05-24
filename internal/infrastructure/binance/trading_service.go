package binance

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
	if user.BinanceAPIKey == "" || user.BinanceSecretKey == "" {
		return "", fmt.Errorf("user '%s' does not have binance api access", user.ID.Hex())
	}

	client := binance.NewClient(user.BinanceAPIKey, user.BinanceSecretKey)

	exchangeInfo, err := client.NewExchangeInfoService().Symbols(tradeSymbol + balanceSymbol).Do(ctx)
	if err != nil {
		return "", fmt.Errorf("could not get exchange info from binance, err: %w", err)
	}

	if !isAvailable(exchangeInfo, tradeSymbol) {
		return "", domain.ErrTradingNotAvailable
	}

	order, err := client.NewCreateOrderService().
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

	boughtAmount := 0.0
	for i := range order.Fills {
		p, _ := strconv.ParseFloat(order.Fills[i].Price, 64)
		q, _ := strconv.ParseFloat(order.Fills[i].Quantity, 64)
		c, _ := strconv.ParseFloat(order.Fills[i].Commission, 64)

		if order.Fills[i].CommissionAsset == balanceSymbol {
			c = c / p
		} else if order.Fills[i].CommissionAsset != tradeSymbol {
			c = 0
			fmt.Printf("invalid commission asset %s for %s\n", tradeSymbol, order.Fills[i].CommissionAsset)
		}

		boughtAmount += q - c
	}

	return fmt.Sprintf("%f", boughtAmount), nil
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

	client := binance.NewClient(user.BinanceAPIKey, user.BinanceSecretKey)

	exchangeInfo, err := client.NewExchangeInfoService().Symbols(tradeSymbol + balanceSymbol).Do(ctx)
	if err != nil {
		return 0, fmt.Errorf("could not get exchange info from binance, err: %w", err)
	}

	precision := getPrecision(exchangeInfo, tradeSymbol)

	if precision > 0 {
		a := strings.Split(tradeAmount, ".")
		if len(a) == 2 && len(a[1]) > precision {
			tradeAmount = a[0] + "." + a[1][:precision]
		}
	} else {
		tradeAmount = strings.Split(tradeAmount, ".")[0]
	}

	order, err := client.NewCreateOrderService().
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
