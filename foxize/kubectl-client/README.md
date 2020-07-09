# Tools

aws cli
eksctl
kubectl

kube-prompt:
https://github.com/c-bata/kube-prompt

# Namespace
```
kubectl config set-context --current --namespace=pro
```

# Kube-config
Está en ~/.kube
```
kubectl config get-contexts
```
```
aws eks update-kubeconfig --name EKS-cluster-foxize
```

# Accessing to cluster insides
```
kubectl port-forward service/my-nginx 7000:80
```
Next you can access directly: curl http://127.0.0.1:7000/

# Nodes CPU/Mem usage
kubectl top node