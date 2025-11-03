package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	// Alanlara 'p' üzerinden erişiriz (this.Name değil, p.Name)
	fmt.Printf("Merhaba, benim adım %s\n", p.Name)
}

func main() {
	// Bir 'Person' struct'ı oluşturmanın yolu:
	p1 := Person{
		Name: "Ayşe",
		Age:  30,
	}

	// Metodu çağırmak C#'taki ile aynı:
	p1.Greet() // Çıktı: Merhaba, benim adım Ayşe

	// Başka bir oluşturma yolu (sırayla)
	p2 := Person{"Ali", 25}
	p2.Greet() // Çıktı: Merhaba, benim adım Ali
}