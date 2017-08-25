package dvm

const DIMENSION = 2

type Point struct {
    Left float64
    Top  float64
}

func (p Point) Multiply(multiplier float64) {
    p.Left *= multiplier
    p.Top *= multiplier
}

func (p Point) ToArray() []float64 {
    return []float64{p.Left, p.Top}
}
