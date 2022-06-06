// cat - concatenate and print files
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(w io.Writer, r io.Reader) error {
	if _, err := io.Copy(w, r); err != nil {
		return err
	}
	return nil
}

func run(w io.Writer, r io.Reader, args []string) error {
	if len(args) == 0 {
		if err := cat(w, r); err != nil {
			return err
		}
		return nil
	}
	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		if err := cat(w, file); err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

func main() {
	flag.Parse()
	if err := run(os.Stdout, os.Stdin, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "cat: %v\n", err)
		os.Exit(1)
	}
}
