package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"terraform-provider-td/internal/tdclient"
)

var (
	_ datasource.DataSource              = &foldersDataSource{}
	_ datasource.DataSourceWithConfigure = &foldersDataSource{}
)

func NewFoldersDataSource() datasource.DataSource {
	return &foldersDataSource{}
}

type foldersDataSource struct {
	client *tdclient.Client
}

type foldersDataSourceModel struct {
	Folders []foldersModel `tfsdk:"folders"`
}

type foldersModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	ParentFolderID types.String `tfsdk:"parent_folder_id"`
}

func (d *foldersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
    tflog.Info(ctx, "Configure Treasure Data Client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*tdclient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *tdclient.Client, got %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (d *foldersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_folders"
}

func (d *foldersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"folders": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"parent_folder_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *foldersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state foldersDataSourceModel

	folders, error := d.client.GetFolders()
	if error != nil {
		resp.Diagnostics.AddError("Failed to get folders", error.Error())
		return
	}

	for _, folder := range folders.Data {
		state.Folders = append(state.Folders, foldersModel{
			ID:             types.StringValue(folder.ID),
			Name:           types.StringValue(folder.Attributes.Name),
			Description:    types.StringValue(folder.Attributes.Description),
			ParentFolderID: types.StringValue(folder.Relationships.Parent.Data.Id),
		})
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
