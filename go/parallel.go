// +build OMIT

package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// sumFiles start goroutings to walk the directory at root and digest each regular file
// one file one goroutine, they all send the result to result channel, if done is closed
// sumFile abandon its work

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// for each regular file, start a gorouting do the sum work
	// send the result to channel c, send the error of filepath.Walk on errc

	c := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			wg.Add(1)

			go func() {
				select {
				case <-done:
				default:
					data, err := ioutil.ReadFile(path)
					c <- result{path, md5.Sum(data), err}
				}

				wg.Done()
			}()

			select {
			case <-done:
				return errors.New("walk canceled")
			default:
				return nil
			}
		})

		go func() {
			wg.Wait()
			close(c)
		}()

		errc <- err
	}()

	return c, errc
}

// MD5ALL reads all the files which rooted at root and returns a map
// (key is file path name, value is its md5 sum of the file's contents)
// if the directory walk fails or any read operation fails, use defer to close done
// channel. In that case all the inflight readfile and do md5sum goroutine is stopped.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root) // first stage sumFiles

	m := make(map[string][md5.Size]byte)

	for r := range c {
		if r.err != nil {
			// early return
			return nil, r.err
		}

		m[r.path] = r.sum
	}

	if err := <-errc; err != nil {
		return nil, err
	}

	return m, nil

}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	m, err := MD5All(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string

	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	fmt.Println(len(paths))
	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}

}
