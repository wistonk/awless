package api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Infra struct {
	*ec2.EC2
}

func NewInfra(sess *session.Session) *Infra {
	return &Infra{ec2.New(sess)}
}

func (inf *Infra) Instances() (interface{}, error) {
	return inf.DescribeInstances(&ec2.DescribeInstancesInput{})
}

func (inf *Infra) Vpcs() (interface{}, error) {
	return inf.DescribeVpcs(&ec2.DescribeVpcsInput{})
}

func (inf *Infra) Subnets() (interface{}, error) {
	return inf.DescribeSubnets(&ec2.DescribeSubnetsInput{})
}

func (inf *Infra) Regions() (interface{}, error) {
	return inf.DescribeRegions(&ec2.DescribeRegionsInput{})
}

func (inf *Infra) Vpc(id string) (interface{}, error) {
	input := &ec2.DescribeVpcsInput{
		VpcIds: []*string{aws.String(id)},
	}

	return inf.DescribeVpcs(input)
}
