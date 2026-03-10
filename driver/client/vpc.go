package client

import (
	"fmt"

	tccommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vpcapi "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

type VPCClient struct {
	client *vpcapi.Client
}

func GetVPCClient(credential *tccommon.Credential, region string) (*VPCClient, error) {
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vpc.tencentcloudapi.com"
	client, err := vpcapi.NewClient(credential, region, cpf)
	if err != nil {
		return nil, err
	}

	return &VPCClient{client: client}, nil
}

func (v VPCClient) GetVPCs() (*vpcapi.DescribeVpcsResponse, error) {
	request := vpcapi.NewDescribeVpcsRequest()

	response, err := v.client.DescribeVpcs(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (v VPCClient) GetSubnets() (*vpcapi.DescribeSubnetsResponse, error) {
	request := vpcapi.NewDescribeSubnetsRequest()

	response, err := v.client.DescribeSubnets(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (v VPCClient) GetSecurityGroups() (*vpcapi.DescribeSecurityGroupsResponse, error) {
	request := vpcapi.NewDescribeSecurityGroupsRequest()

	response, err := v.client.DescribeSecurityGroups(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (v VPCClient) CreateVPC(name, cidrBlock string) (string, error) {
	request := vpcapi.NewCreateVpcRequest()
	request.VpcName = &name
	request.CidrBlock = &cidrBlock

	response, err := v.client.CreateVpc(request)
	if err != nil {
		return "", fmt.Errorf("failed to create VPC: %w", err)
	}

	if response.Response == nil || response.Response.Vpc == nil || response.Response.Vpc.VpcId == nil {
		return "", fmt.Errorf("create VPC returned empty response")
	}

	return *response.Response.Vpc.VpcId, nil
}

func (v VPCClient) CreateSubnet(vpcID, name, cidrBlock, zone string) (string, error) {
	request := vpcapi.NewCreateSubnetRequest()
	request.VpcId = &vpcID
	request.SubnetName = &name
	request.CidrBlock = &cidrBlock
	request.Zone = &zone

	response, err := v.client.CreateSubnet(request)
	if err != nil {
		return "", fmt.Errorf("failed to create subnet: %w", err)
	}

	if response.Response == nil || response.Response.Subnet == nil || response.Response.Subnet.SubnetId == nil {
		return "", fmt.Errorf("create subnet returned empty response")
	}

	return *response.Response.Subnet.SubnetId, nil
}

func (v VPCClient) DeleteVPC(vpcID string) error {
	request := vpcapi.NewDeleteVpcRequest()
	request.VpcId = &vpcID

	_, err := v.client.DeleteVpc(request)
	if err != nil {
		return fmt.Errorf("failed to delete VPC %s: %w", vpcID, err)
	}

	return nil
}

func (v VPCClient) DeleteSubnet(subnetID string) error {
	request := vpcapi.NewDeleteSubnetRequest()
	request.SubnetId = &subnetID

	_, err := v.client.DeleteSubnet(request)
	if err != nil {
		return fmt.Errorf("failed to delete subnet %s: %w", subnetID, err)
	}

	return nil
}
