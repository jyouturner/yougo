package datastructures

import (
	"sort"
)

type SortedMap struct {
	keys []string
	data map[string]int
}

func (m *SortedMap) Len() int {
	return len(m.keys)
}

func (m *SortedMap) Less(i, j int) bool {
	return m.keys[i] < m.keys[j]
}

func (m *SortedMap) Swap(i, j int) {
	m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
}

func (m *SortedMap) Set(key string, value int) {
	m.data[key] = value
	m.keys = append(m.keys, key)
	sort.Strings(m.keys)
}

func (m *SortedMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}
