package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCred(t *testing.T) {
	var a string = "abc123"
	out, _ := getCred("app.terraform.io")
	assert.Equal(t, a, out, "The secrets should match")
}
