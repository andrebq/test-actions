package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

var (
	wh          = flag.String("webhook", "", "Webhook endpoint to call")
	ct          = flag.String("contentType", "application/json", "Content type to use")
	body        = flag.String("body", "{}", "what to send as part of the body.")
	printOption = flag.String("print", "all", `Select what should be printed:
- body -> only body
- status -> only status
- any other value -> status / header / body`)
)

func main() {
	flag.Parse()
	res, err := http.Post(*wh, *ct, bytes.NewBufferString(*body))
	if err != nil {
		logrus.WithError(err).Error("Unable to perform http request")
		os.Exit(1)
	}
	printResponse(res)

	switch {
	case res.StatusCode >= 400 && res.StatusCode < 500:
		os.Exit(2)
	case res.StatusCode >= 500:
		os.Exit(3)
	}
}

func printResponse(res *http.Response) {
	defer func(body io.ReadCloser) {
		io.Copy(ioutil.Discard, body)
		body.Close()
	}(res.Body)

	switch *printOption {
	case "body":
		io.Copy(os.Stdout, res.Body)
	case "status":
		io.WriteString(os.Stdout, strconv.Itoa(res.StatusCode))
	case "all":
		io.WriteString(os.Stdout, strconv.Itoa(res.StatusCode))
		os.Stdout.WriteString("\n")

		res.Header.Write(os.Stdout)
		os.Stdout.WriteString("\n")

		io.Copy(os.Stdout, res.Body)
	}
	os.Stdout.WriteString("\n")
}
