package genhash

import (
	"time"

	hashids "github.com/speps/go-hashids"
)

// Generate a random sequence
func Generate() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	ts, _ := h.Encode([]int{int(now.Unix())})

	return ts
}
