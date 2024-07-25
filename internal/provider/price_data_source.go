// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-product/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-product/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PriceDataSource{}
var _ datasource.DataSourceWithConfigure = &PriceDataSource{}

func NewPriceDataSource() datasource.DataSource {
	return &PriceDataSource{}
}

// PriceDataSource is the data source implementation.
type PriceDataSource struct {
	client *sdk.SDK
}

// PriceDataSourceModel describes the data model.
type PriceDataSourceModel struct {
	Active                 types.Bool                          `tfsdk:"active"`
	BillingDurationAmount  types.Number                        `tfsdk:"billing_duration_amount"`
	BillingDurationUnit    types.String                        `tfsdk:"billing_duration_unit"`
	Description            types.String                        `tfsdk:"description"`
	Hydrate                types.Bool                          `tfsdk:"hydrate"`
	ID                     types.String                        `tfsdk:"id"`
	IsCompositePrice       types.Bool                          `tfsdk:"is_composite_price"`
	IsTaxInclusive         types.Bool                          `tfsdk:"is_tax_inclusive"`
	LongDescription        types.String                        `tfsdk:"long_description"`
	NoticeTimeAmount       types.Number                        `tfsdk:"notice_time_amount"`
	NoticeTimeUnit         types.String                        `tfsdk:"notice_time_unit"`
	PriceComponents        *tfTypes.PriceCreatePriceComponents `tfsdk:"price_components"`
	PriceDisplayInJourneys types.String                        `tfsdk:"price_display_in_journeys"`
	PricingModel           types.String                        `tfsdk:"pricing_model"`
	RenewalDurationAmount  types.Number                        `tfsdk:"renewal_duration_amount"`
	RenewalDurationUnit    types.String                        `tfsdk:"renewal_duration_unit"`
	Tax                    types.String                        `tfsdk:"tax"`
	TerminationTimeAmount  types.Number                        `tfsdk:"termination_time_amount"`
	TerminationTimeUnit    types.String                        `tfsdk:"termination_time_unit"`
	Tiers                  []tfTypes.PriceTier                 `tfsdk:"tiers"`
	Type                   types.String                        `tfsdk:"type"`
	Unit                   types.String                        `tfsdk:"unit"`
	UnitAmount             types.Number                        `tfsdk:"unit_amount"`
	UnitAmountCurrency     types.String                        `tfsdk:"unit_amount_currency"`
	UnitAmountDecimal      types.String                        `tfsdk:"unit_amount_decimal"`
	VariablePrice          types.Bool                          `tfsdk:"variable_price"`
}

// Metadata returns the data source type name.
func (r *PriceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_price"
}

