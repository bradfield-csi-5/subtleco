package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	dec_to_hex(os.Args[1])
}

func dec_to_hex(s string) {
	hexMap := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "BAD BAD VERY BAD")
		return
	}
	big := i / 16
	little := i % 16
	fmt.Printf("0x%s%s\n", hexMap[big], hexMap[little])
}
