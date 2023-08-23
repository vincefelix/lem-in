package lem_in

import (
	Func "lem_in/tools"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	test := []struct {
		Case     []string
		Expected []string
	}{
		{[]string{"f3", "f11", "f1", "f9", "f8"}, []string{"f1", "f3", "f8", "f9", "f11"}},
		{[]string{"f3", "f11", "f1"}, []string{"f1", "f3", "f11"}},
		{[]string{"f3"}, []string{"f3"}},
		{[]string{"f1", "f1", "f1", "f1", "f1"}, []string{"f1", "f1", "f1", "f1", "f1"}},
		{nil, nil},
	}
	for i, v := range test {
		Func.Sort(v.Case)
		if strings.Join(v.Case, " ") != strings.Join(v.Expected, " ") {
			t.Fatalf("case n: %v - expected: %v - got: %v", i+1, v.Expected, v.Case)
		} else {
			t.Log("test succeeded !")
		}
	}
}

func TestGetnumber(t *testing.T) {
	test := []struct {
		Case     string
		Expected int
	}{
		{"L1", 1},
		{"L1-rtuhgcv", 1},
		{"L1-00", 1},
		{"L100-00-0", 100},
		{"Lo100-a", 0},
		{"500", 0},
	}

	for i, v := range test {
		if Func.GetNumber(v.Case) != v.Expected {
			t.Fatalf("case n: %v - expected: %v - got: %v", i+1, v.Expected, Func.GetNumber(v.Case))
		} else {
			t.Log("test succeeded !")
		}
	}
}

func TestSequence(t *testing.T) {
	object1 := Func.Ant{
		{Path: []string{"1", "3", "4", "0"}, Ant_nbr: 2, Passing_order: []int{1, 3}},
		{Path: []string{"1", "2", "5", "6", "0"}, Ant_nbr: 1, Passing_order: []int{2}},
	}

	object2 := Func.Ant{
		{Path: []string{"1", "3", "4", "0"}, Ant_nbr: 2, Passing_order: []int{1, 2, 3, 4}},
	}

	seq1 := make(map[int][]string)
	seq1[0] = []string{"L1-3", "L2-2"}
	seq1[1] = []string{"L1-4", "L3-3", "L2-5"}
	seq1[2] = []string{"L1-0", "L3-4", "L2-6"}
	seq1[3] = []string{"L3-0", "L2-0"}

	seq2 := make(map[int][]string)
	seq2[0] = []string{"L1-3"}
	seq2[1] = []string{"L1-4", "L2-3"}
	seq2[2] = []string{"L1-0", "L2-4", "L3-3"}
	seq2[3] = []string{"L2-0", "L3-4", "L4-3"}
	seq2[4] = []string{"L3-0", "L4-4"}
	seq2[5] = []string{"L4-0"}

	test := []struct {
		Case     Func.Ant
		Expected map[int][]string
	}{
		{object1, seq1},
		{object2, seq2},
	}

	for i, v := range test {
		tempcase := v.Case.Sequence()
		tempxted := v.Expected
		for j, k := range tempcase {
			if strings.Join(k, " ") != strings.Join(tempxted[j], " ") {
				t.Fatalf("case n: %v - expected: %v - got: %v", i+1, v.Expected, v.Case.Sequence())
			} else {
				t.Log("test succeeded !")
			}
		}
	}
}

func TestAnt_per_path(t *testing.T) {
	object1 := Func.Ant{
		{Path: []string{"1", "3", "4", "0"}, Ant_nbr: 2, Passing_order: []int{1, 3}},
		{Path: []string{"1", "2", "5", "6", "0"}, Ant_nbr: 1, Passing_order: []int{2}},
	}
	pathtab1 := [][]string{
		{"1", "3", "4", "0"},
		{"1", "2", "5", "6", "0"},
	}
	antsize1 := 3

	object2 := Func.Ant{
		{Path: []string{"1", "3", "4", "0"}, Ant_nbr: 2, Passing_order: []int{1, 2, 3, 4}},
	}
	pathtab2 := [][]string{
		{"1", "3", "4", "0"},
	}
	antsize2 := 4

	test := []struct {
		Case     Func.Ant
		Pathtab  [][]string
		antsize  int
		Expected [][]int
	}{
		{object1, pathtab1, antsize1, [][]int{{1, 3}, {2}}},
		{object2, pathtab2, antsize2, [][]int{{1, 2, 3, 4}}},
	}

	for i, v := range test {
		v.Case.Ant_per_path(v.antsize, v.Pathtab)
		for j, k := range v.Case[i].Passing_order {
			if k != v.Expected[i][j] {
				t.Fatalf("case n: %v - expected: %v - got: %v", i+1, v.Expected, v.Case[i].Passing_order)
			} else {
				t.Log("test succeeded !")
			}
		}

	}

}
