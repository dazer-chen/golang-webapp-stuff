package restful

import (
	"testing"
)

func TestRouteBuilder_PathParameter(t *testing.T) {
	p := &Parameter{&ParameterData{Name: "name", Description: "desc"}}
	p.AllowMultiple(true)
	p.DataType("int")
	p.Required(true)
	values := map[string]string{"a": "b"}
	p.AllowableValues(values)
	p.bePath()

	b := new(RouteBuilder)
	b.Param(p)
	r := b.Build()
	if !r.ParameterDocs[0].Data().AllowMultiple {
		t.Error("AllowMultiple invalid")
	}
	if r.ParameterDocs[0].Data().DataType != "int" {
		t.Error("dataType invalid")
	}
	if !r.ParameterDocs[0].Data().Required {
		t.Error("required invalid")
	}
	if r.ParameterDocs[0].Data().Kind != PATH_PARAMETER {
		t.Error("kind invalid")
	}
	if r.ParameterDocs[0].Data().AllowableValues["a"] != "b" {
		t.Error("allowableValues invalid")
	}
}

func TestRouteBuilder(t *testing.T) {
	json := "application/json"
	b := new(RouteBuilder)
	b.Path("/routes").Method("HEAD").Consumes(json).Produces(json)
	r := b.Build()
	if r.Path != "/routes" {
		t.Error("path invalid")
	}
	if r.Produces[0] != json {
		t.Error("produces invalid")
	}
	if r.Consumes[0] != json {
		t.Error("consumes invalid")
	}
}
