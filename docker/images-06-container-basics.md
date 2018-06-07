# Image Creation, Management, and Registry

## Container Basics - Running, Attaching to, and Executing Commands in Containers

Although the basics of running containers are not emphasized on the exam,
it is assumed you have the ability to run and understand the options for
running containers. Let's take a moment and be sure we have a good idea how
to start, run, attach to and work with containers before going forward.

## Running Images/Containers

+ `docker images` to see images
+ `docker ps` to see our running images
+ `docker ps -a` to see what processes have stopped and not removed
+ `docker run centos:6` to create a container based on this image
+ whatever is associated with this image to start will do so
+ `docker ps` shows you nothing running
+ `docker ps -a` shows you that a container _was_ running
+ since we didn't specify anything, `/bin/bash` started up and closed
+ `docker run -i -t` for interactive (attach to stdin); t is current terminal
+ `docker run -it` will also work
+ notice after running we are now interactively running
+ `exit` will then leave the session and close the container
+ `docker run -it --name testcontainer centos:6 /bin/bash` and you'll be in a prompt
+ `exit` and then see in `docker ps -a` that it ran
+ `docker run -it --name testcontainer centos:6 /bin/bash` see you cannot run it
+ `docker rm` to container name/id that you pass in
+ `docker rm `docker ps -a -q`` to delete all your containers
+ `docker run rm -it --rm --name testcontainer centos:6 /bin/bash` to remove after ended
+ `docker run -it --rm --env MYVAR=whatever --name testcontainer centos:6 /bin/bash`
+ `echo $MYVAR` and see that it was passed through
+ `docker run -it --rm --name testcontainer --env MYVAR=whatever  --env MYVAR2=test``

## Running Container Detached

+ `docker run -d --rm --env MYVAR=wahtever --env MYVAR2=whaaa --name testcontainer centos:6 /bin/bash`
+ `docker pull httpd` for the official apache image
+ `docker run -d httpd` to run it in deatched mode
+ `docker ps` to see it runing

## Not Attached to Container

+ `docker attach` to something will exit if we try to exit
+ `docker exec -it distracted_kirch /bin/bash`

## Difference Between Attach and Exec

+ `attach` will attach to exact instance and app
+ `exec` will execute anything; and will exit the container not the process