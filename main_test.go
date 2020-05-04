package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCred(t *testing.T) {
	var a string = "abc123"
	out, err := getCred("app.terraform.i")
	if assert.NotNil(t, err) {
		assert.Equal(t, a, out, "The secrets should match")
	}
}
