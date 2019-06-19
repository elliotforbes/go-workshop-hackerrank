package main

import (
	"testing"
	"fmt"
)

func TestSockMerchant(t *testing.T) {
	var n int32
	n = 9
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	result := SockMerchant(n, ar)
	if result == 3 {
		fmt.Println("Test Passed")
	} else {
		t.Fail()
	}
} 

func TestSockMerchantExtensive(t *testing.T) {
	var tests = []struct {
		n    int32
		ar	[]int32
        expected int32
    }{
        {9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 3},
        {2, []int32{10, 20, }, 0},
        {6, []int32{10, 10, 10, 10, 10, 10,}, 3},
    }

    for _, test := range tests {
        if output := SockMerchant(test.n, test.ar); output != test.expected {
            t.Error("Test Failed: ", test.n, test.ar, test.expected, output)
        }
    }
}