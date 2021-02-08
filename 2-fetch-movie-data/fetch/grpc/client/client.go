package client

import (
	"context"
	"os"

	"github.com/stockbit/2-fetch-movie-data/fetch/grpc/pb"
	"google.golang.org/grpc"
)

var loggerConn *grpc.ClientConn

func ConnectLogger() (err error) {
	loggerConn, err = grpc.Dial(os.Getenv("LOGGER_DNS"), grpc.WithInsecure())
	return err
}

func Logging(logMsg *pb.Log) error {
	_, err := pb.NewLoggingClient(loggerConn).InsertLog(context.Background(), logMsg)
	return err
}
