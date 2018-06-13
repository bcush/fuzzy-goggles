# Golang AWS SDK v2

## Getting Started

```golang
package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
)
```
+ `package main` - Is your starting point as usual.

+ `aws` - 	Package aws provides the core SDK's utilities and shared types.

+ `endpoints` -	Package endpoints provides the types and functionality for defining regions and endpoints, as well as querying those definitions.

+ `external` - I guess this is for external credential providers?