package utils_test

import (
	"strings"
	"testing"

	"github.com/TeaOSLab/EdgeAdmin/internal/utils"
)

func TestStringsStream_Filter(t *testing.T) {
	var stream = utils.NewStringsStream([]string{"a", "b", "1", "2", "", "png", "a"})
	stream.Filter(func(item string) bool {
		return len(item) > 0
	})
	t.Log(stream.Result())
	stream.Map(func(item string) string {
		return "." + item
	})
	t.Log(stream.Result())
	stream.Unique()
	t.Log(stream.Result())
	stream.Map(strings.ToUpper, strings.ToLower)
	t.Log(stream.Result())
}
