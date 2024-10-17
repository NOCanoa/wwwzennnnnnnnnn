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
	outDir := "./out"
	files, err := ioutil.ReadDir(outDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		filePath := filepath.Join(outDir, file.Name())
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		req, err := http.NewRequest("PUT", fmt.Sprintf("https://%s.bunnycdn.com/%s/%s", region, storageZoneName, file.Name()), bytes.NewReader(fileBytes))
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

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("upload failed with status code %d", resp.StatusCode)
		}

		log.Printf("Uploaded %s to BunnyCDN\n", file.Name())
	}

	return nil
}

func main() {
	storageZoneName := "your_storage_zone_name"
	region := "storage"
	apiKey := "your_api_key"
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
