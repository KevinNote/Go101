package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var hc = http.Client{}

const ENDPOINT = "http://127.0.0.1:11451"

func uri(path string) string {
	return ENDPOINT + path
}

func testTime() {
	bs := httpGet(uri("/currentTime"))
	fmt.Println(string(bs))
	t, err := time.Parse(time.RFC3339, string(bs))
	panicNotNil(err)
	fmt.Println(t)
	dif := time.Now().Sub(t)
	if dif.Minutes() > 1 {
		panic("time difference is too big")
	}
}

type addReq struct {
	A int `json:"a"`
	B int `json:"b"`
}

type addStrReq struct {
	A string `json:"a"`
	B string `json:"b"`
}

type SumResponse struct {
	Sum int `json:"sum"`
}

func toJson(v any) []byte {
	bs, err := json.Marshal(v)
	panicNotNil(err)
	return bs
}

func testAdd() {
	a := rand.Int()
	b := rand.Int()
	req := addReq{A: a, B: b}
	bs := httpPost(uri("/add"), toJson(req))
	verifyAddResult(bs, a+b)
}

func testAddStr() {
	a := rand.Int()
	b := rand.Int()
	req := addStrReq{A: strconv.Itoa(a), B: strconv.Itoa(b)}
	bs := httpPost(uri("/addString"), toJson(req))
	verifyAddResult(bs, a+b)
}

func verifyAddResult(bs []byte, ans int) {
	var resp SumResponse
	err := json.Unmarshal(bs, &resp)
	panicNotNil(err)
	if resp.Sum != ans {
		panic("sum is not correct")
	}
}

func main() {
	testTime()
	testAdd()
	testAddStr()
	fmt.Println("all tests passed")
}

func httpGet(url string) []byte {
	resp, err := hc.Get(url)
	panicNotNil(err)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("status code not 200")
	}
	bs, err := io.ReadAll(resp.Body)
	panicNotNil(err)
	return bs
}

func httpPost(url string, bodyBs []byte) []byte {
	body := bytes.NewReader(bodyBs)
	resp, err := hc.Post(url, "application/json", body)
	panicNotNil(err)
	if resp.StatusCode != http.StatusOK {
		panic("status code not 200")
	}
	bs, err := io.ReadAll(resp.Body)
	panicNotNil(err)
	return bs
}

func panicNotNil(e error) {
	if e != nil {
		panic(e)
	}
}
