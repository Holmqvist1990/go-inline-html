package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	dest   string
	source string
)

func main() {
	parseFlags()
	for _, filename := range htmlPaths(source) {
		var (
			fullPath   = fmt.Sprintf("./%v/%v", source, filename)
			htmlBytes  = bytesFromFile(fullPath)
			lookup     = lookupFrom(filename)
			oldContent = bytesFromFile(dest)
			start, end = getStartAndEnd(oldContent, lookup)
			newContent = oldContent[start:end]
			html       = append(append([]byte("[]byte(`"), htmlBytes...), []byte("`)\n")...)
			final      = bytes.Replace(oldContent, newContent, html, 1)
		)
		err := ioutil.WriteFile(dest, []byte(final), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}
	os.Exit(0)
}

func parseFlags() {
	flag.StringVar(&dest, "dest", "", "destination Go file that contains variables to be filled with HTML")
	flag.StringVar(&source, "source", "", "folder with HTML files, named as [variable].html")
	flag.Parse()
	if dest == "" {
		log.Fatal("argument 'destination file' cannot be empty")
	}
	if source == "" {
		log.Fatal("argument 'source file' cannot be empty")
	}
}

func htmlPaths(source string) []string {
	paths := []string{}
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if !strings.Contains(name, ".html") {
			return nil
		}
		paths = append(paths, name)
		return nil
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("htmlPaths: %v", err))
	}
	return paths
}

func bytesFromFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func lookupFrom(filename string) string {
	lookup := filename[:strings.Index(filename, ".html")]
	return fmt.Sprintf("%v = ", lookup)
}

func getStartAndEnd(destBytes []byte, lookup string) (int, int) {
	start := bytes.Index(destBytes, []byte(lookup))
	if start < 0 {
		log.Fatal(fmt.Sprintf("could not find index for %v in %v", lookup, dest))
	}
	start = start + len(lookup)
	end := bytes.Index(destBytes[start:], []byte("\n"))
	if end < 0 {
		log.Fatal("could not find last")
	}
	end = start + end + 1
	return start, end
}
