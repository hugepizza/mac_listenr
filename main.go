package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	ms := getMacAddrs()
	http.DefaultServeMux.HandleFunc("/mac", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "*")
		resp.Write([]byte(strings.Join(ms, ",")))
	})
	fmt.Println("运行中 访问植提桥后台时请保持运行 ...")
	http.ListenAndServe(":9999", http.DefaultServeMux)
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}
