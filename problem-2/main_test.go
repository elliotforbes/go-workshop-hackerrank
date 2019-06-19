package main

import (
	"testing"
)

func TestGetMoneySpent(t *testing.T) {
	keyboards := []int32{3, 1,}
	drives := []int32{5, 2, 8, }
	result := GetMoneySpent(keyboards, drives, 10) 
	if result != 9 {
		t.Error("Unexpected Result")
	}
}