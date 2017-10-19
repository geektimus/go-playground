package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWritter struct{}

func (logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	res, err := http.Get("http://goodsmileshop.com/en/CATEGORY-ROOT/Nendoroid/c/133?site=goodsmile-global&lang=en")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	lw := logWritter{}

	io.Copy(lw, res.Body)

}
