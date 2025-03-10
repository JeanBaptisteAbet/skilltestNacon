package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"skilltestnacon/api"
	pb "skilltestnacon/api/grpcserver/liveevents"
	grpcliveeventserver "skilltestnacon/api/grpcserver/server"
	"skilltestnacon/api/httphandler"
	"skilltestnacon/constant"
	"skilltestnacon/database"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	db, err := database.InitDB("./app.db")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterLiveEventsServiceServer(grpcServer, &grpcliveeventserver.Server{DB: db})

	hs := httphandler.Handler{
		DB:      db,
		Context: context.Background(),
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", os.Getenv(constant.API_PORT)),
		Handler: h2c.NewHandler(http.DefaultServeMux, &http2.Server{}),
	}

	http.Handle("/",
		api.Protect(
			grpcServer, os.Getenv(constant.GRPC_API_KEY),
		),
	)
	http.Handle("/events",
		api.Protect(
			http.HandlerFunc(hs.HandleEvents), os.Getenv(constant.HTTP_API_KEY),
		),
	)
	http.Handle("/events/{id}",
		api.Protect(
			http.HandlerFunc(hs.HandleEvent), os.Getenv(constant.HTTP_API_KEY),
		),
	)

	log.Println("starting server")
	log.Println(s.ListenAndServe())
}
