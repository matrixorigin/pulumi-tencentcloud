// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OrgShareUnitMember struct {
	pulumi.CustomResourceState

	// Shared unit region.
	Area pulumi.StringOutput `pulumi:"area"`
	// Shared member list. Up to 10 items.
	Members OrgShareUnitMemberMemberArrayOutput `pulumi:"members"`
	// Shared unit ID.
	UnitId pulumi.StringOutput `pulumi:"unitId"`
}

// NewOrgShareUnitMember registers a new resource with the given unique name, arguments, and options.
func NewOrgShareUnitMember(ctx *pulumi.Context,
	name string, args *OrgShareUnitMemberArgs, opts ...pulumi.ResourceOption) (*OrgShareUnitMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Area == nil {
		return nil, errors.New("invalid value for required argument 'Area'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.UnitId == nil {
		return nil, errors.New("invalid value for required argument 'UnitId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgShareUnitMember
	err := ctx.RegisterResource("tencentcloud:Organization/orgShareUnitMember:OrgShareUnitMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgShareUnitMember gets an existing OrgShareUnitMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgShareUnitMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgShareUnitMemberState, opts ...pulumi.ResourceOption) (*OrgShareUnitMember, error) {
	var resource OrgShareUnitMember
	err := ctx.ReadResource("tencentcloud:Organization/orgShareUnitMember:OrgShareUnitMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgShareUnitMember resources.
type orgShareUnitMemberState struct {
	// Shared unit region.
	Area *string `pulumi:"area"`
	// Shared member list. Up to 10 items.
	Members []OrgShareUnitMemberMember `pulumi:"members"`
	// Shared unit ID.
	UnitId *string `pulumi:"unitId"`
}

type OrgShareUnitMemberState struct {
	// Shared unit region.
	Area pulumi.StringPtrInput
	// Shared member list. Up to 10 items.
	Members OrgShareUnitMemberMemberArrayInput
	// Shared unit ID.
	UnitId pulumi.StringPtrInput
}

func (OrgShareUnitMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgShareUnitMemberState)(nil)).Elem()
}

type orgShareUnitMemberArgs struct {
	// Shared unit region.
	Area string `pulumi:"area"`
	// Shared member list. Up to 10 items.
	Members []OrgShareUnitMemberMember `pulumi:"members"`
	// Shared unit ID.
	UnitId string `pulumi:"unitId"`
}

// The set of arguments for constructing a OrgShareUnitMember resource.
type OrgShareUnitMemberArgs struct {
	// Shared unit region.
	Area pulumi.StringInput
	// Shared member list. Up to 10 items.
	Members OrgShareUnitMemberMemberArrayInput
	// Shared unit ID.
	UnitId pulumi.StringInput
}

func (OrgShareUnitMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgShareUnitMemberArgs)(nil)).Elem()
}

type OrgShareUnitMemberInput interface {
	pulumi.Input

	ToOrgShareUnitMemberOutput() OrgShareUnitMemberOutput
	ToOrgShareUnitMemberOutputWithContext(ctx context.Context) OrgShareUnitMemberOutput
}

func (*OrgShareUnitMember) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgShareUnitMember)(nil)).Elem()
}

func (i *OrgShareUnitMember) ToOrgShareUnitMemberOutput() OrgShareUnitMemberOutput {
	return i.ToOrgShareUnitMemberOutputWithContext(context.Background())
}

func (i *OrgShareUnitMember) ToOrgShareUnitMemberOutputWithContext(ctx context.Context) OrgShareUnitMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitMemberOutput)
}

// OrgShareUnitMemberArrayInput is an input type that accepts OrgShareUnitMemberArray and OrgShareUnitMemberArrayOutput values.
// You can construct a concrete instance of `OrgShareUnitMemberArrayInput` via:
//
//	OrgShareUnitMemberArray{ OrgShareUnitMemberArgs{...} }
type OrgShareUnitMemberArrayInput interface {
	pulumi.Input

	ToOrgShareUnitMemberArrayOutput() OrgShareUnitMemberArrayOutput
	ToOrgShareUnitMemberArrayOutputWithContext(context.Context) OrgShareUnitMemberArrayOutput
}

type OrgShareUnitMemberArray []OrgShareUnitMemberInput

func (OrgShareUnitMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgShareUnitMember)(nil)).Elem()
}

func (i OrgShareUnitMemberArray) ToOrgShareUnitMemberArrayOutput() OrgShareUnitMemberArrayOutput {
	return i.ToOrgShareUnitMemberArrayOutputWithContext(context.Background())
}

