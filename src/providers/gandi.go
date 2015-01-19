// Gandi.net API (XML-RPC) provider
package providers

import (
	"bytes"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"io"
	"io/ioutil"
	"net/http"
)

const gandiURL = "https://rpc.gandi.net/xmlrpc/"
const gandiOTEURL = "https://rpc.ote.gandi.net/xmlrpc/"

type GandiProvider struct {
	ApiKey string
	Domain string
	OTE    bool
	URL    string
}

func Gandi(api, domain string) (GandiProvider, error) {
	var gandi GandiProvider
	gandi.ApiKey = api
	gandi.Domain = domain
	gandi.OTE = true
	gandi.URL = gandiURL
	if gandi.OTE {
		gandi.URL = gandiOTEURL
	}

	err := gandi.Auth()
	return gandi, err
}

func (g GandiProvider) Auth() error {
	req := struct {
		Apikey string
	}{
		g.ApiKey,
	}
	buf, _ := xml.EncodeClientRequest("domain.tld.list", &req)
	fmt.Println(string(buf))
	reply, err := post(g.URL, buf)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(reply)
	fmt.Println(data)
	return err
}

func (g GandiProvider) Update(record string, ip string) (error, bool) {

	return nil, false
}

func post(url string, request []byte) (io.ReadCloser, error) {
	resp, err := http.Post(url, "text/xml", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Body, err
}
