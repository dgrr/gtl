package gtl

import (
	"github.com/dgrr/gtl"
)

type Tree[T any] struct {
	data T

	name string
	path []string

	depth int

	nodes []*Tree[T]
}

func (tree *Tree[T]) Trees() []*Tree[T] {
	return tree.nodes
}

func (tree *Tree[T]) Name() string {
	return tree.name
}

func (tree *Tree[T]) Path() []string {
	return tree.path
}

func (tree *Tree[T]) Data() T {
	return tree.data
}

// Depth returns the absolute depth of the node in the tree.
func (tree *Tree[T]) Depth() int {
	return tree.depth
}

// Get takes the lowest level data from the path, and returns
// the depth in which the data is and the data.
//
// If no data was found, depth = -1, and data = nil.
//
// Get is useful for getting the deepest object in the path, for example:
// tree.Set(2, "a","b")
// tree.Get("a","b","c")
// As `c` doesn't exist, Get will return depth = 1 and data = 2.
//
// To get the exact data from the path use Fetch.
func (tree *Tree[T]) Get(path ...string) (depth int, opt gtl.Optional[T]) {
	depth = -1

	lastTree := tree.getLastTree(path...)
	if lastTree != nil {
		depth = lastTree.depth
		opt.Set(lastTree.data)
	}

	return
}

// Fetch returns the data from the exact path. If the path doesn't exist nil will be returned.
func (tree *Tree[T]) Fetch(path ...string) (opt gtl.Optional[T]) {
	lastTree := tree.getLastTree(path...)
	if lastTree != nil {
		// len - 1 because the depth starts on 0
		if lastTree.depth == len(path)-1 {
			opt.Set(lastTree.data)
		}
	}

	return
}

// GetTree gets the latest node from the path.
//
// It works like Get but fetching the Tree instead of the depth and data.
func (tree *Tree[T]) GetTree(path ...string) *Tree[T] {
	return tree.getLastTree(path...)
}

// Range will travel the tree using a Depth-first search algo.
func (tree *Tree[T]) Range(fn func(*Tree[T]) bool, path ...string) {
	tree.rangeOver(fn, -1, path...)
}

// RangeAll will travel all the nodes from the tree using a Depth-first search algo.
func (tree *Tree[T]) RangeAll(fn func(*Tree[T]) bool) {
	for _, child := range tree.nodes {
		child.travel(-1, fn)
	}
}

func (tree *Tree[T]) rangeOver(fn func(*Tree[T]) bool, maxDepth int, path ...string) {
	nn := tree.getLastTree(path...)
	if nn == nil {
		return
	}

	// inner node
	for _, in := range nn.nodes {
		if !in.travel(maxDepth, fn) {
			break
		}
	}
}

// RangeLimit limits the range to a maximum depth.
//
// The order in which RangeLimit travels the paths is the same as Range does.
func (tree *Tree[T]) RangeLimit(fn func(*Tree[T]) bool, maxDepth int) {
	tree.rangeOver(fn, maxDepth)
}

// RangeLevel will range over a specific level of the tree.
func (tree *Tree[T]) RangeLevel(fn func(*Tree[T]) bool, level int) {
	for _, child := range tree.nodes {
		child.rangeLevel(fn, 0, level)
	}
}

func (tree *Tree[T]) rangeLevel(fn func(*Tree[T]) bool, depth, level int) bool {
	if level-depth == 0 {
		return fn(tree)
	}

	for _, child := range tree.nodes {
		if !child.rangeLevel(fn, depth+1, level) {
			return false
		}
	}

	return true
}

func (tree *Tree[T]) travel(maxDepth int, fn func(*Tree[T]) bool) bool {
	if maxDepth > -1 && tree.depth >= maxDepth {
		return true
	}

	for _, nn := range tree.nodes {
		if !nn.travel(maxDepth, fn) {
			return false
		}
	}

	return fn(tree)
}

func (tree *Tree[T]) getLastTree(path ...string) *Tree[T] {
	if len(path) == 0 {
		return tree
	}

	for _, newTree := range tree.nodes {
		// if the path is found, advance the index, otherwise, look for the path in the next Tree
		if newTree.name == path[0] {
			nLastTree := newTree.getLastTree(path[1:]...)
			if nLastTree != nil {
				return nLastTree
			}
		}
	}

	return tree
}

func (tree *Tree[T]) Set(data T, path ...string) {
	tree.set(data, 0, []string{}, path...)
}

func (tree *Tree[T]) set(data T, depth int, cumPath []string, path ...string) {
	if len(path) == 0 {
		tree.data = data
		return
	}

	cumPath = append(cumPath, path[0])

	// iterate over the nodes
	for _, newTree := range tree.nodes {
		// if the node was found, then continue iterating
		if newTree.name == path[0] {
			newTree.set(data, depth+1, cumPath, path[1:]...)
			return
		}
	}

	// if not found, create the node
	newTree := &Tree[T]{
		name:  path[0],
		depth: depth,
		path:  cumPath,
	}
	tree.nodes = append(tree.nodes, newTree)

	// and try to go until the end
	newTree.set(data, depth+1, cumPath, path[1:]...)
}

func (tree *Tree[T]) Del(path ...string) {
	tree.del(path...)
}

func (tree *Tree[T]) del(path ...string) bool {
	for i, newTree := range tree.nodes {
		if newTree.name == path[0] {
			if len(path) == 1 {
				tree.nodes = append(tree.nodes[:i], tree.nodes[i+1:]...)

				return true
			}

			if newTree.del(path[1:]...) {
				return true
			}
		}
	}

	return false
}
