package hashkey

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func MakeHashKey(str string, date time.Time) string {
	k := fmt.Sprintf("%s%s", str, date.Format(time.RFC3339))
	hasher := md5.New()
	hasher.Write([]byte(k))
	return hex.EncodeToString(hasher.Sum(nil))
}
