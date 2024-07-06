package irstats_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/metroshica/irstats"
	"github.com/metroshica/irstats/ref"
	"github.com/stretchr/testify/assert"
)

func TestCarsDriven(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(irstats.URLPathCarsDriven, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[ 1, 10, 20 ]`)
	})

	result, _, err := client.CarsDriven(ref.String("2000"))
	assert.NoError(t, err)

	expected := &irstats.CarsDriven{1, 10, 20}
	assert.Equal(t, expected, result)
}
