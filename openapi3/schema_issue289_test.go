package openapi3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIssue289(t *testing.T) {
	spec := []byte(`
components:
  schemas:
    Server:
      properties:
        address:
          oneOf:
          - $ref: "#/components/schemas/ip-address"
          - $ref: "#/components/schemas/domain-name"
        name:
          type: string
      type: object
    domain-name:
      maxLength: 10
      minLength: 5
      pattern: "((([a-zA-Z0-9_]([a-zA-Z0-9\\-_]){0,61})?[a-zA-Z0-9]\\.)*([a-zA-Z0-9_]([a-zA-Z0-9\\-_]){0,61})?[a-zA-Z0-9]\\.?)|\\."
      type: string
    ip-address:
      pattern: "^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"
      type: string
openapi: "3.0.1"
info:
  version: 1.0.0
  title: title
paths: {}
`[1:])

	loader := NewLoader()
	doc, err := loader.LoadFromData(spec)
	require.NoError(t, err)

	err = doc.Validate(loader.Context)
	require.NoError(t, err)

	err = doc.Components.Schemas["Server"].Value.VisitJSON(map[string]any{
		"name":    "kin-openapi",
		"address": "127.0.0.1",
	})
	require.ErrorIs(t, err, ErrOneOfConflict)
}
