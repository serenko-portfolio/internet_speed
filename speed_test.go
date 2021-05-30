package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkConnection_Ookla(t *testing.T) {
	_, _, err := CheckConnection("Ookla")
	assert.Nil(t, err)
}

func Test_checkConnection_Fast(t *testing.T) {
	_, _, err := CheckConnection("Fast")
	assert.Nil(t, err)
}

func Test_checkConnection_WrongProvider(t *testing.T) {
	_, _, err := CheckConnection("Google")
	assert.NotNil(t, err)
}
