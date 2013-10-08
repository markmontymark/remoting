package ui

import (
)

type Rater interface {
	Rate() int
}
type Rating struct {
	Ided
	Rating int
}

/*
func (this *Rating) SetRating(r int) {
	if r < 0 {
		this.Rating = 0
	} else if r > 5 {
		this.Rating = 5
	}
	this.Rating = r
}
*/
