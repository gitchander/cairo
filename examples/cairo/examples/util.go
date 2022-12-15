package main

import (
	"errors"
	"log"
	"math"
	"os"
)

func DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func RadToDeg(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

func makeDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New("file is not dir")
		}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
