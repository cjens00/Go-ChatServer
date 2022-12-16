package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/netip"
	"os"
)

const (
	locale = "en-US"
)

func ReadCodesFromJSON() {
	fCodesJSON, err := os.ReadFile(fmt.Sprintf("./codes/codes_%s.json", locale))
	if err != nil {
		log.Fatal(fmt.Sprintf("Fatal Error: Cannot read file [%s].", fmt.Sprintf("./codes/codes_%s.json", locale)))
	}

	var codesJSON map[string]interface{}
	json.Unmarshal(fCodesJSON, &codesJSON)
	fmt.Println(codesJSON)

	// json.NewDecoder
	// bytes := []byte()
	// var dat map[string]interface{}
}

func CreateUDPAddress(listenaddr string, port uint16) *net.UDPAddr {
	addr, err := netip.ParseAddr(listenaddr)
	if err != nil {
		log.Fatal("Error initializing server. Reason: invalid IP address.")
	}
	return net.UDPAddrFromAddrPort(netip.AddrPortFrom(addr, port))
}

type server struct {
	*net.UDPAddr
}

func New(addr string, port uint16) server {
	new := server{
		CreateUDPAddress(addr, port),
	}
	return new
}
