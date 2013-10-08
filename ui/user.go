package ui

import (
	"fmt"
)

type User struct {
	Ided
   Worker
   Name string
}


func (this User) Work() {
	fmt.Printf("doing work \n")
}
	
