package utils_test

import (
	"fmt"
	"gokratos/utils"
	"testing"
	"time"
)

func Test_GetRandom(t *testing.T) {
	secret := utils.GetRandom(64)

	t.Logf("secret(length: %d): %v", len(secret), secret)
}

func TestTime(t *testing.T) {
	dt := time.Unix(1704259198000/1000, 0)

	fmt.Println(dt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"))
}