func (i OrgShareUnitMemberArray) ToOrgShareUnitMemberArrayOutputWithContext(ctx context.Context) OrgShareUnitMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitMemberArrayOutput)
}

// OrgShareUnitMemberMapInput is an input type that accepts OrgShareUnitMemberMap and OrgShareUnitMemberMapOutput values.
// You can construct a concrete instance of `OrgShareUnitMemberMapInput` via:
//
//	OrgShareUnitMemberMap{ "key": OrgShareUnitMemberArgs{...} }
type OrgShareUnitMemberMapInput interface {
	pulumi.Input

	ToOrgShareUnitMemberMapOutput() OrgShareUnitMemberMapOutput
	ToOrgShareUnitMemberMapOutputWithContext(context.Context) OrgShareUnitMemberMapOutput
}

type OrgShareUnitMemberMap map[string]OrgShareUnitMemberInput

func (OrgShareUnitMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgShareUnitMember)(nil)).Elem()
}

func (i OrgShareUnitMemberMap) ToOrgShareUnitMemberMapOutput() OrgShareUnitMemberMapOutput {
	return i.ToOrgShareUnitMemberMapOutputWithContext(context.Background())
}

func (i OrgShareUnitMemberMap) ToOrgShareUnitMemberMapOutputWithContext(ctx context.Context) OrgShareUnitMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgShareUnitMemberMapOutput)
}

type OrgShareUnitMemberOutput struct{ *pulumi.OutputState }

func (OrgShareUnitMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgShareUnitMember)(nil)).Elem()
}

func (o OrgShareUnitMemberOutput) ToOrgShareUnitMemberOutput() OrgShareUnitMemberOutput {
	return o
}

func (o OrgShareUnitMemberOutput) ToOrgShareUnitMemberOutputWithContext(ctx context.Context) OrgShareUnitMemberOutput {
	return o
}

// Shared unit region.
func (o OrgShareUnitMemberOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgShareUnitMember) pulumi.StringOutput { return v.Area }).(pulumi.StringOutput)
}

// Shared member list. Up to 10 items.
func (o OrgShareUnitMemberOutput) Members() OrgShareUnitMemberMemberArrayOutput {
	return o.ApplyT(func(v *OrgShareUnitMember) OrgShareUnitMemberMemberArrayOutput { return v.Members }).(OrgShareUnitMemberMemberArrayOutput)
}

// Shared unit ID.
func (o OrgShareUnitMemberOutput) UnitId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgShareUnitMember) pulumi.StringOutput { return v.UnitId }).(pulumi.StringOutput)
}

type OrgShareUnitMemberArrayOutput struct{ *pulumi.OutputState }

func (OrgShareUnitMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgShareUnitMember)(nil)).Elem()
}

func (o OrgShareUnitMemberArrayOutput) ToOrgShareUnitMemberArrayOutput() OrgShareUnitMemberArrayOutput {
	return o
}

func (o OrgShareUnitMemberArrayOutput) ToOrgShareUnitMemberArrayOutputWithContext(ctx context.Context) OrgShareUnitMemberArrayOutput {
	return o
}

func (o OrgShareUnitMemberArrayOutput) Index(i pulumi.IntInput) OrgShareUnitMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgShareUnitMember {
		return vs[0].([]*OrgShareUnitMember)[vs[1].(int)]
	}).(OrgShareUnitMemberOutput)
}

type OrgShareUnitMemberMapOutput struct{ *pulumi.OutputState }

func (OrgShareUnitMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgShareUnitMember)(nil)).Elem()
}

func (o OrgShareUnitMemberMapOutput) ToOrgShareUnitMemberMapOutput() OrgShareUnitMemberMapOutput {
	return o
}

func (o OrgShareUnitMemberMapOutput) ToOrgShareUnitMemberMapOutputWithContext(ctx context.Context) OrgShareUnitMemberMapOutput {
	return o
}

func (o OrgShareUnitMemberMapOutput) MapIndex(k pulumi.StringInput) OrgShareUnitMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgShareUnitMember {
		return vs[0].(map[string]*OrgShareUnitMember)[vs[1].(string)]
	}).(OrgShareUnitMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitMemberInput)(nil)).Elem(), &OrgShareUnitMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitMemberArrayInput)(nil)).Elem(), OrgShareUnitMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgShareUnitMemberMapInput)(nil)).Elem(), OrgShareUnitMemberMap{})
	pulumi.RegisterOutputType(OrgShareUnitMemberOutput{})
	pulumi.RegisterOutputType(OrgShareUnitMemberArrayOutput{})
	pulumi.RegisterOutputType(OrgShareUnitMemberMapOutput{})
}