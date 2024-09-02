package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request to /time endpoint")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, time.Now().String())
	})

	slog.Info("starting server 2 on port :8891")

	err := http.ListenAndServe(":8891", nil)
	if err != nil {
		slog.Error("application 2 finished with error:", err)
	}
}
