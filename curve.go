package ecies

import "crypto/elliptic"

// GetECPointByteLength Get the length of bytes of a point on the curve
func GetECPointByteLength(curve elliptic.Curve) int {
	byteLen := (curve.Params().BitSize + 7) / 8
	return 1 + 2*byteLen
}
