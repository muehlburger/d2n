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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const dateFormat string = "2006-01-02T15.04.05"

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Renames files using modification date information.",
	RunE: func(cmd *cobra.Command, args []string) error {
		src, err := cmd.Flags().GetStringSlice("src")
		if err != nil {
			return err
		}
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}
		return rename(src, verbose)
	},
}

func init() {
	renameCmd.Flags().StringSliceP("src", "s", []string{"."}, "Paths to source directories")
	renameCmd.Flags().BoolP("verbose", "v", false, "Show verbose progress messages")
	rootCmd.AddCommand(renameCmd)
}

func rename(roots []string, verbose bool) error {
	// Traverse the file tree.
	paths := make(chan string)
	go func() {
		for _, root := range roots {
			err := walkDir(root, paths)
			if err != nil {
				fmt.Printf("d2n %s: %v\n", root, err)
			}
		}
		close(paths)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles int64
	var names []string
loop:
	for {
		select {
		case name, ok := <-paths:
			if !ok {
				break loop // paths was closed
			}
			nfiles++
			names = append(names, name)
		case <-tick:
			fmt.Printf("%d files\n", nfiles)
		}
	}

	for _, f := range names {
		if err := Rename(f); err != nil {
			return err
		}
	}
	return nil
}

// Rename renames files from src to dst.
func Rename(src string) error {
	f, err := os.Stat(src)
	if err != nil {
		return err
	}

	dst := fmt.Sprintf("%s%s", f.ModTime().Format(dateFormat), strings.ToLower(filepath.Ext(src)))
	dst = filepath.Join(filepath.Dir(src), dst)
	log.Printf("%s -> %s", src, dst)
	return os.Rename(src, dst)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the absolute path of each found file on paths.
func walkDir(dir string, paths chan<- string) error {
	fi, err := os.Stat(dir)
	if !fi.IsDir() {
		path, err := filepath.Abs(filepath.Join(dir))
		if err != nil {
			log.Fatal(err)
		}
		paths <- path
		return nil
	}
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", dir, err)
	}
	for _, fi := range fis {
		if !fi.IsDir() {
			path, err := filepath.Abs(filepath.Join(dir, fi.Name()))
			if err != nil {
				log.Fatal(err)
			}
			paths <- path
		} else {
			walkDir(filepath.Join(dir, fi.Name()), paths)
		}
	}
	return nil
}
