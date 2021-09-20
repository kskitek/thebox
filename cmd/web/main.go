package main

import (
	"context"
	"embed"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

const defaultPort = "8080"

//go:embed static/*
var staticDir embed.FS

func main() {
	static, _ := fs.Sub(staticDir, "static")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(static))))
	http.HandleFunc("/test", testHandler)

	port := getPort()
	log.Println("Starting thebox on port " + port)
	panic(http.ListenAndServe(port, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("magic-cookie")
	if cookie == "" || cookie != getMagicCookie() {
		log.Println("bad cookie")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

	if err := openTheBox(r.Context()); err != nil {
		log.Println("[ERR] " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func openTheBox(ctx context.Context) error {
	log.Println("openning the box")
	client, err := pubsub.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return err
	}
	defer client.Close()

	topic := client.Topic(os.Getenv("TOPIC"))
	msg := &pubsub.Message{Data: []byte("openpls")}
	_, err = topic.Publish(ctx, msg).Get(ctx)
	return err

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
