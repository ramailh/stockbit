package server

import (
	"context"
	"log"

	"github.com/ramailh/stockbit/2-fetch-movie-data/logger/grpc/pb"
	"github.com/ramailh/stockbit/2-fetch-movie-data/logger/model"
	"google.golang.org/grpc"
)

type logServer struct {
}

func (lgs logServer) InsertLog(ctx context.Context, req *pb.Log) (*pb.Response, error) {
	logger := model.Logger{
		Method:     req.Method,
		ClientIP:   req.ClientIp,
		Latency:    float64(req.Latency),
		Path:       req.Path,
		StatusCode: int(req.StatusCode),
		TimeStamp:  req.Timestamp,
	}

	if err := logger.Insert(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.Response{Message: "success", Status: true}, nil
}

func NewLogServer() *grpc.Server {
	srv := grpc.NewServer()
	pb.RegisterLoggingServer(srv, &logServer{})
	return srv
}
