package helper

import (
	"image-service/internal/config"
	"math/rand"
	"strings"
)

func GenerateRandString(resultLength int, int64Seeder int64) string {
	var output strings.Builder
	rand.Seed(int64Seeder)
	for i := 0; i < resultLength; i++ {
		c := rand.Intn(len(config.IMAGE_NAME_CHARSET))
		r := config.IMAGE_NAME_CHARSET[c]

		output.WriteString(string(r))
	}

	return output.String()
}
