package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	twitter "github.com/eOracle/go-twitter/v2"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

/*
*

	In order to run, the user will need to provide the bearer token and the list of tweet ids.

*
*/
func main() {
	token := flag.String("token", "", "twitter API token")
	query := flag.String("query", "", "query")
	flag.Parse()

	client := &twitter.Client{
		Authorizer: authorize{
			Token: *token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	opts := twitter.SpacesSearchOpts{
		SpaceFields: []twitter.SpaceField{twitter.SpaceFieldHostIDs, twitter.SpaceFieldTitle},
	}

	fmt.Println("Callout to spaces search callout")

	spaceResponse, err := client.SpacesSearch(context.Background(), *query, opts)
	if err != nil {
		log.Panicf("spaces search error: %v", err)
	}

	enc, err := json.MarshalIndent(spaceResponse, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(enc))
}
