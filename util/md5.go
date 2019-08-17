package util

import "crypto/md5"

func EncodeMd5(value string) string {
	pswMd5 := md5.New()
	pswMd5.Write([]byte(value))
	return string(pswMd5.Sum(nil))
}