package main

import (
	"fmt"
	"github.com/awslabs/goformation/cloudformation"
)

func main() {

	subnetParam := map[string]string{
		"Type":        "String",
		"Description": "Subnet Parameter",
	}

	// Create a new CloudFormation template
	template := cloudformation.NewTemplate()

	template.Parameters["Subnet1Parameter"] = subnetParam
	template.Parameters["Subnet2Parameter"] = subnetParam

	template.Resources["LoadBalancer1"] = &cloudformation.AWSElasticLoadBalancingLoadBalancer{
		Listeners: []cloudformation.AWSElasticLoadBalancingLoadBalancer_Listeners{
			cloudformation.AWSElasticLoadBalancingLoadBalancer_Listeners{
				InstancePort:     "1",
				LoadBalancerPort: "65535",
				Protocol:         "TCP",
			},
		},
		Scheme:           "internal",
		LoadBalancerName: "svc-limit-temp-lb-1",
		Subnets: []string{
			"subnet-598eec20",
			"subnet-e16b10bb",
		},
	}

	template.Resources["LoadBalancer2"] = &cloudformation.AWSElasticLoadBalancingV2LoadBalancer{
		Name:   "svc-limit-temp-lb-2",
		Scheme: "internal",
		Subnets: []string{
			"subnet-598eec20",
			"subnet-e16b10bb",
		},
	}

	template.Resources["LoadBalancer3"] = &cloudformation.AWSElasticLoadBalancingV2LoadBalancer{
		Name:   "svc-limit-temp-lb-3",
		Type:   "network",
		Scheme: "internal",
		Subnets: []string{
			"subnet-598eec20",
			"subnet-e16b10bb",
		},
	}

	// Let's do YAML
	y, err := template.YAML()
	if err != nil {
		fmt.Printf("Failed to generate YAML: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(y))
	}
}
