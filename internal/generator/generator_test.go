package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	builder := strings.Builder{}
	b := sha256.Sum256(nil)
	str := hex.EncodeToString(b[:])
	fmt.Printf("str: %v\n", str)
	for i := 0; i < len(str); i += 2 {
		end := i + 2
		if end > len(str) {
			end = len(str)
		}
		toParse := str[i:end]
		fmt.Printf("toParse: %v\n", toParse)
		num, _ := strconv.ParseInt(str[i:end], 16, 64)
		fmt.Printf("num: %v\n", num)
		cursor := int(num) % len(_CHARS)
		builder.WriteByte(_CHARS[cursor])
	}

	fmt.Printf("builder.String(): %v\n", builder.String())
}
