package main

import (
	"crypto/fips140"
	"crypto/sha256"
	"crypto/sha3"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

var backend = "none"

var doServe = flag.Bool("serve", false, "Start the HTTP server")

func main() {
	flag.Parse()

	fmt.Printf("fips140.Enabled(): %v\n", fips140.Enabled())
	if data, err := resp("value", "/sum?v=value"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
	if !*doServe {
		return
	}
	http.HandleFunc("/sum", sumHandler)
	http.ListenAndServe(":8080", nil)
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("v")

	data, err := resp(v, r.URL.RequestURI())
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = w.Write(data); err != nil {
		log.Fatalln(err)
	}
}

func resp(v, uri string) ([]byte, error) {
	s := sha256.New()
	if _, err := s.Write([]byte(uri)); err != nil {
		return nil, err
	}
	sumStr := hex.EncodeToString(s.Sum(nil))

	t := time.Now().UTC().Format(time.RFC3339Nano)
	if _, err := s.Write([]byte(t)); err != nil {
		return nil, err
	}
	sumStr2 := hex.EncodeToString(s.Sum(nil))

	hashBytes := sha3.SumSHAKE256([]byte(sumStr2+t), 64)
	sumStr2 = hex.EncodeToString(hashBytes[:])

	data, err := json.MarshalIndent(struct {
		Request string
		Version string
		Backend string
		Type    string
		Value   string
		Sum     string
		Value2  string
		Sum2    string
	}{
		uri,
		runtime.Version(),
		backend,
		fmt.Sprintf("%T", s),
		v,
		sumStr,
		v + t,
		sumStr2,
	}, "", "  ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
