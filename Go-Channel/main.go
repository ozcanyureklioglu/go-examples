package main

import (
	"fmt"
	"time"
)

func fetchApi1(sonucKanali chan string) {
	time.Sleep(1 * time.Second)
	sonucKanali <- "API 1'den Sonuç Geldi"
}

func fetchApi2(sonucKanali chan string) {
	time.Sleep(2 * time.Second)
	sonucKanali <- "API 2'den Sonuç Geldi"
}

func fetchApi3(sonucKanali chan string) {
	time.Sleep(1 * time.Second)
	sonucKanali <- "API 3'ten Sonuç Geldi"
}

func main() {
	// Programın ne kadar süreceğini ölçmek için
	baslangic := time.Now()

	channel := make(chan string)

	fmt.Println("Channel has been created...")

	go fetchApi1(channel)
	go fetchApi2(channel)
	go fetchApi3(channel)

	sonuc1 := <-channel
	sonuc2 := <-channel
	sonuc3 := <-channel

	fmt.Println("--- Toplanan Sonuçlar ---")
	fmt.Println(sonuc1)
	fmt.Println(sonuc2)
	fmt.Println(sonuc3)

	fmt.Printf("Toplam süre: %v\n", time.Since(baslangic))
}
