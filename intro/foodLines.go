package intro

import "fmt"

const MaxLines = 100

func shortestLineIndex(lines [MaxLines]int, n int) int {
	shortest := 0

	for j := 1; j < n; j++ {
		if lines[j] < lines[shortest] {
			shortest = j
		}
	}
	return shortest
}

func solve(lines [MaxLines]int, n int, m int) {
	for i := 0; i < m; i++ {
		shortest := shortestLineIndex(lines, n)
		fmt.Println(lines[shortest])
		lines[shortest]++
	}
}

func Solution() {
	var lines [MaxLines]int
	var n, m int
	_, err := fmt.Scanf("%d%d", &n, &m)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &lines[i])
		if err != nil {
			return
		}
	}
	solve(lines, n, m)
}
