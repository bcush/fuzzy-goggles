# Orchestration

## Increase and Decrease the Number of Replicas in a Service

There are a couple of ways to handle scaling in a running service.
We can use one command for scaling replicas in a single service
and an equivalent that can do both that AND scale multiple services
at a same time.

* `docker service ls`
* `docker node ls`
* `docker images`
* `docker pull nginx` because we want to try a few scenarios
* `docker service create --name testweb -p 80:80 httpd`
* once we create this, we should have one replica running
* `docker service ps testweb`
* this shows that we have one instance or replica running
* `docker service update --replicas 3 --detach==false testweb`
* `docker service ps testweb` and we see all three running across all nodes
* what happens if we have a bunch of diff services
* `docker service create --name testnginx -p 5901:80 nginx`
* `elinks http://tcox1:5901` and we see the nginx site come up
* `docker service ls` we can see that we have two services running
* `testweb` and `testnginx` running w/ diff replicas
* can i update both w/ a single command?
* `docker service update --replicas 5 testweb testnginx`
* it fails because docker service update requires 1 argument
* `docker service scale` is equivalent to the previous for single/multiple services
* `docker service scale --deatch=false testnginx=3`
* this will look and feel like exactly like using update
* it will be considered converged meaning it's occured across all
* with `docker service scale --deatch=false testnginx=4 testweb=2`

* you will see this on the exam, so understand that
* `docker service scale` and `docker service update replicas` on a single
  service is the same command
* the only way to modify the number of replicas on more than one
  service at a time is to use `docker service scale`