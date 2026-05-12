package camera

import (
    "github.com/VigiaLink/agent/internal/config"
)

type Manager struct {
    Cameras []Camera
}

func NewManager(cfg []config.Camera) *Manager {
    cameras := make([]Camera, 0)

    for _, c := range cfg {
        cameras = append(cameras, Camera{
            ID:     c.ID,
            Name:   c.Name,
            RTSP:   c.RTSP,
            Status: "offline",
        })
    }

    return &Manager{
        Cameras: cameras,
    }
}