package utils

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math/rand"
)

type Edge struct {
	U int
	V int
}

func KargerMinCut(edges []Edge, vertices int) mapset.Set[Edge] {
	dsu, v, e := Initialize(vertices), vertices, len(edges)

	for v > 2 {
		// Getting a random integer in the range [0, e-1].
		i := rand.Intn(e)
		set1, set2 := dsu.find(edges[i].U), dsu.find(edges[i].V)
		if set1 != set2 {
			dsu.union(edges[i].U, edges[i].V)
			v--
		}
	}
	cutset := mapset.NewSet[Edge]()
	for _, edge := range edges {
		set1, set2 := dsu.find(edge.U), dsu.find(edge.V)
		if set1 != set2 {
			cutset.Add(edge)
		}
	}
	return cutset
}
