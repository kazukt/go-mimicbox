package main

import (
	"bytes"
	"errors"
	"testing"
	"testing/iotest"
)

func TestCat(t *testing.T) {
	want := "testdata"
	rb := new(bytes.Buffer)
	wb := new(bytes.Buffer)
	if _, err := rb.WriteString(want); err != nil {
		t.Fatal(err)
	}

	if err := cat(wb, rb); err != nil {
		t.Fatal(err)
	}

	if wb.String() != want {
		t.Errorf("\ngot:\t%v\nwant:\t%v", wb, want)
	}
}

func TestCatError(t *testing.T) {
	want := errors.New("readError")
	rb := iotest.ErrReader(want)
	wb := new(bytes.Buffer)

	if err := cat(wb, rb); err != want {
		t.Errorf("\ngot:\t%v\nwant:\t%v", err, want)
	}

	if wb.String() != "" {
		t.Fatal("wb is written")
	}
}

func TestRunFiles(t *testing.T) {}

func TestRunFilesError(t *testing.T) {}

func TestRunPipe(t *testing.T) {}
