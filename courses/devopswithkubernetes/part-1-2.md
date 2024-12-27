# Introduction to Debugging

Commands:

```bash
# Create a deployment with `apply` command from internet
kubectl apply -f https://raw.githubusercontent.com/kubernetes-hy/material-example/master/app1/manifests/deployment.yaml

# Describe a deployment
kubectl describe deployment hashgenerator-dep

# Describe a pod
kubectl describe pod hashgenerator-dep-75bdcc94c-whwsm

# Print pod logs
kubectl logs hashgenerator-dep-75bdcc94c-whwsm
```

Alternatively, you can use [Lens "The Kubernetes IDE"](https://k8slens.dev/) to interact with your Kubernetes workload in real time.
