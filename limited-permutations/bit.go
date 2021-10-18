package main

import "fmt"

const P = false

type order struct {
	u int
	v int
}

var now []int

func printAll(ni, status int, ways [][]int) {
	if status == 0 {
		fmt.Println(now)
		return
	}
	for _, v := range ways[status] {
		to := status ^ (1 << v)
		now[ni] = v + 1
		printAll(ni+1, to, ways)
	}
}

func Permutations(n int, ords []order) int64 {
	e := make([][]bool, n)
	for i := range e {
		e[i] = make([]bool, n)
	}
	for _, ord := range ords {
		e[ord.v-1][ord.u-1] = true
	}

	m := 1 << n
	f := make([]int64, m)
	ways := make([][]int, m)
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
			if add && f[t] > 0 {
				f[s] += f[t]
				ways[s] = append(ways[s], i)
			}
		}
	}
	now = make([]int, n)
	if P {
		printAll(0, m-1, ways)
	}
	return f[m-1]
}

func main() {
	n := 12
	ords := make([]order, 0)
	for i := 1; i < n; i++ {
		ords = append(ords, order{i, n})
	}
	fmt.Println(Permutations(n, ords))
}
