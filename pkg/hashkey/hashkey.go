package hashkey

import (
	"crypto/md5"
	"encoding/hex"
)

func MakeHashKey(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
