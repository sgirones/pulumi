// *** WARNING: this file was generated by the Coconut IDL Compiler (CIDLC).  ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/coconut/pkg/resource"
    "github.com/pulumi/coconut/pkg/tokens"
    "github.com/pulumi/coconut/pkg/util/contract"
    "github.com/pulumi/coconut/pkg/util/mapper"
    "github.com/pulumi/coconut/sdk/go/pkg/cocorpc"
)

/* RPC stubs for Model resource provider */

// ModelToken is the type token corresponding to the Model package type.
const ModelToken = tokens.Type("aws:apigateway/model:Model")

// ModelProviderOps is a pluggable interface for Model-related management functionality.
type ModelProviderOps interface {
    Check(ctx context.Context, obj *Model) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Model) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Model, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Model, new *Model, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Model, new *Model, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// ModelProvider is a dynamic gRPC-based plugin for managing Model resources.
type ModelProvider struct {
    ops ModelProviderOps
}

// NewModelProvider allocates a resource provider that delegates to a ops instance.
func NewModelProvider(ops ModelProviderOps) cocorpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &ModelProvider{ops: ops}
}

func (p *ModelProvider) Check(
    ctx context.Context, req *cocorpc.CheckRequest) (*cocorpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *ModelProvider) Name(
    ctx context.Context, req *cocorpc.NameRequest) (*cocorpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &cocorpc.NameResponse{Name: obj.Name}, nil
}

func (p *ModelProvider) Create(
    ctx context.Context, req *cocorpc.CreateRequest) (*cocorpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &cocorpc.CreateResponse{
        Id:   string(id),
    }, nil
}

func (p *ModelProvider) Get(
    ctx context.Context, req *cocorpc.GetRequest) (*cocorpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &cocorpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *ModelProvider) InspectChange(
    ctx context.Context, req *cocorpc.ChangeRequest) (*cocorpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("contentType") {
            replaces = append(replaces, "contentType")
        }
        if diff.Changed("restAPI") {
            replaces = append(replaces, "restAPI")
        }
        if diff.Changed("modelName") {
            replaces = append(replaces, "modelName")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &cocorpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *ModelProvider) Update(
    ctx context.Context, req *cocorpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ModelProvider) Delete(
    ctx context.Context, req *cocorpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ModelToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ModelProvider) Unmarshal(
    v *pbstruct.Struct) (*Model, resource.PropertyMap, mapper.DecodeError) {
    var obj Model
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Model structure(s) */

// Model is a marshalable representation of its corresponding IDL type.
type Model struct {
    Name string `json:"name"`
    ContentType string `json:"contentType"`
    RestAPI resource.ID `json:"restAPI"`
    Schema interface{} `json:"schema"`
    ModelName *string `json:"modelName,omitempty"`
    Description *string `json:"description,omitempty"`
}

// Model's properties have constants to make dealing with diffs and property bags easier.
const (
    Model_Name = "name"
    Model_ContentType = "contentType"
    Model_RestAPI = "restAPI"
    Model_Schema = "schema"
    Model_ModelName = "modelName"
    Model_Description = "description"
)


