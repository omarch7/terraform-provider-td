package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-td/internal/tdclient"
	"terraform-provider-td/internal/tdclient/models"
)

var (
	_ resource.Resource              = &folderResource{}
	_ resource.ResourceWithConfigure = &folderResource{}
)

func NewFolderResource() resource.Resource {
	return &folderResource{}
}

type folderResource struct {
	client *tdclient.Client
}

type folderResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	ParentFolderID types.String `tfsdk:"parent_folder_id"`
	CreatedAt      types.String `tfsdk:"created_at"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
	CreatedBy      types.String `tfsdk:"created_by"`
	UpdatedBy      types.String `tfsdk:"updated_by"`
}

func (r *folderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*tdclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *tdclient.Client, got %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *folderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_folder"
}

func (r *folderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (r *folderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan folderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var folder = models.Folder{
		ID:   plan.ID.String(),
		Type: "folder-segment",
		Attributes: models.FolderAttributes{
			Name:        plan.Name.String(),
			Description: plan.Description.String(),
		},
		Relationships: models.FolderRelationships{
			ParentFolder: models.Relationship{
				Data: models.RelationshipData{
					ID:   plan.ParentFolderID.String(),
					Type: "folder-segment",
				},
			},
		},
	}

	createdFolder, err := r.client.CreateFolder(folder)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating folder",
			fmt.Sprintf("Could not create folder, unexpected error: %s", err),
		)
		return
	}

	plan.ID = types.StringValue(createdFolder.ID)
	plan.CreatedAt = types.StringValue(createdFolder.Attributes.CreatedAt)
	plan.UpdatedAt = types.StringValue(createdFolder.Attributes.UpdatedAt)
	plan.CreatedBy = types.StringValue(createdFolder.Relationships.CreatedBy.Data.ID)
	plan.UpdatedBy = types.StringValue(createdFolder.Relationships.UpdatedBy.Data.ID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *folderResource) Read(_ context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *folderResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *folderResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
