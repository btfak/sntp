package main

import (
	"netevent" 
	"netapp"
)
 
func main(){
	var handler = netapp.GetHandler()
	netevent.Reactor.ListenUdp(123, handler)
	netevent.Reactor.Run()
}