// Gandi.net API (XML-RPC) provider
package providers

const gandiURL = "https://rpc.ote.gandi.net/xmlrpc/"

type GandiProvider struct {
	ApiKey string
	Domain string
	OTE    bool
}

func Gandi(api, domain string) (GandiProvider, error) {
	var gandi GandiProvider
	gandi.ApiKey = api
	gandi.Domain = domain
	gandi.OTE = true
	return gandi, nil
}

func (g GandiProvider) Update(record string, ip string) (error, bool) {

	return nil, false
}
