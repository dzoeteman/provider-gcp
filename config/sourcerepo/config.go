package sourcerepo

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {

	p.AddResourceConfigurator("google_sourcerepo_repository_iam_member", func(r *config.Resource) {
		// The reference should be manually inferred, but this is not yet activated for GCP
		r.References["repository"] = config.Reference{
			Type: "Repository",
		}
	})

}
