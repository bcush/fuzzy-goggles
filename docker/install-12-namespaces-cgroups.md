# Installation and Configuration

Understanding what namespaces and cgroups are will further your overall
understanding of how Docker (and containers in general) 'do what they do'
behind the scenes. This lesson will give you an overview of both.

## Docker and Namespaces

+ Provide `isolation` so that other pieces of the system remain unaffected
by whatever is within that namespace.

+ Uses namespaces of various kinds to provide isolation that containers
need to remain *portable* and refrain from affecting the remainder
of the host system.

### Namespace Types

+ Process ID

+ Mount

+ IPC

+ User (currently experimental support for)

+ Network

## Docker and Control Groups

+ Control Groups provide resource limitation and reporting capability within
the container space. They allow granular control over what the host resources
are allocated to the container(s) and when they are allocated.

### Common Control Groups

+ CPU

+ Memory

+ Network Bandwidth

+ Disk

+ Priority

## Summary

`Namespaces` provide isolation of the container itself from the underlying host.
`Control Groups` give you the ability to allocate things to that container.
Isolated regarding namespace, you still can control what
underlying resources are granted via CPU and memory. You can control by a
certain amount of CPU or memory that you will allow your underlying container
or service to use.