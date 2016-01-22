package compress

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func ZipName(filename string) string {
	base := filepath.Base(filename)
	return fmt.Sprintf("%s.zip", base)
}

func Compress(zipname string, files ...string) {
	buf := &bytes.Buffer{}
	w := zip.NewWriter(buf)

	next := make(chan string, len(files))
	out := make(chan *Reading, 0)
	done := make(chan bool, 1)
	wg := &sync.WaitGroup{}

	Send(next, files...)
	ReadParallel(w, next, out, done)
	ZipParallel(out, done, wg)

	wg.Add(len(files))
	fmt.Printf("Waiting for %d files to compress\n", len(files))
	wg.Wait()

	err := ioutil.WriteFile(zipname, buf.Bytes(), 0644)
	if err != nil {
		log.Println(err)
	}
	err = w.Close()
	if err != nil {
		log.Println(err)
	}
}

func Send(next chan string, files ...string) {
	for _, f := range files {
		next <- f
	}
}

type Reading struct {
	name   string
	reader io.ReadCloser
	writer *zip.Writer
}

func ReadParallel(w *zip.Writer, next chan string, out chan *Reading, done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			case absfile := <-next:
				fmt.Printf("reading: %s\n", absfile)
				name, reader, err := read(absfile)
				if err != nil {
					fmt.Println(err)
					continue
				}
				out <- &Reading{
					name:   name,
					reader: reader,
					writer: w,
				}
			}
		}
	}()
}

func ZipParallel(next chan *Reading, done chan bool, wg *sync.WaitGroup) {
	go func() {
		for {
			select {
			case <-done:
				return
			case r := <-next:
				err := CompressFile(r.reader, r.writer, r.name)
				if err != nil {
					fmt.Println(err)
				}
				err = r.reader.Close()
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}
		}
	}()
}

func CompressFile(reader io.ReadCloser, w *zip.Writer, name string) error {
	f, err := w.Create(name)
	if err != nil {
		return err
	}

	fmt.Printf("compressing: %s\n", name)
	n, err := io.Copy(f, reader)
	if n == 0 || err != nil {
		return err
	}
	return nil
}

func read(filename string) (string, io.ReadCloser, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", nil, err
	}
	return filename, f, nil
}
