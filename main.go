package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func checkVPNip(ips []string) {
	for _, ip := range ips {
		resp, err := http.Get("https://proxycheck.io/v2/" + ip + "?vpn=1")
		if err != nil {
			log.Println(err.Error())
		}
		defer resp.Body.Close()
		pageBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		} else {
			if strings.Contains(string(pageBody), "\"type\": \"VPN\"") {
				fmt.Println(ip)
			}
		}
	}
}

func checkAppleIP(ips []string) {
	for _, ip := range ips {
		resp, err := http.Get("https://proxycheck.io/v2/" + ip + "?vpn=1&asn=1")
		if err != nil {
			log.Println(err.Error())
		}
		defer resp.Body.Close()
		pageBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		} else {
			if strings.Contains(string(pageBody), "\"provider\": \"Apple Inc.\"") {
				fmt.Println(ip)
			}
		}
	}
}

func main() {
	var IPs []string
	var durFlag = flag.Int("d", 60, "duration for capturing the packets.")
	var interfaceFlag = flag.String("i", "", "interface to capture the packets from.")

	flag.Parse()

	fmt.Println("")
	fmt.Println("APPLE VPN CONNECTION ESCAPE POC")
	fmt.Println("Developed by 0x9747")
	fmt.Println("")

	if *interfaceFlag != "" {
		os.Chdir("C:\\Program Files\\Wireshark")
		newDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("#################")
			fmt.Println("Changed to new directory:", newDir)
		}

		fmt.Println("Running tshark for", strconv.Itoa(*durFlag), "seconds", "on", *interfaceFlag)
		fmt.Println("#################")
		fmt.Println("")
		totalDuration := "duration:" + strconv.Itoa(*durFlag)
		out, err := exec.Command(`tshark.exe`, `-a`, totalDuration, `-i`, *interfaceFlag).Output()
		if err != nil {
			log.Fatal(err.Error())
		} else {
			var re = regexp.MustCompile(`(?m) .* (.+?\..+?\..+?\..+?) → (.+?\..+?\..+?\..+?) `)
			for _, i := range re.FindAllStringSubmatch(string(out), -1) {
				IPs = append(IPs, i[1])
				IPs = append(IPs, i[2])
			}

			if len(unique(IPs)) > 0 {
				fmt.Println("Unqiue IP Addresses Discovered:", unique(IPs))
				fmt.Println("")
				fmt.Println("Checking for VPN-related IP Addresses:")
				checkVPNip(unique(IPs))
				fmt.Println("")
				fmt.Println("Apple IP Addresses Leaked:")
				checkAppleIP(unique(IPs))
			} else {
				fmt.Println("Looks like no IP addresses were captured :( How about generating some traffic on the device and trying again?")
			}
		}
	} else {
		fmt.Println("Looks like you didn't provided any interface names! Try again.")
		fmt.Println("Example:  .\\ios-poc.exe -d 30 -i \"Local Area Connection* 2\"")
	}
	fmt.Println("")
}