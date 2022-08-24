package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func ip_records(domain string) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println()
		fmt.Printf("[IP] Error : " + err.Error())
	}
	fmt.Println()
	for _, ip := range ips {
		fmt.Printf("[IP] %s\n", ip)
	}
}

func mx_records(domain string) {
	mxs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println()
		fmt.Printf("[MX] Error : " + err.Error())
	}
	fmt.Println()
	for _, mx := range mxs {
		fmt.Printf("[MX] %v\t%s\n", mx.Pref, mx.Host)
	}
}

func ns_records(domain string) {
	nss, err := net.LookupNS(domain)
	if err != nil {
		fmt.Println()
		fmt.Printf("[NS] Error : " + err.Error())
	}
	fmt.Println()
	for _, ns := range nss {
		fmt.Printf("[NS] %s\n", ns.Host)
	}
}

func txt_records(url string) {
	txts, err := net.LookupTXT(url)
	if err != nil {
		fmt.Println("[TXT] Error : " + err.Error())
	}
	fmt.Println()
	for _, txt := range txts {
		if strings.Contains(txt, "v=DMARC1") {
			fmt.Printf("[DMARC] %s\n", txt)
			break
		}
		if strings.Contains(txt, "v=spf") {
			fmt.Printf("[SPF] %s\n", txt)
			break
		}
		fmt.Printf("[TXT] %s\n", txt)
	}
}

func cname_records(domain string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Println()
		fmt.Println("[CNAME] Error : " + err.Error())
	} else {
		fmt.Printf("\n[CNAME] %s\n", cname)
	}
}

func main() {

	if len(os.Args) < 2 || len(os.Args) != 2 {
		fmt.Println("Incorrect usage of code-base to query DNS data!")
		os.Exit(0)
	}

	domain := os.Args[1]

	fmt.Printf("Getting DNS Records for %s ...\n", domain)

	ip_records(domain)
	mx_records(domain)
	ns_records(domain)
	cname_records(domain)
	txt_records(domain)
	txt_records("_dmarc." + domain)
}
