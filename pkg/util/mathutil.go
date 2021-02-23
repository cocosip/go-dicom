package util

import "math"

func IsByteOdd(v byte) bool {
	return (v & 1) == 1
}

func IsInt16Odd(v int16) bool {
	return (v & 1) == 1
}

func IsUInt16Odd(v uint16) bool {
	return (v & 1) == 1
}

func IsIntOdd(v int) bool {
	return (v & 1) == 1
}

func IsUIntOdd(v uint) bool {
	return (v & 1) == 1
}

func IsByteEven(v byte) bool {
	return (v & 1) == 0
}

func IsInt16Even(v int16) bool {
	return (v & 1) == 0
}

func IsUInt16Even(v uint16) bool {
	return (v & 1) == 0
}

func IsIntEven(v int) bool {
	return (v & 1) == 0
}

func IsUIntEven(v uint) bool {
	return (v & 1) == 0
}

func IsNearlyZero(v float64) bool {
	return math.Abs(v) < 0.000000001
}





