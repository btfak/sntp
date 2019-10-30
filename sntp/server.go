//package sntp
//author: btfak.com
//create: 2013-9-25

package sntp

import (
	"errors"
	"time"
)

const (
	LI_NO_WARNING      = 0
	LI_ALARM_CONDITION = 3
	VN_FIRST           = 1
	VN_LAST            = 4
	MODE_CLIENT        = 3
	FROM_1900_TO_1970  = 2208988800
)

// Serve
// check the request format.
// get time from local and respond.
func Serve(req []byte) ([]byte, error) {
	if validFormat(req) {
		res := generate(req)
		return res, nil
	}
	return []byte{}, errors.New("invalid format.")
}

// validFormat
// check the first byte,include:
// 	LN:must be 0 or 3
// 	VN:must be 1,2,3 or 4
//	Mode:must be 3
func validFormat(req []byte) bool {
	var l = req[0] >> 6
	var v = (req[0] << 2) >> 5
	var m = (req[0] << 5) >> 5
	if (l == LI_NO_WARNING) || (l == LI_ALARM_CONDITION) {
		if (v >= VN_FIRST) && (v <= VN_LAST) {
			if m == MODE_CLIENT {
				return true
			}
		}
	}
	return false
}

// unix time: the number of seconds elapsed since January 1, 1970 UTC
// npt time: the number of seconds elapsed since January 1, 1900 UTC
// 1900~1970
func unix2ntp(u int64) int64 {
	return u + FROM_1900_TO_1970
}

func ntp2unix(n int64) int64 {
	return n - FROM_1900_TO_1970
}

// int2bytes
// format int number to four bytes.
// big endian.
func int2bytes(i int64) []byte {
	var b = make([]byte, 4)
	h1 := i >> 24
	h2 := (i >> 16) - (h1 << 8)
	h3 := (i >> 8) - (h1 << 16) - (h2 << 8)
	h4 := byte(i)
	b[0] = byte(h1)
	b[1] = byte(h2)
	b[2] = byte(h3)
	b[3] = byte(h4)
	return b
}

var Offset_days int64 = 0

// generate
/*
	  Field Name              Request    Reply
      ----------------------------------------------------------
      LI                      0 or 3     0
      VN                      1-4        copied from request
      Mode                    3          4
      Stratum                 ignore     1
      Poll                    ignore     copied from request
      Precision               ignore     -log2 server significant bits
      Root Delay              ignore     0
      Root Dispersion         ignore     0
      Reference Identifier    ignore     source ident
      Reference Timestamp     ignore     time of last radio update
	  Originate Timestamp     ignore     copied from transmit timestamp
      Receive Timestamp       ignore     time of day
      Transmit Timestamp      (see text) time of day
*/
func generate(req []byte) []byte {
	var second = unix2ntp(time.Now().Unix() + (60 * 60 * 24 * Offset_days))
	var fraction = unix2ntp(int64(time.Now().Nanosecond()))
	var res = make([]byte, 48)
	var vn = req[0] & 0x38
	res[0] = vn + 4
	res[1] = 1
	res[2] = req[2]
	res[3] = 0xEC
	res[12] = 0x4E
	res[13] = 0x49
	res[14] = 0x43
	res[15] = 0x54
	copy(res[16:20], int2bytes(second)[0:])
	copy(res[24:32], req[40:48])
	copy(res[32:36], int2bytes(second)[0:])
	copy(res[36:40], int2bytes(fraction)[0:])
	copy(res[40:48], res[32:40])
	return res
}
