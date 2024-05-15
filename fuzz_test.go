//go:build go1.18
// +build go1.18

package ngap_test

import (
	"testing"

	"github.com/machi12/ngap"
)

func FuzzNGAP(f *testing.F) {
	f.Fuzz(func(t *testing.T, d []byte) {
		//nolint:errcheck // fuzzing code
		ngap.Decoder(d)
	})
}
