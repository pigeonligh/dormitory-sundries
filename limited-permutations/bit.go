package main

import "fmt"

type order struct {
	u int
	v int
}

func Permutations(n int, ords []order) int64 {
	e := make([][]bool, n)
	for i := range e {
		e[i] = make([]bool, n)
	}
	for _, ord := range ords {
		e[ord.u-1][ord.v-1] = true
	}

	m := 1 << n
	f := make([]int64, m)
	f[0] = 1
	for s := 1; s < m; s++ {
		for i := 0; i < n; i++ {
			if s&(1<<i) == 0 {
				continue
			}
			t := s ^ (1 << i)
			add := true
			for j := 0; j < n; j++ {
				if t&(1<<j) == 0 {
					continue
				}
				if e[i][j] {
					add = false
				}
			}
			if add {
				f[s] += f[t]
			}
		}
	}
	return f[m-1]
}

func main() {
	n := 16
	ords := make([]order, 0)
	for i := 1; i < n; i++ {
		ords = append(ords, order{i, n})
	}
	fmt.Println(Permutations(n, ords))
}
