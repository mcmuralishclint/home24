package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Return nil Title if empty string is passed in link
func TestGetTitleEmptyLink(t *testing.T) {
	title, err := GetTitle("")
	assert.Nil(t, err)
	assert.Equal(t, title, "")
}

// Return nil if correct link is passed
func TestGetTitle(t *testing.T) {
	title, err := GetTitle("https://www.facebook.com/")
	assert.Nil(t, err)
	assert.NotEqual(t, title, "")
	assert.NotNil(t, title, "")
}

// Return nil Title if empty string is passed in link
func TestGetHeadingEmptyLink(t *testing.T) {
	headings, err := GetHeadings("")
	assert.Nil(t, err)
	assert.Equal(t, headings, make(map[string]int))
}

// // Return nil if correct link is passed
func TestGetHeading(t *testing.T) {
	headings, err := GetHeadings("https://www.facebook.com/")
	assert.Nil(t, err)
	assert.NotNil(t, headings)
}

// Return nil Links if empty string is passed in link
func TestGetLinksEmptyLink(t *testing.T) {
	err := GetLinks("")
	assert.Nil(t, err)
}

// // Return nil if correct link is passed
func TestGetLinks(t *testing.T) {
	err := GetLinks("https://www.facebook.com/")
	assert.Nil(t, err)
}
