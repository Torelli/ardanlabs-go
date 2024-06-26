package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	name, repos, err := githubInfo(ctx, "torelli")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Hi, my name is %s and I have %d public repos\n", name, repos)
	}
}

// githubInfo returns name and number of public repos for login
func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	// resp, err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}
	defer resp.Body.Close()
	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// var r reply
	var r struct { // anonymous struct
		Name         string
		Public_Repos int
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.Public_Repos, nil
}

/*
type reply struct {
	Name         string
	Public_Repos int
}


JSON 		<-> Go
true/false 	<-> true/false
string 		<-> string
null 		<-> nil
number 		<-> float64, float32, int8, int15, int32, int64, int, uint8, ...
array 		<-> []any ([]interface{})
object 		<-> map[string]any, struct

enconding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
