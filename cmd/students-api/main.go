package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/atindraraut/crudgo/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoadConfig()
	//database setup

	//setup routes
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World! from go project"))
	})
	//setup server
	server := &http.Server{
		Addr:    cfg.HTTPServer.ADDR,
		Handler: router,
	}
	slog.Info("Starting server...", slog.String("address", cfg.HTTPServer.ADDR))
	done:= make(chan os.Signal, 1)
	signal.Notify(done,os.Interrupt,syscall.SIGTERM, syscall.SIGINT)
	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("failed to start server: %s", err.Error())
		}
	}()
	<-done

	slog.Info("Shutting down server...")
	ctx,cancel :=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server",slog.String("error", err.Error()))
	} else {
		slog.Info("Server stopped gracefully")
	}
}
