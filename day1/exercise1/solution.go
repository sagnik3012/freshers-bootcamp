package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	NumRows int `json : "rows"`
	NumCols int `json : "cols"`
	Data [][] int `json :"data"`
}

func (m Matrix) getNumCols() int{
	return m.NumCols
}

func (m Matrix) getNumRows() int{
	return m.NumRows
}



func (m *Matrix) set( row , col , value int) {
	
  m.Data[row][col] = value
  return
}

func (m1 *Matrix) addMatrices(m2 Matrix) Matrix{

	rows := m1.NumRows
	cols := m1.NumCols
	for row := 0 ; row < rows ; row++{
		for col := 0 ; col < cols ; col ++ {
			m1.Data[row][col] = m1.Data[row][col] + m2.Data[row][col]
		}
	}
	return *m1
}

func matrixToJson(m Matrix) error {
	matrix , err := json.MarshalIndent(m,""," ")
	if err != nil {
		return err
	}
	fmt.Println(string(matrix))
	return nil
}

func main(){

	mat1 := Matrix{2,3,[][] int {{1,2,3},{4,5,6}},}
	mat2 := Matrix{2,3,[][] int {{4,5,6},{1,2,3}},}

	mat1.addMatrices(mat2)
	fmt.Println("no. of rows in mat1 = ",mat1.getNumRows())
	fmt.Println("no. of cols in mat2 = ",mat2.getNumCols())

	mat1.set(1,1,100)

	err := matrixToJson(mat1)
	if err != nil{
		fmt.Println("Error in converting matrix to json ",err)
	}
}

