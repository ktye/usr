// Mkinit creates an cpio newc archive of a single file
//
// Usage:
//	mkinit
//
// This is equivalent to do
//	echo init | cpio -o -H newc | gzip > initfs.gz
//
// The input file must be called 'init' and the archive is
// written to initfs.gz
//
// Mkinit ignores the file attributes, sets modes to 0555
// and owner to root.
//
// Note: the result is an initramfs with the file 'init' in the root
// directory instead of the linux kernel default /sbin/init
package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {

	r, err := os.Open("init")
	fatal(err)
	defer r.Close()

	files := []Record{Record{
		ReaderAt: r,
	}}

	out, err := os.Create("initfs.gz")
	fatal(err)
	defer out.Close()

	zw := gzip.NewWriter(out)
	defer zw.Close()

	cw := Newc.Writer(zw)
	err = WriteRecords(cw, files)
	fatal(err)

	err = WriteTrailer(cw)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
