/*
** Copyright 2014 Edward Walker
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http ://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
**
** Description: Useful functions used in various parts of the library
** @author: Ed Walker
 */
package libSvm

import (
	"bufio"
	"sort"
)

const TAU float64 = 1e-12

func absi(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func mini(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxi(a, b int) int {
	if b < a {
		return a
	} else {
		return b
	}
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxf(a, b float64) float64 {
	if b < a {
		return a
	} else {
		return b
	}
}

func MapToSnode(m map[int]float64) []Node {

	keys := make([]int, len(m))
	var i int = 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Ints(keys) // We MUST do this to ensure that we add snodes in ascending key order!
	// Just iterating over the map does not ensure the keys are returned in ascending order.

	x := make([]Node, len(m))

	for i, k := range keys {
		x[i] = Node{Index: k, Value: m[k]}
	}
	return x
}

func SnodeToMap(x []Node) map[int]float64 {

	m := make(map[int]float64)

	for i := 0; x[i].Index != -1; i++ {
		m[x[i].Index] = x[i].Value
	}

	return m
}

/**
 * Reads a complete line with bufio.Reader and returns it as a string
 * Attribution: http://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
 * Thank you Malcolm!
 */
func readline(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
