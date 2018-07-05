# Orchestration

## Apply Node Labels for Task Placement

Node labels are a really cool way for you to control how (and where)
your Docker services run in your swarm. This lesson will show you how
to use a label to apply service constraints to indicate where the
replicas for your service should run.

* `docker node ls` on the manager of the swarm
* you will see 3 nodes, a manager designated by leader and two worker nodes
* by doing a `docker node inspect` you can get all the info on that node
* this comes back in json format
* `docker node inspect --pretty [id]`
* labels can control resource deployments with names, for example
* you can update one of your nodes, adding a label with a value,
  and you can use it when you are creating a service to designate
  where the service goes
* `docker node ls` and we pick #3
* `docker node update --label-add mynode=testnode [id]`
* `docker node ls` doesn't look any diff
* `docker node inspect --pretty [id] | more` will show us a value
  that we didn't see before

## How do I Use this?

* we are going to be using constraint, which let's use node id,
  engine label, etc to determine how and where service runs
* we only want you to run it on this matching node
* `docker service ls`
* `docker service create --name constraints -p 80:80 --constraint `node.labels.mynode == testnode` --replicas=3 httpd`
* `docker service ls` will show us the three replicas running
* `docker service ps constraints` indicates 1, 2, and 3 are running on tcox3
* `docker service create --name normal nginx`
* `docker service ps normal` runs on tcox2
* `docker service ls`
* `docker service update --replicas 3 normal`
* `docker service ls`
* `docker service ps normal`
* `docker service ps constraints`