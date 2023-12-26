package utils

type Arr []int

type DisjointSet struct {
	parent Arr
	size   Arr
}

func Initialize(sz int) *DisjointSet {
	parent, size := make(Arr, sz), make(Arr, sz)
	dsu := DisjointSet{parent, size}
	for i := 0; i < sz; i++ {
		parent[i], size[i] = i, 1
	}
	return &dsu
}

func (dsu *DisjointSet) find(item int) int {
	if item == dsu.parent[item] {
		return item
	}
	res := dsu.find(dsu.parent[item])
	dsu.parent[item] = res
	return res
}

func (dsu *DisjointSet) union(v1, v2 int) {
	p1, p2 := dsu.find(v1), dsu.find(v2)
	if p1 == p2 {
		return
	}
	size1, size2 := dsu.size[v1], dsu.size[v2]
	// Make the item having higher tree size the parent of the one having lower size
	if size1 <= size2 {
		dsu.parent[p1] = p2
		dsu.size[p2] += dsu.size[p1]
	} else {
		dsu.parent[p2] = p1
		dsu.size[p1] += dsu.size[p2]
	}
}
