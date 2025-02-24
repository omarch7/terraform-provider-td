package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var (
	_ datasource.DataSource              = &treasureDataSource{}
	_ datasource.DataSourceWithConfigure = &treasureDataSource{}
)

func NewTreasureDataSource() datasource.DataSource {
	return &treasureDataSource{}
}

type treasureDataSource struct{}

func (d *treasureDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
    if req.ProviderData != nil {
        return
    }
}

func (d *treasureDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_td"
}

func (d *treasureDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (d *treasureDataSource) Read(_ context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}
