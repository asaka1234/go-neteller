package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// BuildSignatureUtils provides signature generation functionality
type BuildSignatureUtils struct{}

// NewBuildSignatureUtils creates a new instance (though we'll mostly use static methods)
func NewBuildSignatureUtils() *BuildSignatureUtils {
	return &BuildSignatureUtils{}
}

// GetGtAuthenticationHeader generates authentication header
func GetGtAuthenticationHeader(request map[string]interface{}, requestSignatureList []string, merchantSecret string) string {
	concatenatedString := getConcatenatedString(request, requestSignatureList) + merchantSecret
	return generateSignature(concatenatedString)
}

func getConcatenatedString(data map[string]interface{}, requestSignatureList []string) string {
	var builder strings.Builder

	for _, key := range requestSignatureList {
		if value, exists := data[key]; exists && value != nil {
			builder.WriteString(fmt.Sprintf("%v", value))
		}
	}

	return builder.String()
}

func generateSignature(input string) string {
	hasher := sha512.New384()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
