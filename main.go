package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func main() {
	// Here add links to the videos
	links := [...]string{
		//"https://www.youtube.com/watch?v=9c5yPIQ3LQI",
		//"https://www.youtube.com/watch?v=qacxYIsOuKY",
	}

	thumbnail_postfix := 1
	for _, link := range links {
		videoID := extractVideoID(link)
		if videoID == "" {
			log.Fatalf("Failed to extract video ID from URL: %s\n", link)
		}

		thumbnailURL := fmt.Sprintf("https://i.ytimg.com/vi/%s/hqdefault.jpg", videoID)
		// You have to have tor up and running
		//this commands will downlad it and start the tor:
		/*
			sudo pacman -S tor
			sudo systemctl start tor
			sudo systemctl enable tor

		*/
		torProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
		if err != nil {
			log.Fatalf("Failed to parse the proxy URL: %v\n", err)
		}

		dialer, err := proxy.FromURL(torProxyURL, proxy.Direct)
		if err != nil {
			log.Fatalf("Failed to obtain proxy dialer: %v\n", err)
		}

		transport := &http.Transport{
			Dial: dialer.Dial,
		}

		client := &http.Client{
			Transport: transport,
		}

		resp, err := client.Get(thumbnailURL)
		if err != nil {
			log.Fatalf("Failed to issue the GET request: %v", err)
		}
		defer resp.Body.Close()

		err = os.MkdirAll("pictures", 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v\n", err)
		}

		filename := fmt.Sprintf("pictures/thumbnail%d.jpg", thumbnail_postfix)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file to store the image: %v", err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatalf("Failed to write the response to file: %v", err)
		}

		fmt.Printf("Operation successful: Saved to %s\n", filename)

		thumbnail_postfix++
	}
}

func extractVideoID(link string) string {
	u, err := url.Parse(link)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return ""
	}

	query := u.Query()
	return query.Get("v")
}

