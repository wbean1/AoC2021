package day16

import (
	"fmt"
	"testing"
)

var testCases = map[string]int{
	"8A004A801A8002F478":             16,
	"620080001611562C8802118E34":     12,
	"C0015000016115A2E0802F182340":   23,
	"A0016C880162017C3686B18A3D4780": 31,
}

func TestSumVersionPackets(t *testing.T) {
	for input, expected := range testCases {
		fmt.Printf("parsing string: %s\n", input)
		got := sumVersionPackets(Input(input))
		if int(got) != expected {
			t.Errorf("wrong Sum of Version Packets.  got: %d, expected: %d", got, expected)
		} else {
			fmt.Printf("yay! got: %d, expected: %d\n", got, expected)
		}
	}
}
