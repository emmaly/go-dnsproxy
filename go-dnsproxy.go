
package main

import (
	"fmt"
	"github.com/miekg/dns"
	"os"
)

func main() {
	dns.HandleFunc(".", proxyServe)
	dns.HandleFunc("companyname.local.", companyServe)

	failure := make(chan error, 1)

	go func(failure chan error) {
		failure <- dns.ListenAndServe("127.0.0.1:53", "tcp", nil)
	}(failure)

	go func(failure chan error) {
		failure <- dns.ListenAndServe("127.0.0.1:53", "udp", nil)
	}(failure)

	fmt.Println(<-failure)
	os.Exit(1)
}

func proxyServe(w dns.ResponseWriter, req *dns.Msg) {
	if req.MsgHdr.Response == true { // supposed responses sent to us are bogus
		return
	}

	c := new(dns.Client)
	m, err := c.Exchange(req, "8.8.8.8:53")

	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(m)
	}
}

func companyServe(w dns.ResponseWriter, req *dns.Msg) {
	if req.MsgHdr.Response == true { // supposed responses sent to us are bogus
		return
	}

	c := new(dns.Client)
	m, err := c.Exchange(req, "10.20.30.40:53")

	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(m)
	}
}
