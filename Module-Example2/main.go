// Bu dosyanın çalıştırılabilir ana paket olduğunu belirtiyoruz
package main

import (
    "fmt"
    // Oluşturduğumuz 'hesap' paketini "import" ediyoruz.
    // 'projem' -> bizim go.mod'daki modül adımız
    // '/hesap'  -> modül içindeki alt klasör/paket adı
    "calculator/calc" 
)

// Programın başlangıç noktası
func main() {
    sayi1 := 10
    sayi2 := 20

    // 'hesap' paketindeki 'Topla' fonksiyonumuzu çağırıyoruz
    // (PaketAdı.FonksiyonAdı şeklinde)
    sonuc := calc.Topla(sayi1, sayi2)

    fmt.Printf("%d ve %d toplamı: %d\n", sayi1, sayi2, sonuc)

    // Bu satır hata verir, çünkü 'cikar' fonksiyonu küçük harfle başlar
    // ve dışarıya açık (exported) değildir:
    // sonucCikar := hesap.cikar(20, 5) // Derleme Hatası!
}