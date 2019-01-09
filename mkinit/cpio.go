package main

// Adapted from:
// github.com/u-root/u-root/pkg/cpio/cpio.go
//
// Copyright 2013-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"io"
	"os"
	"time"
)

var (
	formatMap = make(map[string]RecordFormat)
	Debug     = func(string, ...interface{}) {}
)

// Record represents a CPIO record, which represents a Unix file.
type Record struct {
	io.ReaderAt // ReaderAt contains the content of this CPIO record.
	Info        // Info is metadata describing the CPIO record.

	// Info is metadata describing the CPIO record.
	RecPos  int64  // Where in the file this record is
	RecLen  uint64 // How big the record is.
	FilePos int64  // Where in the CPIO the file's contents are.
}

func (r Record) String() string {
	return "init" // The input file name is hard-coded.
}

// Info holds metadata about files.
type Info struct {
	Ino      uint64
	Mode     uint64
	UID      uint64
	GID      uint64
	NLink    uint64
	MTime    uint64
	FileSize uint64
	Dev      uint64
	Major    uint64
	Minor    uint64
	Rmajor   uint64
	Rminor   uint64
	Name     string
}

func (i Info) String() string {
	return fmt.Sprintf("%s: Ino %d Mode %#o UID %d GID %d NLink %d MTime %v FileSize %d Major %d Minor %d Rmajor %d Rminor %d",
		i.Name,
		i.Ino,
		i.Mode,
		i.UID,
		i.GID,
		i.NLink,
		time.Unix(int64(i.MTime), 0).UTC(),
		i.FileSize,
		i.Major,
		i.Minor,
		i.Rmajor,
		i.Rminor)
}

// A RecordWriter writes one record to an archive.
type RecordWriter interface {
	WriteRecord(Record) error
}

// A RecordFormat gives readers and writers for dealing with archives from io
// objects.
//
// CPIO files have a number of records, of which newc is the most widely used
// today.
type RecordFormat interface {
	Writer(w io.Writer) RecordWriter
}

// Format returns the RecordFormat with that name, if it exists.
func Format(name string) (RecordFormat, error) {
	op, ok := formatMap[name]
	if !ok {
		return nil, fmt.Errorf("%q is not in cpio format map %v", name, formatMap)
	}
	return op, nil
}

func modeFromLinux(mode uint64) os.FileMode {
	return 0555
}
