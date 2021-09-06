
# Observe without GitOps

## Infrastructure

```bash
cd infrastructure/
```

### 1. Namespaces

```bash
# Create required Namespaces
kustomize build 20_namespaces/ | kubectl apply -f -

# Check namespaces
kubectl get namespaces
```

### 2. Monitoring

```bash
# Deploy monitoring
kustomize build 30_monitoring/10_prometheus/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 30_monitoring/10_prometheus/20_operator/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/30_ns-rbac/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/40_prometheus/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n monitoring'
```

```bash
kustomize build 30_monitoring/20_kube-state-metrics/ | kubectl apply -f - && \
kustomize build 30_monitoring/30_node-exporter/ | kubectl apply -f -
```

### 3. Logging

```bash
# Deploy logging
kustomize build 40_logging/10_loki/ | kubectl apply -f - && \
kustomize build 40_logging/20_vector/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n logging'
```

### 4. Tracing

```bash
# Deploy tracing
kustomize build 50_tracing/10_jaeger/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 50_tracing/10_jaeger/20_operator/ | kubectl apply -f - && \
kustomize build 50_tracing/10_jaeger/30_jaeger/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n tracing'
```

### 5. Dashboards

```bash
# Deploy dashboards
kustomize build 60_dashboards/10_grafana/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n dashboards'
```

### 6. Databases

```bash
# Deploy databases
kustomize build 70_databases/postgresql/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n apps'
```

### 7. Brokers

```bash
# Deploy brokers
kustomize build 80_brokers/kafka/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 80_brokers/kafka/20_strimzi-operator/ | kubectl apply -f -
sleep 15 && \
kustomize build 80_brokers/kafka/30_kafka/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/40_topics/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/50_monitoring/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n apps'
```

### 8. Port forwarding

```bash
# Prometheus
kubectl port-forward svc/prometheus 9090 -n monitoring

# Check everyting is properly configured (targets and some metrics like "kube_pod_info")

# Jaeger
kubectl port-forward svc/jaeger-query 16686 -n tracing

# Check is running

# Leave Grafana port-forwarding open, we will use it a lot
kubectl port-forward svc/grafana 3000 -n dashboards

# Check everyting is properly configured (data sources and dashboards)
```

---

## Applications

### 1. Standalone application

```bash
# Deploy
kustomize build 90_apps/10_standalone/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=standalone -f -n apps
```

### 2. gRPC applications

```bash
# Deploy
kustomize build 90_apps/20_grpc/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=grpc-server -f -n apps
kubectl logs -l app=grpc-client -f -n apps
```

### 3. HTTP applications

```bash
# Deploy
kustomize build 90_apps/30_http/ | kubectl apply -f -

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=http-server -f -n apps
kubectl logs -l app=http-client -f -n apps

# Port forward
kubectl port-forward svc/http-client 8080 -n apps

# Now MANUAL make some requests using Postman
```

### 4. Broker applications

```bash
# Deploy
kustomize build 90_apps/40_broker/ | kubectl apply -f -

# Deploy
kubectl apply -f 100_gitops/8_apps.yaml

# Check pods
watch 'kubectl get pods -n apps'

# Check logs
kubectl logs -l app=kafka-producer -f -n apps
kubectl logs -l app=kafka-consumer -f -n apps
```
