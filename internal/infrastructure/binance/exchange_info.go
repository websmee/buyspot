package binance

import (
	"strings"

	"github.com/adshao/go-binance/v2"
)

func getPrecision(exchangeInfo *binance.ExchangeInfo, symbol string) int {
	for i := range exchangeInfo.Symbols {
		if exchangeInfo.Symbols[i].BaseAsset == symbol {
			lotSizeFilter := exchangeInfo.Symbols[i].LotSizeFilter()
			return strings.Index(lotSizeFilter.StepSize, "1") - 1
		}
	}

	return -1
}

func isAvailable(exchangeInfo *binance.ExchangeInfo, symbol string) bool {
	for i := range exchangeInfo.Symbols {
		if exchangeInfo.Symbols[i].BaseAsset == symbol {
			hasMarket := false
			for _, ot := range exchangeInfo.Symbols[i].OrderTypes {
				if ot == string(binance.OrderTypeMarket) {
					hasMarket = true
					break
				}
			}

			return exchangeInfo.Symbols[i].IsSpotTradingAllowed &&
				exchangeInfo.Symbols[i].QuoteOrderQtyMarketAllowed &&
				hasMarket
		}
	}

	return false
}
