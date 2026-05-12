package main

import (
	"log"

	"github.com/VigiaLink/agent/internal/app"
	"github.com/VigiaLink/agent/internal/logger"
	"github.com/VigiaLink/agent/internal/config"
	"github.com/VigiaLink/agent/internal/discovery"
)

func main() {
    log.Println("starting auto discovery")

    found := discovery.Discover()

    log.Printf("found devices: %d\n", len(found))

    for _, device := range found {
        logger.Log.Info().
            Str("ip", device.IP).
            Str("rtsp", device.RTSP).
            Msg("auto discovered camera")
    }

    cfg, err := config.Load("configs/config.json")
    if err != nil {
        log.Fatal(err)
    }

    app.Start(cfg)
}