package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/pranjalbarnwal/olx-api/internal/config"
	"github.com/pranjalbarnwal/olx-api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	fmt.Println("OLX server starting...")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server running at PORT %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
