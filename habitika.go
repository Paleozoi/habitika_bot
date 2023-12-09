package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func login_habitika(u string, p string) {
	url := "https://habitica.com/api/v3/user/auth/local/login"

	data := []byte(fmt.Sprintf(`{"username":"%v", "password":"%v"}`, u, p))
	fmt.Println(string(data))
	payload := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Add("accept", "plain/text")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
