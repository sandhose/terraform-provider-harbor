package harbor

import (
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/sandhose/terraform-provider-harbor/api/client"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	testAccProvider  *schema.Provider
	testAccProviders map[string]terraform.ResourceProvider
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"harbor": testAccProvider,
	}
}

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

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("HARBOR_HOST"); v == "" {
		t.Log(v)
		t.Fatal("HARBOR_HOST must be set for acceptance tests")
	}

	if v := os.Getenv("HARBOR_USERNAME"); v == "" {
		t.Log(v)
		t.Fatal("HARBOR_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("HARBOR_PASSWORD"); v == "" {
		t.Log(v)
		t.Fatal("HARBOR_PASSWORD must be set for acceptance tests")
	}
}
