//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"flag"
	"github.com/briantobin/sntp/netapp"
	"github.com/briantobin/sntp/netevent"
)

func main() {
	var port = flag.Int("p", 123, "NTP port")
	flag.Parse()

	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(*port, handler)
	netevent.Reactor.Run()
}
