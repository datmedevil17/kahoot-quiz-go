package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GeneratePIN() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(100000 + rand.Intn(900000)) // 6-digit PIN
}
