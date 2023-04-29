package domain

import "time"

type NewsArticle struct {
	URL       string               `bson:"url"`
	ImageURL  string               `bson:"image_url"`
	Source    string               `bson:"source"`
	Symbols   []string             `bson:"symbols"`
	Topics    []string             `bson:"topics"`
	Sentiment NewsArticleSentiment `bson:"sentiment"`
	Title     string               `bson:"title"`
	Content   string               `bson:"content"`
	Summary   string               `bson:"summary"`
	Created   time.Time            `bson:"created"`
	Views     int64                `bson:"views"`
}

type NewsArticleSentiment string

const (
	NewsArticleSentimentNeutral  NewsArticleSentiment = "Neutral"
	NewsArticleSentimentPositive NewsArticleSentiment = "Positive"
	NewsArticleSentimentNegative NewsArticleSentiment = "Negative"
)

type NewsPeriod string

const (
	NewsPeriodLastHour  NewsPeriod = "LastHour"
	NewsPeriodLastMonth NewsPeriod = "LastMonth"
)
