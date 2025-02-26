package datacatalog

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_data_catalog_entry_group", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
}
