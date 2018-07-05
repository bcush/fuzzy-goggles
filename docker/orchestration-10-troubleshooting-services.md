# Orchestration

## Lecture: Identify the Steps Needed to Troubleshoot a Service Not Deploying

Troubleshooting is a skill, like any other, learned through experience. You have gained some throughout the course, let's focus on the most common things you can think through when doing so.

### Docker Swarm - Service Troubleshooting

whenever you are having issues with a service not deploying, there are
several steps you can try (all of which you have seen throughout this course at one time or another)
* `docker node ls` are all the nodes available and show that way in your cluster?
* `docker service ps [servicename]` is the service partially running? is the problem specific to the number of replicas?
* `docker service inspect [servicename]` (with/without --format for output) did you apply labels and od they show? did you deploy your service with a constraint but you mismatched the value and the one used on launch?
* `docker ps` is your cluster a `locked` cluster that needs a key for operation (on restart)?
* did you recently restore a swarm configuration and the expected service is not starting? make sure you re-initialize the new manager you restored with the --force-new-cluster so that it is not attempting to contact previous nodes
* have you recently updated docker? make sure you are using the official docker repositories or the setup script from docker.com to avoid version issues or feature unavailability.

### Troubleshooting Requires Experience

as in every technology, the ability to troubleshoot problems will rely on your experience with how the software behaves normally; in addition to docker issues, be sure you can rule out:

* selinux issues - try `setenforce 0` to put selinux into `passive` mode and retry your task
* permissions - if you are launching a service as an unpriviledged user, be sure resources (like voumes/mounts) that you may be using are available to your UID/GID
* cpu/mem - when constraining your containers to a host, be sure that it has the necessaryresources to meet the needs of the service
* routing - make sure that endpoints using the service are on the same network segement or have the necessary routing to get there
* firewall - be sure the appropriate ports and protocols are enabled to the destination IPs