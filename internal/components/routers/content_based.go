package routers

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"

	"bridgedl/k8s"
	"bridgedl/translation"
)

type ContentBased struct{}

var (
	_ translation.Decodable    = (*ContentBased)(nil)
	_ translation.Translatable = (*ContentBased)(nil)
	_ translation.Addressable  = (*ContentBased)(nil)
)

// Spec implements translation.Decodable.
func (*ContentBased) Spec() hcldec.Spec {
	// NOTE(antoineco): see the following implementation to get a sense of
	// how HCL blocks map to hcldec.Specs and cty.Types:
	// https://pkg.go.dev/github.com/hashicorp/terraform@v0.14.7/configs/configschema#Block.DecoderSpec
	return &hcldec.ObjectSpec{
		"route": &hcldec.BlockSetSpec{
			TypeName: "route",
			Nested: &hcldec.ObjectSpec{
				"attributes": &hcldec.AttrSpec{
					Name:     "attributes",
					Type:     cty.Map(cty.String),
					Required: true,
				},
				"to": &hcldec.AttrSpec{
					Name:     "to",
					Type:     k8s.DestinationCty,
					Required: true,
				},
			},
			MinItems: 1,
		},
	}
}

// Manifests implements translation.Translatable.
func (*ContentBased) Manifests(id string, config cty.Value) []interface{} {
	panic("not implemented")
}

// Address implements translation.Addressable.
func (*ContentBased) Address(id string) cty.Value {
	return k8s.NewDestination("eventing.knative.dev/v1", "Broker", k8s.RFC1123Name(id))
}