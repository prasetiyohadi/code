# First Deploy

## What are microservices?

For this course we'll choose the definition set by Sam Newman in Building Microservices: "Microservices are small, autonomous services that work together". The opposite of a microservice is a service that is self-contained and independent called a Monolith.

## What is Kubernetes?

"Kubernetes (K8s) is an open-source system for automating deployment, scaling, and management of containerized applications. It groups containers that make up an application into logical units for easy management and discovery." - kubernetes.io

## First Deploy

Materials:

* [Repository](https://github.com/kubernetes-hy/material-example)

### Deployment

Commands:

```bash
# Get cluster info
kubectl cluster-info

# Create a deployment
kubectl create deployment hashgenerator-dep --image=jakousa/dwk-app1

# Describe pod object
kubectl explain pod

# View pods
kubectl get pods

# View deployments
kubectl get deployments
```

Definitions:

* Pod: Pod is an abstraction around one or more containers.
* Deployment: A Deployment resource takes care of deployment. It's a way to tell Kubernetes what container you want, how they should be running and how many of them should be running.

### Declarative configuration with YAML

Commands:

```bash
# Create a deployment
kubectl create deployment hashgenerator-dep --image=jakousa/dwk-app1

# Scale a deployment
kubectl scale deployment/hashgenerator-dep --replicas=4

# Update image of a deployment
kubectl set image deployment/-dep dwk-app1=jakousa/dwk-app1:b7fc18de2376da80ff0cfc72cf581a9f94d10e64

# Delete a deployment
kubectl delete deployment hashgenerator-dep
```

Files:

* [deployment.yaml](material-example/app1/manifests/deployment.yaml)

Commands:

```bash
# Create a deployment with `apply` command
kubectl apply -f material-example/app1/manifests/deployment.yaml

# Delete a deployment from file
kubectl delete -f material-example/app1/manifests/deployment.yaml

# Recreate a deployment with `apply` command from internet
kubectl apply -f https://raw.githubusercontent.com/kubernetes-hy/material-example/master/app1/manifests/deployment.yaml
```
