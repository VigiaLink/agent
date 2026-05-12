package discovery

import (
    "fmt"
    "strings"
    "sync"
)

func ScanSubnet(cidr string) []string {
    jobs := make(chan string, 256)
    results := make(chan string, 256)

    var wg sync.WaitGroup

    worker := func() {
        defer wg.Done()

        for ip := range jobs {
            for _, port := range CameraPorts {
                if PortOpen(ip, port) {
                    results <- ip
                    break
                }
            }
        }
    }

    for i := 0; i < 64; i++ {
        wg.Add(1)
        go worker()
    }

    parts := strings.Split(cidr, ".")

    if len(parts) < 4 {
        return []string{}
    }

    base := fmt.Sprintf(
        "%s.%s.%s",
        parts[0],
        parts[1],
        parts[2],
    )

    for i := 1; i <= 254; i++ {
        jobs <- fmt.Sprintf("%s.%d", base, i)
    }

    close(jobs)

    go func() {
        wg.Wait()
        close(results)
    }()

    active := []string{}

    for ip := range results {
        active = append(active, ip)
    }

    return active
}