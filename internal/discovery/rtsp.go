package discovery

import "fmt"

var RTSPPaths = []string{
    "/test",
    "/Streaming/Channels/101",
    "/cam/realmonitor?channel=1&subtype=0",
    "/h264/ch1/main/av_stream",
    "/live/ch00_0",
}

func BuildRTSP(ip string, username string, password string) []string {
    urls := []string{}

    for _, path := range RTSPPaths {
        url := fmt.Sprintf(
            "rtsp://%s:%s@%s:554%s",
            username,
            password,
            ip,
            path,
        )

        urls = append(urls, url)
    }

    return urls
}