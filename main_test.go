package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Return nil Title if empty string is passed in link
func TestGetTitleEmptyLink(t *testing.T) {
	err := GetTitle("")
	assert.Nil(t, err)
}

// Return nil if correct link is passed
func TestGetTitle(t *testing.T) {
	err := GetTitle("https://www.facebook.com/")
	assert.Nil(t, err)
}

// Return nil Title if empty string is passed in link
func TestGetHeadingEmptyLink(t *testing.T) {
	err := GetHeadings("")
	assert.Nil(t, err)
}

// Return nil if correct link is passed
func TestGetHeading(t *testing.T) {
	err := GetHeadings("https://www.facebook.com/")
	assert.Nil(t, err)
}
