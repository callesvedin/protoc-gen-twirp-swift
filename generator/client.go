package generator

import (
	"bytes"
	_ "embed" // Used by go:embed
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"unicode"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"google.golang.org/protobuf/proto"

	plugin_go "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
)

//go:embed templates/TwirpClient.swift
var apiTemplate string

type model struct {
	Name         string
	Primitive    bool
	Fields       []modelField
	CanMarshal   bool
	CanUnmarshal bool
}

type modelField struct {
	Name          string
	Type          string
	InternalType  string
	JSONName      string
	JSONType      string
	IsMessage     bool
	IsRepeated    bool
	IsMap         bool
	MapKeyField   *modelField
	MapValueField *modelField
}

type service struct {
	Name                  string
	ClassName             string
	ServicePathPrefixName string
	Package               string
	Methods               []serviceMethod
}

type serviceMethod struct {
	Name       string
	Path       string
	InputArg   string
	InputType  string
	OutputType string
}

func newAPIContext() apiContext {
	ctx := apiContext{}
	ctx.modelLookup = make(map[string]*model)

	return ctx
}

type apiContext struct {
	Models      []*model
	Services    []*service
	Imports     []importPath
	ClientName  string
	modelLookup map[string]*model
}

type importPath struct {
	Path string
}

func (ctx *apiContext) addModel(m *model) {
	ctx.Models = append(ctx.Models, m)
	ctx.modelLookup[m.Name] = m
}

func (ctx *apiContext) applyImports(d *descriptor.FileDescriptorProto) {
	var deps []importPath

	if len(ctx.Services) > 0 {
		deps = append(deps, importPath{"Foundation"})
		deps = append(deps, importPath{"SwiftProtobuf"})
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
		// deps = append(deps, importPath{fullPath})
	}
	ctx.Imports = deps
}

func getClassName(d *descriptor.FileDescriptorProto, pkg string, sName string) string {
	if d.Options.SwiftPrefix != nil {
		return *d.Options.SwiftPrefix + sName
	}
	return strings.ToUpper(pkg[:1]) + pkg[1:] + sName
}

func getClientName(d *descriptor.FileDescriptorProto, pkg string) string {
	if d.Options.SwiftPrefix != nil {
		return *d.Options.SwiftPrefix + "Client"
	}
	return strings.ToUpper(pkg[:1]) + pkg[1:] + "Client"
}

// CreateClientAPI generates the client api from the proto file
func CreateClientAPI(d *descriptor.FileDescriptorProto, generator *generator.Generator) (*plugin_go.CodeGeneratorResponse_File, error) {
	ctx := newAPIContext()
	pkg := d.GetPackage()
	ctx.ClientName = getClientName(d, pkg)

	// Parse all Services for generating swift method interfaces and default client implementations
	for _, s := range d.GetService() {
		serviceName := s.GetName()
		service := &service{
			Name:                  serviceName,
			ClassName:             getClassName(d, pkg, serviceName),
			Package:               pkg,
			ServicePathPrefixName: strings.ToLower(serviceName[0:1]) + serviceName[1:] + "ServicePrefix",
		}

		for _, m := range s.GetMethod() {
			m.GetInputType()
			methodPath := m.GetName()
			methodName := strings.ToLower(methodPath[0:1]) + methodPath[1:]
			in := removePkg(m.GetInputType())
			arg := strings.ToLower(in[0:1]) + in[1:]
			in = getTypeName(d, m.GetInputType())

			method := serviceMethod{
				Name:       methodName,
				Path:       methodPath,
				InputArg:   arg,
				InputType:  in,
				OutputType: getTypeName(d, m.GetOutputType()),
			}

			service.Methods = append(service.Methods, method)
		}
		ctx.Services = append(ctx.Services, service)
	}

	ctx.applyImports(d)

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

func getTypeName(d *descriptor.FileDescriptorProto, s string) string {
	if d.Options.SwiftPrefix != nil {
		return *d.Options.SwiftPrefix + removePkg(s)
	}
	return toSnake(s[1:])
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
