package client

import (
	"bytes"
	"fmt"
	"strings"
)

type SchemaCollection struct {
	Data []Schema
}

type Schemas struct {
	schemasByPath map[string]map[string]*Schema
	versions      []APIVersion
	schemas       []*Schema
	errors        []error
}

func NewSchemas() *Schemas {
	return &Schemas{
		schemasByPath: map[string]map[string]*Schema{},
	}
}

func (s *Schemas) Err() error {
	return NewErrors(s.errors)
}

func (s *Schemas) AddSchema(schema *Schema) *Schemas {
	schema.Type = "schema"
	if schema.ID == "" {
		s.errors = append(s.errors, fmt.Errorf("ID is not set on schema: %v", schema))
		return s
	}
	if schema.Version.Path == "" || schema.Version.Group == "" || schema.Version.Version == "" {
		s.errors = append(s.errors, fmt.Errorf("version is not set on schema: %s", schema.ID))
		return s
	}
	if schema.PluralName == "" {
		schema.PluralName = GuessPluralName(schema.ID)
	}
	if schema.CollectionMethods == nil {
		schema.CollectionMethods = []string{"GET", "POST"}
	}
	if schema.ResourceMethods == nil {
		schema.ResourceMethods = []string{"GET", "PUT", "DELETE"}
	}

	schemas, ok := s.schemasByPath[schema.Version.Path]
	if !ok {
		schemas = map[string]*Schema{}
		s.schemasByPath[schema.Version.Path] = schemas
		s.versions = append(s.versions, schema.Version)
	}

	if _, ok := schemas[schema.ID]; !ok {
		schemas[schema.ID] = schema
		s.schemas = append(s.schemas, schema)
	}

	return s
}

func (s *Schemas) SchemasForVersion(version APIVersion) map[string]*Schema {
	return s.schemasByPath[version.Path]
}

func (s *Schemas) Versions() []APIVersion {
	return s.versions
}

func (s *Schemas) Schemas() []*Schema {
	return s.schemas
}

func (s *Schemas) Schema(version *APIVersion, name string) *Schema {
	var (
		path string
	)

	if strings.Contains(name, "/") {
		idx := strings.LastIndex(name, "/")
		path = name[0:idx]
		name = name[idx+1:]
	} else if version != nil {
		path = version.Path
	} else {
		path = "core"
	}

	schemas, ok := s.schemasByPath[path]
	if !ok {
		return nil
	}

	schema := schemas[name]
	if schema != nil {
		return schema
	}

	for _, check := range schemas {
		if check.PluralName == name {
			return check
		}
	}

	return nil
}

type multiErrors struct {
	errors []error
}

func NewErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}
	return &multiErrors{
		errors: errors,
	}
}

func (m *multiErrors) Error() string {
	buf := bytes.NewBuffer(nil)
	for _, err := range m.errors {
		if buf.Len() > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(err.Error())
	}

	return buf.String()
}
