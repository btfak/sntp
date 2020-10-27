// package netapp
// Modifications copyright (C) 2020 Fabio A. Takeuchi

package netapp

import (
	"net"
	"strconv"
	"strings"

	"github.com/fakiot/sntp/netevent"
	"github.com/fakiot/sntp/sntp"
)

var handler *Handler

type Handler struct {
	netevent.UdpHandler
}

func GetHandler() *Handler {
	if handler == nil {
		handler = new(Handler)
	}
	return handler
}

// DatagramReceived
// every udp request trigger it.
func (p *Handler) DatagramReceived(data []byte, addr net.Addr) {
	res, err := sntp.Serve(data)
	if err == nil {
		ip, port := spliteAddr(addr.String())
		p.UdpWrite(string(res), ip, port)
	}
}

func spliteAddr(addr string) (string, int) {
	ip := strings.Split(addr, ":")[0]
	port := strings.Split(addr, ":")[1]
	p, _ := strconv.Atoi(port)
	return ip, p
}
