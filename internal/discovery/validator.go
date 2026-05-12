package discovery

import (
	"context"
	"os/exec"
	"time"

	"github.com/VigiaLink/agent/internal/logger"
)

func ValidateRTSP(url string) bool {
    ctx, cancel := context.WithTimeout(
        context.Background(),
        5*time.Second,
    )

    defer cancel()

    cmd := exec.CommandContext(
        ctx,
        "ffprobe",
        "-rtsp_transport",
        "tcp",
        "-v",
        "quiet",
        url,
    )

    logger.Log.Info().
    Str("rtsp", url).
    Msg("testing rtsp")

    err := cmd.Run()

    return err == nil
}