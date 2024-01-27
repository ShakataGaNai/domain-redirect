package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var redir_domain string

func main() {
	env_domain, ok := os.LookupEnv("REDIRECT_DOMAIN")
	if !ok {
		log.Fatalf("Environment variable REDIRECT_DOMAIN must be set!")
	}
	redir_domain = env_domain
	http.HandleFunc("/", redirect)
	listenAddr, ok := os.LookupEnv("LISTEN_ADDR")
	if !ok {
		listenAddr = ":8080"
	}
	log.Printf("Listening at %s, redirecting to %s", listenAddr, redir_domain)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("Cannot listen to port: %v\n", err)
	}
}

func redirect(writer http.ResponseWriter, request *http.Request) {
	domains := strings.Split(request.Host, ".")
	subdomain := strings.Join(domains[0:len(domains)-2], "") // cut off the tld and domain
	path := request.RequestURI
	redir := fmt.Sprintf("https://%s.%s%s", subdomain, redir_domain, path)
	fmt.Printf("Got a request: %s looking for %s; redirecting to %s\n", request.Host, path, redir)
	writer.Header().Add("Location", redir)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}
