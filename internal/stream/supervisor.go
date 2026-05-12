package stream

import (
    "time"

	"github.com/VigiaLink/agent/internal/camera"
    "github.com/VigiaLink/agent/internal/config"
    "github.com/VigiaLink/agent/internal/ffmpeg"
    "github.com/VigiaLink/agent/internal/logger"
)

func StartCameraLoop(cam camera.Camera, server config.Server) {
    go func() {
        for {
            logger.Log.Info().
                Str("camera", cam.Name).
                Msg("starting stream")

            cmd, err := ffmpeg.Start(
                cam.RTSP,
                server.Host,
                server.Port,
                cam.ID,
            )

            if err != nil {
                logger.Log.Error().Err(err).Msg("failed to start ffmpeg")
                time.Sleep(5 * time.Second)
                continue
            }

            err = cmd.Wait()

            logger.Log.Error().
                Err(err).
                Str("camera", cam.Name).
                Msg("stream stopped")

            time.Sleep(5 * time.Second)
        }
    }()
}