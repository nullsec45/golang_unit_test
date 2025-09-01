package helper

import "testing"


func HelloWorld(name string) string {
	return "Hello " + name
}

func BenchmarkTable(b *testing.B){
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "Rama",
			request: "Rama",
		},
		{	
			name: "Fajar",
			request: "Fajar",
		},
	}

	for _, benchmark := range benchmarks{
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++{
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchMarkSub(b *testing.B){
	b.Run("Rama", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Rama")
		}
	})
	b.Run("Fajar", func(b *testing.B){
		for i := 0; i < b.N; i++{
			HelloWorld("Fajar")
		}
	}) 
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rama")
	}
}

func BenchmarkHelloWorldFajar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fajar")
	}
}



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
