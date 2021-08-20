
# Instructions (without GitOps)

## Preparation

### Repo

git clone git@github.com:bygui86/workshop-tage-2021.git

cd workshop-tage-2021/

### Kubernetes cluster

### 1. GKE cluster

`TODO`

### 2. CIVO cluster

`TODO`

### 3. Minikube

make start-docker-minikube

make start-hyperkit-minikube

make start-virtualbox-minikube (slower in provisioning and container image download)

---

## Infrastructure

cd infrastructure/

### 1. Namespaces

kustomize build 20_namespaces/ | kubectl apply -f -

### 2. Monitoring

kustomize build 30_monitoring/10_prometheus/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 30_monitoring/10_prometheus/20_operator/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/30_ns-rbac/ | kubectl apply -f - && \
kustomize build 30_monitoring/10_prometheus/40_prometheus/ | kubectl apply -f -

kustomize build 30_monitoring/20_kube-state-metrics/ | kubectl apply -f - && \
kustomize build 30_monitoring/30_node-exporter/ | kubectl apply -f -

### 3. Logging

kustomize build 40_logging/10_loki/ | kubectl apply -f - && \
kustomize build 40_logging/20_vector/ | kubectl apply -f -

### 4. Tracing

kustomize build 50_tracing/10_jaeger/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 50_tracing/10_jaeger/20_operator/ | kubectl apply -f - && \
kustomize build 50_tracing/10_jaeger/30_jaeger/ | kubectl apply -f -

### 5. Grafana

kustomize build 60_grafana/ | kubectl apply -f -

### 6. Port forwarding

kubectl port-forward svc/prometheus 9090 -n monitoring
kubectl port-forward svc/grafana 3000 -n monitoring
kubectl port-forward svc/jaeger-query 16686 -n tracing

---

## Applications

### Standalone

kustomize build 90_apps/10_standalone/ | kubectl apply -f -

### gRPC

kustomize build 90_apps/20_grpc/grpc-client/ | kubectl apply -f - && \
kustomize build 90_apps/20_grpc/grpc-server/ | kubectl apply -f -

### HTTP

kustomize build 70_databases/postgresql/ | kubectl apply -f -

kustomize build 90_apps/30_http/http-client-db/ | kubectl apply -f - && \
kustomize build 90_apps/30_http/http-server-db/ | kubectl apply -f -

kubectl port-forward svc/http-client-db 8080 -n apps

`MANUAL` make some requests using Postman

### Broker

kustomize build 80_brokers/kafka/10_crds/ | kubectl apply -f - && \
sleep 5 && \
kustomize build 80_brokers/kafka/20_strimzi-operator/ | kubectl apply -f -

kustomize build 80_brokers/kafka/30_kafka/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/40_topics/ | kubectl apply -f - && \
kustomize build 80_brokers/kafka/50_monitoring/ | kubectl apply -f -

kustomize build 90_apps/40_broker/kafka-consumer/ | kubectl apply -f - && \
kustomize build 90_apps/40_broker/kafka-producer/ | kubectl apply -f -

### Cleanup

kustomize build 90_apps/10_standalone/ | kubectl delete -f - && \

kustomize build 90_apps/20_grpc/grpc-client/ | kubectl delete -f - && \
kustomize build 90_apps/20_grpc/grpc-server/ | kubectl delete -f - && \

kustomize build 90_apps/30_http/http-client-db/ | kubectl delete -f - && \
kustomize build 90_apps/30_http/http-server-db/ | kubectl delete -f - && \
kustomize build 70_databases/postgresql/ | kubectl delete -f -

kustomize build 90_apps/40_broker/kafka-consumer/ | kubectl delete -f - && \
kustomize build 90_apps/40_broker/kafka-producer/ | kubectl delete -f - && \
kustomize build 80_brokers/kafka/50_monitoring/ | kubectl delete -f - && \
kustomize build 80_brokers/kafka/40_topics/ | kubectl delete -f - && \
kustomize build 80_brokers/kafka/30_kafka/ | kubectl delete -f - && \
kustomize build 80_brokers/kafka/20_strimzi-operator/ | kubectl delete -f - && \
kustomize build 80_brokers/kafka/10_crds/ | kubectl delete -f -
