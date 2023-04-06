package admin

import (
	"context"
	"fmt"

	"websmee/buyspot/internal/domain"
	"websmee/buyspot/internal/usecases"
)

type NewsUpdater struct {
	secretKey       string
	assetRepository usecases.AssetRepository
	newsRepository  usecases.NewsRepository
	newsService     usecases.NewsService
}

func NewNewsUpdater(
	secretKey string,
	assetRepository usecases.AssetRepository,
	newsRepository usecases.NewsRepository,
	newsService usecases.NewsService,
) *NewsUpdater {
	return &NewsUpdater{
		secretKey,
		assetRepository,
		newsRepository,
		newsService,
	}
}

func (r *NewsUpdater) Update(ctx context.Context, secretKey string) error {
	if secretKey != r.secretKey {
		return domain.ErrForbidden
	}

	assets, err := r.assetRepository.GetAvailableAssets(ctx)
	if err != nil {
		return fmt.Errorf("could not get available assets, err: %w", err)
	}

	for i := range assets {
		news, err := r.newsService.GetNews(ctx, []string{assets[i].Symbol}, domain.NewsPeriodLastMonth)
		if err != nil {
			return fmt.Errorf("could not get last month %s news, err: %w", assets[i].Symbol, err)
		}

		for i := range news {
			if err := r.newsRepository.CreateOrUpdate(ctx, &news[i]); err != nil {
				return fmt.Errorf("could not save imported news article, err: %w", err)
			}
		}
	}

	return nil
}
