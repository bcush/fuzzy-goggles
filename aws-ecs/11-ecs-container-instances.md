# Amazon ECS Container Instances

+ an ecs `container` instance is an ec2 instance running amazon container agent, registered into a cluster

+ when you run a `task`, your ec2 launch type `tasks` are placed on your active container instances

## Container Instance Concepts

+ instance must be running `ecs container agent` to join a cluster
+ the `agent` makes calls to amazon ecs and thus needs an iam role
+ if any `task` needs external connectivity, you can map their network ports to ports on your host ec2 container instance
+ container instances need external network access to reach ecs service endpoint; either a public ip address or a nat gateway
+ ec2 type determines resources available to cluster
+ because each container instance has a unique state stored locally on the instance

1. do not degregister an instance and re-register in a different cluster; terminate and launch new
2. you cannot change a container instance type; terminate and launch new

## Container Instance Lifecycle

+ `agent` registers instance into cluster; instance reports as `active` and agent connection as `true`

+ stopping an instance remains as `active` but connection as `false`

+ changing the status to `draining` will cause new tasks to not get placed on the container instance and service tasks are removed

+ if you `deregister` or `terminate` a container instance; the status changes to `inactive`

## Check the Instance Role for Your Account

+ need `ecsInstanceRole`