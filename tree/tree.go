package tree

// Node is one assembled tree node: a reference to the source item
// plus its (already linked) children.
type Node[I comparable, V any] struct {
	ID       I
	Value    V
	Level    int
	Children []*Node[I, V]
}

// BuildTree assembles a forest from a flat list of items, in the spirit
// of hutool's TreeUtil.list2Tree.
//
// id and parent extract the linking keys from each item; items whose
// parent key equals rootID become forest roots. Children are attached
// regardless of input order (parents may appear after their children).
//
// Items that are unreachable from any root — their parent key names a
// missing node, or they participate in a parent cycle — are returned
// separately as orphans, with Level left at its zero value.
func BuildTree[I comparable, V any](items []V, rootID I, id func(V) I, parent func(V) I) ([]*Node[I, V], []V) {
	nodes := make([]*Node[I, V], 0, len(items))
	byID := make(map[I]*Node[I, V], len(items))
	for _, it := range items {
		n := &Node[I, V]{ID: id(it), Value: it}
		nodes = append(nodes, n)
		byID[n.ID] = n
	}

	var roots []*Node[I, V]
	for _, n := range nodes {
		p := parent(n.Value)
		if class, ok := byID[p]; ok {
			// parent present: attach regardless of order
			class.Children = append(class.Children, n)
		}
		if p == rootID {
			roots = append(roots, n)
		}
	}

	visited := make(map[*Node[I, V]]bool, len(nodes))
	var visit func(n *Node[I, V])
	visit = func(n *Node[I, V]) {
		if visited[n] {
			return
		}
		visited[n] = true
		for _, c := range n.Children {
			visit(c)
		}
	}
	var assignLevels func(n *Node[I, V], level int)
	assignLevels = func(n *Node[I, V], level int) {
		n.Level = level
		for _, c := range n.Children {
			assignLevels(c, level+1)
		}
	}
	for _, r := range roots {
		visit(r)
		assignLevels(r, 0)
	}

	var orphans []V
	for _, n := range nodes {
		if !visited[n] {
			orphans = append(orphans, n.Value)
		}
	}
	return roots, orphans
}