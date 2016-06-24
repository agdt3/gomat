package main

import "errors"

type Key struct {
	Row int
	Col int
}

type MatrixMapFloat64 struct {
	R      int
	C      int
	matrix map[Key]float64
}

func NewMatrixMapFloat64(rn, cn int) *MatrixMapFloat64 {
	m := new(MatrixMapFloat64)
	m.R = rn
	m.C = cn
	m.matrix = make(map[Key]float64)
	return m
}

func (m *MatrixMapFloat64) Get(r, c int) (float64, error) {
	if r < 1 || r > m.R || c < 1 || c > m.C {
		return 0, errors.New("row or col out of bounds")
	}

	// Note: This is a zero-fill matrix
	// If the key DNE, go returns the zero-value for that type
	// For float64, the zero value is conveniently 0.0
	return m.matrix[Key{r, c}], nil
}

func (m *MatrixMapFloat64) Set(r, c int, v float64) error {
	if r < 1 || r > m.R || c < 1 || c > m.C {
		return errors.New("row or col out of bounds")
	}

	m.matrix[Key{r, c}] = v
	return nil
}

func (m *MatrixMapFloat64) Add(m2 *MatrixMapFloat64) error {
	if m.R != m2.R || m.C != m2.C {
		return errors.New("m1 dimensions must match m2 dimensions")
	}

	for k, v := range m2.matrix {
		v2, _ := m.Get(k.Row, k.Col)
		if v2 == 0 {
			m.Set(k.Row, k.Col, v)
		} else {
			m.Set(k.Row, k.Col, v+v2)
		}
	}

	return nil
}

func (m *MatrixMapFloat64) MultiplyConstant(c float64) {
	for k, v := range m.matrix {
		m.matrix[k] = c * v
	}
}

func (m *MatrixMapFloat64) Multiply(m2 *MatrixMapFloat64) error {
	if m.C != m2.R {
		return errors.New("m col has to equal m2 row")
	}

	var sum float64
	for k, v := range m.matrix {
		sum = 0
		for k2, v2 := range m2.matrix {
			if k.Col == k2.Row {
				sum += v * v2
			}
		}
		m.Set(k.Row, k.Col, sum)
	}

	m.C = m2.C
	return nil
}

func (m *MatrixMapFloat64) Transpose() {
	// TODO: Implement in-place transpose if possible
	new_matrix := make(map[Key]float64)
	for k, v := range m.matrix {
		new_matrix[Key{k.Col, k.Row}] = v
	}

	temp := m.R
	m.R = m.C
	m.C = temp
	m.matrix = new_matrix
}
