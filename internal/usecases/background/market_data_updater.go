package background

import "websmee/buyspot/internal/usecases"

type MarketDataUpdater struct {
	assetRepository  usecases.AssetRepository
	marketDataStream usecases.MarketDataStream
}
