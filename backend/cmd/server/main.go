package main

import (
    "flag"
    "log"

    "B2CTF/backend/internal/bootstrap"
)

func main() {
    configPath := flag.String("config", "configs/config.yaml", "config file path")
    flag.Parse()

    if err := bootstrap.Run(*configPath); err != nil {
        log.Fatalf("server exit with error: %v", err)
    }
}
