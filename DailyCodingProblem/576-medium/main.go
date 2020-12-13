package main

import (
	"fmt"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

/*
Problem statement:
    Write a function, throw_dice(N, faces, total), that determines how many ways
    it is possible to throw N dice with some number of faces each to get a
    specific total.
Example:
    throw_dice(3, 6, 7) should equal 15
*/

/*
Solver encapsulates static parameters and memoizes subproblem answers
    Faces: # faces per die
    M: Once calculated, Solve(n, t) is stored in M[n][t]
*/
type Solver struct {
	Faces int
	M     map[int]map[int]int
}

// NewSolver constructor
func NewSolver(f int) *Solver {
	return &Solver{Faces: f, M: map[int]map[int]int{}}
}

// Solve subproblem
func (s *Solver) Solve(n int, total int) int {
	if n > total || n*s.Faces < total {
		return 0
	}

	// Check if memoized
	if v, found := s.check(n, total); found {
		return v
	}

	v := 0
	if total < s.Faces+n {
		// Calculate directly with combinatorics
		v = combin.Binomial(total-1, n-1)
	} else {
		// Recurse subproblems
		for i := 1; i <= s.Faces; i++ {
			v += s.Solve(n-1, total-i)
		}
	}

	s.memoize(n, total, v)
	return v
}

// check for memoized subproblem solution
func (s *Solver) check(n int, total int) (int, bool) {
	if m, okay := s.M[n]; okay {
		if v, okay := m[total]; okay {
			return v, true
		}
	}
	return 0, false
}

// memoize subproblem solution
func (s *Solver) memoize(n int, total int, v int) {
	if m, okay := s.M[n]; okay {
		m[total] = v
	} else {
		m := map[int]int{total: v}
		s.M[n] = m
	}
}

func throwDice(n int, faces int, total int) int {
	s := NewSolver(faces)
	return s.Solve(n, total)
}

func main() {
	n, faces, total := 3, 6, 7
	if len(os.Args) == 4 {
		n, _ = strconv.Atoi(os.Args[1])
		faces, _ = strconv.Atoi(os.Args[2])
		total, _ = strconv.Atoi(os.Args[3])
	}
	fmt.Println(throwDice(n, faces, total))
}
