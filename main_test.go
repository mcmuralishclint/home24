package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Return nil Title if empty string is passed in link
func TestGetTitleEmptyString(t *testing.T) {
	err := GetHeadings("")
	assert.Nil(t, err)
}

// Return nil if correct link is passed
func TestGetTitle(t *testing.T) {
	err := GetHeadings("https://www.facebook.com/")
	assert.Nil(t, err)
}
