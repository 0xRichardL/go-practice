package main

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	max := math.MinInt
	m_idx := -1
	for i := range gas {
		gas[i] = gas[i] - cost[i]
		if max < gas[i] {
			m_idx = i
			max = gas[i]
		}
	}
	tank := 0
	for i := m_idx; i != m_idx-1; i = (i + 1) % len(gas) {
		tank += gas[i]
		if tank < 0 {
			return -1
		}
	}
	return m_idx
}
