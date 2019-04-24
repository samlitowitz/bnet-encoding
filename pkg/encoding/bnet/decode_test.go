package bnet

import (
	"testing"

	"github.com/go-test/deep"
)

func TestUnmarshalSupportedTypes(t *testing.T) {
	var out supportedTypes

	err := Unmarshal(supportedTypesTests.encoded, &out)
	if err != nil {
		panic(err)
	}

	if diff := deep.Equal(supportedTypesTests.decoded, out); diff != nil {
		t.Error(diff)
	}
}
