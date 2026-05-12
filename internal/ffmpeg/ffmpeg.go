package ffmpeg

import (
    "fmt"
    "os"
    "os/exec"
)

func Start(rtsp string, host string, port int, streamID string) (*exec.Cmd, error) {
    output := fmt.Sprintf(
        "srt://%s:%d?streamid=publish:%s",
        host,
        port,
        streamID,
    )

    cmd := exec.Command(
        "ffmpeg",
        "-rtsp_transport", "tcp",
        "-i", rtsp,
        "-c:v", "libx264",
        "-preset", "veryfast",
        "-tune", "zerolatency",
        "-f", "mpegts",
        output,
    )

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Start()
    if err != nil {
        return nil, err
    }

    return cmd, nil
}