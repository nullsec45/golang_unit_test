package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fajar")

	if result != "Hello Fajar" {
		//	error
		panic(`Result is not "Hello Fajar" `)
	}
}

func TestHelloWorldEntong(t *testing.T) {
	result := HelloWorld("Entong")

	if result != "Hello Entong" {
		//	error
		panic(`Result is not "Hello Entong" `)
	}
}
