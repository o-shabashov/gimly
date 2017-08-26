package dvm

import "errors"

type DistortionVectorMatrix struct {
    VectorMatrix [][]DistortionVector
}

func (d DistortionVectorMatrix) SetFromDistortionMatrix(distortionMatrix []float64) []DistortionVector {
    chunksMatrix, err := ArrayChunk(distortionMatrix, 2*DIMENSION)
    if err != nil {
        panic(err)
    }

    vectors := []DistortionVector{}

    for _, chunk := range chunksMatrix {
        vectors = append(vectors, DistortionVector{
            Start: Point{Left: chunk[0], Top: chunk[1]},
            End:   Point{Left: chunk[2], Top: chunk[3]},
        })
    }

    vectorMatrix := [][]DistortionVector{}

    // TODO нужен пример запроса
    //prevVectorStartTop := vectors[0].Start.Top
    //
    //for _, vector := range vectors {
    //    index := len(vectorMatrix) - 1
    //
    //    if vector.Start.Top == prevVectorStartTop {
    //        vectorMatrix[index] = append(vectorMatrix[index], []DistortionVector{vector}...)
    //    } else {
    //        vectorMatrix = append(vectorMatrix, []DistortionVector{})
    //        vectorMatrix[index + 1] = append(vectorMatrix[index + 1], []DistortionVector{vector}...)
    //    }
    //
    //    prevVectorStartTop = vector.Start.Top
    //}

    d.VectorMatrix = vectorMatrix

    return vectors // INT исправить в тестах, не должен возвращать вектора, проверять d.VectorMatrix
}

func (d DistortionVectorMatrix) GetDistortionMatrix() (numbers []float64) {
    for _, row := range d.VectorMatrix {
        for _, vector := range row {
            numbers = append(numbers, vector.ToArray()...)
        }
    }

    return
}

func (d DistortionVectorMatrix) Multiply(multiplier float64) {
    for _, vectorRow := range d.VectorMatrix {
        for _, vector := range vectorRow {
            vector.Multiply(multiplier)
        }
    }
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

func (d DistortionVectorMatrix) Clone(){
    for row, vectorRow := range d.VectorMatrix {
        for column, vector :=range vectorRow {
            d.VectorMatrix[row][column] = vector.Clone()
        }
    }
}

// Возвращает массив массивов, в каждом по chunkSize элементов.
func ArrayChunk(data []float64, chunkSize int) (result [][]float64, err error) {
    for i := 0; i < len(data); i += chunkSize {
        end := i + chunkSize

        // Если элементов не хватило на массив, то он будет создан частично.
        // Например, если элементов всего 5, а chunkSize = 3, то результат будет [[0,1,2] [3,4]]
        if end > len(data) {
            end = len(data)
        }

        if len(data[i:end]) < chunkSize {
            return result, errors.New("result array is smaller than ChunkSize")
        }

        result = append(result, data[i:end])
    }

    return
}

// INT TODO нужен пример запроса
func SplitMatrix(matrix [][]DistortionVector, rowSize int, columnSize int) (parts []float64) {
    amountRows := len(matrix)
    amountColumns := len(matrix[0])

    if amountRows < 2 || amountColumns < 2 {
        panic("Matrix does not matches the selected row and column size")
    }

    for row := 0; row < len(matrix); row = row + (rowSize - 1) {
        for column := 0; column < len(matrix[row]); column = column + (columnSize - 1) {
            if amountRows - (rowSize - 1) > row && amountColumns - (columnSize - 1) > column {
            parts = append(parts, SubMatrix(matrix, row, column, rowSize, columnSize)...)
            }
        }
    }
    return
}

// INT TODO нужен пример запроса
func SubMatrix(matrix [][]DistortionVector, startRow int, startColumn int, endRow int, endColumn int) (subMatrix []float64){
    subMatrix = matrix[startRow:endRow]

    for row, rowItems := range subMatrix {
     subMatrix[row] = rowItems[startColumn:endColumn]
    }

    return
}