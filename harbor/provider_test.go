package harbor

import (
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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

	err := rawProvider.Configure(terraform.NewResourceConfigRaw(raw))
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
