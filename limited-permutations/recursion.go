package main

import "fmt"

type order struct {
	u int
	v int
}

var now []int
var pool []int
var edges [][]int
var degree []int

func find(nlen, plen int) int {
	if plen == 0 {
		if nlen+1 == len(edges) {
			return 1
		}
		return 0
	}
	ret := 0
	for i := 0; i < plen; i++ {
		ni := pool[i]
		pool[i], pool[plen-1] = pool[plen-1], pool[i]
		nplen := plen - 1
		for _, v := range edges[ni] {
			degree[v]--
			if degree[v] == 0 {
				pool[nplen] = v
				nplen++
			}
		}

		now[nlen] = ni
		ret += find(nlen+1, nplen)

		for _, v := range edges[ni] {
			degree[v]++
		}
		pool[plen-1] = pool[i]
		pool[i] = ni
	}
	return ret
}

func Permutations(n int, ords []order) int {
	edges = make([][]int, n+1)
	degree = make([]int, n+1)
	for i := 1; i <= n; i++ {
		edges[i] = make([]int, 0)
		degree[i] = 0
	}
	for _, ord := range ords {
		edges[ord.u] = append(edges[ord.u], ord.v)
		degree[ord.v]++
	}
	now = make([]int, n)
	pool = make([]int, n)
	plen := 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 {
			pool[plen] = i
			plen++
		}
	}
	return find(0, plen)
}

func main() {
	fmt.Println(Permutations(12, []order{
		{3, 5},
		{8, 4},
	}))
}
