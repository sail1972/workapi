package workapi

import (
	"bytes"

	"github.com/json-iterator/go"
)

// DataToReader 数据转换
func DataToReader(data interface{}) (*bytes.Reader, error) {
	j, err := jsoniter.Marshal(&data)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(j), nil
}
