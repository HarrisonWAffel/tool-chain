package net

import (
	"fmt"
	"net"
)

func GetAllHostsOnNetwork(subnet string) map[string]string {
	out := make(map[string]string)
	for i := 0; i < 255; i++ {
		n := fmt.Sprintf(subnet+"."+"%d", i)
		ptr, _ := net.LookupAddr(n)
		for _, ptrvalue := range ptr {
			out[n] = ptrvalue
		}
	}
	return out
}