// Schema defines the schema for the data source.
func (r *PriceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Price DataSource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the price can be used for new purchases.`,
			},
			"billing_duration_amount": schema.NumberAttribute{
				Computed:    true,
				Description: `The billing period duration`,
			},
			"billing_duration_unit": schema.StringAttribute{
				Computed:    true,
				Description: `The billing period duration unit. must be one of ["weeks", "months", "years"]`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: `A brief description of the price.`,
			},
			"hydrate": schema.BoolAttribute{
				Optional:    true,
				Description: `Hydrates entities in relations when passed true`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"is_composite_price": schema.BoolAttribute{
				Computed:    true,
				Description: `The flag for prices that contain price components.`,
			},
			"is_tax_inclusive": schema.BoolAttribute{
				Computed:    true,
				Description: `Specifies whether the price is considered ` + "`" + `inclusive` + "`" + ` of taxes or not.`,
			},
			"long_description": schema.StringAttribute{
				Computed:    true,
				Description: `A detailed description of the price. This is shown on the order document and order table. Multi-line supported.`,
			},
			"notice_time_amount": schema.NumberAttribute{
				Computed:    true,
				Description: `The notice period duration`,
			},
			"notice_time_unit": schema.StringAttribute{
				Computed:    true,
				Description: `The notice period duration unit. must be one of ["weeks", "months", "years"]`,
			},
			"price_components": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"dollar_relation": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"tags": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
									Description: `An arbitrary set of tags attached to the composite price - component relation`,
								},
								"entity_id": schema.StringAttribute{
									Computed:    true,
									Description: `The id of the price component`,
								},
							},
						},
					},
				},
				Description: `A set of [price](/api/pricing#tag/simple_price_schema) components that define the composite price.`,
			},
			"price_display_in_journeys": schema.StringAttribute{
				Computed:    true,
				Description: `Defines the way the price amount is display in epilot journeys. must be one of ["show_price", "show_as_starting_price", "show_as_on_request"]`,
			},
			"pricing_model": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Describes how to compute the price per period. Either ` + "`" + `per_unit` + "`" + `, ` + "`" + `tiered_graduated` + "`" + ` or ` + "`" + `tiered_volume` + "`" + `.` + "\n" +
					`- ` + "`" + `per_unit` + "`" + ` indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity` + "\n" +
					`- ` + "`" + `tiered_graduated` + "`" + ` indicates that the unit pricing will be computed using tiers attribute. The customer pays the price per unit in every range their purchase rises through.` + "\n" +
					`- ` + "`" + `tiered_volume` + "`" + ` indicates that the unit pricing will be computed using tiers attribute. The customer pays the same unit price for all purchased units.` + "\n" +
					`- ` + "`" + `tiered_flatfee` + "`" + ` While similar to tiered_volume, tiered flat fee charges for the same price (flat) for the entire range instead using the unit price to multiply the quantity.` + "\n" +
					`` + "\n" +
					`must be one of ["per_unit", "tiered_volume", "tiered_graduated", "tiered_flatfee"]`,
			},
			"renewal_duration_amount": schema.NumberAttribute{
				Computed:    true,
				Description: `The renewal period duration`,
			},
			"renewal_duration_unit": schema.StringAttribute{
				Computed:    true,
				Description: `The renewal period duration unit. must be one of ["weeks", "months", "years"]`,
			},
			"tax": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
			},
			"termination_time_amount": schema.NumberAttribute{
				Computed:    true,
				Description: `The termination period duration`,
			},
			"termination_time_unit": schema.StringAttribute{
				Computed:    true,
				Description: `The termination period duration unit. must be one of ["weeks", "months", "years"]`,
			},
			"tiers": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"display_mode": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["hidden", "on_request"]`,
						},
						"flat_fee_amount": schema.NumberAttribute{
							Computed: true,
						},
						"flat_fee_amount_decimal": schema.StringAttribute{
							Computed: true,
						},
						"unit_amount": schema.NumberAttribute{
							Computed: true,
						},
						"unit_amount_decimal": schema.StringAttribute{
							Computed: true,
						},
						"up_to": schema.NumberAttribute{
							Computed: true,
						},
					},
				},
				MarkdownDescription: `Defines an array of tiers. Each tier has an upper bound, an unit amount and a flat fee.` + "\n" +
					``,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `One of ` + "`" + `one_time` + "`" + ` or ` + "`" + `recurring` + "`" + ` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase. must be one of ["one_time", "recurring"]`,
			},
			"unit": schema.StringAttribute{
				Computed:    true,
				Description: `The unit of measurement used for display purposes and possibly for calculations when the price is variable.`,
			},
			"unit_amount": schema.NumberAttribute{
				Computed:    true,
				Description: `The unit amount in cents to be charged, represented as a whole integer if possible.`,
			},
			"unit_amount_currency": schema.StringAttribute{
				Computed:    true,
				Description: `Three-letter ISO currency code, in lowercase.`,
			},
			"unit_amount_decimal": schema.StringAttribute{
				Computed:    true,
				Description: `The unit amount in cents to be charged, represented as a decimal string with at most 12 decimal places.`,
			},
			"variable_price": schema.BoolAttribute{
				Computed:    true,
				Description: `The flag for prices that can be influenced by external variables such as user input.`,
			},
		},
	}
}

func (r *PriceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PriceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PriceDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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

	hydrate := new(bool)
	if !data.Hydrate.IsUnknown() && !data.Hydrate.IsNull() {
		*hydrate = data.Hydrate.ValueBool()
	} else {
		hydrate = nil
	}
	priceID := data.ID.ValueString()
	request := operations.GetPriceRequest{
		Hydrate: hydrate,
		PriceID: priceID,
	}
	res, err := r.client.Price.GetPrice(ctx, request)
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
	if !(res.Price != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPrice(res.Price)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
