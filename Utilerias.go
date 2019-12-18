package negocio

import (
	"time"
)

const (
	workerIdBits   = uint64(10)
	sequenceBits   = uint64(12)
	workerIdShift  = sequenceBits
	timestampShift = sequenceBits + workerIdBits
	sequenceMask   = int64(-1) ^ (int64(-1) << sequenceBits)

	// Tue, 21 Mar 2006 20:50:14.000 GMT
	twepoch = int64(1288834974657)
)

var sequence int64
var lastTimestamp int64

type GUID int64

func GeneraGuid(workerId int64) (GUID) {
	ts := time.Now().UnixNano() / 1e6
	if ts < lastTimestamp {
		return 0
	}
	if lastTimestamp == ts {
		sequence = (sequence + 1) & sequenceMask
		if sequence == 0 {
			return 0
		}
	} else {
		sequence = 0
	}
	lastTimestamp = ts

	id := ((ts - twepoch) << timestampShift) |
		(workerId << workerIdShift) |
		sequence
	return GUID(id)
}

