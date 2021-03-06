package bnet_test

import (
	"github.com/samlitowitz/bnet-encoding/pkg/encoding/bnet"
	"testing"

	"github.com/go-test/deep"
)

func TestMarshalSupportedTypes(t *testing.T) {
	out, err := bnet.Marshal(supportedTypesTests.decoded)
	if err != nil {
		panic(err)
	}

	if diff := deep.Equal(supportedTypesTests.encoded, out); diff != nil {
		t.Error(diff)
	}
}
