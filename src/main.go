//package main
//author: Lubia Yang
//create: 2013-9-25
//about: www.lubia.me

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