package getDummyRows

import (
	"math/rand"
	"time"
)

const MAX_VARIANTS = 20
const VARIANT_CHANCE float32 = 0.5

func RandomlyHasVariants() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if r.Float32() >= VARIANT_CHANCE {
		return true
	}

	return false
}

func RandomVariantNum() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(MAX_VARIANTS) + 1
}
