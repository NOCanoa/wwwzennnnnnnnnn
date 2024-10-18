package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFilesToBunnyCDN(storageZoneName, region, apiKey string) error {
	distDir := "./dist"
	return filepath.Walk(distDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileBytes, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(distDir, path)
			if err != nil {
				return err
			}
			req, err := http.NewRequest("PUT", fmt.Sprintf("https://%s.bunnycdn.com/%s/%s", region, storageZoneName, relPath), bytes.NewReader(fileBytes))
			if err != nil {
				return err
			}
			req.Header.Set("AccessKey", apiKey)
			req.Header.Set("Content-Type", "application/octet-stream")
			req.Header.Set("Accept", "application/json")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
				return fmt.Errorf("upload failed with status code %d", resp.StatusCode)
			}
			log.Printf("Uploaded %s to BunnyCDN\n", relPath)
		}
		return nil
	})
}

func main() {
	storageZoneName := "no"
	region := "storage"
	apiKey := "hehe"

	if len(os.Args) >= 3 {
		apiKey = os.Args[1]
	}
	if len(os.Args) >= 3 {
		storageZoneName = os.Args[2]
	}

	if err := uploadFilesToBunnyCDN(storageZoneName, region, apiKey); err != nil {
		log.Fatal(err)
	}
}
