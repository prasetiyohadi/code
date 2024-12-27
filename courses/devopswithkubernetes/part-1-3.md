# Introduction to Networking

## Connecting from outside of the cluster

Commands:

```bash
# Create a deployment of simple networking application
kubectl apply -f https://raw.githubusercontent.com/kubernetes-hy/material-example/master/app2/manifests/deployment.yaml

# View pods
kubectl get pods

# Forward the pod's service port to localhost
kubectl port-forward hashresponse-dep-57bcc888d7-dj5vk 3003:3000
```

## What is a Service?

Definitions:

* Service: Service resource will take care of serving the application to connections from outside of the cluster.

Files:

* [nodeport_service.yaml](./material-example/app2/manifests/nodeport_service.yaml)

Commands:

```bash
# Create a service with kind NodePort
kubectl apply -f ./material-example/app2/manifests/nodeport_service.yaml

# View services
kubectl get services

# Access the service
curl localhost:30080
```

## What is an Ingress?

Definitions:

* Ingress: Incoming Network Access resource Ingress is a completely different type of resource from Services. If you've got your OSI model memorized, it works in layer 7 while services work on layer 4.

Files:

* [service.yaml](./material-example/app2/manifests/service.yaml)
* [ingress.yaml](./material-example/app2/manifests/ingress.yaml)

Commands:

```bash
# Delete existing service
kubectl delete -f ./material-example/app2/manifests/nodeport_service.yaml

# Create a service with kind ClusterIP for Ingress
kubectl apply -f ./material-example/app2/manifests/service.yaml

# Create an ingress
kubectl apply -f ./material-example/app2/manifests/ingress.yaml

# View services and ingresses
kubectl get svc,ing

# Access the application
curl http://dwk-app2.localdev.me
```
