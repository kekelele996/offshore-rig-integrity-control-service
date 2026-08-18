package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/api"
	"github.com/kekelele996/offshore-rig-integrity-control-service/internal/service"
)

func main() {
	addr := os.Getenv("RIG_INTEGRITY_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	system := service.NewSystem(time.Now)
	server := api.NewServer(system)
	log.Printf("rig integrity control listening on %s", addr)
	if err := http.ListenAndServe(addr, server.Routes()); err != nil {
		log.Fatal(err)
	}
}
