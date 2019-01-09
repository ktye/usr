package main

// Adapted from:
// github.com/u-root/u-root/pkg/cpio/util.go
//
// Copyright 2013-2017 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
)

// Trailer is the name of the trailer record.
const Trailer = "TRAILER!!!"

// TrailerRecord is the last record in any CPIO archive.
var TrailerRecord = StaticRecord(nil, Info{Name: Trailer})

// StaticRecord returns a record with the given contents and metadata.
func StaticRecord(contents []byte, info Info) Record {
	info.FileSize = uint64(len(contents))
	return Record{
		ReaderAt: bytes.NewReader(contents),
		Info:     info,
	}
}

// StaticFile returns a normal file record.
func StaticFile(name string, content string, perm uint64) Record {
	return StaticRecord([]byte(content), Info{
		Name: name,
		Mode: 0x8000 | perm,
	})
}

// DedupWriter is a RecordWriter that does not write more than one record with
// the same path.
//
// There seems to be no harm done in stripping duplicate names when the record
// is written, and lots of harm done if we don't do it.
type DedupWriter struct {
	rw RecordWriter

	// alreadyWritten keeps track of paths already written to rw.
	alreadyWritten map[string]struct{}
}

// NewDedupWriter returns a new deduplicating rw.
func NewDedupWriter(rw RecordWriter) RecordWriter {
	return &DedupWriter{
		rw:             rw,
		alreadyWritten: make(map[string]struct{}),
	}
}

// WriteRecord implements RecordWriter.
//
// If rec.Name was already seen once before, it will not be written again and
// WriteRecord returns nil.
func (dw *DedupWriter) WriteRecord(rec Record) error {
	if _, ok := dw.alreadyWritten[rec.Name]; ok {
		return nil
	}
	dw.alreadyWritten[rec.Name] = struct{}{}
	return dw.rw.WriteRecord(rec)
}

// WriteRecords writes multiple records to w.
func WriteRecords(w RecordWriter, files []Record) error {
	for _, f := range files {
		if err := w.WriteRecord(f); err != nil {
			return fmt.Errorf("WriteRecords: writing %q got %v", f.Info.Name, err)
		}
	}
	return nil
}

// WriteTrailer writes the trailer record.
func WriteTrailer(w RecordWriter) error {
	return w.WriteRecord(TrailerRecord)
}
