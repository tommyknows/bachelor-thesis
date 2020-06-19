package tree

type Tree interface {
	tree()
}

type Leaf int

func (Leaf) tree() {}

type Node struct {
	v           int
	left, right Tree
}

func (Node) tree() {}

func depthAcc(n int, ts []Tree) int {
	if len(ts) == 0 {
		return n
	}

	flatten := func(n Tree) []Tree {
		switch no := n.(type) {
		case Leaf:
			return []Tree{}
		case Node:
			return []Tree{no.left, no.right}
		}
		return nil
	}
	return depthAcc(n+1, concatTree(flatten, ts))
}

func concatTree(f func(n Tree) []Tree, ts []Tree) []Tree {
	return foldr(
		// concat
		func(t Tree, acc []Tree) []Tree { return append(f(t), acc...) },
		[]Tree{}, ts,
	)
}

func depth(t Tree) int {
	return depthAcc(0, []Tree{t})
}

