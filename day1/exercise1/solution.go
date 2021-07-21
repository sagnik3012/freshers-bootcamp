package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	numRows int
	numCols int
	data [][] int
}

func (m Matrix) getNumCols() int{
	return m.numCols
}

func (m Matrix) getNumRows() int{
	return m.numRows
}

func (m *Matrix) set( row , col , value int) {
	m.data[row][col] = value
	return
}
func (m1 *Matrix) addMatrices(m2 Matrix) Matrix{

	rows := m1.numRows
	cols := m1.numCols
	for row := 0 ; row < rows ; row++{
		for col := 0 ; col < cols ; col ++ {
			m1.data[row][col] = m1.data[row][col] + m2.data[row][col]
		}
	}
	return *m1
}
func matrixToJson(m Matrix){
	matrix , err := json.MarshalIndent(m.data,""," ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(matrix))
}




func main(){

	mat1 := Matrix{2,3,[][] int {{1,2,3},{4,5,6}},}
	mat2 := Matrix{2,3,[][] int {{4,5,6},{1,2,3}},}
	mat1.addMatrices(mat2)
	fmt.Println("no. of rows in mat1 = ",mat1.getNumRows())
	fmt.Println("no. of cols in mat2 = ",mat2.getNumCols())
	mat1.set(1,1,100)
	matrixToJson(mat1)
}

