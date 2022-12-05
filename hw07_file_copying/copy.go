package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrInvalidParameters     = errors.New("invalid request parameters")
	ErrCrtFile               = errors.New("couldn't create destination file")
)

func GetLen(f *os.File) (int, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, ErrUnsupportedFile
	}
	return int(stat.Size()), nil
}

func checkLimiters(ot int64, size int) error {
	switch {
	case ot < 0:
		return ErrInvalidParameters
	case int(ot) > size:
		return ErrOffsetExceedsFileSize
	}
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	dest, err := os.Create(toPath)
	if err != nil {
		return ErrCrtFile
	}

	fileLen, err := GetLen(src)
	if err != nil {
		return err
	}

	defer src.Close()
	defer dest.Close()

	if err := checkLimiters(offset, fileLen); err != nil {
		return err
	}

	var bufLen int
	if limit > 0 {
		bufLen = int(limit)
	} else {
		bufLen = fileLen - int(offset)
	}

	src.Seek(offset, io.SeekStart)

	bar := pb.Full.Start64(int64(bufLen))
	barReader := bar.NewProxyReader(src)

	b, err := io.CopyN(dest, barReader, int64(bufLen))
	if err != nil {
		if errors.Is(err, io.EOF) {
			fmt.Printf("EOF copied bytes %d", b)
		} else {
			return errors.New("file copy error")
		}
	}
	bar.Finish()
	return nil
}
