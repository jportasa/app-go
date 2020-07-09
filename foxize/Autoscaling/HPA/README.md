# HPA Horitzontal Pod Autoscaler

Docu: https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/


Download last release:
```
helm pull stable/metrics-server
tar -xvf metrics-server-2.8.8.tgz
```
In values.yaml add in args section:
--kubelet-preferred-address-types=InternalIP

```
helm install metrics-server ./metrics-server \
    --namespace pro \
    --version 2.9.0 
```

Confirm that the metrics API is available
```
 kubectl get --raw "/apis/metrics.k8s.io/v1beta1/nodes"
```

```
cat << EoF > hpa.yaml
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: my-nginx      # Change
  namespace: pro    # Change
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: my-nginx    # Change
  minReplicas: 1        # Change
  maxReplicas: 5       # Change
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 10 # Change
EoF
```


Check HPA is working
```
kubectl get hpa -w
```
```
kubectl describe hpa my-nginx
```

Test the HPA autoscaling:
```
kubectl run -i --tty load-generator --image=busybox /bin/sh
```
```
while true; do wget -q -O - http://my-nginx; done
```
You probably will see <unknown>/50% for 1-2 minutes and then you should be able to see 0%/50%