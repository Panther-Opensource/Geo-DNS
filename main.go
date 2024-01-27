package main

import (
	"Geo-DNS/healthchecks"
	"Geo-DNS/server"
	"Geo-DNS/stores"
	"Geo-DNS/structs"
	"Geo-DNS/util"
	"fmt"
	"os"

	"github.com/ip2location/ip2location-go/v9"
	"github.com/miekg/dns"
)

func main() {
	go healthchecks.StartHealthChecks()
	stores.Config = util.LoadConfig()
	for popIp, popName := range stores.Config.Pops {
		stores.AvailablePops = append(stores.AvailablePops, structs.Pop{
			Ip:       popIp,
			Hostname: popName,
		})
	}

	p, _ := os.Getwd()
	loaded, e := ip2location.OpenDB(fmt.Sprintf("%v/%v", p, "IP2LOCATION-LITE-DB5.BIN"))
	if e != nil {
		fmt.Println("Could not load IP DB")
	}
	stores.Db = loaded

	// Define DNS server configuration
	addr := ":53" // Listen on port 53, the default DNS port
	udpServer := &dns.Server{Addr: addr, Net: "udp"}

	// Set up the DNS handler function
	dns.HandleFunc(".", server.HandleDNSRequest)

	// Start the DNS server
	go func() {
		if err := udpServer.ListenAndServe(); err != nil {
			fmt.Printf("Failed to set up DNS server: %v\n", err)
			os.Exit(1)
		}
	}()

	fmt.Printf("DNS server listening on %s...\n", addr)

	// Wait for an interrupt signal (e.g., Ctrl+C) to gracefully shut down the server
	select {}
}
