# Cleaning Up your Amazon ECS Resources

## Scale Down Services

+ first scale down desired count of tasks

`aws ecs update-service --cluster default --service service_name --desired-count 0 --region us-west-2`

## Delete Services

+ after scaling down; you can delete a service

`aws ecs delete-service --cluster default --service service_name --region us-west-2`

## Deregister Container Instances

+ before you can delete cluster; you must deregister container instances

`aws ecs deregister-container-instance --cluster default --container-instance container_instance_id --region us-west-2 --force`

## Delete a Cluster

+ finally after resources are removed, you can delete cluster

`aws ecs delete-cluster --cluster default --region us-west-2`

