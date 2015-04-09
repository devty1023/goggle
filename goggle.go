package main

import (
	"fmt"
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

func (g *Goggle) Search(query string) []string {
	terms := strings.Split(query, " ")
	for i, t := range terms {
		terms[i] = strings.ToLower(t)
	}

	pos := make([][]int, 0)
	for _, t := range terms {
		pos = append(pos, IndexAll(g.text, t))
	}

	fmt.Println(pos)

	return terms
}

// Like strings.Index(s) except returns
// all indicies
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
func main() {
	fmt.Println("goggle goggle")
	g := NewGoggle("hello World")
	g.Add("goodbye World")
	fmt.Println(g.Search("world hello"))
}
