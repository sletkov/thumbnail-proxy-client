package get

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to make HTTP GET request: %v", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("error: code %d for %s", response.StatusCode, url)
		return
	}

	fileName := extractFileName(url)
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("failed to create file %s: %v", fileName, err)
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		fmt.Printf("failed to copy file data for  %s: %v", url, err)
		return
	}
}

func extractFileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-2] + ".jpg"
}
