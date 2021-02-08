package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ramailh/stockbit/2-fetch-movie-data/logger/grpc/server"
	"github.com/ramailh/stockbit/2-fetch-movie-data/logger/repository/postgres"
	"github.com/subosito/gotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if err := gotenv.Load(env); err != nil {
		log.Println(err)
	}

	if err := postgres.ConnectToPostgres(); err != nil {
		log.Println(err)
	}
}

func main() {
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%v", <-c)
	}()

	go func() {
		lis, err := net.Listen("tcp", ":57600")
		if err != nil {
			errs <- err
		}

		fmt.Println("Running on port 57600")

		srv := server.NewLogServer()
		if err = srv.Serve(lis); err != nil {
			errs <- err
		}
	}()

	log.Fatal(<-errs)
}
