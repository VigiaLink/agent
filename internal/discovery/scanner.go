package discovery

import (
    "net"
)

func LocalSubnets() ([]string, error) {
   interfaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }

    subnets := []string{}

    for _, iface := range interfaces {
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            ipnet, ok := addr.(*net.IPNet)
            if !ok {
                continue
            }

            ipv4 := ipnet.IP.To4()
            if ipv4 == nil {
                continue
            }

            subnet := ipv4.Mask(ipnet.Mask)

            cidr := subnet.String() + "/24"
            if ipv4[0] == 192 || ipv4[0] == 10 {
            subnets = append(subnets, cidr)
        }
        }
    }

    return subnets, nil
}