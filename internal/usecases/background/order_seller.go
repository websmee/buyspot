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
	userRepository          usecases.UserRepository
	balanceService          usecases.BalanceService
	assetRepository         usecases.AssetRepository
	currentPricesRepository usecases.CurrentPricesRepository
	orderRepository         usecases.OrderRepository
	tradingService          usecases.TradingService
	notifier                usecases.Notifier
	logger                  *log.Logger
}

func NewOrderSeller(
	userRepository usecases.UserRepository,
	balanceService usecases.BalanceService,
	assetRepository usecases.AssetRepository,
	currentPricesRepository usecases.CurrentPricesRepository,
	orderRepository usecases.OrderRepository,
	tradingService usecases.TradingService,
	notifier usecases.Notifier,
	logger *log.Logger,
) *OrderSeller {
	return &OrderSeller{
		userRepository,
		balanceService,
		assetRepository,
		currentPricesRepository,
		orderRepository,
		tradingService,
		notifier,
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
			for j := range assets {
				price, err := r.currentPricesRepository.GetPrice(ctx, assets[j].Symbol, balanceSymbols[i])
				if err != nil {
					r.logger.Println(fmt.Errorf(
						"could not get %s%s price, err: %w",
						assets[j].Symbol,
						balanceSymbols[i],
						err,
					))
					continue
				}

				if price == 0 {
					r.logger.Println(fmt.Errorf(
						"could not find %s%s price",
						assets[j].Symbol,
						balanceSymbols[i],
					))
					continue
				}

				orders, err := r.orderRepository.GetActiveOrdersToSell(ctx, balanceSymbols[i], assets[j].Symbol, price)
				if err != nil {
					r.logger.Println(fmt.Errorf(
						"could not get active %s%s orders to sell, err: %w",
						assets[j].Symbol,
						balanceSymbols[i],
						err,
					))
					continue
				}

				for k := range orders {
					user, err := r.userRepository.GetByID(ctx, orders[k].UserID)
					if err != nil {
						r.logger.Println(fmt.Errorf(
							"could not find user by ID = '%s', err: %w",
							orders[k].UserID,
							err,
						))
						continue
					}

					closeAmount, err := r.tradingService.Sell(
						ctx,
						user.ID.Hex(),
						orders[k].ToSymbol,
						orders[k].ToAmount,
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

					if err := r.notify(ctx, user, &orders[k]); err != nil {
						r.logger.Println(fmt.Errorf("could not notify order sell, err: %w", err))
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

func (r *OrderSeller) notify(ctx context.Context, user *domain.User, order *domain.Order) error {
	if user.NotificationsKey == "" {
		return nil
	}

	if order == nil {
		return nil
	}

	return r.notifier.Notify(
		ctx,
		user,
		"ORDER CLOSED",
		fmt.Sprintf("%s: %f", order.ToAssetName, order.CloseAmount-order.FromAmount),
	)
}
