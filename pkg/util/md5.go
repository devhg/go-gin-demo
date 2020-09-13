package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(val string) string {
	hash := md5.New()
	hash.Write([]byte(val))
	return hex.EncodeToString(hash.Sum(nil))
}
