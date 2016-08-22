//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"github.com/btfak/sntp/netapp"
	"github.com/btfak/sntp/netevent"
)

func main() {
	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(123, handler)
	netevent.Reactor.Run()
}
