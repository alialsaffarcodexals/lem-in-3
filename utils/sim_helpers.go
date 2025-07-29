package utils

import (
	"fmt"
	"sort"
	"strings"
)

// moveEvent represents an ant move to a room during simulation.
type moveEvent struct {
	id   int
	room *Room
}

// moveStartedAnts advances ants that have already started along their paths.
func moveStartedAnts(g *Graph, paths [][]*Room, pos []int, started []bool, occupancy map[*Room]int, finished *int, route []int) []moveEvent {
	var evts []moveEvent
	for id := 0; id < len(route); id++ {
		if !started[id] {
			continue
		}
		p := paths[route[id]]
		if pos[id] < len(p)-1 {
			next := p[pos[id]+1]
			if next == g.End || occupancy[next] == 0 {
				if p[pos[id]] != g.Start {
					delete(occupancy, p[pos[id]])
				}
				pos[id]++
				if next != g.End {
					occupancy[next] = id + 1
				} else {
					*finished++
				}
				evts = append(evts, moveEvent{id: id + 1, room: next})
			}
		}
	}
	return evts
}

// startQueuedAnts launches new ants at the beginning of their paths if possible.
func startQueuedAnts(g *Graph, paths [][]*Room, queues [][]int, started []bool, pos []int, occupancy map[*Room]int, finished *int) []moveEvent {
	var evts []moveEvent
	for i, q := range queues {
		if len(q) == 0 {
			continue
		}
		ant := q[0]
		next := paths[i][1]
		if next == g.End || occupancy[next] == 0 {
			started[ant] = true
			pos[ant] = 1
			if next != g.End {
				occupancy[next] = ant + 1
			} else {
				*finished++
			}
			queues[i] = q[1:]
			evts = append(evts, moveEvent{id: ant + 1, room: next})
		}
	}
	return evts
}

// formatEvents converts move events into a formatted string output.
func formatEvents(evts []moveEvent) string {
	sort.Slice(evts, func(i, j int) bool { return evts[i].id < evts[j].id })
	line := make([]string, len(evts))
	for i, e := range evts {
		line[i] = fmt.Sprintf("L%d-%s", e.id, e.room.Name)
	}
	return strings.Join(line, " ")
}
