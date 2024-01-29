package openapi3

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCallbackRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := CallbackRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Callback{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestExampleRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := ExampleRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Example{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestHeaderRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := HeaderRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)
	// Header does not have its own extensions.
}

func TestLinkRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := LinkRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Link{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestParameterRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := ParameterRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Parameter{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestRequestBodyRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := RequestBodyRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &RequestBody{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestResponseRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := ResponseRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Response{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestSchemaRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := SchemaRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &Schema{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}

func TestSecuritySchemeRef_Extensions(t *testing.T) {
	data := []byte(`{"$ref":"#/components/schemas/Pet","something":"integer","x-order":1}`)

	ref := SecuritySchemeRef{}
	err := json.Unmarshal(data, &ref)
	assert.NoError(t, err)

	// captures extension
	assert.Equal(t, "#/components/schemas/Pet", ref.Ref)
	assert.Equal(t, float64(1), ref.Extensions["x-order"])

	// does not capture non-extensions
	assert.Nil(t, ref.Extensions["something"])

	// validation
	err = ref.Validate(context.Background())
	require.EqualError(t, err, "extra sibling fields: [something]")

	err = ref.Validate(context.Background(), ProhibitExtensionsWithRef())
	require.EqualError(t, err, "extra sibling fields: [something x-order]")

	err = ref.Validate(context.Background(), AllowExtraSiblingFields("something"))
	assert.ErrorContains(t, err, "found unresolved ref") // expected since value not defined

	// non-extension not json lookable
	_, err = ref.JSONLookup("something")
	assert.Error(t, err)

	t.Run("extentions in value", func(t *testing.T) {
		ref.Value = &SecurityScheme{Extensions: map[string]interface{}{}}
		ref.Value.Extensions["x-order"] = 2.0

		// prefers the value next to the $ref over the one in the $ref.
		v, err := ref.JSONLookup("x-order")
		assert.NoError(t, err)
		assert.Equal(t, float64(1), v)
	})
}
