package basic

import (
	"sort"
	"testing"
)

func TestDWayHeap_InsertAndDelMin(t *testing.T) {
	for _, d := range []int{2, 3, 4, 5} {
		h := NewDWayHeap[int](d)
		h.Insert(5)
		h.Insert(3)
		h.Insert(7)
		h.Insert(1)
		h.Insert(4)
		h.Insert(2)
		h.Insert(6)

		expected := []int{1, 2, 3, 4, 5, 6, 7}
		for i, want := range expected {
			if h.Size() != len(expected)-i {
				t.Fatalf("d=%d: Size()=%d, want %d", d, h.Size(), len(expected)-i)
			}
			got := h.DelMin()
			if got != want {
				t.Fatalf("d=%d: DelMin() #%d = %d, want %d", d, i, got, want)
			}
		}
		if !h.IsEmpty() {
			t.Fatalf("d=%d: expected empty after draining", d)
		}
	}
}

func TestDWayHeap_Min(t *testing.T) {
	h := NewDWayHeap[int](3)
	h.Insert(10)
	if h.Min() != 10 {
		t.Fatalf("Min()=%d, want 10", h.Min())
	}
	h.Insert(5)
	if h.Min() != 5 {
		t.Fatalf("Min()=%d, want 5", h.Min())
	}
	h.Insert(8)
	if h.Min() != 5 {
		t.Fatalf("Min()=%d, want 5", h.Min())
	}
}

func TestDWayHeap_Duplicates(t *testing.T) {
	h := NewDWayHeap[int](4)
	vals := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	for _, v := range vals {
		h.Insert(v)
	}

	sorted := make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)

	for i, want := range sorted {
		got := h.DelMin()
		if got != want {
			t.Fatalf("DelMin() #%d = %d, want %d", i, got, want)
		}
	}
}

func TestDWayHeap_Strings(t *testing.T) {
	h := NewDWayHeap[string](3)
	h.Insert("banana")
	h.Insert("apple")
	h.Insert("cherry")

	expected := []string{"apple", "banana", "cherry"}
	for i, want := range expected {
		got := h.DelMin()
		if got != want {
			t.Fatalf("DelMin() #%d = %q, want %q", i, got, want)
		}
	}
}

func TestDWayHeap_PanicOnEmptyDelMin(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on DelMin from empty heap")
		}
	}()
	h := NewDWayHeap[int](2)
	h.DelMin()
}

func TestDWayHeap_PanicOnEmptyMin(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on Min from empty heap")
		}
	}()
	h := NewDWayHeap[int](2)
	h.Min()
}

func TestDWayHeap_PanicOnInvalidD(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for d < 2")
		}
	}()
	NewDWayHeap[int](1)
}
