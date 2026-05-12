package heartbeat

import (
    "time"

    "github.com/VigiaLink/agent/internal/logger"
)

func Start(agentID string) {
    go func() {
        for {
            logger.Log.Info().
                Str("agent", agentID).
                Msg("heartbeat")

            time.Sleep(10 * time.Second)
        }
    }()
}