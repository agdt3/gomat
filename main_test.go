package main

import (
	"fmt"
	"testing"
)

/*
func TestConstructor(t *testing.T) {
	m := NewMatrixMapFloat64(1000000, 1000000)
}
*/

func TestGet(t *testing.T) {
	m := NewMatrixMapFloat64(1000000, 1000000)
	m.Set(10000, 50000, 5.0)
	r1, _ := m.Get(10000, 50000)
	r2, _ := m.Get(10001, 50000)
	fmt.Println(r1)
	fmt.Println(r2)
}

func TestConstMultplier(t *testing.T) {
	m := NewMatrixMapFloat64(1000000, 1000000)
	m.Set(10000, 50000, 5.0)
	m.MultiplyConstant(3.0)
	r, _ := m.Get(10000, 50000)
	fmt.Println(r)
}

func TestMultplier(t *testing.T) {
	m := NewMatrixMapFloat64(1000000, 1000000)
	m2 := NewMatrixMapFloat64(1000000, 1000000)
	m.Set(10000, 50000, 5.0)
	m2.Set(50000, 30000, 2.0)
	m.Multiply(m2)
	fmt.Println(m)
}
