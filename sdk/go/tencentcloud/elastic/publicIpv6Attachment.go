// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"errors"
	"github.com/matrixorigin/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicIpv6Attachment struct {
	pulumi.CustomResourceState

	// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
	Ipv6AddressId pulumi.StringOutput `pulumi:"ipv6AddressId"`
	// Whether to keep the Elastic Network Interface bound when unbinding.
	KeepBindWithEni pulumi.BoolPtrOutput `pulumi:"keepBindWithEni"`
	// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
	// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
	// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
	NetworkInterfaceId pulumi.StringPtrOutput `pulumi:"networkInterfaceId"`
	// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
	// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
	// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
	// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
	// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
	PrivateIpv6Address pulumi.StringPtrOutput `pulumi:"privateIpv6Address"`
}

// NewPublicIpv6Attachment registers a new resource with the given unique name, arguments, and options.
func NewPublicIpv6Attachment(ctx *pulumi.Context,
	name string, args *PublicIpv6AttachmentArgs, opts ...pulumi.ResourceOption) (*PublicIpv6Attachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ipv6AddressId == nil {
		return nil, errors.New("invalid value for required argument 'Ipv6AddressId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PublicIpv6Attachment
	err := ctx.RegisterResource("tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicIpv6Attachment gets an existing PublicIpv6Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIpv6Attachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicIpv6AttachmentState, opts ...pulumi.ResourceOption) (*PublicIpv6Attachment, error) {
	var resource PublicIpv6Attachment
	err := ctx.ReadResource("tencentcloud:Elastic/publicIpv6Attachment:PublicIpv6Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicIpv6Attachment resources.
type publicIpv6AttachmentState struct {
	// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
	Ipv6AddressId *string `pulumi:"ipv6AddressId"`
	// Whether to keep the Elastic Network Interface bound when unbinding.
	KeepBindWithEni *bool `pulumi:"keepBindWithEni"`
	// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
	// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
	// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
	// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
	// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
	// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
	// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
	PrivateIpv6Address *string `pulumi:"privateIpv6Address"`
}

type PublicIpv6AttachmentState struct {
	// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
	Ipv6AddressId pulumi.StringPtrInput
	// Whether to keep the Elastic Network Interface bound when unbinding.
	KeepBindWithEni pulumi.BoolPtrInput
	// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
	// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
	// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
	NetworkInterfaceId pulumi.StringPtrInput
	// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
	// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
	// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
	// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
	// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
	PrivateIpv6Address pulumi.StringPtrInput
}

func (PublicIpv6AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpv6AttachmentState)(nil)).Elem()
}

type publicIpv6AttachmentArgs struct {
	// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
	Ipv6AddressId string `pulumi:"ipv6AddressId"`
	// Whether to keep the Elastic Network Interface bound when unbinding.
	KeepBindWithEni *bool `pulumi:"keepBindWithEni"`
	// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
	// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
	// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
	// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
	// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
	// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
	// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
	PrivateIpv6Address *string `pulumi:"privateIpv6Address"`
}

// The set of arguments for constructing a PublicIpv6Attachment resource.
type PublicIpv6AttachmentArgs struct {
	// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
	Ipv6AddressId pulumi.StringInput
	// Whether to keep the Elastic Network Interface bound when unbinding.
	KeepBindWithEni pulumi.BoolPtrInput
	// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
	// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
	// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
	NetworkInterfaceId pulumi.StringPtrInput
	// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
	// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
	// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
	// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
	// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
	PrivateIpv6Address pulumi.StringPtrInput
}

func (PublicIpv6AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicIpv6AttachmentArgs)(nil)).Elem()
}

type PublicIpv6AttachmentInput interface {
	pulumi.Input

	ToPublicIpv6AttachmentOutput() PublicIpv6AttachmentOutput
	ToPublicIpv6AttachmentOutputWithContext(ctx context.Context) PublicIpv6AttachmentOutput
}

func (*PublicIpv6Attachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpv6Attachment)(nil)).Elem()
}

func (i *PublicIpv6Attachment) ToPublicIpv6AttachmentOutput() PublicIpv6AttachmentOutput {
	return i.ToPublicIpv6AttachmentOutputWithContext(context.Background())
}

func (i *PublicIpv6Attachment) ToPublicIpv6AttachmentOutputWithContext(ctx context.Context) PublicIpv6AttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpv6AttachmentOutput)
}

// PublicIpv6AttachmentArrayInput is an input type that accepts PublicIpv6AttachmentArray and PublicIpv6AttachmentArrayOutput values.
// You can construct a concrete instance of `PublicIpv6AttachmentArrayInput` via:
//
//	PublicIpv6AttachmentArray{ PublicIpv6AttachmentArgs{...} }
type PublicIpv6AttachmentArrayInput interface {
	pulumi.Input

	ToPublicIpv6AttachmentArrayOutput() PublicIpv6AttachmentArrayOutput
	ToPublicIpv6AttachmentArrayOutputWithContext(context.Context) PublicIpv6AttachmentArrayOutput
}

type PublicIpv6AttachmentArray []PublicIpv6AttachmentInput

func (PublicIpv6AttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIpv6Attachment)(nil)).Elem()
}

func (i PublicIpv6AttachmentArray) ToPublicIpv6AttachmentArrayOutput() PublicIpv6AttachmentArrayOutput {
	return i.ToPublicIpv6AttachmentArrayOutputWithContext(context.Background())
}

func (i PublicIpv6AttachmentArray) ToPublicIpv6AttachmentArrayOutputWithContext(ctx context.Context) PublicIpv6AttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpv6AttachmentArrayOutput)
}

