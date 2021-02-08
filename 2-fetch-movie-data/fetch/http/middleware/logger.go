package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stockbit/2-fetch-movie-data/fetch/grpc/client"
	"github.com/stockbit/2-fetch-movie-data/fetch/grpc/pb"
)

func LoggerMiddleware(params gin.LogFormatterParams) string {
	err := client.Logging(&pb.Log{
		Method:     params.Method,
		ClientIp:   params.ClientIP,
		Latency:    float32(params.Latency.Seconds()),
		Path:       params.Path,
		StatusCode: int64(params.StatusCode),
		Timestamp:  params.TimeStamp.Unix(),
	})
	if err != nil {
		log.Println(err)
	}

	return ""
}
