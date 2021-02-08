package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stockbit/2-fetch-movie-data/fetch/grpc/client"
	"github.com/stockbit/2-fetch-movie-data/fetch/http/route"
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

	if err := client.ConnectLogger(); err != nil {
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
		rtr := route.NewRouter()
		if err := rtr.Run(":" + os.Getenv("PORT")); err != nil {
			errs <- err
		}
	}()

	log.Fatal(<-errs)
}
