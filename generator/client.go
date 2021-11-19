package generator

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"

	plugin_go "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

type Model struct {
	Name         string
	Primitive    bool
	Fields       []ModelField
	CanMarshal   bool
	CanUnmarshal bool
}

type ModelField struct {
	Name          string
	Type          string
	InternalType  string
	JSONName      string
	JSONType      string
	IsMessage     bool
	IsRepeated    bool
	IsMap         bool
	MapKeyField   *ModelField
	MapValueField *ModelField
}

type Service struct {
	Name                  string
	ServicePathPrefixName string
	Package               string
	Methods               []ServiceMethod
}

func (s Service) ClassName() string {
	return strings.ToUpper(s.Package[:1]) + s.Package[1:] + s.Name
}

type ServiceMethod struct {
	Name       string
	Path       string
	InputArg   string
	InputType  string
	OutputType string
}

func NewAPIContext() APIContext {
	ctx := APIContext{}
	ctx.modelLookup = make(map[string]*Model)

	return ctx
}

type APIContext struct {
	Models      []*Model
	Services    []*Service
	Imports     []Import
	ClientName  string
	modelLookup map[string]*Model
}

type Import struct {
	Path string
}

func (ctx *APIContext) AddModel(m *Model) {
	ctx.Models = append(ctx.Models, m)
	ctx.modelLookup[m.Name] = m
}

func (ctx *APIContext) ApplyImports(d *descriptor.FileDescriptorProto) {
	var deps []Import

	if len(ctx.Services) > 0 {
		deps = append(deps, Import{"Foundation"})
		deps = append(deps, Import{"SwiftProtobuf"})
	}

	for _, dep := range d.Dependency {
		if dep == "google/protobuf/timestamp.proto" {
			continue
		}
		importPath := path.Dir(dep)
		sourceDir := path.Dir(*d.Name)
		sourceComponents := strings.Split(sourceDir, fmt.Sprintf("%c", os.PathSeparator))
		distanceFromRoot := len(sourceComponents)
		for _, pathComponent := range sourceComponents {
			if strings.HasPrefix(importPath, pathComponent) {
				importPath = strings.TrimPrefix(importPath, pathComponent)
				distanceFromRoot--
			}
		}
		fileName := swiftFilename(dep)
		fullPath := fileName
		fullPath = path.Join(importPath, fullPath)
		if distanceFromRoot > 0 {
			for i := 0; i < distanceFromRoot; i++ {
				fullPath = path.Join("..", fullPath)
			}
		}
		// Do not generate model which should do by protoc-gen-swift
		// deps = append(deps, Import{fullPath})
	}
	ctx.Imports = deps
}

func CreateClientAPI(d *descriptor.FileDescriptorProto, generator *generator.Generator) (*plugin_go.CodeGeneratorResponse_File, error) {
	ctx := NewAPIContext()
	pkg := d.GetPackage()
	ctx.ClientName = strings.ToUpper(pkg[:1]) + pkg[1:] + "Client"

	// Parse all Services for generating swift method interfaces and default client implementations
	for _, s := range d.GetService() {
		serviceName := s.GetName()
		service := &Service{
			Name:                  serviceName,
			Package:               pkg,
			ServicePathPrefixName: strings.ToLower(serviceName[0:1]) + serviceName[1:] + "ServicePrefix",
		}

		for _, m := range s.GetMethod() {
			methodPath := m.GetName()
			methodName := strings.ToLower(methodPath[0:1]) + methodPath[1:]
			in := removePkg(m.GetInputType())
			arg := strings.ToLower(in[0:1]) + in[1:]
			in = toSnake(m.GetInputType()[1:])

			method := ServiceMethod{
				Name:       methodName,
				Path:       methodPath,
				InputArg:   arg,
				InputType:  in,
				OutputType: toSnake(m.GetOutputType()[1:]),
			}

			service.Methods = append(service.Methods, method)
		}
		ctx.Services = append(ctx.Services, service)
	}

	ctx.ApplyImports(d)

	fileBytes, err := ioutil.ReadFile("./templates/TwirpClient.swift")

	if err != nil {
		panic(err)
	}
	apiTemplate := string(fileBytes)

	t, err := template.New("client_api").Parse(apiTemplate)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBufferString("")
	err = t.Execute(b, ctx)
	if err != nil {
		return nil, err
	}

	cf := &plugin_go.CodeGeneratorResponse_File{}
	//cf.Name = proto.String(ctx.ClientName+".swift")
	cf.Name = proto.String(swiftModuleFilename(d))
	cf.Content = proto.String(b.String())

	return cf, nil
}

func removePkg(s string) string {
	p := strings.Split(s, ".")
	return p[len(p)-1]
}

func toSnake(s string) string {
	ss := strings.Split(s, ".")

	res := []string{}
	for _, v := range ss {
		// first to upper
		res = append(res, string(byte(unicode.ToUpper(rune(v[0]))))+v[1:])
	}
	return strings.Join(res, "_")
}

func camelCase(s string) string {
	parts := strings.Split(s, "_")

	for i, p := range parts {
		if i == 0 {
			parts[i] = p
		} else {
			parts[i] = strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
		}
	}

	return strings.Join(parts, "")
}
