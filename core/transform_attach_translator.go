package core

import (
	"github.com/hashicorp/hcl/v2"

	"bridgedl/graph"
	"bridgedl/translation"
)

// AttachableTranslatorVertex is implemented by all types used as graph.Vertex
// that can have a translation.BlockTranslator attached.
type AttachableTranslatorVertex interface {
	TranslatorFinder
	AttachTranslator(translation.BlockTranslator)
}

// TranslatorFinder can look up a suitable translation.BlockTranslator in a
// collection of Translators.
type TranslatorFinder interface {
	FindTranslator(*Translators) (translation.BlockTranslator, hcl.Diagnostics)
}

// AttachTranslatorsTransformer is a GraphTransformer that attaches a block
// translator to all graph vertices that support it.
type AttachTranslatorsTransformer struct {
	Translators *Translators
}

var _ GraphTransformer = (*AttachTranslatorsTransformer)(nil)

// Transform implements GraphTransformer.
func (t *AttachTranslatorsTransformer) Transform(g *graph.DirectedGraph) hcl.Diagnostics {
	var diags hcl.Diagnostics

	for _, v := range g.Vertices() {
		attch, ok := v.(AttachableTranslatorVertex)
		if !ok {
			continue
		}

		transl, translDiags := attch.FindTranslator(t.Translators)
		diags = diags.Extend(translDiags)
		attch.AttachTranslator(transl)
	}

	return diags
}
