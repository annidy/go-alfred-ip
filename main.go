package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func ip() string {
	ifaces, _ := net.Interfaces()
	var ip = net.IPv4(127, 0, 0, 1)
	// handle err
	for _, i := range ifaces {
		if strings.HasPrefix(i.Name, "en") {
			addrs, _ := i.Addrs()
			// handle err
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
			}
		}
	}
	return ip.String()
}

func main() {

	ip := ip()

	fmt.Fprint(os.Stdout, ip)

	file, err := os.Open("./info.plist")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newdata := make([]string, 0)
	scanner := bufio.NewScanner(file)
	nextline := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "<key>text</key>") {
			nextline = true
		} else if nextline {
			nextline = false
			re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
			line = re.ReplaceAllString(line, ip)
		}
		newdata = append(newdata, line)
	}
	data := strings.Join(newdata, "\n")
	ioutil.WriteFile("./info.plist", []byte(data), 0644)
	log.Println(data)
}
