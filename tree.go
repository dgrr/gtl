package gtl

type Tree[Key comparable, Value any] struct {
	data Optional[Value]

	name Key
	path []Key

	depth int

	nodes []*Tree[Key, Value]
}

func (tree *Tree[Key, Value]) Trees() []*Tree[Key, Value] {
	return tree.nodes
}

func (tree *Tree[Key, Value]) Name() Key {
	return tree.name
}

func (tree *Tree[Key, Value]) Path() []Key {
	return tree.path
}

func (tree *Tree[Key, Value]) Data() Optional[Value] {
	return tree.data
}

// Depth returns the absolute depth of the node in the tree.
func (tree *Tree[Key, Value]) Depth() int {
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
func (tree *Tree[Key, Value]) Get(path ...Key) (depth int, opt Optional[Value]) {
	depth = -1

	lastTree := tree.getLastTree(path...)
	if lastTree != nil {
		depth = lastTree.depth
		opt = lastTree.data
	}

	return
}

// Fetch returns the data from the specific path.
func (tree *Tree[Key, Value]) Fetch(path ...Key) (opt Optional[Value]) {
	if len(path) == 0 {
		return tree.data
	}

	for _, node := range tree.nodes {
		if node.Name() == path[0] {
			return node.Fetch(path[1:]...)
		}
	}

	return
}

// GetTree gets the latest node from the path.
//
// It works like Get but fetching the Tree instead of the depth and data.
func (tree *Tree[Key, Value]) GetTree(path ...Key) *Tree[Key, Value] {
	return tree.getLastTree(path...)
}

// Range will travel the tree using a Depth-first search algo.
func (tree *Tree[Key, Value]) Range(fn func(*Tree[Key, Value]) bool, path ...Key) {
	tree.rangeOver(fn, -1, path...)
}

// RangeAll will travel all the nodes from the tree using a Depth-first search algo.
func (tree *Tree[Key, Value]) RangeAll(fn func(*Tree[Key, Value]) bool) {
	for _, child := range tree.nodes {
		child.travel(-1, fn)
	}
}

func (tree *Tree[Key, Value]) rangeOver(fn func(*Tree[Key, Value]) bool, maxDepth int, path ...Key) {
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
func (tree *Tree[Key, Value]) RangeLimit(fn func(*Tree[Key, Value]) bool, maxDepth int) {
	tree.rangeOver(fn, maxDepth)
}

// RangeLevel will range over a specific level of the tree.
func (tree *Tree[Key, Value]) RangeLevel(fn func(*Tree[Key, Value]) bool, level int) {
	for _, child := range tree.nodes {
		child.rangeLevel(fn, 0, level)
	}
}

func (tree *Tree[Key, Value]) rangeLevel(fn func(*Tree[Key, Value]) bool, depth, level int) bool {
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

func (tree *Tree[Key, Value]) travel(maxDepth int, fn func(*Tree[Key, Value]) bool) bool {
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

func (tree *Tree[Key, Value]) getLastTree(path ...Key) *Tree[Key, Value] {
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

func (tree *Tree[Key, Value]) Set(data Value, path ...Key) {
	tree.set(data, 0, []Key{}, path...)
}

func (tree *Tree[Key, Value]) set(data Value, depth int, cumPath []Key, path ...Key) {
	if len(path) == 0 {
		tree.data.Set(data)
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
	newTree := &Tree[Key, Value]{
		name:  path[0],
		depth: depth,
		path:  cumPath,
	}
	tree.nodes = append(tree.nodes, newTree)

	// and try to go until the end
	newTree.set(data, depth+1, cumPath, path[1:]...)
}

func (tree *Tree[Key, Value]) Del(path ...Key) {
	tree.del(path...)
}

func (tree *Tree[Key, Value]) del(path ...Key) bool {
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
