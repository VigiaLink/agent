package telemetry

import (
    "time"

    "github.com/VigiaLink/agent/internal/logger"
    "github.com/shirou/gopsutil/v4/cpu"
    "github.com/shirou/gopsutil/v4/mem"
)

func Start() {
    go func() {
        for {
            cpuPercent, _ := cpu.Percent(0, false)
            memStats, _ := mem.VirtualMemory()

            logger.Log.Info().
                Float64("cpu", cpuPercent[0]).
                Float64("memory", memStats.UsedPercent).
                Msg("telemetry")

            time.Sleep(15 * time.Second)
        }
    }()
}