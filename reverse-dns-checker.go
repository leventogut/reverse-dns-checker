package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	domain = flag.String("domain", "", "Domain to resolve")
)

func main() {
	flag.Parse()
	if *domain == "" {
		flag.Usage()
		os.Exit(1)
	}
	if strings.HasSuffix(*domain, ".") {

	} else {
		*domain = *domain + "."
	}
	var match bool
	fmt.Println("Domain: ", *domain)
	aRecords, err := getARecords(*domain)

	if err != nil {
		if strings.Contains(err.Error(), "No address associated with hostname") {
			fmt.Println("No A Record found")
		} else {
			panic(fmt.Errorf("failed to get A records, error: %v", err))
		}

	}
	for _, aRecord := range aRecords {
		fmt.Println("A Record: ", aRecord)
		ptrRecords, err := getPTRRecords(aRecord)
		if err != nil {
			if strings.Contains(err.Error(), "Name or service not known") {
				fmt.Println("No PTR found")
			} else {
				panic(fmt.Errorf("failed to get PTR records, error: %v", err))
			}

		}

		for _, ptrRecord := range ptrRecords {
			fmt.Printf("PTR for %s: %s\n", aRecord, ptrRecord)
			if *domain == ptrRecord {
				match = true
			}
		}

	}
	if match == true {
		fmt.Println("\nMatch exists")
	} else {
		fmt.Println("\nMatch does NOT exists")
	}

}

func getARecords(domain string) (records []string, err error) {
	return net.LookupHost(domain)
}

func getPTRRecords(ip string) (records []string, err error) {
	return net.LookupAddr(ip)
}
