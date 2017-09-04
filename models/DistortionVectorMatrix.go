package models

const DIMENSION = 2

type DistortionVectorMatrix struct {
    VectorMatrix [][]DistortionVector
}

func (d *DistortionVectorMatrix) SetFromDistortionMatrix(distortionMatrix []float64) {
    // Разбиваем плоский массив на куски по 2*DIMENSION
    chunksMatrix, err := ArrayChunk(distortionMatrix, 2*DIMENSION)
    if err != nil {
        panic(err)
    }

    vectors := []DistortionVector{}

    // Берём первые четыре элемента каждого chunksMatrix, создаём новую структуру DistortionVector и укладываем в массив
    for _, chunk := range chunksMatrix {
        vectors = append(vectors, DistortionVector{
            Start: Point{Left: chunk[0], Top: chunk[1]},
            End:   Point{Left: chunk[2], Top: chunk[3]},
        })
    }

    prevVectorStartTop := vectors[0].Start.Top

    // Группировка векторов по координате Start.Top - каждый такой вектор укладывается в свой массив
    for _, vector := range vectors {
        index := len(d.VectorMatrix) - 1
        if index < 0 {
            index = 0
        }

        if len(d.VectorMatrix) == 0 {
            d.VectorMatrix = append(d.VectorMatrix, []DistortionVector{})
        }

        if vector.Start.Top == prevVectorStartTop {
            d.VectorMatrix[index] = append(d.VectorMatrix[index], []DistortionVector{vector}...)
        } else {
            d.VectorMatrix = append(d.VectorMatrix, []DistortionVector{})
            d.VectorMatrix[index+1] = append(d.VectorMatrix[index+1], []DistortionVector{vector}...)
        }

        prevVectorStartTop = vector.Start.Top
    }
}

func (d DistortionVectorMatrix) GetDistortionMatrix() (numbers []float64) {
    for _, row := range d.VectorMatrix {
        for _, vector := range row {
            numbers = append(numbers, vector.ToArray()...)
        }
    }

    return
}

func (d *DistortionVectorMatrix) Multiply(multiplier float64) {
    // Потому что в Go нельзя в цикле брать ссылку на оригинальные данные, как, например, в PHP &$variable
    vm := [][]DistortionVector{}
    vr := []DistortionVector{}

    for _, vectorRow := range d.VectorMatrix {
        for _, vector := range vectorRow {
            vector.Multiply(multiplier)
            vr = append(vr, vector)
        }
        vm = append(vm, vr)
    }
    d.VectorMatrix = vm
}

// Несколько геттеров, чтобы не дублировать код
func (d DistortionVectorMatrix) GetLeft() float64 {
    return d.GetFirstPoint().Start.Left
}
func (d DistortionVectorMatrix) GetTop() float64 {
    return d.GetFirstPoint().Start.Top
}
func (d DistortionVectorMatrix) GetWidth() float64 {
    return d.GetLastPoint().Start.Left - d.GetLeft()
}
func (d DistortionVectorMatrix) GetHeight() float64 {
    return d.GetLastPoint().Start.Top - d.GetTop()
}
func (d DistortionVectorMatrix) GetFirstPoint() DistortionVector {
    return d.VectorMatrix[0][0]
}
func (d DistortionVectorMatrix) GetLastPoint() DistortionVector {
    lastRow := d.VectorMatrix[len(d.VectorMatrix)-1]
    return lastRow[len(lastRow)-1]
}

// Подразбиение исходной матрицы на меньшие матрицы, 2х2
func SplitMatrix(matrix [][]DistortionVector, rowSize int, columnSize int) (parts [][][]DistortionVector) {
    amountRows := len(matrix)
    amountColumns := len(matrix[0])

    if amountRows < 2 || amountColumns < 2 {
        panic("Matrix does not matches the selected row and column size")
    }

    for row := 0; row < len(matrix); row = row + (rowSize - 1) {
        for column := 0; column < len(matrix[row]); column = column + (columnSize - 1) {
            if amountRows-(rowSize-1) > row && amountColumns-(columnSize-1) > column {
                parts = append(parts, subMatrix(matrix, row, column, rowSize, columnSize))
            }
        }
    }
    return
}

// Вспомогательная функция, вынесена для упрощения чтения кода
func subMatrix(matrix [][]DistortionVector, startRow int, startColumn int, endRow int, endColumn int) ([][]DistortionVector) {
    tempMatrix := make([][]DistortionVector, len(matrix))
    copy(tempMatrix, matrix)
    subMatrix := tempMatrix[startRow:endRow]

    for row, rowItems := range subMatrix {
        subMatrix[row] = rowItems[startColumn:endColumn+startColumn]
    }

    return subMatrix
}

type Point struct {
    Left float64
    Top  float64
}

func (p *Point) Multiply(multiplier float64) {
    p.Left *= multiplier
    p.Top *= multiplier
}
func (p Point) ToArray() []float64 {
    return []float64{p.Left, p.Top}
}

type DistortionVector struct {
    Start Point
    End   Point
}

func (d DistortionVector) ToArray() []float64 {
    return append(d.Start.ToArray(), d.End.ToArray()...)
}
func (d *DistortionVector) Multiply(multiplier float64) {
    d.Start.Multiply(multiplier)
    d.End.Multiply(multiplier)
}
