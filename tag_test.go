package tag_test

import (
	"testing"

	"github.com/oidc-proxy-ecosystem/go-tag"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Test  string `default:"test" validate:"Custom" tag_test:"value" tag2:"value=test"`
	Test2 string `default:"test" validate:"Custom" tag_test:"value" tag2:"value=test"`
	test2 string `default:"test" validate:"Custom" tag_test:"value" tag2:"value=test"`
}

func TestTag(t *testing.T) {
	test := &TestStruct{}
	tag.MustRegister("tag_test", "tag2")
	tagStructs, err := tag.New(test)
	assert.Equal(t, 3, len(tagStructs))
	assert.NoError(t, err)
	assert.Equal(t, "Test", tagStructs[0].Name)
	assert.Equal(t, "default", tagStructs[0].Tags[0].Key)
	assert.Equal(t, "validate", tagStructs[0].Tags[1].Key)
	assert.Equal(t, "tag_test", tagStructs[0].Tags[2].Key)
	assert.Equal(t, "tag2", tagStructs[0].Tags[3].Key)
	assert.Equal(t, "test", tagStructs[0].Tags[0].Value)
	assert.Equal(t, "Custom", tagStructs[0].Tags[1].Value)
	assert.Equal(t, "value", tagStructs[0].Tags[2].Value)
	assert.Equal(t, "value=test", tagStructs[0].Tags[3].Value)

	assert.Equal(t, "Test2", tagStructs[1].Name)
	assert.Equal(t, "default", tagStructs[1].Tags[0].Key)
	assert.Equal(t, "validate", tagStructs[1].Tags[1].Key)
	assert.Equal(t, "tag_test", tagStructs[1].Tags[2].Key)
	assert.Equal(t, "tag2", tagStructs[1].Tags[3].Key)
	assert.Equal(t, "test", tagStructs[1].Tags[0].Value)
	assert.Equal(t, "Custom", tagStructs[1].Tags[1].Value)
	assert.Equal(t, "value", tagStructs[1].Tags[2].Value)
	assert.Equal(t, "value=test", tagStructs[1].Tags[3].Value)

	assert.Equal(t, "test2", tagStructs[2].Name)
	assert.Equal(t, "default", tagStructs[2].Tags[0].Key)
	assert.Equal(t, "validate", tagStructs[2].Tags[1].Key)
	assert.Equal(t, "tag_test", tagStructs[2].Tags[2].Key)
	assert.Equal(t, "tag2", tagStructs[2].Tags[3].Key)
	assert.Equal(t, "test", tagStructs[2].Tags[0].Value)
	assert.Equal(t, "Custom", tagStructs[2].Tags[1].Value)
	assert.Equal(t, "value", tagStructs[2].Tags[2].Value)
	assert.Equal(t, "value=test", tagStructs[2].Tags[3].Value)
}
