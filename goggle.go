package goggle

import (
	"strings"
)

type Goggle struct {
	text string
}

func NewGoggle(text string) *Goggle {
	_text := strings.ToLower(text)
	return &Goggle{_text}
}

func (g *Goggle) Add(text string) *Goggle {
	_text := strings.ToLower(text)
	g.text = g.text + " " + _text
	return g
}

// Return index of the min value
func FindMin(xs []int) (idx int) {
	idx = 0
	for i, v := range xs {
		if v < xs[idx] {
			idx = i
		}
	}
	return
}

// Return index of the max value
func FindMax(xs []int) (idx int) {
	idx = 0
	for i, v := range xs {
		if v > xs[idx] {
			idx = i
		}
	}
	return
}

func FindSpan(xs []int) int {
	if len(xs) == 0 {
		return -1
	}

	min := FindMin(xs)
	max := FindMax(xs)
	return xs[max] - xs[min]
}

// Returns index of the next larger value in the slice
// Assumes slice is sorted
func FindNextIndex(xs []int, v int) (idx int) {
	idx = -1
	for i, x := range xs {
		if x > v {
			idx = i
			break
		}
	}
	return
}

func FindMinSpan(pos [][]int) []int {
	mins := make([]int, 0)
	for _, d := range pos {
		mins = append(mins, d[0])
	}
	for cur_span, min := FindSpan(mins), FindMin(mins); ; {

		replace_min := FindNextIndex(pos[min], mins[min])

		if replace_min == -1 {
			break
		}

		_mins := mins
		_mins[min] = pos[min][replace_min]

		if cur_span < FindSpan(_mins) {
			break
		}

		mins = _mins
		cur_span = FindSpan(mins)
		min = FindMin(mins)
	}

	return mins
}

func (g *Goggle) Search(query string) []string {
	terms := strings.Split(query, " ")
	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}

	pos := make([][]int, 0)
	for _, t := range terms {
		pos = append(pos, IndexAll(g.text, t))
	}

	return terms
}

// Like strings.Index(s) except returns all indices
func IndexAll(s, sep string) []int {
	idx := make([]int, 0)
	for i, prev := strings.Index(s, sep), 0; i != -1; {
		i = i + prev
		idx = append(idx, i)
		prev = i + len(sep)
		i = strings.Index(s[prev:], sep)
	}
	return idx
}
