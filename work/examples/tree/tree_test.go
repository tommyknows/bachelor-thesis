package tree

import "testing"

func TestTree(t *testing.T) {
	tr := Node{2, Leaf(2), Node{3, Leaf(4), Leaf(5)}}

	tr2 := Node{1, tr, tr}

	t.Logf("%v", depth(tr))
	t.Logf("%v", depth(tr2))
}
