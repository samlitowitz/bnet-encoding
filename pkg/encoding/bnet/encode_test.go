package bnet

import (
	"testing"

	"github.com/go-test/deep"
)

func TestMarshalSupportedTypes(t *testing.T) {
	out, err := Marshal(supportedTypesTests.decoded)
	if err != nil {
		panic(err)
	}

	if diff := deep.Equal(supportedTypesTests.encoded, out); diff != nil {
		t.Error(diff)
	}
}
