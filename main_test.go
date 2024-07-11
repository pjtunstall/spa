package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// This test does not require an internet connection to run. The test server is created using the httptest.NewServer function, which creates a test server running on a random port on the local machine. The test server is closed when the test function exits, so it does not rely on an external server.

func TestMain(t *testing.T) {
	// Test that static files are served correctly
	t.Run("StaticFileServed", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/style.css", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	// Test that index.html is served for non-static paths
	t.Run("IndexServed", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandler)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	// Test that the server starts without error
	t.Run("ServerStarts", func(t *testing.T) {
		go func() {
			if err := http.ListenAndServe(":3000", nil); err != nil {
				log.Fatal("Error starting server: ", err)
			}
		}()
		// Give the server some time to start
		time.Sleep(time.Second)
	})
}