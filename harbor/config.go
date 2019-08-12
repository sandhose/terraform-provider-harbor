package harbor

import (
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/httpclient"
)

type Config struct {
	Host     string
	Username string
	Password string
}

type CombinedConfig struct {
	client *apiclient.Harbor
	auth   runtime.ClientAuthInfoWriter
}

func (c *Config) Config() (*apiclient.Harbor, error) {
	cfg := apiclient.DefaultTransportConfig().
		WithHost(c.Host)
	cli := httpclient.New()
	cli.Transport = logging.NewTransport("Harbor", cli.Transport)

	runtime := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, cli)
	auth := httptransport.BasicAuth(c.Username, c.Password)
	runtime.DefaultAuthentication = auth
	client := apiclient.New(runtime, strfmt.Default)
	return client, nil
}
