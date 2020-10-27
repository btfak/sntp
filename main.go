// package main
// Modifications copyright (C) 2020 Fabio A. Takeuchi

package main

import (
	"time"

	"github.com/fakiot/sntp/netapp"
	"github.com/fakiot/sntp/netevent"
)

type test struct{}

// Sync method to return example datetime
func (v *test) Sync() time.Time {
	return time.Date(2005, 1, 1, 1, 1, 1, 0, time.Local)
}

func main() {
	var handler = netapp.GetHandler()
	// Uncomment this to see the example. Result datetime is 2005-01-01 01:01:01
	// sntp.SetSyncFunc(&test{})
	netevent.Reactor.ListenUdp(123, handler)
	netevent.Reactor.Run()
}
