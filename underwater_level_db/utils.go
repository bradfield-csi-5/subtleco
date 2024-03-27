package main

func BEQ(a, b []byte) bool {
	return string(a) == string(b)
}

func BLT(a, b []byte) bool {
	return string(a) < string(b)
}
