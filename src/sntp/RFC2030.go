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

func Serve(req []byte) ([]byte, error) {
	var r_f_int = unix2ntp(time.Now().Unix())
	var r_f_dec = unix2ntp(int64(time.Now().Nanosecond()))
	var r_c_int, r_c_dec = r_f_int, r_f_dec
	if validFormat(req) {
		var res = make([]byte, 48)
		var vn = req[0] & 0x38
		res[0] = vn + 4
		res[1] = 1
		res[2] = req[2]
		res[3] = 0xEC
		res[12] = 49
		res[13] = 0x4E
		res[14] = 49
		res[15] = 52
		for k, v := range int2bytes(r_f_int) {
			res[16+k] = v
		}
		copy(res[24:32], req[40:48])
		for k, v := range int2bytes(r_c_int) {
			res[32+k] = v
		}
		for k, v := range int2bytes(r_c_dec) {
			res[36+k] = v
		}
		var tt_int = unix2ntp(time.Now().Unix())
		var tt_dec = unix2ntp(int64(time.Now().Nanosecond()))
		for k, v := range int2bytes(tt_int) {
			res[40+k] = v
		}
		for k, v := range int2bytes(tt_dec) {
			res[44+k] = v
		}
		return res, nil
	}
	return []byte{}, errors.New("invalid format.")
}

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

func unix2ntp(u int64) int64 {
	return u + FROM_1900_TO_1970
}

func ntp2unix(n int64) int64 {
	return n - FROM_1900_TO_1970
}

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
