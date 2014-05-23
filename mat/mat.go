/*
package mat implements basic matrix operations on real-valued matrices. The 
subpackage cmat/ implements the operations for complex-valued matrices. In the
interest of providing clean interfaces several non-trivial optimizations based
on parameter-spamming and algorithm selection have been ignored. These
clunkier, optimized interfaces can be found in the optmat/ and optcmat/
subpackages.
*/
package mat

// Matrix represents a two-dimensional rectangluar array of real values.
type Matrix struct { }

// MatrixError is the error type used by package mat. *MatrixError implements
// the error interface.
type MatrixError struct { }

func (err *MatrixError) Error() string {
	return ""
}

// type assertion
var _ error = &MatrixError{}

// Initialization functions

// New returns a matrix with the given dimensions where all elements are
// initialized to zero.
//
// If height or width is non-positive, an error will be returned.
func New(width, height int) (*Matrix, error) {
	return nil, nil
}

// Identity returns a square matrix with the given width which contains ones
// down its diagonal and zeroes everywhere else.
//
// If height or width is non-positive, an error will be returned.
func Identity(width int) (*Matrix, error) {
	return nil, nil
}

// FromArray converts an slice of floats to a matrix with the given dimensions.
// The element with zero-indexed coordinates (x, y) in the matrix will be the
// same as that at index values[y * width + x].
//
// If width * height != len(values) or if height or width is non-positive, an
// error will be returned.
func FromArray(width, height int, values []float64) (*Matrix, error) {
	return nil, nil
}

// FromGrid conversts a 2D slice of floats to a matrix with the same
// dimensions. The element with zero-indexed coordinates of (x, y) in the
// matrix will be the same as that at index values[y][x].
//
// If any two rows in data have different lengths, or if len(values) == 0 or
// len(values[0]) == 0, an error will be returned.
func FromGrid(values [][]float64) (*Matrix, error) {
	return nil, nil
}

// Utility functions

// Equal returns true if every element in the two given arrays is equal to
// within the library precision fraction (as defined in num/convergence.go) and
// false otherwise.
func Equal(m1, m2 *Matrix) bool {
	return false
}

// Compatible returns true if the two given matrices have the same shapes and
// false otherwise.
func Compatible(m1, m2 *Matrix) bool {
	return false
}

// MultCompatible returns true if the two given matrices can be multiplied
// together and false otherwise.
func MultCompatible(m1, m2 *Matrix) bool {
	return false
}

// Height returns the height of the matrix.
func (m *Matrix) Height() int {
	return -1
}

// Width returns the width of the matrix.
func (m *Matrix) Width() int {
	return -1
}

// Get returns the element of the matrix with coordinates (x, y).
//
// Get and Set are unique in that they panic upon erroneous input instead of
// returning an error. This is done to retain interface consistency with slice
// accessing.
func (m *Matrix) Get(x, y int) float64 {
	return 0.0
}

// Set changes the element in the matrix with coordinates (x, y) so that it
// has the given value. 
//
// Get and Set are unique in that they panic upon erroneous input instead of
// returning an error. This is done to retain interface consistency with slice
// accessing.
func (m *Matrix) Set(x, y int, val float64) {
	return
}

// Print prints the contents of the matrix as a comma-separated array of
// arrays. Each row in the matrix is given its own line.
func (m *Matrix) Print() {
	return
}

// Printf prints the contents of of the matrix as a comma-separated array of
// arrays where each element is formatted according to the given format string.
// Each row in the matrix is given its own line.
func (m *Matrix) Printf(fmt string) {
	return
}