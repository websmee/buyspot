package example

import (
	"context"
	"time"

	"websmee/buyspot/internal/domain"
)

type NewsRepository struct {
}

func NewNewsRepository() *NewsRepository {
	return &NewsRepository{}
}

func (r *NewsRepository) GetFreshNewsBySymbol(_ context.Context, _ string, _ time.Time) ([]domain.NewsArticle, error) {
	return []domain.NewsArticle{
		{
			Sentiment: domain.NewsArticleSentimentNeutral,
			Title:     "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
			Content:   "Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain. Ark Invest on $1.4M BTC, Possible Julian Assange-Linked File on Bitcoin Blockchain",
			Created:   time.Now().Add(-1 * time.Hour),
			Views:     15678,
		},
		{
			Sentiment: domain.NewsArticleSentimentPositive,
			Title:     "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
			Content:   "IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed. IMF Report on El Salvador's Bitcoin Adoption: Risks Averted, but Transparency Needed",
			Created:   time.Now().Add(-12 * time.Hour),
			Views:     25678,
		},
		{
			Sentiment: domain.NewsArticleSentimentNegative,
			Title:     "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
			Content:   "Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown. Daily Raids on Crypto Farms as Abkhazia Intensifies Mining Crackdown",
			Created:   time.Now().Add(-48 * time.Hour),
			Views:     178,
		},
	}, nil
}
