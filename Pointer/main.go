package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) Display() {
	fmt.Printf("--- Ürün Bilgisi (Kopya) ---\n")
	fmt.Printf("İsim: %s\n", p.Name)
	fmt.Printf("Fiyat: %.2f TL\n", p.Price)
	fmt.Printf("Stok: %d\n", p.Stock)
	fmt.Println("----------------------------")
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.Price = newPrice
	fmt.Printf(">>> Fiyat güncellendi: %s\n", p.Name)
}

func main() {
	p1 := Product{
		Name:  "Phone",
		Price: 34.5,
		Stock: 5,
	}
	fmt.Println("### Orijinal Ürün Oluşturuldu ###")
	p1.Display()
	p1.UpdatePrice(27500.00)

	fmt.Println("\n### Fiyat Güncellemesi Sonrası ###")
	p1.Display()
	fmt.Printf("\n'main' içindeki p1'in son hali: %+v\n", p1)
}
