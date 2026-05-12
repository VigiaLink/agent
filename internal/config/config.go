package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    AgentID string   `json:"agent_id"`
    Server  Server   `json:"server"`
    Cameras []Camera `json:"cameras"`
}

type Server struct {
    Host      string `json:"host"`
    Port      int    `json:"port"`
    WebSocket string `json:"websocket"`
}

type Camera struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    RTSP string `json:"rtsp"`
}

func Load(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var cfg Config

    err = json.Unmarshal(data, &cfg)
    if err != nil {
        return nil, err
    }

    return &cfg, nil
}