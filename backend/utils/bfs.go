package utils

import (
	"github.com/Workiva/go-datastructures/queue"
	"slices"
)

const (
	NOT_VISITED = -2	
	START = -1
)

func restorePath(prev map[int]int, start, goal int) []int {
	var path []int
	for at := goal; at != START; at = prev[at] {
		path = append(path, at)
		if at == NOT_VISITED {
			return nil
		}
	}
	slices.Reverse(path)
	return path
}

func BFS(graph map[int][]int, start int, isGoal func(id int) bool) ([]int, error) {
	var q queue.Queue
	q.Put(start)

	prev := make(map[int]int)
	for v := range graph {
		prev[v] = NOT_VISITED
	}

	prev[start] = START

	for !q.Empty() {
		currentSlice, err := q.Get(1)
		if err != nil {
			return nil, err
		}
		current := currentSlice[0].(int)
		if isGoal(current) {
			return restorePath(prev, start, current), nil
		}
		for _, u := range graph[current] {
			if prev[u] >= START {
				continue
			}
			prev[u] = current
			q.Put(u)
		}
	}

	return nil, nil
}