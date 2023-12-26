package day25

import (
	"advent-of-code/utils"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"strings"
)

var id = 0
var ids = make(map[string]int)

func Solve() {
	lines := utils.ReadLines("/day25/input")
	adj, edges := make(map[int]mapset.Set[int]), make([]utils.Edge, 0)

	for _, line := range lines {
		src, dstr, _ := strings.Cut(line, ": ")
		dest := strings.Split(dstr, " ")
		for _, d := range dest {
			edge := addLink(adj, src, d)
			edges = append(edges, edge)
		}
	}

	mincut, vertices := mapset.NewSet[utils.Edge](), id
	for mincut.Cardinality() != 3 {
		mincut = utils.KargerMinCut(edges, vertices)
	}

	for rem := range mincut.Iter() {
		adj[rem.U].Remove(rem.V)
		adj[rem.V].Remove(rem.U)
	}

	queue, visited := arrayqueue.New(), mapset.NewSet[int]()
	queue.Enqueue(0)
	visited.Add(0)

	for !queue.Empty() {
		vertex, _ := queue.Dequeue()
		for n := range adj[vertex.(int)].Iter() {
			if !visited.Contains(n) {
				visited.Add(n)
				queue.Enqueue(n)
			}
		}
	}

	ans := visited.Cardinality() * (vertices - visited.Cardinality())
	fmt.Println("Answer: ", ans) // 547080
}

func getId(src string) int {
	if _, hs := ids[src]; !hs {
		ids[src] = id
		id++
	}
	return ids[src]
}

func addLink(adj map[int]mapset.Set[int], src, dest string) utils.Edge {
	sid, did := getId(src), getId(dest)
	if _, hs := adj[sid]; !hs {
		adj[sid] = mapset.NewSet[int]()
	}
	if _, hd := adj[did]; !hd {
		adj[did] = mapset.NewSet[int]()
	}
	adj[sid].Add(did)
	adj[did].Add(sid)
	return utils.Edge{U: sid, V: did}
}
