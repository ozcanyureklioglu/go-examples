// Bu dosyanın 'hesap' paketine ait olduğunu belirtiyoruz
package calc

// Topla fonksiyonu. Adı büyük harfle (T) başladığı için
// başka paketlerden (mesela 'main' paketinden) erişilebilir.
func Topla(a int, b int) int {
	return a + b
}

// bu fonksiyon küçük harfle başladığı için 'private' gibidir,
// sadece 'hesap' paketi içinden erişilebilir.
func cikar(a int, b int) int {
	return a - b
}
