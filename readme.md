# Matrix

##### Creating Matrices
```Go
m1 := matrix.New(3, 4)  // 3x4
m2 := matrix.New(1, 2)  // 1x2

data := []float{1.0, 2.0, 3.0}
m3 := matrix.FromArray(data, true)  // 3x1
```

##### Matrix Math
```Go
// Add
m1.Add(m2)  // or m4 := matrix.Add(m1, m2)

// Subtract
m1.Subtract(m2)  // or m4 := matrix.Subtract(m1, m2)

// Multiply
m1.Multiply(m2)  // or m4 := matrix.Multiply(m1, m2)

// Dot Product
m1.Dot(m2)

// Scalar Multiplication
m1.Scalar(3.0)

// Transpose matrix
m4 := m1.Transpose()
```

##### Helper functions
```Go
// Map
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-1.0*x))
}

m1.Map(Sigmoid)  // Each value is mapped to passed function

// Getters & Setters
// All matrix rows and columns begin at 1
n := m1.Get(1,2)  // Get value at row 1, column 2
r := m1.GetRow(1)  // returns a float64 slice
c := m1.GetCol(1)  // returns a float64 slice

m1.Set(1, 2, 5.0)  // Set value 5.0 at row 1, column 2

r := []float64{2.0, 3.0, 4.0, 5.0}
m1.SetRow(1, r)  // length of data must match length of matrix row

c := []float64{3.0, 4.0, 5.0}
m1.SetCol(1, c)  // length of data must match length of matrix column

// Randomize values in matrix
m1.Randomize(-10.0, 10.0)  // values will be between -10 and 10
```