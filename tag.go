package tag

import (
	"errors"
	"reflect"
)

var tags = []string{"default", "validate"}
var mpTag = map[string]struct{}{
	"default":  {},
	"validate": {},
}
var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

// Register is getter struct tag
func Register(tagNames ...string) error {
	for _, tagName := range tagNames {
		if _, ok := mpTag[tagName]; !ok {
			tags = append(tags, tagName)
			mpTag[tagName] = struct{}{}
		}
	}
	return nil
}

// MustRegister is wrapper Register of ignore at error
func MustRegister(tagName ...string) {
	_ = Register(tagName...)
}

type Tag struct {
	Key   string
	Value string
}

type StructTag struct {
	Name  string
	Field reflect.Value
	Tags  []Tag
}

// New is spec is pointer struct
func New(spec interface{}) ([]StructTag, error) {
	s := reflect.ValueOf(spec)
	if s.Kind() != reflect.Ptr {
		return nil, ErrInvalidSpecification
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, ErrInvalidSpecification
	}
	rTyp := s.Type()
	infos := make([]StructTag, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fType := rTyp.Field(i)
		strTag := StructTag{
			Name:  fType.Name,
			Field: field,
			Tags:  []Tag{},
		}
		for _, tag := range tags {
			value := fType.Tag.Get(tag)
			if value == "" {
				continue
			}
			strTag.Tags = append(strTag.Tags, Tag{
				Key:   tag,
				Value: value,
			})
		}
		infos[i] = strTag
	}
	return infos, nil
}
