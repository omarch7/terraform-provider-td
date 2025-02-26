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
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"parent_folder_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *folderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan folderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var folder = models.Folder{
		Type: "folder-segment",
		Attributes: models.FolderAttributes{
			Name:        plan.Name.ValueString(),
			Description: plan.Description.ValueString(),
		},
		Relationships: models.FolderRelationships{
			ParentFolder: models.Relationship{
				Data: models.RelationshipData{
					ID:   plan.ParentFolderID.ValueString(),
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

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *folderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state folderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	folder, err := r.client.GetFolder(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading folder",
			fmt.Sprintf("Could not read folder, unexpected error: %s", err),
		)
		return
	}

	state.ID = types.StringValue(folder.ID)
	state.Name = types.StringValue(folder.Attributes.Name)
	state.Description = types.StringValue(folder.Attributes.Description)
	state.ParentFolderID = types.StringValue(folder.Relationships.ParentFolder.Data.ID)
	state.CreatedAt = types.StringValue(folder.Attributes.CreatedAt)
	state.UpdatedAt = types.StringValue(folder.Attributes.UpdatedAt)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *folderResource) Update(_ context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *folderResource) Delete(_ context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
