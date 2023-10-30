package gmac

import (
	"crypto/cipher"
)

// CalculateGMAC calculates the Galois Message Authentication Code (GMAC)
func CalculateGMAC(block cipher.Block, nonce, data []byte) []byte {
	mode := cipher.NewGCM(block)
	gmac := mode.Seal(nil, nonce, data, nil)
	return gmac
}