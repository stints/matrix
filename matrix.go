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
	return &Mat{
		rows: rows,
		cols: cols,
		data: make([]float64, rows*cols),
	}
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

func (m *Mat) size() int {
	return m.rows * m.cols
}

func (m *Mat) index(row int, col int) int {
	return (row-1)*m.cols + col - 1
}

func (m *Mat) Randomize(min float64, max float64) {
	rand.Seed(time.Now().UnixNano())
	high := max - min
	for i := 0; i < m.size(); i++ {
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

func (m *Mat) Add(m2 *Mat) {
	if checkSize(m, m2, false) {
		panic("Rows and columns must match.")
	}
	for i := 0; i < m.size(); i++ {
		m.data[i] = m.data[i] + m2.data[i]
	}
}

func (m *Mat) Subtract(m2 *Mat) {
	if checkSize(m, m2, false) {
		panic("Rows and columns must match.")
	}
	for i := 0; i < m.size(); i++ {
		m.data[i] = m.data[i] - m2.data[i]
	}
}

func (m *Mat) Scalar(scalar float64) {
	for i := 0; i < m.size(); i++ {
		m.data[i] = m.data[i] * scalar
	}
}

func (m *Mat) Multiply(m2 *Mat) *Mat {
	if checkSize(m, m2, true) {
		panic("The columns of the first matrix must match the rows of the second matrix.")
	}
	r := New(m.rows, m2.cols)
	for i := 1; i <= r.rows; i++ {
		for j := 1; j <= r.cols; j++ {
			var value float64
			for k := 1; k <= m.cols; k++ {
				value += m.data[m.index(i, k)] * m2.data[m2.index(k, j)]
			}
			r.data[r.index(i, j)] = value
		}
	}
	return r
}

func (m *Mat) Transpose() *Mat {
	r := New(m.cols, m.rows)
	for i := 1; i <= m.cols; i++ {
		r.SetRow(i, m.GetCol(i))
	}
	return r
}

func (m *Mat) Map(fn MappingFunc) {
	for i := 0; i < m.size(); i++ {
		m.data[i] = fn(m.data[i])
	}
}

// Static

func Subtract(m1 *Mat, m2 *Mat) *Mat {
	if checkSize(m1, m2, false) {
		panic("Rows and columns must match.")
	}

	r := New(m1.rows, m1.cols)

	for i := 0; i < r.size(); i++ {
		r.data[i] = m1.data[i] - m2.data[i]
	}

	return r
}

func Map(m1 *Mat, fn MappingFunc) *Mat{
	r := New(m1.rows, m1.cols)

	for i := 0; i < r.size(); i++ {
		r.data[i] = fn(m1.data[i])
	}

	return r
}

func checkSize(m1 *Mat, m2 *Mat, mult bool) bool {
	if mult {
		return m1.cols == m2.rows
	}
	return m1.rows == m2.rows && m1.cols == m2.cols
}
