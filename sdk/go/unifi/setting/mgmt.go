// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package setting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `setting.Mgmt` manages settings for a unifi site.
type Mgmt struct {
	pulumi.CustomResourceState

	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrOutput `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site pulumi.StringOutput `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrOutput `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys MgmtSshKeyArrayOutput `pulumi:"sshKeys"`
}

// NewMgmt registers a new resource with the given unique name, arguments, and options.
func NewMgmt(ctx *pulumi.Context,
	name string, args *MgmtArgs, opts ...pulumi.ResourceOption) (*Mgmt, error) {
	if args == nil {
		args = &MgmtArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Mgmt
	err := ctx.RegisterResource("unifi:setting/mgmt:Mgmt", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMgmt gets an existing Mgmt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMgmt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MgmtState, opts ...pulumi.ResourceOption) (*Mgmt, error) {
	var resource Mgmt
	err := ctx.ReadResource("unifi:setting/mgmt:Mgmt", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mgmt resources.
type mgmtState struct {
	// Automatically upgrade device firmware.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled *bool `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys []MgmtSshKey `pulumi:"sshKeys"`
}

type MgmtState struct {
	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrInput
	// SSH key.
	SshKeys MgmtSshKeyArrayInput
}

func (MgmtState) ElementType() reflect.Type {
	return reflect.TypeOf((*mgmtState)(nil)).Elem()
}

type mgmtArgs struct {
	// Automatically upgrade device firmware.
	AutoUpgrade *bool `pulumi:"autoUpgrade"`
	// The name of the site to associate the settings with.
	Site *string `pulumi:"site"`
	// Enable SSH authentication.
	SshEnabled *bool `pulumi:"sshEnabled"`
	// SSH key.
	SshKeys []MgmtSshKey `pulumi:"sshKeys"`
}

// The set of arguments for constructing a Mgmt resource.
type MgmtArgs struct {
	// Automatically upgrade device firmware.
	AutoUpgrade pulumi.BoolPtrInput
	// The name of the site to associate the settings with.
	Site pulumi.StringPtrInput
	// Enable SSH authentication.
	SshEnabled pulumi.BoolPtrInput
	// SSH key.
	SshKeys MgmtSshKeyArrayInput
}

func (MgmtArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mgmtArgs)(nil)).Elem()
}

type MgmtInput interface {
	pulumi.Input

	ToMgmtOutput() MgmtOutput
	ToMgmtOutputWithContext(ctx context.Context) MgmtOutput
}

func (*Mgmt) ElementType() reflect.Type {
	return reflect.TypeOf((**Mgmt)(nil)).Elem()
}

func (i *Mgmt) ToMgmtOutput() MgmtOutput {
	return i.ToMgmtOutputWithContext(context.Background())
}

func (i *Mgmt) ToMgmtOutputWithContext(ctx context.Context) MgmtOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MgmtOutput)
}

// MgmtArrayInput is an input type that accepts MgmtArray and MgmtArrayOutput values.
// You can construct a concrete instance of `MgmtArrayInput` via:
//
//	MgmtArray{ MgmtArgs{...} }
type MgmtArrayInput interface {
	pulumi.Input

	ToMgmtArrayOutput() MgmtArrayOutput
	ToMgmtArrayOutputWithContext(context.Context) MgmtArrayOutput
}

type MgmtArray []MgmtInput

func (MgmtArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mgmt)(nil)).Elem()
}

func (i MgmtArray) ToMgmtArrayOutput() MgmtArrayOutput {
	return i.ToMgmtArrayOutputWithContext(context.Background())
}

func (i MgmtArray) ToMgmtArrayOutputWithContext(ctx context.Context) MgmtArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MgmtArrayOutput)
}

// MgmtMapInput is an input type that accepts MgmtMap and MgmtMapOutput values.
// You can construct a concrete instance of `MgmtMapInput` via:
//
//	MgmtMap{ "key": MgmtArgs{...} }
type MgmtMapInput interface {
	pulumi.Input

	ToMgmtMapOutput() MgmtMapOutput
	ToMgmtMapOutputWithContext(context.Context) MgmtMapOutput
}

type MgmtMap map[string]MgmtInput

func (MgmtMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mgmt)(nil)).Elem()
}

func (i MgmtMap) ToMgmtMapOutput() MgmtMapOutput {
	return i.ToMgmtMapOutputWithContext(context.Background())
}

func (i MgmtMap) ToMgmtMapOutputWithContext(ctx context.Context) MgmtMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MgmtMapOutput)
}

type MgmtOutput struct{ *pulumi.OutputState }

func (MgmtOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mgmt)(nil)).Elem()
}

func (o MgmtOutput) ToMgmtOutput() MgmtOutput {
	return o
}

func (o MgmtOutput) ToMgmtOutputWithContext(ctx context.Context) MgmtOutput {
	return o
}

// Automatically upgrade device firmware.
func (o MgmtOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Mgmt) pulumi.BoolPtrOutput { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

// The name of the site to associate the settings with.
func (o MgmtOutput) Site() pulumi.StringOutput {
	return o.ApplyT(func(v *Mgmt) pulumi.StringOutput { return v.Site }).(pulumi.StringOutput)
}

// Enable SSH authentication.
func (o MgmtOutput) SshEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Mgmt) pulumi.BoolPtrOutput { return v.SshEnabled }).(pulumi.BoolPtrOutput)
}

// SSH key.
func (o MgmtOutput) SshKeys() MgmtSshKeyArrayOutput {
	return o.ApplyT(func(v *Mgmt) MgmtSshKeyArrayOutput { return v.SshKeys }).(MgmtSshKeyArrayOutput)
}

type MgmtArrayOutput struct{ *pulumi.OutputState }

func (MgmtArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mgmt)(nil)).Elem()
}

func (o MgmtArrayOutput) ToMgmtArrayOutput() MgmtArrayOutput {
	return o
}

func (o MgmtArrayOutput) ToMgmtArrayOutputWithContext(ctx context.Context) MgmtArrayOutput {
	return o
}

func (o MgmtArrayOutput) Index(i pulumi.IntInput) MgmtOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mgmt {
		return vs[0].([]*Mgmt)[vs[1].(int)]
	}).(MgmtOutput)
}

type MgmtMapOutput struct{ *pulumi.OutputState }

func (MgmtMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mgmt)(nil)).Elem()
}

func (o MgmtMapOutput) ToMgmtMapOutput() MgmtMapOutput {
	return o
}

func (o MgmtMapOutput) ToMgmtMapOutputWithContext(ctx context.Context) MgmtMapOutput {
	return o
}

func (o MgmtMapOutput) MapIndex(k pulumi.StringInput) MgmtOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mgmt {
		return vs[0].(map[string]*Mgmt)[vs[1].(string)]
	}).(MgmtOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MgmtInput)(nil)).Elem(), &Mgmt{})
	pulumi.RegisterInputType(reflect.TypeOf((*MgmtArrayInput)(nil)).Elem(), MgmtArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MgmtMapInput)(nil)).Elem(), MgmtMap{})
	pulumi.RegisterOutputType(MgmtOutput{})
	pulumi.RegisterOutputType(MgmtArrayOutput{})
	pulumi.RegisterOutputType(MgmtMapOutput{})
}
