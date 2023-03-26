package background

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type OrderSeller struct {
	userRepository   usecases.UserRepository
	balanceService   usecases.BalanceService
	assetRepository  usecases.AssetRepository
	pricesService    usecases.PricesService
	orderRepository  usecases.OrderRepository
	converterService usecases.ConverterService
	logger           *log.Logger
}

func NewOrderSeller(
	userRepository usecases.UserRepository,
	balanceService usecases.BalanceService,
	assetRepository usecases.AssetRepository,
	pricesService usecases.PricesService,
	orderRepository usecases.OrderRepository,
	converterService usecases.ConverterService,
	logger *log.Logger,
) *OrderSeller {
	return &OrderSeller{
		userRepository,
		balanceService,
		assetRepository,
		pricesService,
		orderRepository,
		converterService,
		logger,
	}
}

func (r *OrderSeller) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Minute).Do(func() {
		r.logger.Println("selling orders")

		assets, err := r.assetRepository.GetAvailableAssets(ctx)
		if err != nil {
			r.logger.Println(fmt.Errorf("could not get assets to sell orders, err: %w", err))
			return
		}

		balanceTickers, err := r.balanceService.GetAvailableTickers(ctx)
		if err != nil {
			r.logger.Println(fmt.Errorf("could not get available balance tickers, err: %w", err))
			return
		}

		for i := range balanceTickers {
			currentPrices, err := r.pricesService.GetCurrentPrices(ctx, balanceTickers[i])
			if err != nil {
				r.logger.Println(fmt.Errorf(
					"could not get current prices for ticker '%s', err: %w",
					balanceTickers[i],
					err,
				))
				continue
			}

			for j := range assets {
				tickerCurrentPrice, ok := currentPrices.PricesByTickers[assets[j].Ticker]
				if !ok {
					r.logger.Println(fmt.Errorf(
						"could not get current price for ticker '%s'",
						assets[j].Ticker,
					))
					continue
				}

				orders, err := r.orderRepository.GetActiveOrdersToSell(ctx, balanceTickers[i], assets[j].Ticker, tickerCurrentPrice)
				if err != nil {
					r.logger.Println(fmt.Errorf(
						"could not get active %s orders to sell for %s, err: %w",
						assets[j].Ticker,
						balanceTickers[i],
						err,
					))
					continue
				}

				for k := range orders {
					user, err := r.userRepository.GetUserByID(ctx, orders[k].UserID)
					if err != nil {
						r.logger.Println(fmt.Errorf(
							"could not find user by ID = '%s', err: %w",
							orders[k].UserID,
							err,
						))
						continue
					}

					closeAmount, err := r.converterService.Convert(
						ctx,
						user,
						orders[k].ToAmount,
						orders[k].ToTicker,
						balanceTickers[i],
					)
					if err != nil {
						r.logger.Println(fmt.Errorf(
							"could not convert %s to %s, err: %w",
							orders[k].ToTicker,
							balanceTickers[i],
							err,
						))
						continue
					}

					orders[k].CloseAmount = closeAmount
					orders[k].CloseTicker = balanceTickers[i]
					orders[k].Updated = time.Now()
					orders[k].Status = domain.OrderStatusClosed
					if err := r.orderRepository.SaveOrder(ctx, &orders[k]); err != nil {
						r.logger.Println(fmt.Errorf("could not save order after conversion, err: %w", err))
						continue
					}
				}
			}
		}

		r.logger.Println("orders sold")
	})

	s.StartAsync()

	return err
}