// PublicIpv6AttachmentMapInput is an input type that accepts PublicIpv6AttachmentMap and PublicIpv6AttachmentMapOutput values.
// You can construct a concrete instance of `PublicIpv6AttachmentMapInput` via:
//
//	PublicIpv6AttachmentMap{ "key": PublicIpv6AttachmentArgs{...} }
type PublicIpv6AttachmentMapInput interface {
	pulumi.Input

	ToPublicIpv6AttachmentMapOutput() PublicIpv6AttachmentMapOutput
	ToPublicIpv6AttachmentMapOutputWithContext(context.Context) PublicIpv6AttachmentMapOutput
}

type PublicIpv6AttachmentMap map[string]PublicIpv6AttachmentInput

func (PublicIpv6AttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIpv6Attachment)(nil)).Elem()
}

func (i PublicIpv6AttachmentMap) ToPublicIpv6AttachmentMapOutput() PublicIpv6AttachmentMapOutput {
	return i.ToPublicIpv6AttachmentMapOutputWithContext(context.Background())
}

func (i PublicIpv6AttachmentMap) ToPublicIpv6AttachmentMapOutputWithContext(ctx context.Context) PublicIpv6AttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpv6AttachmentMapOutput)
}

type PublicIpv6AttachmentOutput struct{ *pulumi.OutputState }

func (PublicIpv6AttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpv6Attachment)(nil)).Elem()
}

func (o PublicIpv6AttachmentOutput) ToPublicIpv6AttachmentOutput() PublicIpv6AttachmentOutput {
	return o
}

func (o PublicIpv6AttachmentOutput) ToPublicIpv6AttachmentOutputWithContext(ctx context.Context) PublicIpv6AttachmentOutput {
	return o
}

// Elastic IPv6 unique ID, EIPv6 unique ID is like eipv6-11112222.
func (o PublicIpv6AttachmentOutput) Ipv6AddressId() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicIpv6Attachment) pulumi.StringOutput { return v.Ipv6AddressId }).(pulumi.StringOutput)
}

// Whether to keep the Elastic Network Interface bound when unbinding.
func (o PublicIpv6AttachmentOutput) KeepBindWithEni() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PublicIpv6Attachment) pulumi.BoolPtrOutput { return v.KeepBindWithEni }).(pulumi.BoolPtrOutput)
}

// Elastic Network Interface ID to bind. Elastic Network Interface ID is like eni-11112222. NetworkInterfaceId and
// InstanceId cannot be specified simultaneously. The Elastic Network Interface ID can be queried by logging in to the
// console, or obtained through the networkInterfaceId in the return value of the DescribeNetworkInterfaces interface.
func (o PublicIpv6AttachmentOutput) NetworkInterfaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpv6Attachment) pulumi.StringPtrOutput { return v.NetworkInterfaceId }).(pulumi.StringPtrOutput)
}

// The intranet IPv6 to bind. If NetworkInterfaceId is specified, PrivateIPv6Address must also be specified, which means
// that the EIP is bound to the specified private network IP of the specified Elastic Network Interface. Also ensure that
// the specified PrivateIPv6Address is an intranet IPv6 on the specified NetworkInterfaceId. The intranet IPv6 of the
// specified Elastic Network Interface can be queried by logging in to the console, or obtained through the
// Ipv6AddressSet.Address in the return value of the DescribeNetworkInterfaces interface.
func (o PublicIpv6AttachmentOutput) PrivateIpv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpv6Attachment) pulumi.StringPtrOutput { return v.PrivateIpv6Address }).(pulumi.StringPtrOutput)
}

type PublicIpv6AttachmentArrayOutput struct{ *pulumi.OutputState }

func (PublicIpv6AttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicIpv6Attachment)(nil)).Elem()
}

func (o PublicIpv6AttachmentArrayOutput) ToPublicIpv6AttachmentArrayOutput() PublicIpv6AttachmentArrayOutput {
	return o
}

func (o PublicIpv6AttachmentArrayOutput) ToPublicIpv6AttachmentArrayOutputWithContext(ctx context.Context) PublicIpv6AttachmentArrayOutput {
	return o
}

func (o PublicIpv6AttachmentArrayOutput) Index(i pulumi.IntInput) PublicIpv6AttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicIpv6Attachment {
		return vs[0].([]*PublicIpv6Attachment)[vs[1].(int)]
	}).(PublicIpv6AttachmentOutput)
}

type PublicIpv6AttachmentMapOutput struct{ *pulumi.OutputState }

func (PublicIpv6AttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicIpv6Attachment)(nil)).Elem()
}

func (o PublicIpv6AttachmentMapOutput) ToPublicIpv6AttachmentMapOutput() PublicIpv6AttachmentMapOutput {
	return o
}

func (o PublicIpv6AttachmentMapOutput) ToPublicIpv6AttachmentMapOutputWithContext(ctx context.Context) PublicIpv6AttachmentMapOutput {
	return o
}

func (o PublicIpv6AttachmentMapOutput) MapIndex(k pulumi.StringInput) PublicIpv6AttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicIpv6Attachment {
		return vs[0].(map[string]*PublicIpv6Attachment)[vs[1].(string)]
	}).(PublicIpv6AttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpv6AttachmentInput)(nil)).Elem(), &PublicIpv6Attachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpv6AttachmentArrayInput)(nil)).Elem(), PublicIpv6AttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpv6AttachmentMapInput)(nil)).Elem(), PublicIpv6AttachmentMap{})
	pulumi.RegisterOutputType(PublicIpv6AttachmentOutput{})
	pulumi.RegisterOutputType(PublicIpv6AttachmentArrayOutput{})
	pulumi.RegisterOutputType(PublicIpv6AttachmentMapOutput{})
}
