package MapDirectories

import (
	"fmt"
	"log"
	"os"
)

// This package navigate throgh the filesystem of operating system
// In this case we have to naviate in Windows OS

// Function to view directory permissions
func hasPermission(dirPath string) bool {
	// Read directorie and see permissions
	_, err := os.ReadDir(dirPath)
	return err == nil
}

// Mapping all files in the Windows operating system.
// The idea is to map all files in all directories.
// After that, the idea to reduce the cost of I/O calls is to use Afero framework.
// https://www.github.com/spf13/afero
// Afero is a filesystem in RAM memory, which reduces the cost of I/O calls.
func ReadAllDir(dirPath string) {
	// View directory permissions
	if !hasPermission(dirPath) {
		fmt.Printf("No se pudo acceder al directorio: %s\n", dirPath)
		return
	}
	// Get directories and files in directorie path
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	// This loop retrieves all files from a directory and displays them in the terminal
	for _, file := range files {
		// If file is a directory
		if file.IsDir() {
			fmt.Println(file.Name())                                 // Print the name os current file
			subDirPath := fmt.Sprintf("%s/%s", dirPath, file.Name()) // Get directory name
			ReadAllDir(subDirPath)                                   // Recurisve function to subdirectory
		}
	}
}
