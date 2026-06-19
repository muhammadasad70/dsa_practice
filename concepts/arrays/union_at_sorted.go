package main

import "fmt"

func union_at_sorted(){
	A:=[]int{3,4,5,6,10}
	B:=[]int{2,4,5,7,12}
	C:=[]int{}

	i:=0
	j:=0

	for i<len(A) && j<len(B){
		if A[i]<B[j]{
			C=append(C, A[i])
			i++
		}else if B[j]<A[i]{
			C=append(C, B[i])
			j++
		}else if A[i]==B[j]{
			C=append(C, A[i])
			i++
			j++
		}
	}
	for ;i<len(A);i++{
		C=append(C, A[i])
	}
	for ;j<len(B);j++{
		C=append(C, B[j])
	}
	fmt.Println("After union ")
	for _,k:=range C{
		fmt.Println(k)
	}
}