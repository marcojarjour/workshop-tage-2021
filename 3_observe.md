
# Observe without GitOps

## Infrastructure

```bash
cd infrastructure/
```

### 1. Namespaces

```bash
kustomize build 20_namespaces/ | kubectl apply -f -
```

### 2. Monitoring

```bash
kustomize build 30_monitoring/10_prometheus/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 30_monitoring/10_prometheus/20_operator/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/30_ns-rbac/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/40_prometheus/ | kubectl apply -f -
```

```bash
kustomize build 30_monitoring/20_kube-state-metrics/ | kubectl apply -f - && \
kustomize build 30_monitoring/30_node-exporter/ | kubectl apply -f -
```

### 3. Logging

```bash
kustomize build 40_logging/10_loki/ | kubectl apply -f - && \
kustomize build 40_logging/20_vector/ | kubectl apply -f -
```

### 4. Tracing

```bash
kustomize build 50_tracing/10_jaeger/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 50_tracing/10_jaeger/20_operator/ | kubectl apply -f - && \
kustomize build 50_tracing/10_jaeger/30_jaeger/ | kubectl apply -f -
```

### 5. Grafana

```bash
kustomize build 60_grafana/ | kubectl apply -f -
```

### 6. Port forwarding

```bash
kubectl port-forward svc/prometheus 9090 -n monitoring
kubectl port-forward svc/grafana 3000 -n monitoring
kubectl port-forward svc/jaeger-query 16686 -n tracing
```

---

## Applications

### Standalone

```bash
kustomize build 90_apps/10_standalone/ | kubectl apply -f -
```

### gRPC

```bash
kustomize build 90_apps/20_grpc/grpc-client/ | kubectl apply -f - && \
kustomize build 90_apps/20_grpc/grpc-server/ | kubectl apply -f -
```

### HTTP

```bash
kustomize build 70_databases/postgresql/ | kubectl apply -f -
```

```bash
kustomize build 90_apps/30_http/http-client-db/ | kubectl apply -f - && \
kustomize build 90_apps/30_http/http-server-db/ | kubectl apply -f -
```

```bash
kubectl port-forward svc/http-client-db 8080 -n apps

# Now MANUAL make some requests using Postman
```

### Broker

```bash
kustomize build 80_brokers/kafka/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 80_brokers/kafka/20_strimzi-operator/ | kubectl apply -f -
```

```bash
kustomize build 80_brokers/kafka/30_kafka/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/40_topics/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/50_monitoring/ | kubectl apply -f -
```

```bash
kustomize build 90_apps/40_broker/kafka-consumer/ | kubectl apply -f - && \
kustomize build 90_apps/40_broker/kafka-producer/ | kubectl apply -f -
```
