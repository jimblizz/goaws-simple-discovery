package godisco

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"errors"
)

type Godisco struct {}

var tagKey = "Service"
var region = "eu-west-1"

func (gd *Godisco) SetTagKey(newTagKey string) {
	tagKey = newTagKey
}

func (gd *Godisco) SetRegion(newRegion string) {
	region = newRegion
}

func (gd *Godisco) GetIPs(service string) (ips []string, err error) {

	svc := ec2.New(session.New(), aws.NewConfig().WithRegion(region))

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:" + tagKey),
				Values: []*string{
					aws.String(service),
				},
			},
		},
	}
	resp, ec2err := svc.DescribeInstances(params)

	if ec2err != nil {
		err = ec2err
		return
	}

	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			ips = append(ips, *instance.PrivateIpAddress)
		}
	}

	if len(ips) == 0 {
		err = errors.New("No IPs found for the service " + service)
	}

	return
}

func (gd *Godisco) GetFirstIp(service string) (ip string, err error) {
	ips, ipsErr := gd.GetIPs(service)

	if ipsErr != nil {
		err = ipsErr
	} else {
		ip = ips[0]
	}
	return
}