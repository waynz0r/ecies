package ecies

import (
	"crypto/elliptic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurve(t *testing.T) {
	for _, test := range []struct {
		Name  string
		Curve elliptic.Curve
		Len   int
	}{
		{
			Name:  "P-256",
			Curve: elliptic.P256(),
			Len:   65,
		},
		{
			Name:  "P-384",
			Curve: elliptic.P384(),
			Len:   97,
		},
		{
			Name:  "P-521",
			Curve: elliptic.P521(),
			Len:   133,
		},
	} {
		assert.Equal(t, test.Name, test.Curve.Params().Name)
		assert.Equal(t, test.Len, GetECPointByteLength(test.Curve))
	}
}
