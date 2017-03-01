// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"net/url"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, u := range os.Args[1:] {
		go fetch(u, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(u string, ch chan<-string) {
	start := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	ct := time.Now()
	urlString, _ := url.Parse(u)
	fname := urlString.Host + "_" fmt.Sprintf("%d:%d-%d-%s-%d.html", ct.Hour(), ct.Minute(), ct.Day(), ct.Month(), ct.Year())
	file, err := os.Create(fname)
	if err != nil {
		ch <- fmt.Sprint("while creating file %s: %v", fname, err)
		return
	}
	written, err := io.Copy(file, resp.Body)
	resp.Body.Close() //don't leak reseources
	file.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s, $d bytes written: %v", u, written, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs file %s created", secs, fname)
}

