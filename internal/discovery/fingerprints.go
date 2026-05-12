package discovery

import "strings"

func Fingerprint(ip string) string {
    if strings.Contains(ip, "192.168") {
        return "Unknown"
    }

    return "Unknown"
}