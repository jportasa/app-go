# Nodegroups
https://eksctl.io/usage/managing-nodegroups/

## Create nodegroups
```
eksctl create nodegroup --config-file=cluster-foxize.yml
```
```
eksctl create nodegroup --config-file=cluster-foxize.yml --include=front
```

# Scale number of nodes in a nodegroup
```
eksctl get nodegroup --cluster=EKS-cluster-foxize --name=front
```
```
eksctl scale nodegroup --cluster=EKS-cluster-foxize --nodes=0 --name=front
```
# Drain Nodegroup
```
eksctl drain nodegroup --cluster=EKS-cluster-foxize --name=front
```

# Autoscaling

## HPA (Horitzontal Pod Autoscaling)
