package codes

import (
	"fmt"
)

type errCode int

const (
	ErrFileSizeExceeded errCode = iota + 1000
	ErrFileParseFail
	ErrReadFileFail
	ErrFileConversionFail
	ErrFileStoreFail
	ErrFetchingFile
	ErrIdNotfound
	ErrKeyNotFound
)

var errCodes = map[errCode]string{
	ErrFileSizeExceeded:   "file size exceeded",
	ErrFileParseFail:      "failed to parse file",
	ErrReadFileFail:       "failed to read file",
	ErrFileConversionFail: "failed to convert file",
	ErrFileStoreFail:      "unable to store file",
	ErrFetchingFile:       "failed to fetch file",
	ErrIdNotfound:         "id not found",
	ErrKeyNotFound:        "unable to find this Uuid",
}

func GetErr(code errCode) string {
	x, ok := errCodes[code]
	if !ok {
		return ""
	}
	return fmt.Sprintf("%d: %s", code, x)
}
