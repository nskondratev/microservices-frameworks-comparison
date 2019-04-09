package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"sync"
)

type Hasher struct {
	hf hash.Hash
	mu *sync.Mutex
}

func New () *Hasher {
	return &Hasher{
		hf: sha256.New(),
		mu: &sync.Mutex{},
	}
}

func (h *Hasher) Sum(input *string) (string, error) {
	h.mu.Lock()
	h.hf.Reset()
	if _, err := h.hf.Write([]byte(*input)); err != nil {
		return "", err
	}
	res := hex.EncodeToString(h.hf.Sum(nil))
	h.mu.Unlock()
	return res, nil
}
