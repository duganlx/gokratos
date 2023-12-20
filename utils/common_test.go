package utils_test

import (
	"gokratos/utils"
	"testing"
)

func Test_GetRandom(t *testing.T) {
	secret := utils.GetRandom(64)

	t.Logf("secret(length: %d): %v", len(secret), secret)
}
