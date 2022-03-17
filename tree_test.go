package gtl

import (
	"strings"
	"testing"
)

func TestTreeKeyInt(t *testing.T) {
	var tree Tree[int, string]

	tree.Set("hello", 1, 2, 3)
	tree.Set("world", 1, 2)
	tree.Set("idk", 1)

	data := tree.Fetch(1, 2, 3)
	if data.Get() != "hello" {
		t.Fatalf("expected hello, got `%s`", data.Get())
	}

	data = tree.Fetch(1, 2)
	if data.Get() != "world" {
		t.Fatalf("expected world, got %s", data.Get())
	}
}

func TestTree(t *testing.T) {
	var tree Tree[string, int]

	tree.Set(1, strings.Split("d/c/b/a", "/")...)
	tree.Set(2, strings.Split("d/c/b/a", "/")...)
	tree.Set(8, strings.Split("d/c/b", "/")...)
	tree.Set(3, strings.Split("d/c", "/")...)
	tree.Set(4, "d")

	depth, data := tree.Get(strings.Split("d/c/b/a", "/")...)
	if !data.HasValue() || depth != 3 {
		t.Fatalf("failed, expecting depth 3, got %d", depth)
	}

	if data.Get() != 2 {
		t.Fatal("n is not 2")
	}

	depth, data = tree.Get("d", "c")
	if !data.HasValue() || depth != 1 {
		t.Fatal("failed 2")
	}

	if data.Get() != 3 {
		t.Fatal("n is not 3")
	}

	depth, data = tree.Get("d")
	if !data.HasValue() || depth != 0 {
		t.Fatal("failed 3")
	}

	if data.Get() != 4 {
		t.Fatal("n is not 4")
	}

	depth, data = tree.Get("d", "c", "b", "a", "z", "s")
	if !data.HasValue() || depth != 3 {
		t.Fatal("failed 4")
	}

	if data.Get() != 2 {
		t.Fatal("n is not 2")
	}

	tree.Del("d", "c", "b", "a")

	data = tree.Fetch("d", "c", "b", "a")
	if data.HasValue() {
		t.Fatal("Data has not been removed from `a`")
	}

	data = tree.Fetch("d", "c", "b")
	if !data.HasValue() {
		t.Fatal("Data has been removed from `b`")
	}

	tree.Del("d", "c")

	data = tree.Fetch("d", "c", "b")
	if data.HasValue() {
		t.Fatal("Data has not been removed from `c`")
	}

	data = tree.Fetch("d", "c")
	if data.HasValue() {
		t.Fatal("Data has not been removed from `c`")
	}

	data = tree.Fetch("d")
	if !data.HasValue() {
		t.Fatal("Data has been removed from `d`")
	}
}

func TestTreeTravel(t *testing.T) {
	var tree Tree[string, int]

	tree.Set(1, strings.Split("d/c/b/a", "/")...)
	tree.Set(2, strings.Split("d/c/b", "/")...)
	tree.Set(3, strings.Split("d/c", "/")...)
	tree.Set(4, "d")
	tree.Set(80, "d", "c", "x")

	rangeCount := 0
	tree.RangeAll(func(node *Tree[string, int]) bool {
		rangeCount++

		return true
	})

	if rangeCount != 5 {
		t.Fatalf("Unexpected range: %d<>5", rangeCount)
	}

	tree.RangeLimit(func(node *Tree[string, int]) bool {
		if node.Name() != "d" {
			t.Fatal("expected d only")
		}

		return true
	}, 1)

	iter := 0
	tree.RangeLevel(func(node *Tree[string, int]) bool {
		iter++
		if node.Name() != "d" {
			t.Fatalf("Expected d, got %s", node.Name())
		}
		return true
	}, 0)
	if iter != 1 {
		t.Fatalf("unexpected iterations: %d", iter)
	}

	countLevel := 0
	tree.RangeLevel(func(node *Tree[string, int]) bool {
		countLevel++
		if node.Name() != "c" {
			t.Fatalf("Expected b, got %s", node.Name())
		}
		return true
	}, 1) // should iterate over c's child
	if countLevel != 1 {
		t.Fatalf("Unexpected count level: %d", countLevel)
	}

	countLevel = 0
	tree.RangeLevel(func(node *Tree[string, int]) bool {
		countLevel++
		if countLevel == 1 && node.Name() != "b" {
			t.Fatalf("Expected b, got %s", node.Name())
		}
		if countLevel == 2 && node.Name() != "x" {
			t.Fatalf("Expected x, got %s", node.Name())
		}
		if countLevel > 2 {
			t.Fatalf("Iterated over childs more than two times")
		}
		return true
	}, 2) // should iterate over c's child
	if countLevel != 2 {
		t.Fatalf("Unexpected count level: %d", countLevel)
	}
	tree.Del("d", "c", "x")

	count := 0
	tree.RangeLimit(func(node *Tree[string, int]) bool {
		switch count {
		case 0:
			if node.Name() != "c" {
				t.Fatal("expected c only")
			}
		case 1:
			if node.Name() != "d" {
				t.Fatal("expected d only")
			}
		}

		count++

		return true
	}, 2)

	order := []string{"a", "b", "c", "d"}
	datas := []int{1, 2, 3, 4}

	idx := 0
	tree.Range(func(n *Tree[string, int]) bool {
		if n.Name() != order[idx] {
			t.Fatalf("unexpected element. Expected %s got %s", order[idx], n.Name())
		}

		if n.Data().Get() != datas[idx] {
			t.Fatalf("unexpected data. Expected %d got %d", datas[idx], n.Data().Get())
		}

		path := n.Path()
		for i, j := idx, len(path)-1; i < len(path); i++ {
			if path[i] != order[j] {
				t.Fatalf("unexpected path: %d: %s <> %s", idx, path[i], order[j])
			}

			j--
		}

		idx++

		return true
	})
	if idx != len(order) {
		t.Fatal("Not travelled enough")
	}

	idx = 0
	tree.Range(func(n *Tree[string, int]) bool {
		if n.Name() != order[idx] {
			t.Fatalf("unexpected element. Expected %s got %s", order[idx], n.Name())
		}

		if n.Data().Get() != datas[idx] {
			t.Fatalf("unexpected data. Expected %d got %d", datas[idx], n.Data().Get())
		}

		idx++

		return true
	})
	if idx == 3 {
		t.Fatal("expecting 3 iterations")
	}
}
