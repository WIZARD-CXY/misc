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

// walkFile starts a gorouting to walk the directory at root and send the path
// of each regular file on the string channel. If done is closed, walkFile abandons its work
func walkFile(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)

		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

// digester reads from paths and send md5sum to the chan c until either paths is closed
// or done is closed
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		select {
		case <-done:
			return
		default:
			data, err := ioutil.ReadFile(path)
			c <- result{path, md5.Sum(data), err}
			//c <- result{path, [md5.Size]byte{}, nil}
		}
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFile(done, root)

	//Start a fixed number of goroutings to read and digest files
	c := make(chan result)

	var wg sync.WaitGroup
	const numDigesters = 8

	wg.Add(numDigesters)

	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err

		}

		m[r.path] = r.sum
	}

	// check walkFile error
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

	//fmt.Println(len(paths))
	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}
