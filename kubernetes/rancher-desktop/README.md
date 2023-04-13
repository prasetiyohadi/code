# Rancher-desktop

Original website: https://rancherdesktop.io/

## Setup

### Setup NGINX Ingress Controller

Prerequisites:

1. Uncheck `Enable Traefik` from `Kubernetes Settings` page to disable Traefik.
2. Deploy the NGINX ingress controller via `helm` or `kubectl`
3. Wait for the ingress pods to come up and running
4. Create a sample deployment and the associated service
5. Create an ingress resource
6. Forward a local port to the ingress controller

Commands:

```bash
# Use helm to deploy NGINX ingress controller
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace

# Use helmfile to deploy NGINX ingress controller
helmfile apply

# Wait for the ingress pods to come up and running
kubectl get pods --namespace=ingress-nginx

# Create a sample deployment and the associated service
kubectl create deployment demo --image=nginx --port=80
kubectl expose deployment demo

# Create an ingress resource
kubectl create ingress demo-localhost --class=nginx --rule="demo.localdev.me/*=demo:80"

# Forward a local port to the ingress controller (make sure you allow attaching to port < 1000 in sysctl.conf)
kubectl port-forward --namespace=ingress-nginx service/ingress-nginx-controller 80:80

# Try to access the demo application
curl http://demo.localdev.me
```
