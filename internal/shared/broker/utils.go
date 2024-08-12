package broker_utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"math/rand"
)

func GetCookies() string {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	num := now % 1000
	num2 := (now/1000 - 1420070400) | (num&0x1FFFFFF)<<32 | 0x4200000000000000
	text := CreateHardId()[:17]
	num3 := now/1000 - 86400
	return fmt.Sprintf("_fz_uniq=%d;uniq=%d;age=%d;tid=%s;", num2, num2, num3, text)
}

func CreateHardId() string {
	now := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(now))
	buffer := make([]byte, 256)
	for i := 0; i < 256; i++ {
		buffer[i] = byte(rng.Uint32() >> 16 & 0xFF)
	}
	hash := md5.Sum(buffer)
	hash[0] = 0
	for i := 1; i < 16; i++ {
		hash[0] += hash[i]
	}
	return hex.EncodeToString(hash[:])
}

func GenerateSignature(company string, version string) []byte {

	data := fmt.Sprintf("company=%s&code=%s", company, version)

	// First hash
	hash1 := md5.Sum([]byte(data))

	// Second hash
	md5Instance := md5.New()
	md5Instance.Write(hash1[:16])

	array2 := []byte{
		61, 123, 21, 22, 214, 234, 187, 52,
		217, 214, 99, 227, 98, 62, 27, 215,
		251, 220, 174, 244, 87, 59, 223, 53,
		127, 168, 207, 11, 190, 173, 146, 127,
	}

	md5Instance.Write(array2)

	finalHash := md5Instance.Sum(nil)

	return finalHash
}
