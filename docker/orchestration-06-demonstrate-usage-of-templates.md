# Orchestration

## Lecture: Demonstrate the Usage of Templates with `docker service create`

Templates give us greater flexibility over
a number of options when we create a service.
This lesson will show you how to set various
options using templating and how to display
those values after the fact.

* templates can only be used with certain values
* specifically, hostnames, mounts, env variables
* there are also a certain num of the template ids
* service id, service name, service labels, task id, task name, task something

* `docker service create --name myweb --hostname="{{.Node.ID}}-{{.Service.Name}}" --detach=false httpd`
* `docker service ps --no-trunc myweb`
* shows myweb.1 running service
* `docker inspect --format="{{.Config.Hostname}}" myweb.1.[whatver id of the node]`
* `docker inspect myweb.[id here] > output.json`
* cat output.json which will give us all the info about the specific instance
