# Amazon ECS

+ Highly scalable, fast, container management service
+ Easy to run, stop, and manage Docker containers on a cluster

## Features

The following sectiosn describe individual elements

### Containers and Images

+ a Docker container is a standardized unit of software development
+ images are built from a `Dockerfile`
+ images are stored in a `registry`

### Task Definitions

+ `task definition` is JSON describing one or more containers up to 10
+ this description is used to form your application; a blueprint
+ parameters, launch type, port to open, data volumes, etc.

```json
{
    "family": "webserver",
    "containerDefinitions": [
        {
            "name": "web",
            "image": "nginx",
            "memory": "100",
            "cpu": "99"
        },
    ],
    "requiresCompatibilities": [
        "FARGATE"
    ],
    "networkMode": "awsvpc",
    "memory": "512",
    "cpu": "256",
}
```

### Tasks and Scheduling

+ a `task` is the instantiation of a `task definition` within a cluster
+ you then specify the number of `tasks` that will run on your cluster
+ amazon ecs _places_ these `tasks` within your cluster; orchestrates
+ several schedules available, maintain a specific number, etc.

### Clusters

+ a `task` is run inside a `cluster` - a logical grouping of resources

### Container Agent

+ the `container agent` runs on each infrastructure within ecs cluster
+ sends info about running tasks and resource utilization to ecs
+ responsible for starts/stops when a request is received

### Summary

`dockerfile -> image -> registry -> task definition -> task -> cluster`