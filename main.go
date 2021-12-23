package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func ip() []Item {
	ifaces, _ := net.Interfaces()
	ips := make([]Item, 0)
	// handle err
	for _, i := range ifaces {
		if strings.HasPrefix(i.Name, "en") {
			addrs, _ := i.Addrs()
			// handle err
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					ones, _ := v.Mask.Size()
					if ones < 64 {
						ips = append(ips, Item{v.IP.String(), v.IP.String(), i.Name})
					}
				}
			}
		}
	}
	return ips
}

type Item struct {
	Title    string `json:"title"`
	Arg      string `json:"arg"`
	Subtitle string `json:"subtitle"`
}

type Filter struct {
	Items []Item `json:"items"`
}

func main() {

	ip := ip()

	filter := Filter{ip}

	js, _ := json.Marshal(filter)
	jss := string(js)

	fmt.Fprint(os.Stdout, jss)
	log.Println(jss)
}
