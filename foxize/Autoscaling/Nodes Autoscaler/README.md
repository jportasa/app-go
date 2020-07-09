Docu: https://github.com/helm/charts/tree/master/stable/cluster-autoscaler

```
helm install cluster-autoscaler stable/cluster-autoscaler \
    --set cloudProvider=aws \
    --set awsRegion=eu-west-1 \
    --set autoDiscovery.clusterName=EKS-cluster-foxize \
    --set rbac.create=true
```
```
helm delete cluster-autoscaler
```

# Node utilitzation

```
kubectl get nodes --no-headers | awk '{print $1}' | xargs -I {} sh -c 'echo {}; kubectl describe node {} | grep Allocated -A 5 | grep -ve Event -ve Allocated -ve percent -ve -- ; echo'
```