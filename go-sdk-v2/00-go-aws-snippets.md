## Creating a Client

```golang
import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

...

func getEC2Client(region string) (*ec2.EC2, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	...
	cfg.Region = endpoints.UsEast1RegionID
}
```

+ `LoadDefaultAWSConfig` load environment variables, shared cred file and config
+ now create a new `config` object and set it equal to `ec2Svc`

```golang
ec2Svc ::= ec2.New(cfg)
```

+ now that we have the client, you can send api calls to aws services
+ first, create a request input

```golang
filter := ec2.Filter{
	Name: aws.String("name"),
	Values: []string{ubuntuImageSearch},
}

filters := []ec2.Filter{filter}
dii := ec2.DescribeImagesInput{
	Filters: filters,
}
```

+ now pass a pointer the input paramater to a `request` method
+ what is a `request` method?
+ calls to the api are performed through the use of a method, suffixed with `Request` at the end
+ previously v1 used `ec2svc.DescribeImages` and with v2 we use `ec2svc.DescribeImagesRequest`
+ `ec2svc.DescribeImagesRequest` creates a request object that can be used to pass optional parameters, like `Context` to control timeouts

```golang
...
dir := client.DescribeImagesRequest(&dii)
...
```

+ we can modify the request before its sent to allow you to add the cancelable `Context`, i.e.

```golang
ctx := context.Background()
ctx, cancelFn := context.WithTimeout(ctx, time.Minute*5)
defer cancelFn()
dir.SetContext(ctx)
...
```

+ now with the request objects, we can call the aws api like

```golang
...
amis, err := dir.Send()
if err != nil {
...
}
```

+ the result value and error are returned by `Send`
+ the error is your typical Go error, and the result is a pointer to the `Output` object