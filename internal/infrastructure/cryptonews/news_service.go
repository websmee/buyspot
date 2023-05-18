package cryptonews

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"websmee/buyspot/internal/domain"
)

const apiBaseURL = "https://cryptonews-api.com/api/v1"

type NewsService struct {
	token string
}

func NewNewsService(token string) *NewsService {
	return &NewsService{token}
}

type NewsArticleResponse struct {
	Data       []NewsArticle `json:"data"`
	TotalPages int           `json:"total_pages"`
}

type NewsArticle struct {
	NewsUrl    string   `json:"news_url"`
	ImageUrl   string   `json:"image_url"`
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	SourceName string   `json:"source_name"`
	Date       string   `json:"date"`
	Topics     []string `json:"topics"`
	Sentiment  string   `json:"sentiment"`
	Type       string   `json:"type"`
	Tickers    []string `json:"tickers"`
}

func (s *NewsService) GetNews(ctx context.Context, symbols []string, period domain.NewsPeriod) ([]domain.NewsArticle, error) {
	var news []domain.NewsArticle
	for i := 0; i < len(symbols); i = i + 20 {
		j := i + 20
		if i+20 > len(symbols) {
			j = len(symbols)
		}

		data, err := s.makeRequest(ctx, symbols[i:j], period)
		if err != nil {
			return nil, fmt.Errorf("could not make news api request, err: %w", err)
		}

		var response NewsArticleResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, fmt.Errorf("could not unmarshal news api response, err: %w", err)
		}

		news = append(news, apiResponseToNewsArticles(response.Data)...)
	}

	return news, nil
}

func (s *NewsService) makeRequest(ctx context.Context, symbols []string, period domain.NewsPeriod) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiBaseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create http request, err: %w", err)
	}

	req.URL.RawQuery = fmt.Sprintf(
		"tickers=%s&items=100&page=1&period=%s&token=%s",
		strings.Join(symbols, ","),
		periodToAPIParam(period),
		s.token,
	)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not make http request, err: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http request failed with status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read http response, err: %w", err)
	}

	return data, nil
}

func periodToAPIParam(period domain.NewsPeriod) string {
	switch period {
	case domain.NewsPeriodLastHour:
		return "last60min"
	case domain.NewsPeriodLastMonth:
		return "last30days"
	default:
		return "last60min"
	}
}

func apiResponseToNewsArticles(apiNewsArticles []NewsArticle) []domain.NewsArticle {
	var news []domain.NewsArticle
	for _, apiNewsArticle := range apiNewsArticles {
		date, err := time.Parse(time.RFC1123Z, apiNewsArticle.Date)
		if err != nil {
			date = time.Now()
		}

		news = append(news, domain.NewsArticle{
			URL:       apiNewsArticle.NewsUrl,
			ImageURL:  apiNewsArticle.ImageUrl,
			Source:    apiNewsArticle.SourceName,
			Symbols:   apiNewsArticle.Tickers,
			Topics:    apiNewsArticle.Topics,
			Sentiment: apiResponseToSentiment(apiNewsArticle.Sentiment),
			Title:     apiNewsArticle.Title,
			Content:   apiNewsArticle.Text,
			Created:   date,
		})
	}

	return news
}

func apiResponseToSentiment(apiSentiment string) domain.NewsArticleSentiment {
	switch apiSentiment {
	case "Positive":
		return domain.NewsArticleSentimentPositive
	case "Negative":
		return domain.NewsArticleSentimentNegative
	default:
		return domain.NewsArticleSentimentNeutral
	}
}
