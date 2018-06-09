# Lecture: Describe and Display How Image Layers Work

When creating an image, a number of `intermediate` layers are completed
that, when aggregated together, form the image itself. In this lesson,
we talk about that process and how to display the details of each layer.

## Image History

+ `docker images` to see images
+ `docker image history mybuild:v3` to show you the image build history
+ Docker uses a UnionFS, allowing multiple branches to overlay to one FS
+ each layer has it's own information
+ `docker image history mybuild:v3 --no-trunc`
+ `history` tells us what the sizes of each layer are
+ in general, the aggregate should be the size of the image
+ `docker images mybuild:v3` to see file size; and then add up `history`
  it should be the same

## Finding Volumes, Images

+ `/var/lib/docker` is usually where this is stored