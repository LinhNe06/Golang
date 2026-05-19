package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchGames(url string, dataChan chan<- string, errorChan chan<- error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		errorChan <- err
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorChan <- err
		return
	}

	dataChan <- string(body[:500])
}

func main() {
	// Thay bằng API Key thật lấy từ rawg.io/apikeys
	apiKey := "YOUR_API_KEY_HERE"

	apiUrl := fmt.Sprintf("https://api.rawg.io/api/games?key=%s", apiKey)

	dataChan := make(chan string)
	errorChan := make(chan error)

	timeoutChan := time.After(5000 * time.Millisecond)

	fmt.Println("[SYSTEM] Đang kích hoạt Goroutine để gọi RAWG API...")
	go fetchGames(apiUrl, dataChan, errorChan)

	select {
	case data := <-dataChan:
		fmt.Println("[SUCCEED] Dữ liệu từ API:", data)

	case err := <-errorChan:
		fmt.Println("[ERROR] Không thể lấy dữ liệu:", err)

	case <-timeoutChan:
		fmt.Println("[TIMEOUT] Đã quá 5 giây, hủy kết nối!")
	}
}
