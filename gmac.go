package gmac

import (
	"crypto/cipher"
)

// CalculateGMAC calculates the Galois Message Authentication Code (GMAC)
func CalculateGMAC(block cipher.Block, nonce, data []byte) ([]byte, error) {
	// Create an AEAD using the provided block
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Calculate the GMAC using the AEAD object
	gmac := aead.Seal(nil, nonce, nil, data)
	return gmac, nil
}
