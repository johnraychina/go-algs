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

func TestHashTableGetMissing(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	if ht.Get("nonexistent") != nil {
		t.Fatal("expected nil for missing key")
	}
}

func TestHashTableSize(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	if ht.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ht.Size())
	}
	ht.Put("a", &Person{"a", 1, "111"})
	ht.Put("b", &Person{"b", 2, "222"})
	if ht.Size() != 2 {
		t.Fatalf("expected size 2, got %d", ht.Size())
	}
}

func TestHashTableOverwrite(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	ht.Put("Alice", &Person{"Alice", 20, "111"})
	ht.Put("Alice", &Person{"Alice", 25, "999"})

	if ht.Size() != 1 {
		t.Fatalf("expected size 1 after overwrite, got %d", ht.Size())
	}
	p := ht.Get("Alice")
	if p.age != 25 || p.phone != "999" {
		t.Fatalf("expected updated value, got age=%d phone=%s", p.age, p.phone)
	}
}

func TestHashTableDeleteAndGet(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	ht.Put("Alice", &Person{"Alice", 20, "111"})
	ht.Put("Bob", &Person{"Bob", 22, "222"})
	ht.Put("Charlie", &Person{"Charlie", 18, "333"})

	ht.Delete("Bob")
	if ht.Size() != 2 {
		t.Fatalf("expected size 2 after delete, got %d", ht.Size())
	}
	if ht.Get("Bob") != nil {
		t.Fatal("expected nil after deleting Bob")
	}
	if ht.Get("Alice") == nil || ht.Get("Charlie") == nil {
		t.Fatal("other keys should still exist")
	}
}

func TestHashTableDeleteMissing(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	ht.Put("Alice", &Person{"Alice", 20, "111"})
	ht.Delete("nobody")
	if ht.Size() != 1 {
		t.Fatalf("size should remain 1, got %d", ht.Size())
	}
}

func TestHashTableDeleteAll(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	keys := []String{"a", "b", "c", "d", "e"}
	for _, k := range keys {
		ht.Put(k, &Person{name: k})
	}
	for _, k := range keys {
		ht.Delete(k)
	}
	if ht.Size() != 0 {
		t.Fatalf("expected size 0 after deleting all, got %d", ht.Size())
	}
	for _, k := range keys {
		if ht.Get(k) != nil {
			t.Fatalf("expected nil for key %s after deletion", k)
		}
	}
}

func TestHashTableEnsureCapacity(t *testing.T) {
	ht := NewHashTable[String, *Person](2)
	for i := range 100 {
		k := String(fmt.Sprintf("key_%d", i))
		ht.Put(k, &Person{name: k, age: i})
	}
	if ht.Size() != 100 {
		t.Fatalf("expected size 100, got %d", ht.Size())
	}
	for i := range 100 {
		k := String(fmt.Sprintf("key_%d", i))
		p := ht.Get(k)
		if p == nil {
			t.Fatalf("missing key %s after resize", k)
		}
		if p.age != i {
			t.Fatalf("key %s: expected age %d, got %d", k, i, p.age)
		}
	}
}

func TestHashTableKeys(t *testing.T) {
	ht := NewHashTable[String, *Person](4)
	ht.Put("x", &Person{name: "x"})
	ht.Put("y", &Person{name: "y"})
	ht.Put("z", &Person{name: "z"})

	keys := ht.Keys()
	if len(keys) != 3 {
		t.Fatalf("expected 3 keys, got %d", len(keys))
	}
	seen := map[String]bool{}
	for _, k := range keys {
		seen[k] = true
	}
	for _, want := range []String{"x", "y", "z"} {
		if !seen[want] {
			t.Fatalf("missing key %s in Keys()", want)
		}
	}
}
