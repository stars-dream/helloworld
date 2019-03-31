package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if hello() != "Hello world" {
		t.Error("Testing error!")
	}
}
