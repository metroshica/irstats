package irstats_test

import (
	"testing"

	"github.com/metroshica/irstats"
	"github.com/stretchr/testify/assert"
)

func TestWithUserAgent(t *testing.T) {
	c := &irstats.Client{}
	err := irstats.WithUserAgent("My User Agent")(c)
	assert.NoError(t, err)
	assert.Equal(t, "My User Agent", c.GetUserAgent())
}
func TestGetUserAgent(t *testing.T) {
	c := &irstats.Client{}
	err := irstats.WithUserAgent("My User Agent")(c)
	assert.NoError(t, err)
	assert.Equal(t, "My User Agent", c.GetUserAgent())
}