package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"pubg/prishthbhagah"

	"github.com/rs/cors"
)

func main() {
	router := prishthbhagah.NewRouter()

	// Serve the form.html file
	router.Handle("GET", "/list", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
		// prishthbhagah.ServeFile(w, req, "./static/form.html")
		resp, err := http.Get("http://localhost:8080/list")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		// Check if the response status code is OK (200)
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error:", resp.Status)
			return
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Print the JSON response
		fmt.Println("JSON response:", string(body))

		// Alternatively, if you want to parse the JSON response into a map
		var responseData map[string]interface{}
		if err := json.Unmarshal(body, &responseData); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		// Print the parsed JSON data
		fmt.Println("Parsed JSON data:", responseData)
		prishthbhagah.RespondJSON(w, responseData, http.StatusOK)
	})

	// Handle the set endpoint
	// Handle the set endpoint
	router.Handle("POST", "/set", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
		// Parse JSON request body
		var requestData map[string]string
		if err := json.NewDecoder(req.Body).Decode(&requestData); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Access the data from the request
		key := requestData["key"]
		value := requestData["value"]
		ttl := requestData["ttl"]

		// Validate key, value, and ttl
		if key == "" || value == "" || ttl == "" {
			http.Error(w, "Key, value, and ttl are required", http.StatusBadRequest)
			return
		}

		// Make the internal POST request
		resp, err := http.PostForm("http://localhost:8080/set", url.Values{
			"key":   {key},
			"value": {value},
			"ttl":   {ttl},
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to make internal POST request: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Check if the response status code is OK (200)
		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("Unexpected response status: %s", resp.Status), http.StatusInternalServerError)
			return
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading response body: %v", err), http.StatusInternalServerError)
			return
		}

		// Print the response
		fmt.Println("Response:", string(body))

		// Respond to the client
		prishthbhagah.RespondJSON(w, map[string]string{"message": "Key-value pair set successfully"}, http.StatusOK)
	})

	// Add CORS middleware
	c := cors.AllowAll()

	// Start server with CORS middleware
	handler := c.Handler(router)
	http.ListenAndServe(":6969", handler)
}
