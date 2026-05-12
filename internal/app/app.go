package app

import (
    "github.com/VigiaLink/agent/internal/camera"
    "github.com/VigiaLink/agent/internal/config"
    "github.com/VigiaLink/agent/internal/heartbeat"
    "github.com/VigiaLink/agent/internal/logger"
    "github.com/VigiaLink/agent/internal/stream"
    "github.com/VigiaLink/agent/internal/telemetry"
    "github.com/VigiaLink/agent/internal/websocket"
)

func Start(cfg *config.Config) {
    logger.Init()

    logger.Log.Info().Msg("starting monitra agent")

    telemetry.Start()
    heartbeat.Start(cfg.AgentID)
    websocket.Start(cfg.Server.WebSocket)

    manager := camera.NewManager(cfg.Cameras)

    for _, cam := range manager.Cameras {
        stream.StartCameraLoop(cam, cfg.Server)
    }

    select {}
}