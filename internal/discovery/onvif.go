package discovery

import (
    "fmt"
    "net/http"
    "time"
)

func HasONVIF(ip string) bool {
    client := http.Client{
        Timeout: 3 * time.Second,
    }

    // url := fmt.Sprintf(
    //     "http://%s/onvif/device_service",
    //     ip,
    // )
    url := fmt.Sprintf(
        "rtsp://%s:8554%s",
        ip,
    )

    resp, err := client.Get(url)
    if err != nil {
        return false
    }

    defer resp.Body.Close()

    return resp.StatusCode == 200
}