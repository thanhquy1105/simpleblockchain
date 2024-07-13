package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type NeighborNode struct {
	HostPort    string `json:"host_port"`
	StakeBlance int64  `json:"stake_balance"`
	NodeAddress string `json:"node_address"`
}

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	fmt.Printf("%s\n", target)
	return true
}

// 192.168.0.10:5000
// 192.168.0.11:5000
// 192.168.0.12:5000
// 192.168.0.10:5001
// 192.168.0.10:5002
// 192.168.0.10:5003
// pattern for ipv4
var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeightbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) ([]NeighborNode, []NeighborNode) {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)

	if m == nil {
		return nil, nil
	}

	prefixHost := m[1]
	lastIp, _ := strconv.Atoi(m[len(m)-1])

	neighbors := make([]NeighborNode, 0)
	allNodes := make([]NeighborNode, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {

				endpoint := fmt.Sprintf("http://%s/info", guessTarget)
				resp, _ := http.Get(endpoint)
				if resp.StatusCode == 200 {
					var bcResp NeighborNode
					decoder := json.NewDecoder(resp.Body)
					_ = decoder.Decode(&bcResp)

					ne := NeighborNode{
						HostPort:    guessTarget,
						NodeAddress: bcResp.NodeAddress,
						StakeBlance: bcResp.StakeBlance,
					}
					neighbors = append(neighbors, ne)
				}

			}
		}
	}

	if IsFoundHost(myHost, myPort) {
		endpoint := fmt.Sprintf("http://%s/info", address)
		resp, _ := http.Get(endpoint)
		if resp.StatusCode == 200 {
			var bcResp NeighborNode
			decoder := json.NewDecoder(resp.Body)
			_ = decoder.Decode(&bcResp)

			ne := NeighborNode{
				HostPort:    address,
				NodeAddress: bcResp.NodeAddress,
				StakeBlance: bcResp.StakeBlance,
			}
			allNodes = append(neighbors, ne)
		}

	}

	return neighbors, allNodes
}

func GetHost() string {
	// hostname, err := os.Hostname()
	// if err != nil {
	// 	return "127.0.0.1"
	// }
	// fmt.Println(hostname)
	// address, err := net.LookupHost(hostname)
	// if err != nil {
	// 	return "127.0.0.1"
	// }
	// fmt.Println(address)
	// return address[0]

	return "127.0.0.1"

}
