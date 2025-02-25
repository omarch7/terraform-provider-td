package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-td/internal/tdclient"
)

var (
	_ provider.Provider = &treasureDataProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &treasureDataProvider{
			version: version,
		}
	}
}

type treasureDataProviderModel struct {
	Host   types.String `tfsdk:"host"`
	Apikey types.String `tfsdk:"apikey"`
}

type treasureDataProvider struct {
	version string
}

func (p *treasureDataProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "td"
	resp.Version = p.version
}

func (p *treasureDataProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The host of the treasure data",
				Optional:    true,
			},
			"apikey": schema.StringAttribute{
				Description: "The apikey of the treasure data",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *treasureDataProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configure Treasure Data Client")
	var config treasureDataProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Treasure Data API Host",
			"The provider cannot create the Treasure Data API client as there is an unknown configuration value for the Treasure Data API Host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TRESURE_DATA_HOST environment variable.",
		)
	}

	if config.Apikey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("apikey"),
			"Unknown Treasure Data API Key",
			"The provider cannot create the Treasure Data API client as there is an unknown configuration value for the Treasure Data API Key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TRESURE_DATA_APIKEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("TREASURE_DATA_HOST")
	apikey := os.Getenv("TREASURE_DATA_APIKEY")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Apikey.IsNull() {
		apikey = config.Apikey.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Treasure Data API Host",
			"The provider cannot create the Treasure Data API client as there is a missing or empty value for the Treasure Data API Host. "+
				"Set the host value in the configuration or use the TREASURE_DATA_HOST environment variable."+
				"If either is already set, ensure that the value is not empty.",
		)
	}

	if apikey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("apikey"),
			"Missing Treasure Data API Key",
			"The provider cannot create the Treasure Data API client as there is a missing or empty value for the Treasure Data API Key. "+
				"Set the apikey value in the configuration or use the TREASURE_DATA_APIKEY environment variable."+
				"If either is already set, ensure that the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "treasure_data_host", host)
	ctx = tflog.SetField(ctx, "treasure_data_apikey", apikey)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "treasure_data_apikey")

	tflog.Debug(ctx, "Creating Treasure Data API Client")

	client, err := tdclient.NewClient(&host, &apikey)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Treasure Data API Client",
			"An unexpected error occurred when creating the Treasure Data API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Treasure Data Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Treasure Data Client", map[string]any{"success": true})
}

func (p *treasureDataProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewFoldersDataSource,
		NewParentSegmentsDataSource,
	}
}

func (p *treasureDataProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
