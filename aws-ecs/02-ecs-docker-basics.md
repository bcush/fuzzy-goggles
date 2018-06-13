## Docker Basics for ECS

+ docker uses `images` in `task definitions` to launch `containers` on `ec2 instances`

### Create a Docker Image

+ `task definitions` use `images` to launch `containers` on the `container instances` in your `clusters`

### Push Your Image to ECR (optional)

### Register a `task defition` with `hello-world` image

1. Create `hello-world-task-def.json`:

```json
{
    "family": "hello-world",
    "containerDefinitions": [
        {
            "name": "hello-world",
            "image": "aws_account_id.dkr.ecr.us-east-1.amazonaws.com/hello-world",
            "cpu": 10,
            "memory": 500,
            "portMappings": [
                {
                    "containerPort": 80,
                    "hostPort": 80
                }
            ],
            "entryPoint": [
                "/usr/sbin/apache2",
                "-D",
                "FOREGROUND"
            ],
            "essential": true
        }
    ]
}
```
Note: replace `repositoryUri` with yours from ECR.

2. Register a `task defition`:

`aws ecs register-task-definition --cli-input-json file://hello-world-task-def.json`

3. Run the `task definition`:

`aws ecs run-task --task-definition hello-world`

