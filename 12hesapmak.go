package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("lüten seçim yapınız ")
	fmt.Println("1-Toplama \n" +
		"2-çıkarma \n" +
		"3-çarpma \n" +
		"4-bölme \n" +
		"5 -Mod alma ")

	fmt.Print("lütfen hangi işlemi yapmak istediğinizi giriniz (1,2,3,4,5) =   ")
	scanner.Scan()
	secim := scanner.Text()

	fmt.Print("1. sayiyi giriniz =  ")
	scanner.Scan()
	sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("2. sayiyi giriniz =  ")
	scanner.Scan()
	sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)

	if secim == "1" {
		fmt.Printf("toplam = %f", sayi1+sayi2)
	} else if secim == "2" {
		fmt.Printf("çıkarma = %f", sayi1+sayi2)
	} else if secim == "3" {
		fmt.Printf("çarpma = %f", sayi1*sayi2)
	} else if secim == "4" {
		fmt.Printf("bölme =  %f", sayi1/sayi2)
	} else if secim == "5" {
		modalma := int(sayi1) % int(sayi2)
		fmt.Printf("mod alma = %d", modalma)
	} else {
		fmt.Print("lütfen geçerli bir secim yapınız .")
	}
}
