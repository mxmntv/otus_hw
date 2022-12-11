package hw10programoptimization

import (
	"archive/zip"
	"testing"
)

func BenchmarkGetDomainStat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, err := zip.OpenReader("testdata/users.dat.zip")
		if err != nil {
			b.Fatal(err)
		}

		data, err := r.File[0].Open()
		if err != nil {
			b.Fatal(err)
		}

		_, err = GetDomainStat(data, "com")
		if err != nil {
			b.Fatal(err)
		}
	}
}
