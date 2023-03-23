package domain

import "time"

type NewsArticle struct {
	Sentiment NewsArticleSentiment `json:"sentiment"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	Created   time.Time            `json:"created"`
	Views     int64                `json:"views"`
}

type NewsArticleSentiment string

const (
	NewsArticleSentimentNeutral  NewsArticleSentiment = "NEUTRAL"
	NewsArticleSentimentPositive NewsArticleSentiment = "POSITIVE"
	NewsArticleSentimentNegative NewsArticleSentiment = "NEGATIVE"
)
