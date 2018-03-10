package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type response struct {
	Kind        string
	ID          string
	OutputLabel string
	OutputMulti []struct {
		Label string
		Score string
	}
	OutputValue float64
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter text > ")
	for s.Scan() {
		res, err := do(s.Text())
		if err != nil {
			panic(err)
		}
		fmt.Printf("This looks like %s to me", res.OutputLabel)
		fmt.Println("Enter text > ")
	}
}

func do(query string) (*response, error) {

	values := url.Values{
		"model":  []string{"Language Detection"},
		"Phrase": []string{query},
	}

	var result response

	res, err := http.PostForm("http://try-prediction.appspot.com/predict", values)

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
