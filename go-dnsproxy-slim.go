package main

import (
	"github.com/miekg/dns"
	"os"
)

func main() {
	dns.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) { serve(w, req, "8.8.8.8:53") })
	dns.HandleFunc("companyname.local.", func(w dns.ResponseWriter, req *dns.Msg) { serve(w, req, "10.20.30.40:53") })
	dns.HandleFunc("example.com.", func(w dns.ResponseWriter, req *dns.Msg) { serve(w, req, "10.10.10.10:53") })
	go func() {
		dns.ListenAndServe("127.0.0.1:53", "tcp", nil)
		os.Exit(1)
	}()
	dns.ListenAndServe("127.0.0.1:53", "udp", nil)
	os.Exit(1)
}

func serve(w dns.ResponseWriter, req *dns.Msg, host string) {
	c := new(dns.Client)
	m, err := c.Exchange(req, host)
	if err == nil {
		w.Write(m)
	}
}