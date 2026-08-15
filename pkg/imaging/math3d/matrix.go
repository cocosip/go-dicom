// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package math3d

import (
	"fmt"
	"math"
)

// Matrix4 is a row-major 4x4 homogeneous transformation matrix.
type Matrix4 [16]float64

// IdentityMatrix4 returns a 4x4 identity matrix.
func IdentityMatrix4() Matrix4 {
	return Matrix4{0: 1, 5: 1, 10: 1, 15: 1}
}

// Set assigns the value at row and column. Invalid indexes panic, matching
// ordinary Go array indexing.
func (m *Matrix4) Set(row, column int, value float64) {
	m[row*4+column] = value
}

// TransformPoint applies the homogeneous transform to point.
func (m Matrix4) TransformPoint(point Point3) Point3 {
	x := m[0]*point.X + m[1]*point.Y + m[2]*point.Z + m[3]
	y := m[4]*point.X + m[5]*point.Y + m[6]*point.Z + m[7]
	z := m[8]*point.X + m[9]*point.Y + m[10]*point.Z + m[11]
	w := m[12]*point.X + m[13]*point.Y + m[14]*point.Z + m[15]
	if math.Abs(w) > epsilon && math.Abs(w-1) > epsilon {
		x, y, z = x/w, y/w, z/w
	}
	return Point3{X: x, Y: y, Z: z}
}

// Inverse returns the matrix inverse using Gauss-Jordan elimination.
func (m Matrix4) Inverse() (Matrix4, error) {
	var augmented [4][8]float64
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			augmented[row][column] = m[row*4+column]
		}
		augmented[row][row+4] = 1
	}

	for column := 0; column < 4; column++ {
		pivot := column
		for row := column + 1; row < 4; row++ {
			if math.Abs(augmented[row][column]) > math.Abs(augmented[pivot][column]) {
				pivot = row
			}
		}
		if math.Abs(augmented[pivot][column]) <= epsilon {
			return Matrix4{}, fmt.Errorf("matrix is singular")
		}
		augmented[column], augmented[pivot] = augmented[pivot], augmented[column]
		divisor := augmented[column][column]
		for index := 0; index < 8; index++ {
			augmented[column][index] /= divisor
		}
		for row := 0; row < 4; row++ {
			if row == column {
				continue
			}
			factor := augmented[row][column]
			for index := 0; index < 8; index++ {
				augmented[row][index] -= factor * augmented[column][index]
			}
		}
	}

	var inverse Matrix4
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			inverse[row*4+column] = augmented[row][column+4]
		}
	}
	return inverse, nil
}
