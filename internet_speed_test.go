package internet_speed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkConnection_Ookla(t *testing.T) {
	_, _, err := GetInternetSpeed("Ookla")
	assert.Nil(t, err)
}

func Test_checkConnection_Fast(t *testing.T) {
	_, _, err := GetInternetSpeed("fastStruct")
	assert.Nil(t, err)
}

func Test_checkConnection_WrongProvider(t *testing.T) {
	_, _, err := GetInternetSpeed("Google")
	assert.NotNil(t, err)
}

func BenchmarkConnectionOokla(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := GetInternetSpeed("Ookla")
		assert.Nil(b, err)
	}
}
func BenchmarkConnectionFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := GetInternetSpeed("fastStruct")
		assert.Nil(b, err)
	}
}
