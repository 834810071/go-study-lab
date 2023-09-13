package main

import (
	"testing"
)

func TestNewWorker(t *testing.T) {
	ewer := NewWorker(&EnvironmentalSanitationWorker{})
	ewer.Daily()
	programmer := NewWorker(&Programmer{})
	programmer.Daily()
}
