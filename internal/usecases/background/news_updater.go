package background

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type NewsUpdater struct {
	assetRepository usecases.AssetRepository
	newsRepository  usecases.NewsRepository
	newsService     usecases.NewsService
	summarizer      usecases.Summarizer
	logger          *log.Logger
}

func NewNewsUpdater(
	assetRepository usecases.AssetRepository,
	newsRepository usecases.NewsRepository,
	newsService usecases.NewsService,
	summarizer usecases.Summarizer,
	logger *log.Logger,
) *NewsUpdater {
	return &NewsUpdater{
		assetRepository,
		newsRepository,
		newsService,
		summarizer,
		logger,
	}
}

func (u *NewsUpdater) Run(ctx context.Context) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(15 * time.Minute).Do(func() {
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

		var wg sync.WaitGroup
		for i := range news {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				exists, err := u.newsRepository.IsArticleExists(ctx, &news[i])
				if err != nil {
					u.logger.Println(fmt.Errorf("could not check if news article exists, err: %w", err))
					return
				}
				if exists {
					return
				}

				summary, err := u.summarizer.GetSummary(ctx, news[i].URL)
				if err != nil {
					u.logger.Println(fmt.Errorf("could not summarize news artice, err: %w", err))
				}
				news[i].Summary = summary

				if err := u.newsRepository.CreateOrUpdate(ctx, &news[i]); err != nil {
					u.logger.Println(fmt.Errorf("could save imported news article, err: %w", err))
				}
			}(i)
		}
		wg.Wait()

		u.logger.Println("news updated")
	})

	s.StartAsync()

	return err
}
