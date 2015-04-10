package goggle

import (
	"testing"
)

func TestFindMinMaxSpan(t *testing.T) {
	cases := []struct {
		xs                            []int
		want_max, want_min, want_span int
	}{
		{[]int{}, 0, 0, -1},
		{[]int{3, 2, 1}, 0, 2, 2},
	}
	for _, c := range cases {
		got_min := FindMin(c.xs)
		got_max := FindMax(c.xs)
		got_span := FindSpan(c.xs)
		if got_min != c.want_min &&
			got_max != c.want_max &&
			got_span != c.want_span {
			t.Errorf("FindMin(%q) == %q, want %q", c.xs, got_min, c.want_min)
			t.Errorf("FindMax(%q) == %q, want %q", c.xs, got_max, c.want_max)
			t.Errorf("FindSpan(%q) == %q, want %q", c.xs, got_span, c.want_span)
		}

	}
}

func TestFindNextIndex(t *testing.T) {
	cases := []struct {
		xs      []int
		x, want int
	}{
		{[]int{}, 0, -1},
		{[]int{1, 2, 3}, 2, 2},
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 2, 3}, 3, -1},
	}
	for _, c := range cases {
		got := FindNextIndex(c.xs, c.x)
		if got != c.want {
			t.Errorf("FindNextIndex(%q, %q) == %q, want %q", c.xs, c.x, got, c.want)
		}

	}
}

func TestFindMinSpan(t *testing.T) {
	cases := []struct {
		pos  [][]int
		want []int
	}{
		{[][]int{[]int{0}}, []int{0}},
		{[][]int{[]int{1, 2}, []int{3, 4}, []int{5, 6}}, []int{2, 3, 5}},
		{[][]int{[]int{3, 10}, []int{1, 7}, []int{6}}, []int{10, 7, 6}},
	}
	for _, c := range cases {
		got := FindMinSpan(c.pos)
		for i, g := range got {
			if g != c.want[i] {
				t.Errorf("FindNextIndex(%v) == %v, want %v", c.pos, got, c.want)
			}
		}

	}
}

func TestAllIndex(t *testing.T) {
	cases := []struct {
		s, sep string
		want   []int
	}{
		{"argarg", "arg", []int{0, 3}},
		{"argarg", "foo", []int{-1}},
		{"foobar", "oba", []int{2}},
		{"obobo", "obo", []int{0}},
	}
	for _, c := range cases {
		got := IndexAll(c.s, c.sep)
		for i, g := range got {
			if g != c.want[i] {
				t.Errorf("IndexAll(%q, %q) == %q, want %q", c.s, c.sep, got, c.want)
			}
		}
	}
}

/*
func TestSearch(t *testing.T) {
	g := Goggle{"foo bar tar bar car boo"}
	cases := []struct {
		s    string
		want []string
	}{
		{"tar car", []string{"tar bar car"}},
	}
	for _, c := range cases {
		got := g.Search(c.s)
		for i, g := range got {

			if g != c.want[i] {
				t.Errorf("Search(%q) == %q, want %q", c.s, got, c.want)
			}
		}

	}

}
*/
