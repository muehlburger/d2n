package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// setup creates a testfile and sets creation and modification times.
func setup(t *testing.T) (*os.File, func()) {
	t.Parallel()

	mTime := time.Date(2016, time.February, 20, 5, 52, 3, 0, time.Local)
	aTime := time.Date(2016, time.February, 20, 5, 52, 0, 0, time.Local)

	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	teardown := func() {
		f.Close()
		os.Remove(f.Name())
	}

	err = os.Chtimes(f.Name(), aTime, mTime)
	return f, teardown
}

func TestRename(t *testing.T) {
	f, teardown := setup(t)
	defer teardown()

	n := Rename(f.Name())

	got := filepath.Base(n)
	exp := "2016-02-20T05.52.03"
	if exp != got {
		t.Fatalf("wanted: %v; got: %v", exp, got)
	}
}
