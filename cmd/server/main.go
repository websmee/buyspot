package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	httpAPI "websmee/buyspot/internal/api/http"
	"websmee/buyspot/internal/infrastructure/mock"
	redisInfra "websmee/buyspot/internal/infrastructure/redis"
	"websmee/buyspot/internal/usecases"
)

func main() {
	ctx := context.Background()

	logger := newLogger("[MAIN]")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	currentSpotsRepository := redisInfra.NewCurrentSpotsRepository(client)
	spotReader := usecases.NewSpotReader(currentSpotsRepository)
	spotMaker := usecases.NewSpotMaker(
		currentSpotsRepository,
		mock.NewMarketDataRepository(),
		newLogger("[SPOT MAKER]"),
	)

	router := gin.Default()
	router.Use(httpAPI.CORSMiddleware())
	httpHandlersLogger := newLogger("[HTTP HANDLER]")
	httpAPI.AddBalanceHandlers(ctx, router)
	httpAPI.AddSpotHandlers(ctx, router, spotReader, httpHandlersLogger)

	if err := spotMaker.Run(ctx); err != nil {
		logger.Fatalln(fmt.Errorf("could not run spot maker, err: %w", err))
	}

	_ = router.Run("localhost:8080")
}

func newLogger(name string) *log.Logger {
	return log.New(os.Stdout, name+" ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
