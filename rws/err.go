package rws

import "errors"

var (
	ErrEmptyMapItems = errors.New("weightedItems map cannot be empty")
	ErrEmptyItems    = errors.New("items slice cannot be empty")
	ErrInvalidWeight = errors.New("weights must be positive integers")
	ErrNullItems     = errors.New("no items available for selection")
)
