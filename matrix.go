package matrix

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type MappingFunc func(x float64) float64

type Mat struct {
	rows, cols int
	data       []float64
}

func New(rows int, cols int) *Mat {
	rand.Seed(time.Now().UnixNano())
	return &Mat{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}
}

func FromArray(data []float64, columnVector bool) *Mat {
	var mat *Mat
	if columnVector {
		mat = New(len(data), 1)
		mat.SetCol(1, data)
	} else {
		mat = New(1, len(data))
		mat.SetRow(1, data)
	}

	return mat
}

func (m *Mat) String() string {
	var b bytes.Buffer
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			b.WriteString(fmt.Sprintf("%f ", m.data[m.index(i+1, j+1)]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (m *Mat) Size() int {
	return m.rows * m.cols
}

func (m *Mat) Shape() (int, int) {
	return m.rows, m.cols
}

func (m *Mat) index(row int, col int) int {
	return (row-1)*m.cols + col - 1
}

func (m *Mat) Randomize(min float64, max float64) {
	high := max - min
	for i := 0; i < m.Size(); i++ {
		m.data[i] = min + rand.Float64()*high
	}
}

func (m *Mat) Set(row int, col int, value float64) {
	m.data[m.index(row, col)] = value
}

func (m *Mat) SetRow(row int, data []float64) {
	startingIndex := m.index(row, 1)
	j := 0
	for i := startingIndex; i < startingIndex+m.cols; i++ {
		m.data[i] = data[j]
		j++
	}
}

func (m *Mat) SetCol(col int, data []float64) {
	j := 0
	for i := 1; i < m.rows+1; i++ {
		m.data[m.index(i, col)] = data[j]
		j++
	}
}

func (m *Mat) Get(row int, col int) float64 {
	return m.data[m.index(row, col)]
}

func (m *Mat) GetRow(row int) []float64 {
	result := make([]float64, m.cols)
	startingIndex := m.index(row, 1)
	j := 0
	for i := startingIndex; i < startingIndex+m.cols; i++ {
		result[j] = m.data[i]
		j++
	}

	return result
}

func (m *Mat) GetCol(col int) []float64 {
	result := make([]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		result[i] = m.data[m.index(i+1, col)]
	}
	return result
}

func (m *Mat) Add(m2 *Mat) *Mat {
	if !checkSize(m, m2, false) {
		panic("Rows and columns must match.")
	}
	for i := 0; i < m.Size(); i++ {
		m.data[i] = m.data[i] + m2.data[i]
	}
	return m
}

func (m *Mat) Subtract(m2 *Mat) *Mat {
	if !checkSize(m, m2, false) {
		panic("Rows and columns must match.")
	}
	for i := 0; i < m.Size(); i++ {
		m.data[i] = m.data[i] - m2.data[i]
	}
	return m
}

func (m *Mat) Scalar(scalar float64) *Mat {
	for i := 0; i < m.Size(); i++ {
		m.data[i] = m.data[i] * scalar
	}
	return m
}

func (m *Mat) Hadamard(m2 *Mat) *Mat {
	if !checkSize(m, m2, false) {
		panic("Rows and columns must match.")
	}
	for i := 0; i < m.Size(); i++ {
		m.data[i] = m.data[i] * m2.data[i]
	}
	return m
}

func (m *Mat) Map(fn MappingFunc) *Mat {
	for i := 0; i < m.Size(); i++ {
		m.data[i] = fn(m.data[i])
	}
	return m
}

func (m *Mat) Multiply(m2 *Mat) *Mat {
	if !checkSize(m, m2, true) {
		panic("The columns of the first matrix must match the rows of the second matrix.")
	}
	r := m.rows
	c := m2.cols
	data := make([]float64, r*c)

	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			var value float64
			for k := 1; k <= m.cols; k++ {
				value += m.data[m.index(i, k)] * m2.data[m2.index(k, j)]
			}
			data[(i-1)*m2.cols+j-1] = value
		}
	}
	m.cols = c
	m.data = data
	return m
}

// Static

func Add(m1 *Mat, m2 *Mat) *Mat {
	if !checkSize(m1, m2, false) {
		panic("Rows and columns must match.")
	}

	r := New(m1.rows, m1.cols)

	for i := 0; i < r.Size(); i++ {
		r.data[i] = m1.data[i] + m2.data[i]
	}

	return r
}

func Subtract(m1 *Mat, m2 *Mat) *Mat {
	if !checkSize(m1, m2, false) {
		panic("Rows and columns must match.")
	}

	r := New(m1.rows, m1.cols)

	for i := 0; i < r.Size(); i++ {
		r.data[i] = m1.data[i] - m2.data[i]
	}

	return r
}

func Hadamard(m1 *Mat, m2 *Mat) *Mat {
	if !checkSize(m1, m2, false) {
		panic("Rows and columns must match.")
	}

	r := New(m1.rows, m1.cols)

	for i := 0; i < r.Size(); i++ {
		r.data[i] = m1.data[i] * m2.data[i]
	}
	return r
}

func Multiply(m1 *Mat, m2 *Mat) *Mat {
	if !checkSize(m1, m2, true) {
		panic("The columns of the first matrix must match the rows of the second matrix.")
	}
	r := New(m1.rows, m2.cols)
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			var value float64
			for k := 1; k <= m1.cols; k++ {
				value += m1.data[m1.index(i, k)] * m2.data[m2.index(k, j)]
			}
			r.data[r.index(i, j)] = value
		}
	}
	return r
}

func Map(m1 *Mat, fn MappingFunc) *Mat {
	r := New(m1.rows, m1.cols)

	for i := 0; i < r.Size(); i++ {
		r.data[i] = fn(m1.data[i])
	}

	return r
}

func Transpose(m1 *Mat) *Mat {
	r := New(m1.cols, m1.rows)
	for i := 1; i <= m1.cols; i++ {
		r.SetRow(i, m1.GetCol(i))
	}
	return r
}

func checkSize(m1 *Mat, m2 *Mat, mult bool) bool {
	if mult {
		return m1.cols == m2.rows
	}
	return m1.rows == m2.rows && m1.cols == m2.cols
}
