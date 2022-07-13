package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pankajyadav2741/ott/pkg/model"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/users", model.HandleUser).Methods("POST")
	myRouter.HandleFunc("/users", model.HandleUser).Methods("DELETE")
	myRouter.HandleFunc("/contents", model.HandleContent).Methods("POST")
	myRouter.HandleFunc("/contents", model.HandleContent).Methods("DELETE")

	srv := &http.Server{
		Handler:      myRouter,
		Addr:         ":5000",
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("error: %w", err)
		}
	}()
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	intChan := make(chan os.Signal, 1)
	signal.Notify(intChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-intChan

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	os.Exit(0)
}
