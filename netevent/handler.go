//package netevent
//about: btfak.com
//create: 2013-7-20

package netevent

import ()

type UdpHandler struct {
	udptransport Transport
}

type TcpHandler struct {
	tcptransport Transport
}

func (p *UdpHandler) SetUdpTransport(transport Transport) {
	p.udptransport = transport
}

func (p *TcpHandler) SetTcpTransport(transport Transport) {
	p.tcptransport = transport
}

func (p *UdpHandler) UdpWrite(data string, addr string, port int) {
	p.udptransport.Write(data, addr, port)
}

func (p *TcpHandler) TcpWrite(data string, addr string, port int) {
	//p.transport.Write(data,addr,port)
}
