package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	app2URL, ok := os.LookupEnv("APP2_URL")
	if !ok {
		slog.Error("missing APP2_URL")
		panic("missing APP2_URL")
	}

	getTimeURL := fmt.Sprintf("%s/time", app2URL)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request to /hello endpoint")
		resp, err := http.DefaultClient.Get(getTimeURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "app2 (%s) is not available: %v", getTimeURL, err)
			return
		}
		defer resp.Body.Close()

		resTime, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "get response time from body: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World at %s", string(resTime))
	})

	slog.Info("starting server on port 8890")
	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		slog.Error("application 1 finished with error:", err)
	}
}
