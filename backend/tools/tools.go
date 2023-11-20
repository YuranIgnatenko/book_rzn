package tools

import (
	"fmt"
	"math/rand"
	"time"
)

func NewRandomTokenFastOrder() string {
	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(999999999999)
	return fmt.Sprintf("%d", token)
}
