package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func SignMessage(priv *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	hashed := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, priv, hashed[:])
	if err != nil {
		return nil, err
	}

	return asn1.Marshal(ecdsaSignature{r, s})
}

func VerifyMessage(pub *ecdsa.PublicKey, message []byte, signature []byte) bool {
	var rs ecdsaSignature

	if _, err := asn1.Unmarshal(signature, &rs); err != nil {
		return false
	}

	hashed := sha256.Sum256(message)
	return ecdsa.Verify(pub, hashed[:], rs.R, rs.S)
}
