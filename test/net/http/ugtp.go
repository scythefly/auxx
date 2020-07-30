package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_")

func newUgtpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ugtp",
		Short: "Run http examples",
		RunE:  ugtpRun,
	}

	return cmd
}

func ugtpRun(*cobra.Command, []string) error {
	proxy, err := url.Parse("http://127.0.0.1:44320")
	if err != nil {
		return err
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client := &http.Client{
		Transport: transport,
	}
	for i := 0; i < 5; i++ {
		go fakePublish(client)
	}

	select {}
}

func randomString(length int) string {
	var out []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		out = append(out, chars[rand.Int()%len(chars)])
	}
	return string(out)
}

func fakePublish(c *http.Client) {
	domain := randomString(5) + ".scythefly.top"
	app := randomString(4)
	name := randomString(7)
	start := fmt.Sprintf("http://127.0.0.1:44320/notify?call=publish&act=start&domain=%s&app=%s&name=%s",
		domain, app, name)
	update := fmt.Sprintf("http://127.0.0.1:44320/notify?call=publish&act=update&domain=%s&app=%s&name=%s",
		domain, app, name)
	done := fmt.Sprintf("http://127.0.0.1:44320/notify?call=publish&act=done&domain=%s&app=%s&name=%s",
		domain, app, name)
	fmt.Printf("http://127.0.0.1:44320/%s/%s\n", app, name)
	fmt.Println(domain)
	var request *http.Request
	var err error

	if request, err = http.NewRequest("GET", start, nil); err != nil {
		return
	}
	if _, err = c.Do(request); err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Minute)

	for {
		if request, err = http.NewRequest("GET", update, nil); err != nil {
			break
		}
		if _, err = c.Do(request); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Minute)
	}

	if request, err = http.NewRequest("GET", done, nil); err != nil {
		return
	}
	if _, err = c.Do(request); err != nil {
		fmt.Println(err)
		return
	}
}
