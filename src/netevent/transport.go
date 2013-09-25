package netevent

import (
	"net"
	"fmt"
	"strconv"
)

type Transport interface {
	Write(data string, addr string, port int)
}

type udpTransport struct {
	conn *net.UDPConn
}

func (p *udpTransport) setConn(conn *net.UDPConn){
	p.conn = conn
}

type tcpTransport struct {
	listener *net.TCPListener
}

func (p *tcpTransport) setListener(listener *net.TCPListener){
	p.listener = listener
}

func (p *udpTransport) Write(data string, addr string, port int){
	laddr, err := net.ResolveUDPAddr("udp",  addr+":"+strconv.Itoa(port))	
	if err != nil {
		fmt.Println("resolve addr err")
		return
	}
	_,er := p.conn.WriteTo([]byte(data),laddr)
	if er != nil {
		fmt.Println(er)
		return
	}
}

func (p *tcpTransport) Write(data string, addr string, port int){
	_, err := net.ResolveUDPAddr("udp",  addr+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("resolve addr err")
		return
	}
}

type unixTransport struct {
	conn *net.UnixConn
}

func (p *unixTransport) setConn(conn *net.UnixConn){
	p.conn = conn
}

func (p *unixTransport) Write(data string, addr string, port int){
	laddr, err := net.ResolveUDPAddr("udp",  addr+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("resolve addr err")
		return
	}
	_,er := p.conn.WriteTo([]byte(data),laddr)
	if er != nil {
		fmt.Println(er)
		return
	}
}
