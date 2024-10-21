// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceLocalDisk,
			TypeName: "aws_storagegateway_local_disk",
			Name:     "Local Disk",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCache,
			TypeName: "aws_storagegateway_cache",
			Name:     "Cache",
		},
		{
			Factory:  resourceCachediSCSIVolume,
			TypeName: "aws_storagegateway_cached_iscsi_volume",
			Name:     "Cached iSCSI Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFileSystemAssociation,
			TypeName: "aws_storagegateway_file_system_association",
			Name:     "File System Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGateway,
			TypeName: "aws_storagegateway_gateway",
			Name:     "Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceNFSFileShare,
			TypeName: "aws_storagegateway_nfs_file_share",
			Name:     "NFS File Share",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSMBFileShare,
			TypeName: "aws_storagegateway_smb_file_share",
			Name:     "SMB File Share",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceStorediSCSIVolume,
			TypeName: "aws_storagegateway_stored_iscsi_volume",
			Name:     "Stored iSCSI Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceTapePool,
			TypeName: "aws_storagegateway_tape_pool",
			Name:     "Tape Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUploadBuffer,
			TypeName: "aws_storagegateway_upload_buffer",
			Name:     "Upload Buffer",
		},
		{
			Factory:  ResourceWorkingStorage,
			TypeName: "aws_storagegateway_working_storage",
			Name:     "Working Storage",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.StorageGateway
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*storagegateway.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))

	return storagegateway.NewFromConfig(cfg,
		storagegateway.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
