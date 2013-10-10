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

