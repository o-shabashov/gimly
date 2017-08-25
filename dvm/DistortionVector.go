package dvm

type DistortionVector struct {
    Start Point
    End   Point
}

func (d DistortionVector) ToArray() []float64 {
    return append(d.Start.ToArray(), d.End.ToArray()...)
}

func (d DistortionVector) Multiply(multiplier float64) {
    d.Start.Multiply(multiplier)
    d.End.Multiply(multiplier)
}
