package genhash

import (
	"time"

	"github.com/speps/go-hashids/v2"
)

// Generate a random sequence for URI
func Generate() string {
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	ts, _ := h.Encode([]int{int(time.Now().Unix())})

	return ts
}