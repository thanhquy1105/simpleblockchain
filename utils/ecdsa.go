package utils

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func NewSignature(R, S *big.Int) *Signature {
	return &Signature{R, S}
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
