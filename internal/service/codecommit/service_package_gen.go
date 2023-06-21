// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codecommit

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	codecommit_sdkv1 "github.com/aws/aws-sdk-go/service/codecommit"
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
			Factory:  DataSourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
		},
		{
			Factory:  DataSourceRepository,
			TypeName: "aws_codecommit_repository",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceApprovalRuleTemplate,
			TypeName: "aws_codecommit_approval_rule_template",
		},
		{
			Factory:  ResourceApprovalRuleTemplateAssociation,
			TypeName: "aws_codecommit_approval_rule_template_association",
		},
		{
			Factory:  ResourceRepository,
			TypeName: "aws_codecommit_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTrigger,
			TypeName: "aws_codecommit_trigger",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeCommit
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*codecommit_sdkv1.CodeCommit, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return codecommit_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
