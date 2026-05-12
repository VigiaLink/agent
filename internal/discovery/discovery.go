package discovery

import (
	"fmt"

	"github.com/VigiaLink/agent/internal/logger"
)

func Discover() []Device {
	devices := []Device{}

	fmt.Println("DISCOVERY STARTED")

	subnets, err := LocalSubnets()
	if err != nil {
		fmt.Println("SUBNET ERROR:", err)
		return devices
	}

	fmt.Println("SUBNETS FOUND:", subnets)

	for _, subnet := range subnets {

		fmt.Println("SCANNING SUBNET:", subnet)

		logger.Log.Info().
			Str("subnet", subnet).
			Msg("scanning subnet")

		hosts := ScanSubnet(subnet)

		fmt.Println("HOSTS FOUND:", len(hosts))

		for _, host := range hosts {

			fmt.Println("HOST:", host)

			logger.Log.Info().
				Str("ip", host).
				Msg("possible device found")

			onvif := HasONVIF(host)

			fmt.Println("ONVIF:", onvif)

			urls := BuildRTSP(
				host,
				"admin",
				"admin",
			)

			for _, url := range urls {

				fmt.Println("TESTING RTSP:", url)

				valid := ValidateRTSP(url)

				fmt.Println("VALID:", valid)

				if valid {

					device := Device{
						IP:           host,
						Manufacturer: Fingerprint(host),
						RTSP:         url,
						ONVIF:        onvif,
						Valid:        true,
					}

					devices = append(devices, device)

					logger.Log.Info().
						Str("ip", host).
						Str("rtsp", url).
						Msg("camera discovered")

					fmt.Println("CAMERA DISCOVERED:", url)

					break
				}
			}
		}
	}

	fmt.Println("DISCOVERY FINISHED")

	return devices
}