package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// setup creates a testfile and sets creation and modification times.
func setup(t *testing.T) (*os.File, func()) {
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
	if err != nil {
		t.Fatalf("could not update modification and creation timestamps: %v", err)
	}
	return f, teardown
}

func TestRenameToModTime(t *testing.T) {
	f, teardown := setup(t)
	defer teardown()

	n := filepath.Base(Rename(f.Name()))

	expected := fmt.Sprintf("%s %s", "2016-02-20T05.52.03", filepath.Base(f.Name()))

	if n != expected {
		t.Errorf("filename should be \"%s\"; got: %v", expected, n)
	}
}
