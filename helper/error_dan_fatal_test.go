package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Entong")
	if result != "Hello Fajar" {
		t.Error("Harusnya Hello Fajar")
	}
	fmt.Println("Ini akan dieksekusi")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Entong")
	if result != "Hello Fajar" {
		t.Fatal("Harusnya Hello Fajar")
	}
	fmt.Println("Ini tidak akan dieksekusi")
}
