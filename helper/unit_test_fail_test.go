package helper

import (
	"fmt"
	"testing"
)

func TestFail(t *testing.T) {
	result := HelloWorld("Salah")

	if result != "Hello Fajar" {
		t.Fail()
	}
	fmt.Println("Test Fail")
}

func TestFailNow(t *testing.T) {
	result := HelloWorld("Fajar")
	if result != "Hello  Fajar" {
		t.FailNow()
	}
	fmt.Println("Test Fail Now")
}
