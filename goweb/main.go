package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Handler function to respond with a colored background
func colorHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is from a browser
	ip := os.Getenv("POD_IP")
	bgColor := os.Getenv("COLOR")
	if bgColor == "" {
		bgColor = "white" // Default color if environment variable is not set
	}
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" || !isBrowser(userAgent) {
		response := fmt.Sprintf("The background color is %s on %v\n", bgColor, ip)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))
		return
	}

	// Get the background color from the environment variable

	// Respond with an HTML page with the specified background color
	response := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Color Background</title>
		</head>
		<body style="background-color: %s;">
			<h1>The background color is %s on %v </h1>
		</body>
		</html>`, bgColor, bgColor, ip)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}

// Function to check if the User-Agent belongs to a browser
func isBrowser(userAgent string) bool {
	// Simple check for common browser signatures in the User-Agent string
	browsers := []string{"Mozilla", "Chrome", "Safari", "Firefox", "Edge", "Opera"}
	for _, browser := range browsers {
		if strings.Contains(userAgent, browser) {
			return true
		}
	}
	return false
}

func main() {
	http.HandleFunc("/", colorHandler)
	port := ":8081"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
