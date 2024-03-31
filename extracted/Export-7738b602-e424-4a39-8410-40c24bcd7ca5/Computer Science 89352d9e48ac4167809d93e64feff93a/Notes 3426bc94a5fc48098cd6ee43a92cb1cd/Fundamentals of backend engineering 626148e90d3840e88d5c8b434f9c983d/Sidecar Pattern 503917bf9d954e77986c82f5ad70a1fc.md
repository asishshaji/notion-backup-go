# Sidecar Pattern

> Thick clients, thicker backends
> 

Every protocol requires a library in an application.

![Untitled](Sidecar%20Pattern%20503917bf9d954e77986c82f5ad70a1fc/Untitled.png)

HTTP/2 client usually have both Http/2 and Http/1. The library makes the decision for you. 

![Untitled](Sidecar%20Pattern%20503917bf9d954e77986c82f5ad70a1fc/Untitled%201.png)

![Untitled](Sidecar%20Pattern%20503917bf9d954e77986c82f5ad70a1fc/Untitled%202.png)

![Untitled](Sidecar%20Pattern%20503917bf9d954e77986c82f5ad70a1fc/Untitled%203.png)

## Why do you need a sidecar pattern?

The client and server can use the basic protocols. Sidecar proxy can use the latest protocols. So if we need to update the protocol, we can directly update the protocols on the sidecar. 

![Untitled](Sidecar%20Pattern%20503917bf9d954e77986c82f5ad70a1fc/Untitled%204.png)

The client and server code still remains the same.