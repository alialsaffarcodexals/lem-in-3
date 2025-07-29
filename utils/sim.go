package utils

type antState struct {
	id   int
	path int
	pos  int
}

func SimulateMulti(g *Graph, paths [][]*Room) []string {
	if len(paths) == 0 {
		return nil
	}
	route := assignPaths(paths, g.Ants)
	queues := make([][]int, len(paths))
	for ant, p := range route {
		queues[p] = append(queues[p], ant)
	}
	pos := make([]int, len(route))
	started := make([]bool, len(route))
	occupancy := map[*Room]int{}
	finished := 0
	var moves []string
	for finished < len(route) {
		evts := moveStartedAnts(g, paths, pos, started, occupancy, &finished, route)
		evts = append(evts, startQueuedAnts(g, paths, queues, started, pos, occupancy, &finished)...)
		if len(evts) > 0 {
			moves = append(moves, formatEvents(evts))
		}
	}
	return moves
}
