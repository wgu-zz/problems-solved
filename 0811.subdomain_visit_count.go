package main

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, cpdomain := range cpdomains {
		f := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(f[0])
		domains := strings.Split(f[1], ".")
		subdomain := ""
		for i := len(domains) - 1; i >= 0; i-- {
			if subdomain == "" {
				subdomain = domains[i]
			} else {
				subdomain = domains[i] + "." + subdomain
			}
			m[subdomain] += count
		}
	}
	res := []string{}
	for key, value := range m {
		res = append(res, strconv.Itoa(value)+" "+key)
	}
	return res
}
