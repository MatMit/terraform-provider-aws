// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sqs

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	sqs_sdkv2 "github.com/aws/aws-sdk-go-v2/service/sqs"
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
			Factory:  dataSourceQueue,
			TypeName: "aws_sqs_queue",
		},
		{
			Factory:  dataSourceQueues,
			TypeName: "aws_sqs_queues",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceQueue,
			TypeName: "aws_sqs_queue",
			Name:     "Queue",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceQueuePolicy,
			TypeName: "aws_sqs_queue_policy",
		},
		{
			Factory:  ResourceQueueRedriveAllowPolicy,
			TypeName: "aws_sqs_queue_redrive_allow_policy",
		},
		{
			Factory:  ResourceQueueRedrivePolicy,
			TypeName: "aws_sqs_queue_redrive_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SQS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*sqs_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return sqs_sdkv2.NewFromConfig(cfg, func(o *sqs_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
