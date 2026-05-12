package discovery

import (
    "fmt"
    "net"
    "time"
)

var CameraPorts = []int{
    554,
    8554,
    80,
    8080,
    8000,
    8899,
}

func PortOpen(ip string, port int) bool {
    address := fmt.Sprintf("%s:%d", ip, port)

    conn, err := net.DialTimeout(
        "tcp",
        address,
        500*time.Millisecond,
    )

    if err != nil {
        return false
    }

    conn.Close()
    return true
}