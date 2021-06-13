package util

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/tianmai777/blog/global"
)

func EncodeMD5(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		global.Log.Panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
