package main

import "fmt"

func main() {
	var celsius float64

	fmt.Println("=== PROGRAM KONVERSI SUHU ===")
	fmt.Print("Masukkan suhu Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273.15

	fmt.Println("---------------------------")
	fmt.Println("Celsius:", celsius)
	fmt.Println("Fahrenheit:", fahrenheit)
	fmt.Println("Kelvin:", kelvin)
}
