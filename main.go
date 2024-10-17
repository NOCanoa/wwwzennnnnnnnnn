package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	start := time.Now()
	fmt.Printf("Start\n")
	ftpPassword := "no-pasword"
	if len(os.Args) >= 3 {
		ftpPassword = os.Args[1]
	}
	ftpUsername := "no-username"
	if len(os.Args) >= 3 {
		ftpUsername = os.Args[2]
	}

	// Define your FTP server credentials
	ftpServer := "storage.bunnycdn.com:21"

	// Connect to the FTP server
	conn, err := ftp.Connect(ftpServer)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Quit()

	// Login to the FTP server
	err = conn.Login(ftpUsername, ftpPassword)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Uploading files to FTP server...")
	// Open the directory you want to upload
	dir, err := os.Open("./out")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()

	// Get a list of files in the directory
	files, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	// Upload each file in the directory
	for _, file := range files {
		if file.IsDir() {
			// Skip subdirectories for now
			continue
		}

		// Open the file
		f, err := os.Open("out/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Upload the file to the FTP server
		err = conn.Stor(file.Name(), f)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Upload complete!")
	elapsed := time.Since(start)
	fmt.Printf("Done in %v\n", elapsed)
	fmt.Printf("Done!!")
}
