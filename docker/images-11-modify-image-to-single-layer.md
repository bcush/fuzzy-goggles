# Lecture: Modify an Image to a Single Layer

To optimize our images and save space on the underlying host,
you may choose to `flatten` the layers of your image to a
single storage layer. We will show you a way to do so in this lesson.

## Officially No Way to Flatten an Image to a Single Layer

+ `docker squash` is an external utility for this; but not on exam
+ we can get an image to a single layer in an optimized manner
+ `docker images` to see images
+ `docker image history mybuild:v3` and see the 8-layers
+ we can't squash it down to one or two layers
+ `docker run mybuild:v3` which just returns folder count
+ `docker export eloquent_thompson > mybuild4.tar`
+ basically, create a container, export the container to a tar
+ now you can move to another system if you want
+ now you reimport and tag as you want, then tag as you like
+ now you have a flattened image, but you lose the history

## Reimport Image

+ `docker import mybuildv4.tar mybuildimport:v5`
+ this method described is the only way to officially flatten