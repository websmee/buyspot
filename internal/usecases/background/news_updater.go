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

type NewsUpdater struct {
	assetRepository usecases.AssetRepository
	newsRepository  usecases.NewsRepository
	newsService     usecases.NewsService
	logger          *log.Logger
}

func NewNewsUpdater(
	assetRepository usecases.AssetRepository,
	newsRepository usecases.NewsRepository,
	newsService usecases.NewsService,
	logger *log.Logger,
) *NewsUpdater {
	return &NewsUpdater{
		assetRepository,
		newsRepository,
		newsService,
		logger,
	}
}

func (u *NewsUpdater) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(time.Hour).Do(func() {
		u.logger.Println("updating news")

		assets, err := u.assetRepository.GetAvailableAssets(ctx)
		if err != nil {
			u.logger.Println(fmt.Errorf("could get available assets, err: %w", err))
			return
		}

		var symbols []string
		for i := range assets {
			symbols = append(symbols, assets[i].Symbol)
		}

		news, err := u.newsService.GetNews(ctx, symbols, domain.NewsPeriodLastHour)
		if err != nil {
			u.logger.Println(fmt.Errorf("could get last hour news, err: %w", err))
			return
		}

		for i := range news {
			if err := u.newsRepository.CreateOrUpdate(ctx, &news[i]); err != nil {
				u.logger.Println(fmt.Errorf("could save imported news article, err: %w", err))
				continue
			}
		}

		u.logger.Println("news updated")
	})

	s.StartAsync()

	return err
}
