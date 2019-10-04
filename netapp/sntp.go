//package netapp
//author: btfak.com
//create: 2013-9-24
//update: 2016-08-22

package netapp

import (
	"github.com/btfak/sntp/netevent"
	"github.com/btfak/sntp/sntp"
	"net"
	"strconv"
	"strings"
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
	delimiterIdx := strings.LastIndex(addr, ":")
	ip := addr[0:delimiterIdx]
	port := addr[delimiterIdx+1:]

	p, _ := strconv.Atoi(port)
	return ip, p
}
