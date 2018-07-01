# Orchestration

## Lecture: Running Replicas vs. Global Services

You can run your service in either replicated or global mode. The
difference is subtle but important; let's talk a bit about that
difference and then demonstrate the behavior in our swarm.

* replicas just means that the service will be replicated a number
  of times to a specific set of nodes, running as many tasks as
  specified by your command
* let's look at our swarm
* `docker service ls`
* `docker images`
* creating a service in replicated mode, meaning that we control
  exactly how many instances of this application happen to running
  across the nodes in our cluster
* `docker service create --name testweb -p 80:80 httpd`
* we are creating a single instance, or replicas, of the httpd
  service running across my cluster
* `docker service ls` we see there is only one replicas running
* `docker ps` tells us it's running on tcox-3
* `docker service scale testweb=3` to go to 3
* `docker service ps testweb` and we have 3 instances running
* `docker service ls`

## We can also run services in Global Mode

* global mode means that I want to run my application across all the nodes
* by this i lose the granular control over the number of replicas

* `docker service create --name testnginx -p 5901:80 --mode global --detached=false nginx`
* creating a global service called testnginx running on port 5901
  redirecting to 80 in detached mode so we can see the deployment
* even though i dont indicate that i want more than one instance running,
  i still get 3 out of 3 tasks; and as such `docker service ls` shows
  both running, and 3 replicas, and the mode being in global mode
* `docker service ps testnginx` and we can see that we are currently
  running across all three of our nodes
* can i scale? `docker service update --replicas 5 testnginx`
* this tell us that replicas can only be used with replicated mode
* `docker service scasle testnginx=1`
* scale can only be used with replicated mode
* so, globally means that regardless, normally when i start a service,
  and i don't specify the number of replicas, it's only going to
  run across 1, unless you give it information on how to replicate itself,
  or how to scale itself
* global means that i will always have a replica running across each of
  my nodes
* this is useful because if i am scaling by adding add'l nodes and 
  then joining them to the cluster, which i can do in a running cluster,
  one i scale w/ additional nodes, then the service wil automatically
  create a service on that node; this is diff with replicas, because
  if i scale an additional node, that doesn't mean i get a new node
* can i update my global service to be replicated?
* `docker service ls`
* `docker service update --mode replicated testnginx` this will fail
* there is no mode update in a service config, once you indicate that you
  want to be in global and by default replicated mode, you cannot
  change back to replicated mode
* you cannot change a replicated service into a global service

* just know that this is the two subtle differences