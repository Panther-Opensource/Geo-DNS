package healthchecks

import (
	"Geo-DNS/stores"
	"Geo-DNS/structs"
	"fmt"
	"net/http"
	"time"
)

func StartHealthChecks() {
	tk := time.NewTicker(time.Second * 10)

	for range tk.C {
		nodes := stores.Config.Pops
		availableNodes := []structs.Pop{}
		for node, name := range nodes {
			start := time.Now().UnixMilli()
			res, e := http.Get(fmt.Sprintf("http://%v/cdn-cgi/health", node))
			if e != nil {
				continue
			}
			ping := int(time.Now().UnixMilli() - start)
			if res.StatusCode == 200 && ping < 1000 {
				availableNodes = append(availableNodes, structs.Pop{
					Ip:       node,
					Hostname: name,
				})
			}
		}
		fmt.Printf("Healthchecks done, %v/%v nodes are available\n", len(nodes), len(availableNodes))
		stores.AvailablePops = availableNodes
	}
}
