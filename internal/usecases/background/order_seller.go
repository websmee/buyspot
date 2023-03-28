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

		balanceSymbols, err := r.balanceService.GetAvailableSymbols(ctx)
		if err != nil {
			r.logger.Println(fmt.Errorf("could not get available balance symbols, err: %w", err))
			return
		}

		for i := range balanceSymbols {
			currentPrices, err := r.pricesService.GetCurrentPrices(ctx, balanceSymbols[i])
			if err != nil {
				r.logger.Println(fmt.Errorf(
					"could not get current prices for symbol '%s', err: %w",
					balanceSymbols[i],
					err,
				))
				continue
			}

			for j := range assets {
				symbolCurrentPrice, ok := currentPrices.PricesBySymbols[assets[j].Symbol]
				if !ok {
					r.logger.Println(fmt.Errorf(
						"could not get current price for symbol '%s'",
						assets[j].Symbol,
					))
					continue
				}

				orders, err := r.orderRepository.GetActiveOrdersToSell(ctx, balanceSymbols[i], assets[j].Symbol, symbolCurrentPrice)
				if err != nil {
					r.logger.Println(fmt.Errorf(
						"could not get active %s orders to sell for %s, err: %w",
						assets[j].Symbol,
						balanceSymbols[i],
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
						orders[k].ToSymbol,
						balanceSymbols[i],
					)
					if err != nil {
						r.logger.Println(fmt.Errorf(
							"could not convert %s to %s, err: %w",
							orders[k].ToSymbol,
							balanceSymbols[i],
							err,
						))
						continue
					}

					orders[k].CloseAmount = closeAmount
					orders[k].CloseSymbol = balanceSymbols[i]
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
