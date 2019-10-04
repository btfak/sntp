//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"github.com/briantobin/sntp/netapp"
	"github.com/briantobin/sntp/netevent"
)

func main() {
	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(6123, handler)
	netevent.Reactor.Run()
}
