package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Some code is extracted by https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

// ReadFile Reads a file and return your content
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// wrap error with a text representing what happens
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		// wrap error with a text representing what happens
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

// ReadConfig returns config, but also wraps the error with message
func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.Wrap(err, "could not read config")
}

func one() error {
	// one call two, two call three and and an error happens
	err := two()
	if err != nil {
		return errors.Wrap(err, "error on two")
	}
	return nil
}

func two() error {
	// something happens is not intuitive
	err := three()
	if err != nil {
		return errors.Wrap(err, "error on three")
	}
	anotherErr := four()
	if anotherErr != nil {
		return errors.Wrap(anotherErr, "error on four")
	}
	return nil
}

func three() error {
	return fmt.Errorf("something happens")
}

func four() error {
	return fmt.Errorf("something happens")
}

func main() {
	// err := one()
	_, err := ReadConfig()
	if err != nil {
		// print complete traceback
		stack := errors.WithStack(err)
		fmt.Printf("%+v", stack)
	}
}
