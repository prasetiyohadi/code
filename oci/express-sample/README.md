# express-sample

[Source](https://github.com/rancher-sandbox/docs.rancherdesktop.io/tree/main/assets/express-sample)

## Commands

```bash
nerdctl -h
nerdctl images
nerdctl --namespace k8s.io build -t expressapp:v1.0 .
kubectl run --image expressapp:v1.0 expressapp
kubectl port-forward pods/expressapp 3000:3000
curl localhost:3000
```
