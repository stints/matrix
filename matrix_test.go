package matrix

import (
	"testing"
)

func TestMatrixCreation(t *testing.T) {
	m1 := New(2, 3)
	if rows, cols := m1.Shape(); m1.Size() != 6 || rows != 2 || cols != 3 {
		t.Errorf("Matrix creation via New failed.")
	}

	data := []float64{1.0, 2.0, 3.0}
	m2 := FromArray(data, true)
	if rows, cols := m2.Shape(); m2.Size() != 3 || rows != 3 || cols != 1 {
		t.Errorf("Matrix creation via FromArray failed.")
	}
}

func TestGetterSetter(t *testing.T) {
	m1 := New(3, 3)
	m1.Set(1, 1, 1.0)
	m1.Set(1, 2, 2.0)
	m1.Set(1, 3, 3.0)

	row2 := []float64{4.0, 5.0, 6.0}
	row3 := []float64{7.0, 8.0, 9.0}
	m1.SetRow(2, row2)
	m1.SetRow(3, row3)

	n := m1.Get(2, 3)
	r := m1.GetRow(1)
	c := m1.GetCol(2)
	if n != 6.0 {
		t.Errorf("Error in get.  Expected %f, Got %f", 6.0, n)
	}
	if !equal(r, []float64{1.0, 2.0, 3.0}) {
		t.Errorf("Error in get row.")
	}
	if !equal(c, []float64{2.0, 5.0, 8.0}) {
		t.Errorf("Errorin get col.")
	}
}

func TestAdd(t *testing.T) {
	m1, m2 := createMatrix()

	m1.Add(m2)

	if m1.Get(1, 1) != -4.0 {
		t.Errorf("Error with adding")
	}
}

func TestSubtract(t *testing.T) {
	m1, m2 := createMatrix()

	m1.Subtract(m2)

	if m1.Get(2, 1) != 2.5 {
		t.Error("Error with subtraction")
	}
}

func TestTranspose(t *testing.T) {
	m1, _ := createMatrix()

	m2 := m1.Transpose()

	if m2.Get(2, 1) != 2.0 {
		t.Errorf("Error with transpose")
	}
}

func TestHaramard(t *testing.T) {
	m1, m2 := createMatrix()

	m1.Hadamard(m2)

	if m1.Get(2, 2) != -4.0 {
		t.Errorf("Error with hadamard")
	}
}

func TestMultiply(t *testing.T) {
	m1, _ := createMatrix()

	m2 := New(2, 1)
	m2.SetCol(1, []float64{-5.0, -6.0})

	m3 := m1.Multiply(m2)

	if m3.Get(1, 1) != -17.0 || m3.Get(2, 1) != -39.0 {
		t.Errorf("Error with multiply")
	}
}

func createMatrix() (*Mat, *Mat) {
	m1 := New(2, 2)
	m2 := New(2, 2)

	m1.SetRow(1, []float64{1.0, 2.0})
	m1.SetRow(2, []float64{3.0, 4.0})
	m2.SetRow(1, []float64{-5.0, -2.0})
	m2.SetRow(2, []float64{0.5, -1})

	return m1, m2
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
