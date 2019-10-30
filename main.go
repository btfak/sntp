//package main
//about: btfak.com
//create: 2013-9-25
//update: 2016-08-22

package main

import (
	"flag"
	"github.com/briantobin/sntp/netapp"
	"github.com/briantobin/sntp/netevent"
	"github.com/briantobin/sntp/sntp"
)

func main() {
	var port = flag.Int("p", 123, "NTP port")
	var offset_days = flag.Int("o", 0, "Time offset (in days)")
	flag.Parse()

	sntp.Offset_days = int64(*offset_days)

	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(*port, handler)
	netevent.Reactor.Run()
}
