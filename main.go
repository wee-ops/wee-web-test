package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	startTime    time.Time
	BuildVersion = "development"
)

func getEnv(env string, defaultValue string) (envValue string) {
	envValue = os.Getenv(env)
	if len(envValue) == 0 {
		envValue = defaultValue
	}
	return
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>Wee Web Test</h1><a href=\"../health\">Health</a>")
}

func handleHealth(w http.ResponseWriter, req *http.Request) {
	envs := make(map[string]string)
	envs["StartTime"] = startTime.String()
	envs["RunningBy"] = time.Since(startTime).String()
	envs["HOST_NAME"] = getEnv("HOSTNAME", "localhost")
	envs["BuildVersion"] = BuildVersion
	w.Header().Add("Content-type", "application/json")

	body, _ := json.Marshal(envs)
	w.Write(body)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	startTime = time.Now()
	httpPort, err := strconv.Atoi(getEnv("HTTP_PORT", "8080"))
	if err == nil && (httpPort < 1024 || httpPort > 65535) {
		err = fmt.Errorf("expected port between 1024 and 65535 - %d", httpPort)
	}
	if err != nil {
		log.Fatalf("Invalid HTTP_PORT - %v", err)
	}

	log.Printf("WeeWebTest \"%s\" - Listening on %v\n", BuildVersion, httpPort)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/health", handleHealth)

	err = http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
