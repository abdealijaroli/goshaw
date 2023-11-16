package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func getImage(statusCode int) ([]byte, error) {
	url := fmt.Sprintf("https://http.cat/%d", statusCode)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d not found", statusCode)
	}

	imageBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		statusCode, err := strconv.Atoi(r.FormValue("status"))
		if err != nil {
			http.Error(w, "Invalid status code", http.StatusBadRequest)
			return
		}

		imageBytes, err := getImage(statusCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Set the appropriate content type for JPEG images
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(imageBytes)
		return
	}

	http.ServeFile(w, r, "static/index.html")
}

func main() {
	http.HandleFunc("/", imageHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe("localhost:8080", nil)
}
