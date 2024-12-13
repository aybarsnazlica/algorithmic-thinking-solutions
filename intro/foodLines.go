package intro

import "fmt"

func findShortestLineIndex(lines []int) int {
	shortest := 0
	for j := 1; j < len(lines); j++ {
		if lines[j] < lines[shortest] {
			shortest = j
		}
	}
	return shortest
}

func solve(lines []int, m int) {
	for i := 0; i < m; i++ {
		shortest := findShortestLineIndex(lines)
		fmt.Println(lines[shortest])
		lines[shortest]++
	}
}

func Solution() {
	var n, m int
	_, err := fmt.Scanf("%d%d", &n, &m)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &lines[i])
		if err != nil {
			fmt.Println("Error reading line value:", err)
			return
		}
	}
	solve(lines, m)
}
