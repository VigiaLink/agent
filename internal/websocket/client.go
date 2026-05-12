package websocket

import (
    "time"

    "github.com/gorilla/websocket"
    "github.com/VigiaLink/agent/internal/logger"
)

func Start(url string) {
    go func() {
        for {
            conn, _, err := websocket.DefaultDialer.Dial(url, nil)
            if err != nil {
                logger.Log.Error().Err(err).Msg("websocket connection failed")
                time.Sleep(5 * time.Second)
                continue
            }

            logger.Log.Info().Msg("websocket connected")

            for {
                _, _, err := conn.ReadMessage()
                if err != nil {
                    logger.Log.Error().Err(err).Msg("websocket disconnected")
                    conn.Close()
                    break
                }
            }
        }
    }()
}