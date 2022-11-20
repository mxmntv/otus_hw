package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	name, err := os.MkdirTemp("./", "copied")
	if err != nil {
		log.Fatal(err)
	}

	t.Run("simple case", func(t *testing.T) {
		dest := "./" + name + "/cinput.txt"
		Copy("./testdata/input.txt", dest, 0, 0)
		src, err := os.Open(dest)
		if err != nil {
			log.Fatal(err)
		}
		i, _ := GetLen(src)
		require.Equal(t, 6617, i)
	})

	t.Run("offset 0, limit 10 case", func(t *testing.T) {
		dest := "./" + name + "/cout_offset0_limit10.txt"
		Copy("./testdata/out_offset0_limit10.txt", dest, 0, 10)
		src, err := os.Open(dest)
		if err != nil {
			log.Fatal(err)
		}
		i, _ := GetLen(src)
		require.Equal(t, 10, i)
	})

	t.Run("offset 0, limit 1000 case", func(t *testing.T) {
		dest := "./" + name + "/cout_offset0_limit1000.txt"
		Copy("./testdata/out_offset0_limit1000.txt", dest, 0, 1000)
		src, err := os.Open(dest)
		if err != nil {
			log.Fatal(err)
		}
		i, _ := GetLen(src)
		require.Equal(t, 1000, i)
	})

	t.Run("offset 0, limit 10000 case", func(t *testing.T) {
		dest := "./" + name + "/cout_offset0_limit10000.txt"
		Copy("./testdata/out_offset0_limit10000.txt", dest, 0, 10000)
		src, err := os.Open(dest)
		if err != nil {
			log.Fatal(err)
		}
		i, _ := GetLen(src)
		fmt.Println(i)
		require.Equal(t, 6617, i)
	})

	t.Run("offset 100, limit 1000 case", func(t *testing.T) {
		dest := "./" + name + "/cout_offset100_limit1000.txt"
		Copy("./testdata/out_offset100_limit1000.txt", dest, 100, 1000)
		src, err := os.Open(dest)
		if err != nil {
			log.Fatal(err)
		}
		i, _ := GetLen(src)
		fmt.Println(i)
		require.Equal(t, 900, i)
	})

	defer os.RemoveAll(name)
}
