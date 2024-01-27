package util

import (
	"Geo-DNS/stores"
	"Geo-DNS/structs"
)

func GetClosestPop(ip string) structs.Pop {
	closest := -1.0
	selected := ""
	selectedName := ""

	for _, pop := range stores.AvailablePops {
		distance := CalculateDistance(pop.Ip, ip)
		if closest == -1.0 {
			closest = distance
			selected = pop.Ip
			selectedName = pop.Hostname
			continue
		}
		if closest > distance {
			selected = pop.Ip
			closest = distance
			selectedName = pop.Hostname
		}
	}

	return structs.Pop{
		Ip:       selected,
		Hostname: selectedName,
	}
}
