package library

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

type BasicAuthTransport struct {
	Username  string
	Password  string
	Transport http.RoundTripper
}

func (b *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(b.Username, b.Password)
	return b.Transport.RoundTrip(req)
}

func InitElasticsearch() {
	address := os.Getenv("ELASTICSEARCH_ADDRESS")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	var err error
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{address},
		Transport: &BasicAuthTransport{
			Username:  username,
			Password:  password,
			Transport: transport,
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func GetESClient() *elasticsearch.Client {
	return client
}
