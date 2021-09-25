package main

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	// "cloud.google.com/go/pubsub"
	"github.com/kskitek/thebox/internal/pubsub"
)

const defaultPort = "8080"

//go:embed static/*
var staticDir embed.FS

func main() {
	static, _ := fs.Sub(staticDir, "static")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(static))))
	http.Handle("/test", cors(cookieAuth(http.HandlerFunc(testHandler))))
	http.Handle("/reset", cors(cookieAuth(http.HandlerFunc(resetHandler))))

	port := getPort()
	log.Println("Starting thebox on port " + port)
	panic(http.ListenAndServe(port, nil))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, magic-cookie")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func cookieAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Header.Get("magic-cookie")
		if cookie == "" || cookie != getMagicCookie() {
			log.Println("bad cookie: " + cookie)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if err := changeBoxState(r.Context(), false); err != nil {
		log.Println("[ERR] " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	testString, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	if len(testString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("testing: " + string(testString))
	if string(testString) != getSecret() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := changeBoxState(r.Context(), true); err != nil {
		log.Println("[ERR] " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func changeBoxState(ctx context.Context, open bool) error {
	pub, err := pubsub.NewPublisher()
	if err != nil {
		return err
	}
	msg := pubsub.MessageClose
	if open {
		msg = pubsub.MessageOpen
	}
	return pub.Publish(ctx, msg)
}

func getPort() string {
	if port := os.Getenv("HTTP_PORT"); port != "" {
		return ":" + port
	}
	return ":" + defaultPort
}

func getMagicCookie() string {
	return os.Getenv("SECRET_COOKIE")
}

func getSecret() string {
	return os.Getenv("SECRET")
}
