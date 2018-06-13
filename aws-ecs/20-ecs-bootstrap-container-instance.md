# Bootstrapping Container Instances with Amazon EC2 User Data

+ the most common userdata is to pass configuration info to the docker daemon and ecs container agent

## Amazon ECS Container Agent

+ `/etc/ecs/ecs.config` is where config data is read

#### Set a Single Configuration Variable

```bash
#!/bin/bash
echo "ECS_CLUSTER=MyCluster" >> /etc/ecs/ecs.config
```

#### Set Multiple Variables

```bash
#!/bin/bash
cat <<'EOF' >> /etc/ecs/ecs.config
ECS_CLUSTER=MyCluster
ECS_ENGINE_AUTH_TYPE=docker
ECS_ENGINE_AUTH_DATA={"https://index.docker.io/v1/":{"username":"my_name","password":"my_password","email":"email@example.com"}}
ECS_LOGLEVEL=debug
EOF
```

## Docker Daemon

+ `cloud-boothook` user data format executes before user data
+ `cloud-boothook` runs at every instance boot