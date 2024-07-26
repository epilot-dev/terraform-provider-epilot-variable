// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TaxResource{}
var _ resource.ResourceWithImportState = &TaxResource{}

func NewTaxResource() resource.Resource {
	return &TaxResource{}
}

// TaxResource defines the resource implementation.
type TaxResource struct {
	client *sdk.SDK
}

// TaxResourceModel describes the resource data model.
type TaxResourceModel struct {
	Active      types.Bool   `tfsdk:"active"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Rate        types.String `tfsdk:"rate"`
	Region      types.String `tfsdk:"region"`
	Type        types.String `tfsdk:"type"`
}

func (r *TaxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax"
}

func (r *TaxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Tax Resource",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rate": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required:    true,
				Description: `must be one of ["DE", "AT", "CH"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"DE",
						"AT",
						"CH",
					),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `must be one of ["VAT", "Custom"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"VAT",
						"Custom",
					),
				},
			},
		},
	}
}

func (r *TaxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *TaxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *TaxResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedTaxCreate()
	res, err := r.client.Tax.CreateTax(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *TaxResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// read.tax.hydrateread.tax.hydrate impedance mismatch: "boolean" != "class"trace=["Tax#create.req"]
	// {"ResponseEnvelope":false,"ResolvedModel":"TaxCreate","Original":{"Name":"TaxCreate","OriginalName":"TaxCreate","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":false,"MustUse":false},{"Type":"component","Identifier":"true","Used":false,"MustUse":false}],"Type":"class","ItemType":null,"Fields":[{"Name":"active","OriginalName":"active","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"active"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"description","OriginalName":"description","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"description"}],"Nullable":false,"Optional":true,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"rate","OriginalName":"rate","Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"rate"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"region","OriginalName":"region","Type":{"Name":"TaxCreate_region","OriginalName":"region","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["DE","AT","CH"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"region"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null},{"Name":"type","OriginalName":"type","Type":{"Name":"TaxCreate_type","OriginalName":"type","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["VAT","Custom"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"type"}],"Nullable":false,"Optional":false,"SerializationMethod":null,"ErrorMessage":false,"Const":null,"Default":null,"IsAdditionalProperties":false,"ParameterIndex":null}],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"shared","IsComponent":true,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{"x-speakeasy-entity":"Tax"},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"IsComponent":true,"AssociatedTypes":[],"Fields":[{"Default":null,"Const":null,"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"_id"}],"IsAdditionalProperties":false,"OriginalName":"_id","Type":{"ResolvedModel":"","Enum":null,"ContextStack":[],"EventStreamEnvelope":false,"ItemType":null,"IsComponent":false,"OriginalName":"","Type":"string","Comments":null,"Discriminator":null,"CircularReference":null,"OutputLocation":"","Examples":[{}],"Fields":[],"Format":"uuid","Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[{}],"Format":"uuid","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Name":"","Extensions":{"x-speakeasy-trace":{"Tax#create.resp.id":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-param-readonly":true,"x-speakeasy-in-get":true},"ResponseEnvelope":false,"Truncated":false,"Validations":{"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null},"Scope":"","AssociatedTypes":[],"Input":false,"Output":false},"Name":"_id","Nullable":false,"Optional":false,"ErrorMessage":false},{"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"active"}],"OriginalName":"active","Optional":false,"Default":null,"Const":null,"ErrorMessage":false,"Type":{"ItemType":null,"Output":false,"Scope":"","ResponseEnvelope":false,"Enum":null,"Discriminator":null,"EventStreamEnvelope":false,"Type":"boolean","Format":"","Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Name":"","AssociatedTypes":[],"Input":false,"ContextStack":[],"OutputLocation":"","Comments":null,"CircularReference":null,"ResolvedModel":"","OriginalName":"","Examples":[],"Validations":{"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null},"IsComponent":false,"Truncated":false,"Extensions":{"x-untouched":true,"x-speakeasy-in-get":true,"x-speakeasy-trace":{"Tax#create.req.active":true},"x-speakeasy-param-computed":true},"Fields":[]},"Name":"active","Nullable":false,"IsAdditionalProperties":false},{"Name":"description","Optional":true,"Default":null,"Const":null,"IsAdditionalProperties":false,"ErrorMessage":false,"OriginalName":"description","Type":{"Name":"","Extensions":{"x-speakeasy-trace":{"Tax#create.req.description":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true},"IsComponent":false,"EventStreamEnvelope":false,"ResolvedModel":"","Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[],"Discriminator":null,"Examples":[],"OriginalName":"","Enum":null,"ResponseEnvelope":false,"Comments":null,"Truncated":false,"CircularReference":null,"Type":"string","Output":false,"Scope":"","Fields":[],"Format":"","Input":false,"ItemType":null,"OutputLocation":"","AssociatedTypes":[],"Validations":{"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null}},"Nullable":false,"Comments":null,"Annotations":[{"FieldName":"description","Ignore":false}]},{"Const":null,"Comments":null,"Annotations":[{"Ignore":false,"FieldName":"rate"}],"ErrorMessage":false,"OriginalName":"rate","Default":null,"Nullable":false,"Optional":false,"IsAdditionalProperties":false,"Type":{"Format":"","Type":"string","Extensions":{"x-speakeasy-trace":{"Tax#create.req.rate":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true},"Enum":null,"IsComponent":false,"ResponseEnvelope":false,"ResolvedModel":"","OutputLocation":"","EventStreamEnvelope":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[],"Name":"","Validations":{"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null},"OriginalName":"","Scope":"","CircularReference":null,"ItemType":null,"Comments":null,"Fields":[],"Input":false,"Output":false,"Truncated":false,"AssociatedTypes":[],"Discriminator":null,"Examples":[]},"Name":"rate"},{"OriginalName":"region","Name":"region","Comments":null,"Annotations":[{"Ignore":false,"FieldName":"region"}],"IsAdditionalProperties":false,"Type":{"AssociatedTypes":[],"Comments":null,"ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Original":{"Name":"TaxCreate_region","OriginalName":"region","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["DE","AT","CH"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"OutputLocation":"models/shared","Format":"","IsComponent":false,"EventStreamEnvelope":false,"Name":"TaxCreate_region","Discriminator":null,"Input":false,"OriginalName":"region","Type":"enum","CircularReference":null,"Validations":{"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null},"ResolvedModel":"TaxCreate","Output":false,"ResponseEnvelope":false,"Examples":[],"ItemType":null,"Scope":"shared","Truncated":false,"Enum":{"Type":{"CircularReference":null,"IsComponent":false,"Scope":"","Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Extensions":{},"Fields":[],"Format":"","ContextStack":[],"Name":"","Comments":null,"EventStreamEnvelope":false,"Enum":null,"ItemType":null,"Truncated":false,"OriginalName":"","Discriminator":null,"Input":false,"ResponseEnvelope":false,"AssociatedTypes":[],"ResolvedModel":"","OutputLocation":"","Type":"string","Examples":[],"Output":false,"Validations":{"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null}},"Open":false,"Names":[],"Values":["DE","AT","CH"]},"Fields":[],"Extensions":{"x-speakeasy-trace":{"Tax#create.req.region":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get":true}},"Nullable":false,"Optional":false,"Default":null,"Const":null,"ErrorMessage":false},{"Type":{"Name":"","Extensions":{"x-speakeasy-match":"id","x-speakeasy-trace":{"Tax#get.req.tax_id":true}},"Comments":{"Summary":"","Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":""},"Discriminator":null,"Enum":null,"AssociatedTypes":[],"Truncated":false,"EventStreamEnvelope":false,"ResponseEnvelope":false,"OriginalName":"","IsComponent":false,"OutputLocation":"","Format":"uuid","Output":false,"Scope":"","ResolvedModel":"","CircularReference":null,"Type":"string","Input":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":{"Summary":"","Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":""},"Input":false,"Output":false,"Extensions":{"x-speakeasy-match":"id"},"Examples":[{}],"Format":"uuid","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"ContextStack":[],"Examples":[{}],"Fields":[],"ItemType":null,"Validations":{"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null}},"Name":"taxId","Nullable":false,"Const":null,"Comments":{"DeprecationMessage":"","Summary":"","Description":"The tax id","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":""},"OriginalName":"","Optional":false,"Default":null,"Annotations":[{"Serialization":"","Style":"simple","Explode":false,"FieldType":{"Truncated":false,"Comments":{"ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":"","Summary":"","Description":"The tax id"},"Extensions":{"x-speakeasy-match":"id"},"EventStreamEnvelope":false,"OriginalName":"","Type":"string","Scope":"","IsComponent":false,"Input":false,"Format":"uuid","OutputLocation":"","ResponseEnvelope":false,"Name":"","Enum":null,"Examples":[{}],"ComplexAny":false,"Fields":[],"AssociatedTypes":[],"Validations":{"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null},"Output":false,"Discriminator":null,"ResolvedModel":"","ContextStack":[],"ItemType":null},"Hidden":false,"ParamType":"pathParam","Name":"taxId"}],"IsAdditionalProperties":false,"ErrorMessage":false},{"Comments":null,"IsAdditionalProperties":false,"Type":{"Output":false,"CircularReference":null,"ResponseEnvelope":false,"Name":"TaxCreate_type","Examples":[],"Discriminator":null,"Fields":[],"AssociatedTypes":[],"Comments":null,"ResolvedModel":"TaxCreate","Input":false,"ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Enum":{"Names":[],"Values":["VAT","Custom"],"Type":{"IsComponent":false,"Name":"","Extensions":{},"Examples":[],"EventStreamEnvelope":false,"OriginalName":"","Type":"string","Enum":null,"Input":false,"ItemType":null,"Discriminator":null,"Output":false,"ResolvedModel":"","Truncated":false,"CircularReference":null,"OutputLocation":"","ContextStack":[],"AssociatedTypes":[],"Fields":[],"Format":"","Validations":{"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null},"ResponseEnvelope":false,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Comments":null,"Scope":""},"Open":false},"IsComponent":false,"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"Original":{"Name":"TaxCreate_type","OriginalName":"type","ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":true,"MustUse":false}],"Type":"enum","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":{"Type":{"Name":"","OriginalName":"","ContextStack":[],"Type":"string","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"Values":["VAT","Custom"],"Names":[],"Open":false,"Format":""},"Scope":"shared","IsComponent":false,"Truncated":false,"Comments":null,"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"models/shared","ResolvedModel":"TaxCreate","EventStreamEnvelope":false,"ResponseEnvelope":false},"OutputLocation":"models/shared","Scope":"shared","Truncated":false,"EventStreamEnvelope":false,"Format":"","OriginalName":"type","Extensions":{"x-untouched":true,"x-speakeasy-in-get":true,"x-speakeasy-trace":{"Tax#create.req.type":true},"x-speakeasy-param-computed":true}},"Name":"type","Nullable":false,"Optional":false,"Default":null,"Const":null,"ErrorMessage":false,"OriginalName":"type","Annotations":[{"Ignore":false,"FieldName":"type"}]}],"ItemType":null,"Truncated":false,"OriginalName":"TaxCreate","OutputLocation":"models/shared","Examples":[],"Output":false,"EventStreamEnvelope":false,"Validations":{"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null},"Extensions":{"x-speakeasy-entity":"Tax","x-speakeasy-trace":{"Tax#create.req":true},"x-speakeasy-param-computed":true,"x-untouched":true,"x-speakeasy-in-get-request":true,"x-speakeasy-in-get":true,"x-speakeasy-root":true},"Scope":"shared","CircularReference":null,"Format":"","Name":"Tax","Type":"class","Discriminator":null,"ContextStack":[{"Type":"refType","Identifier":"Schemas","Used":false,"MustUse":false},{"Type":"refName","Identifier":"TaxCreate","Used":false,"MustUse":false},{"Type":"component","Identifier":"true","Used":false,"MustUse":false}],"Enum":null,"Input":false}
	// {"Examples":[],"ItemType":null,"Original":{"Name":"","OriginalName":"","ContextStack":[],"Type":"boolean","ItemType":null,"Fields":[],"Validations":{"MinItems":null,"MinLength":null,"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null},"AssociatedTypes":[],"Enum":null,"Scope":"","IsComponent":false,"Truncated":false,"Comments":{"Summary":"","Description":"Hydrates entities in relations when passed true","ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":""},"Input":false,"Output":false,"Extensions":{},"Examples":[],"Format":"","Discriminator":null,"ComplexAny":false,"OutputLocation":"","ResolvedModel":"","EventStreamEnvelope":false,"ResponseEnvelope":false},"OriginalName":"","Type":"boolean","Discriminator":null,"Fields":[],"Validations":{"Minimum":null,"MaxItems":null,"MaxLength":null,"Maximum":null,"Pattern":null,"UniqueItems":null,"MinItems":null,"MinLength":null},"ResolvedModel":"","Input":false,"Scope":"","Comments":{"ExternalDocs":null,"ExtendedComments":{},"Deprecated":false,"DeprecationReplacement":"","DeprecationMessage":"","Summary":"","Description":"Hydrates entities in relations when passed true"},"Extensions":{},"ResponseEnvelope":false,"ContextStack":[],"Name":"","IsComponent":false,"Output":false,"Truncated":false,"EventStreamEnvelope":false,"Enum":null,"OutputLocation":"","AssociatedTypes":[],"Format":"","CircularReference":null}
	var hydrate *bool
	taxID := data.ID.ValueString()
	request := operations.GetTaxRequest{
		Hydrate: hydrate,
		TaxID:   taxID,
	}
	res, err := r.client.Tax.GetTax(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *TaxResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	taxCreate := *data.ToSharedTaxCreate()
	taxID := data.ID.ValueString()
	request := operations.UpdateTaxRequest{
		TaxCreate: taxCreate,
		TaxID:     taxID,
	}
	res, err := r.client.Tax.UpdateTax(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.Tax != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTax(res.Tax)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TaxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *TaxResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	taxID := data.ID.ValueString()
	request := operations.DeleteTaxRequest{
		TaxID: taxID,
	}
	res, err := r.client.Tax.DeleteTax(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *TaxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
