Docu ALB: https://github.com/helm/charts/tree/master/incubator/aws-alb-ingress-controller

Docu AWS Ingress: https://github.com/kubernetes-sigs/aws-alb-ingress-controller

# Ingress controller:

Requiere tagear las subnets:
Publicas:
    kubernetes.io/role/elb: (empty value)
Privadas:
    kubernetes.io/role/internal-elb: (empty value)

```
helm install alb-ingress incubator/aws-alb-ingress-controller \
--set clusterName=EKS-cluster-foxize \
--set autoDiscoverAwsRegion=true \
--set autoDiscoverAwsVpcID=true \
--namespace pro
```

# Create self-signed certificate if needed

https://www.selfsignedcertificate.com
