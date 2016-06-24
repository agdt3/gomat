package main

import "testing"

func TestConstructor(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(5000000, 1000000)
	if m.R != 5000000 || m.C != 1000000 {
		t.Error("Get failed to return the correct values")
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(1000000, 1000000)
	m.Set(10000, 50000, 5.0)
	r1, _ := m.Get(10000, 50000)
	r2, _ := m.Get(10001, 50000)

	if r1 != 5.0 || r2 != 0 {
		t.Error("Get failed to return the correct values")
	}
}

func TestConstAdd(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(1000000, 1000000)
	m2 := NewMatrixMapFloat64(1000000, 1000000)

	m.Set(10000, 50000, 5.0)
	m.Set(10000, 50001, 12.0)
	m.Set(10001, 50000, 14.0)

	m2.Set(10000, 50000, 2.0)
	m2.Set(20000, 30000, 4.0)

	m.Add(m2)

	r1, _ := m.Get(10000, 50000)
	r2, _ := m.Get(10000, 50001)
	r3, _ := m.Get(10001, 50000)
	r4, _ := m.Get(20000, 30000)

	if r1 != 7.0 || r2 != 12.0 || r3 != 14.0 || r4 != 4.0 {
		t.Error("Incorrect matrix addition")
	}
}

func TestConstMultplier(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(1000000, 1000000)
	m.Set(10000, 50000, 5.0)
	m.MultiplyConstant(3.0)
	r, _ := m.Get(10000, 50000)

	if r != 15.0 {
		t.Error("Incorrect multiplication by const")
	}
}

func TestMultplierValues(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(1000000, 1000000)
	m2 := NewMatrixMapFloat64(1000000, 1000000)

	m.Set(10, 50000, 5.0)
	m.Set(10000, 50000, 7.0)
	m2.Set(50000, 30000, 2.0)
	m.Multiply(m2)

	r1, _ := m.Get(10, 50000)
	r2, _ := m.Get(10000, 50000)

	if r1 != 10.0 || r2 != 14.0 {
		t.Error("Incorrect multiplication of matricies")
	}
}

func TestMultplierDimension(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(3, 1000000)
	m2 := NewMatrixMapFloat64(1000000, 500)

	m.Multiply(m2)

	if m.R != 3 || m.C != 500 {
		t.Error("Incorrect matrix dimensions")
	}
}

func TestTranspose(t *testing.T) {
	t.Parallel()

	m := NewMatrixMapFloat64(4, 2)

	m.Set(1, 1, 1.0)
	m.Set(1, 2, 2.0)
	m.Set(2, 1, 3.0)
	m.Set(2, 2, 4.0)
	m.Set(3, 1, 5.0)
	m.Set(3, 2, 6.0)
	m.Set(4, 1, 7.0)
	m.Set(4, 2, 8.0)

	m.Transpose()

	r1, _ := m.Get(1, 1)
	r2, _ := m.Get(1, 2)
	r3, _ := m.Get(1, 3)
	r4, _ := m.Get(1, 4)

	if r1 != 1.0 || r2 != 3.0 || r3 != 5.0 || r4 != 7.0 {
		t.Error("Matrix transpose failed")
	}
}
