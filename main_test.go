package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// This test does not require an internet connection to run. The test server is created using the httptest.NewServer function, which creates a test server running on a random port on the local machine. The test server is closed when the test function exits, so it does not rely on an external server.

func TestMain(t *testing.T) {
    // Create a test server with the mainHandler
    ts := httptest.NewServer(http.HandlerFunc(mainHandler))
    defer ts.Close()

    // Test case for serving static files
    t.Run("StaticFileServed", func(t *testing.T) {
        // Create a request to the static file
        req, err := http.NewRequest("GET", ts.URL+"/style.css", nil)
        if err != nil {
            t.Fatal(err)
        }

        // Send the request to the test server
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            t.Fatal(err)
        }
        defer resp.Body.Close()

        // Check the response status code
        if resp.StatusCode != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
        }
    })

    // Test case for serving the index.html file
    t.Run("IndexFileServed", func(t *testing.T) {
        // Create a request to the index.html file
        req, err := http.NewRequest("GET", ts.URL+"/", nil)
        if err != nil {
            t.Fatal(err)
        }

        // Send the request to the test server
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            t.Fatal(err)
        }
        defer resp.Body.Close()

        // Check the response status code
        if resp.StatusCode != http.StatusOK {
            t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
        }
    })
}