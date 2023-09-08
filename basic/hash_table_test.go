package basic

import (
	"fmt"
	"testing"
)

type Person struct {
	name  String
	age   int
	phone string
}

func TestHashTableDelete(t *testing.T) {
	ht := NewHashTable[String, *Person](3)
	p1 := Person{"Alice", 20, "123"}
	p2 := Person{"Bob", 22, "222"}
	p3 := Person{"George", 18, "333"}
	ht.Put(p1.name, &p1)
	ht.Put(p2.name, &p2)
	ht.Put(p3.name, &p3)

	//ht.Delete(p1.name)
	//ht.Delete(p2.name)

	fmt.Println(ht.Keys())

}

func TestHashTablePut(t *testing.T) {
	ht := NewHashTable[String, *Person](3)
	p1 := Person{"Alice", 20, "123"}
	p2 := Person{"Bob", 22, "222"}
	p3 := Person{"George", 18, "333"}
	ht.Put(p1.name, &p1)
	ht.Put(p2.name, &p2)
	ht.Put(p3.name, &p3)

	fmt.Println(p1.name, ":", ht.Get(p1.name).phone)
	fmt.Println(p2.name, ":", ht.Get(p1.name).phone)
	fmt.Println(p3.name, ":", ht.Get(p1.name).phone)

	if ht.Get("nobody") != nil {
		fmt.Println("nobody", ":", ht.Get("nobody").phone)
	}
}

func nearestPowerOfTwo(c uint) uint {
	x := uint(1)
	for i := 1; x+1 < c; i = i * 2 {
		x = x | x<<i
	}
	return x + 1
}

func TestNearestPowerOfTwo(t *testing.T) {
	fmt.Println(nearestPowerOfTwo(uint(3)))
	fmt.Println(nearestPowerOfTwo(uint(4)))
	fmt.Println(nearestPowerOfTwo(uint(7)))
	fmt.Println(nearestPowerOfTwo(uint(31)))
	fmt.Println(nearestPowerOfTwo(uint(50)))
}
