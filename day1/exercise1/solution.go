package main

import (
	"fmt"
)

type Matrix struct{
	num_rows int
	num_cols int
	array [10][10] int
}

func (m Matrix) get_num_cols() int{
	return m.num_cols
}

func (m Matrix) get_num_rows() int{
	return m.num_rows
}

func (m *Matrix) set( row , col , value int) {
	m.array[row][col] = value
	return
}
func (m1 Matrix) addMatrices(m2 Matrix) Matrix{
	var sum = Matrix{}
	rows := m1.num_rows
	cols := m1.num_cols
	for row := 0 ; row < rows ; row++{
		for col := 0 ; col < cols ; col ++ {
			sum.array[row][col] = m1.array[row][col] + m2.array[row][col]
		}
	}

	return sum
}

func main(){
	//var x Matrix [2][3] int
	//c := x.get_num_cols()
	var rows,cols int
	var mat1 = Matrix{}
	var mat2 = Matrix{}
	var mat3 = Matrix{}
	fmt.Println("Enter number of rows :")
	fmt.Scanf("%d",&rows)
	fmt.Println("Enter number of columns : ")
	fmt.Scanf("%d",&cols)
	mat1.num_rows = rows
	mat1.num_cols = cols

	mat2.num_rows = rows
	mat2.num_cols = cols
	mat1.set(2,3,999)
	mat3 = mat1.addMatrices(mat2)
	fmt.Println(mat3.array)
}

