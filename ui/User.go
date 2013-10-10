package ui

import (
	"fmt"
)

type User struct {
	Ided
	Named
   Worker
}


func (this User) Work() {
	fmt.Printf("doing work \n")
}
	
