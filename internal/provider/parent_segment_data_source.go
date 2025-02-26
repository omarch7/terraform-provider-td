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
	_ datasource.DataSource              = &parentSegmentsDataSource{}
	_ datasource.DataSourceWithConfigure = &parentSegmentsDataSource{}
)

func NewParentSegmentsDataSource() datasource.DataSource {
	return &parentSegmentsDataSource{}
}

type parentSegmentsDataSource struct {
	client *tdclient.Client
}

type parentSegmentsDataSourceModel struct {
	ParentSegments []parentSegmentModel `tfsdk:"parent_segments"`
}

type parentSegmentModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	ParentFolderId types.String `tfsdk:"parent_folder_id"`
}

func (d *parentSegmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *parentSegmentsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_parent_segments"
}

func (d *parentSegmentsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"parent_segments": schema.ListNestedAttribute{
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

func (d *parentSegmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state parentSegmentsDataSourceModel

	parent_segments, err := d.client.GetParentSegments()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get parent segments",
			fmt.Sprintf("Error: %s", err),
		)
		return
	}

	for _, parent_segment := range parent_segments.Data {
		state.ParentSegments = append(state.ParentSegments, parentSegmentModel{
			ID:             types.StringValue(parent_segment.ID),
			Name:           types.StringValue(parent_segment.Attributes.Name),
			Description:    types.StringValue(parent_segment.Attributes.Description),
			ParentFolderId: types.StringValue(parent_segment.Relationships.ParentFolder.Data.ID),
		})
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
