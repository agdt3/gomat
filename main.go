package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const SEP = "_"

type Key struct {
	Row int
	Col int
}

type MatrixMapFloat64 struct {
	R int
	C int
	//matrix map[string]float64
	matrix map[Key]float64
}

func NewMatrixMapFloat64(rn, cn int) *MatrixMapFloat64 {
	m := new(MatrixMapFloat64)
	m.R = rn
	m.C = cn
	//m.matrix = make(map[string]float64)
	m.matrix = make(map[Key]float64)
	return m
}

func toKeyStr(r, c int) string {
	return strings.Join(
		[]string{strconv.Itoa(r), strconv.Itoa(c)},
		SEP)
}

func fromKeyString(s string) (r, c int) {
	//
	return 0, 0
}

func (m *MatrixMapFloat64) Get(r, c int) (float64, error) {
	if r < 0 || r > m.R || c < 0 || c > m.C {
		return 0, errors.New("row or col out of bounds")
	}

	//keyStr := toKeyStr(r, c)
	// Note: This is a zero-fill matrix
	// If the key dne, go returns the zero-value for that type
	// For float64, the zero value is 0.0
	//return m.matrix[keyStr], nil
	return m.matrix[Key{r, c}], nil
}

func (m *MatrixMapFloat64) Set(r, c int, v float64) error {
	if r < 0 || r > m.R || c < 0 || c > m.C {
		return errors.New("row or col out of bounds")
	}

	//keyStr := toKeyStr(r, c)
	//m.matrix[keyStr] = v
	m.matrix[Key{r, c}] = v
	return nil
}

func (m *MatrixMapFloat64) MultiplyConstant(c float64) {
	for k, v := range m.matrix {
		m.matrix[k] = c * v
	}
}

func (m *MatrixMapFloat64) Multiply(m2 *MatrixMapFloat64) {
	var sum float64
	for k, v := range m.matrix {
		sum = 0
		for k2, v2 := range m2.matrix {
			if k.Col == k2.Row {
				sum += v * v2
			}
		}
		fmt.Println(sum)
	}
}
