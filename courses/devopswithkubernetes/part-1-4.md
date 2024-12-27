# Introduction to Storage

There are two things that are known to be difficult with Kubernetes:

* [Networking](https://www.youtube.com/watch?v=GgCA2USI5iQ)
* [Storage](https://softwareengineeringdaily.com/2019/01/11/why-is-storage-on-kubernetes-is-so-hard/)

## Simple Volume

`emptyDir` volumes are shared filesystems inside a pod

Cons:

* their lifecycle is tied to a pod, when the pod is destroyed the data is lost
* moving the pod from another node will destroy the contents of the volume as the space is reserved from the node the pod is running on

Pros:

* it may be used as a cache as it persists between container restarts
* it can be used to share files between two containers in a pod

Files:

* [deployment.yaml](./material-example/app3/manifests/deployment.yaml)
* [service.yaml](./material-example/app3/manifests/service.yaml)
* [ingress.yaml](./material-example/app3/manifests/ingress.yaml)

Commands:

```bash
kubectl apply -f ./material-example/app3/manifests/deployment.yaml
kubectl apply -f ./material-example/app3/manifests/service.yaml
kubectl apply -f ./material-example/app3/manifests/ingress.yaml
curl dwk-app3.localdev.me -o image.jpg
md5sum image.jpg
kubect get pods
kubectl exec images-dep-7bfb7955c6-bdb62 -c image-finder -- md5sum files/image.jpg
kubectl exec images-dep-7bfb7955c6-bdb62 -c image-response -- md5sum files/image.jpg
```
