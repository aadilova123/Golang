package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestaddaddYearofStudy(t *testing.T){
	a := 0
	added := addYearofStudy(a)
	assert.Equal(t,1, added)
}