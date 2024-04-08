package utils

const (
	DELETE   = 0x00
	PUT      = 0x01
	MAX_LVL  = 10
	LAST_KEY = "zzzzzz"
	LOG      = "WAL.log"
)

func BEQ(a, b []byte) bool {
	return string(a) == string(b)
}

// Returns true if the first value belongs before the second
func BLT(a, b []byte) bool {
	return string(a) < string(b)
}
