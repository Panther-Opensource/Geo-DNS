package server

import (
	"Geo-DNS/util"
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	ip := strings.Split(w.RemoteAddr().String(), ":")[0]

	fmt.Printf("Received DNS query from %v\n", ip)

	pop := util.GetClosestPop(ip)

	// Loop through each question in the DNS query
	for _, q := range r.Question {
		// For simplicity, respond with a hardcoded IP address for all queries
		answer := dns.A{
			Hdr: dns.RR_Header{
				Name:   q.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    60, // Time to live in seconds
			},
			A: net.ParseIP(pop.Ip),
		}
		fmt.Printf("Responded with: %v\n", pop.Hostname)
		m.Answer = append(m.Answer, &answer)
	}

	// Send the DNS response
	if err := w.WriteMsg(m); err != nil {
		fmt.Printf("Failed to send DNS response: %v\n", err)
	}
}
