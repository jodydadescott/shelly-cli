package types

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func getIntFromKey(input string) (int, error) {

	inputSlice := strings.Split(input, ":")
	if len(inputSlice) != 2 {
		return 0, fmt.Errorf("Input string %s is not valid, missing ':'", input)
	}

	key, err := strconv.Atoi(inputSlice[1])
	if err != nil {
		return 0, fmt.Errorf("Input string %s is not valid, key is not integer", input)
	}

	return key, nil
}

func getSHA256(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}

func getCnonce() string {
	b := make([]byte, 8)
	io.ReadFull(rand.Reader, b)
	return fmt.Sprintf("%x", b)[:16]
}
