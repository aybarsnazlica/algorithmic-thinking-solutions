package intro

import "fmt"

func findShortestLineIndex(lines []int) int {
	shortestIndex := 0
	for j := 1; j < len(lines); j++ {
		if lines[j] < lines[shortestIndex] {
			shortestIndex = j
		}
	}
	return shortestIndex
}

func solve(lines []int, m int) []int {
	results := make([]int, m)
	for i := 0; i < m; i++ {
		shortestIndex := findShortestLineIndex(lines)
		results[i] = lines[shortestIndex]
		lines[shortestIndex]++
	}
	return results
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
	results := solve(lines, m)
	for _, val := range results {
		fmt.Println(val)
	}
}
