package main

import (
	"fmt"
	"math"
)

func printMessage(message string) {
	defer fmt.Println("This is deferred") // Ini akan dieksekusi terakhir
	fmt.Println(message)
	fmt.Println("This is not deferred")
}

func main() {
	printMessage("Hello, Golang!")
}
func (receiver TipeData) namaMetode(parameter1 tipeData1, parameter2 tipeData2) tipeDataKembalian {
	// Blok kode metode
	// ...
	return nilaiKembalian
}

type Rectangle struct {
	width, height float64
}

// Metode untuk menghitung luas Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// Metode untuk mengubah lebar Rectangle
func (r *Rectangle) setWidth(newWidth float64) {
	r.width = newWidth
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width, height float64
}

// Metode untuk menghitung luas Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

rect := Rectangle{width: 5, height: 10}
area := rect.area()
fmt.Println(area) // Output: 50.0

type Circle struct {
	radius float64
}

// Metode dengan receiver non-pointer untuk menghitung luas Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Circle struct {
	radius float64
}

// Metode dengan receiver pointer untuk mengubah nilai Circle
func (c *Circle) setRadius(newRadius float64) {
	c.radius = newRadius
}

var numbers [5]int // Deklarasi array berukuran 5 dengan tipe int
numbers = [5]int{1, 2, 3, 4, 5}

var numbers []int // Deklarasi slice kosong
numbers = []int{1, 2, 3, 4, 5}

var person map[string]int // Deklarasi map dengan kunci bertipe string dan nilai bertipe int
person = map[string]int{"John": 25, "Alice": 30, "Bob": 28}

type Person struct {
	Name string
	Age  int
}

var p1 Person
p1.Name = "John"
p1.Age = 30

var x int = 10
var ptr *int // Deklarasi pointer yang menunjuk ke int

ptr = &x // Assign alamat memori dari x ke pointer ptr
fmt.Println(*ptr) // Output: 10 (nilai dari variabel yang ditunjuk oleh pointer)