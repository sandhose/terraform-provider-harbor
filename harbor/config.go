package harbor

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"
)

type Config struct {
	Host     string
	Username string
	Password string
}

type transport struct {
	transport http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Fix for goharbor/harbor#8197
	resp.Header.Set("Content-Type", "application/json")

	return resp, nil
}

type CombinedConfig struct {
	client *apiclient.Harbor
	auth   runtime.ClientAuthInfoWriter
}

func (c *Config) Config() (*apiclient.Harbor, error) {
	cfg := apiclient.DefaultTransportConfig().
		WithHost(c.Host)
	cli := cleanhttp.DefaultPooledClient()
	cli.Transport = &transport{transport: cli.Transport}
	cli.Transport = logging.NewTransport("Harbor", cli.Transport)

	runtime := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, cli)
	auth := httptransport.BasicAuth(c.Username, c.Password)
	runtime.DefaultAuthentication = auth
	client := apiclient.New(runtime, strfmt.Default)
	log.Printf("[INFO] Harbor Client configured for host: %s", c.Host)
	return client, nil
}
