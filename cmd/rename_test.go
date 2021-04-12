// Copyright © 2019 Herbert Muehlburger <mail@muehlburger.at>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// setup creates a testfile and sets creation and modification times.
func setup(t *testing.T, pattern string) (*os.File, func()) {
	mTime := time.Date(2016, time.February, 20, 5, 52, 3, 0, time.Local)
	aTime := time.Date(2016, time.February, 20, 5, 52, 0, 0, time.Local)

	f, err := ioutil.TempFile("", pattern)
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

func TestRenameWithJpgExtension(t *testing.T) {
	f, teardown := setup(t, "*.jpg")
	defer teardown()

	n := filepath.Base(Rename(f.Name()))

	expected := "2016-02-20T05.52.03.jpg"

	if n != expected {
		t.Errorf("filename should be \"%s\"; got: %v", expected, n)
	}
}

func TestRenameWithoutExtension(t *testing.T) {
	f, teardown := setup(t, "")
	defer teardown()

	n := filepath.Base(Rename(f.Name()))

	expected := "2016-02-20T05.52.03"

	if n != expected {
		t.Errorf("filename should be \"%s\"; got: %v", expected, n)
	}
}

func TestRenameWithExtensionToLowercase(t *testing.T) {
	f, teardown := setup(t, "*.JPG")
	defer teardown()

	n := filepath.Base(Rename(f.Name()))

	expected := "2016-02-20T05.52.03.jpg"

	if n != expected {
		t.Errorf("filename should be \"%s\"; got: %v", expected, n)
	}
}
