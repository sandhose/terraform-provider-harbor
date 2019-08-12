package harbor

import (
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProviderParameters(t *testing.T) {
	rawProvider := Provider()
	raw := map[string]interface{}{
		"host":     "registry.example.com",
		"username": "user",
		"password": "hunter2",
	}
	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	rawProvider.Configure(terraform.NewResourceConfig(rawConfig))
	meta := rawProvider.(*schema.Provider).Meta()
	if meta == nil {
		t.Fatalf("Expected metadata, got nil: err: %s", err)
	}

	client := meta.(*apiclient.Harbor)
	transport := client.Transport.(*httptransport.Runtime)
	if transport.Host != raw["host"] {
		t.Fatalf("Expected %s, got %s", raw["host"], transport.Host)
	}

	// TODO: test basic auth
}
