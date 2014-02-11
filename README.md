sntp
====

### A implementation of NTP server with Golang
##### What this?
- Base on [RFC 2030](http://tools.ietf.org/html/rfc2030)
- Multiple equipments sync time from local
- Design for multiple equipments which can't connect to internet and need synchronization time
- Compatible with [ntpdate](http://www.eecis.udel.edu/~mills/ntp/html/ntpdate.html) service on the linux
- NTP client is fork from [beevik](https://github.com/beevik/ntp/)，a better client

####Usage manual
#####1. install Golang

Please reference  [Go install](https://github.com/astaxie/build-web-application-with-golang/blob/master/ebook/01.1.md) chapter of open-source Golang book "build-web-application-with-golang".

#####2. install Sntp

    $ go get github.com/lubia/sntp

##### More? 
[My Blog](http://www.lubia.me)

##### License
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
