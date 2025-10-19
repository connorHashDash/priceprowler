package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"priceprowler/internal/hmlandreg"
	"priceprowler/internal/output"
)

func main() {
	postcode := Flags()
	if len(postcode) > 1 || len(postcode) < 4 {
		ServerCall(postcode)
	}
	hmlandreg.Init()
	err := output.TrendByPropertyType()
	if err != nil {
		log.Fatal(err)
	}
	output.WholePostCodeTrend()
}

func Flags() string {
	PostCode := flag.String("postcode", "", "Postcode of area to be searched")
	flag.Parse()

	return *PostCode
}

func ServerCall(postcode string) error {
	var url string = fmt.Sprintf("%s%s%s%s", "http://", os.Getenv("ENDPOINT_URL"), "?postcode=", postcode)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Postcode Updated")
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}

	bodyString := string(bodyBytes)

	return fmt.Errorf("%v", bodyString)
}
