// Gandi.net API (XML-RPC) provider
package providers

const gandiURL = "https://rpc.ote.gandi.net/xmlrpc/"

type GandiProvider struct {
	ApiKey string
	OTE    bool
}

func Gandi(api string) GandiProvider {
	var gandi GandiProvider
	gandi.ApiKey = api
	return gandi
}
