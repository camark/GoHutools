package tree

import (
	"testing"
)

// Menu is a flat record used as tree source material.
type Menu struct {
	Code   string
	Parent string
	Title  string
}

func idOf(m *Menu) string    { return m.Code }
func parentOf(m *Menu) string { return m.Parent }

func TestBuildTreeBasic(t *testing.T) {
	items := []*Menu{
		{Code: "A", Parent: "0", Title: "root A"},
		{Code: "B", Parent: "A", Title: "child B"},
		{Code: "C", Parent: "A", Title: "child C"},
		{Code: "D", Parent: "0", Title: "root D"},
		{Code: "E", Parent: "D", Title: "child E"},
	}
	roots, orphans := BuildTree(items, "0", idOf, parentOf)
	if len(orphans) != 0 {
		t.Fatalf("orphans = %v", orphans)
	}
	if len(roots) != 2 {
		t.Fatalf("roots = %d, want 2", len(roots))
	}
	a, d := roots[0], roots[1]
	if a.ID != "A" || a.Level != 0 {
		t.Errorf("root A = %s/%d", a.ID, a.Level)
	}
	if len(a.Children) != 2 || a.Children[0].ID != "B" || a.Children[1].ID != "C" {
		t.Errorf("A children = %v", a.Children)
	}
	if a.Children[0].Level != 1 || a.Children[0].Value.Title != "child B" {
		t.Errorf("B node = %d %s", a.Children[0].Level, a.Children[0].Value.Title)
	}
	if len(d.Children) != 1 || d.Children[0].ID != "E" || d.Children[0].Level != 1 {
		t.Errorf("D children = %v", d.Children)
	}
}

func TestBuildTreeShuffledInput(t *testing.T) {
	// children appear before their parents in the source list
	items := []*Menu{
		{Code: "C", Parent: "B"},
		{Code: "B", Parent: "A"},
		{Code: "A", Parent: "0"},
	}
	roots, orphans := BuildTree(items, "0", idOf, parentOf)
	if len(orphans) != 0 || len(roots) != 1 {
		t.Fatalf("roots=%v orphans=%v", roots, orphans)
	}
	top := roots[0]
	if top.ID != "A" || len(top.Children) != 1 || top.Children[0].ID != "B" ||
		len(top.Children[0].Children) != 1 || top.Children[0].Children[0].ID != "C" {
		t.Errorf("shuffled tree shape wrong: %+v", top)
	}
	if top.Level != 0 || top.Children[0].Level != 1 || top.Children[0].Children[0].Level != 2 {
		t.Error("levels wrong")
	}
}

func TestBuildTreeOrphans(t *testing.T) {
	items := []*Menu{
		{Code: "A", Parent: "0"},
		{Code: "X", Parent: "ghost"}, // parent "ghost" does not exist
		{Code: "B", Parent: "A"},
	}
	roots, orphans := BuildTree(items, "0", idOf, parentOf)
	if len(roots) != 1 || roots[0].ID != "A" {
		t.Fatalf("roots = %v", roots)
	}
	if len(orphans) != 1 || orphans[0].Code != "X" {
		t.Errorf("orphans = %v", orphans)
	}
}

func TestBuildTreeEmptyAndIntKeys(t *testing.T) {
	empty := []int{}
	roots, orphans := BuildTree(empty, 0, func(i int) int { return i }, func(i int) int { return 0 })
	if len(roots) != 0 || len(orphans) != 0 {
		t.Errorf("empty: roots=%v orphans=%v", roots, orphans)
	}

	ints := []int{1, 2, 3} // 1 root, 2 parent=1, 3 parent=2
	roots, _ = BuildTree(ints, 0, func(i int) int { return i }, func(i int) int { return i - 1 })
	if len(roots) != 1 || roots[0].ID != 1 {
		t.Fatalf("int roots = %v", roots)
	}
	if len(roots[0].Children) != 1 || len(roots[0].Children[0].Children) != 1 {
		t.Errorf("int chain wrong")
	}
}

func TestBuildTreeLeafCycleIsOrphan(t *testing.T) {
	// a node whose parent key matches another node that itself is not
	// reachable from root is not orphaned mid-chain — here B->C->B is a
	// cycle; only the root connects. B and C reference each other, neither
	// is a root, and there is no "0" parent, so both end up orphans.
	items := []*Menu{
		{Code: "B", Parent: "C"},
		{Code: "C", Parent: "B"},
	}
	roots, orphans := BuildTree(items, "0", idOf, parentOf)
	if len(roots) != 0 || len(orphans) != 2 {
		t.Errorf("cycle: roots=%v orphans=%v", roots, orphans)
	}
}