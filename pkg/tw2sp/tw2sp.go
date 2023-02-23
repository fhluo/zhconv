package tw2sp

import (
	_ "embed"
	"github.com/bytedance/sonic"
	"github.com/fhluo/hanconv/pkg/hanconv"
)

var (
	//go:embed tw2sp.json
	data []byte
	conv hanconv.Converter
)

func init() {
	if err := sonic.Unmarshal(data, &conv); err != nil {
		panic(err)
	}
}

func Convert(data []byte) []byte {
	return conv.Convert(data)
}

func ConvertString(s string) string {
	return conv.ConvertString(s)
}