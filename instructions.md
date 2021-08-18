
# Instructions (without GitOps)

## Preparation

git clone git@github.com:bygui86/workshop-tage-2021.git

cd workshop-tage-2021/

### Kubernetes cluster

### 1. GKE cluster

`TODO`

### 2. CIVO cluster

`TODO`

### 3. Minikube

make start-docker-minikube
make start-virtualbox-minikube

## Infrastructure

cd infrastructure/

### 1. Namespaces

kustomize build 10_namespaces/ | kubectl apply -f -

### 2. Monitoring

kustomize build 30_monitoring/10_prometheus/10_crds/ | kubectl apply -f -
kustomize build 30_monitoring/10_prometheus/20_operator/ | kubectl apply -f -
kustomize build 30_monitoring/10_prometheus/30_ns-rbac/ | kubectl apply -f -
kustomize build 30_monitoring/10_prometheus/40_prometheus/ | kubectl apply -f -

kustomize build 30_monitoring/20_kube-state-metrics/ | kubectl apply -f -
kustomize build 30_monitoring/30_node-exporter/ | kubectl apply -f -

### 3. Logging

kustomize build 40_logging/10_loki/ | kubectl apply -f -
kustomize build 40_logging/20_vector/ | kubectl apply -f -

### 4. Tracing

kustomize build 50_tracing/10_jaeger/10_crds/ | kubectl apply -f -
kustomize build 50_tracing/10_jaeger/20_operator/ | kubectl apply -f -
kustomize build 50_tracing/10_jaeger/30_jaeger/ | kubectl apply -f -

### 5. Grafana

kustomize build 60_grafana/ | kubectl apply -f -

### 6. Port forwarding

kubectl port-forward svc/prometheus 9090 -n monitoring
kubectl port-forward svc/grafana 3000 -n monitoring

`TILL HERE EVERYTHING TESTED`

### 6. Applications

#### Standalone

kustomize build 100_applications/10_standalone/ | kubectl apply -f -

#### HTTP

kustomize build 100_applications/20_http/http-client/ | kubectl apply -f -
kustomize build 100_applications/20_http/http-server/ | kubectl apply -f -

#### HTTP with DB backbone

`TODO` 70_databases/
kustomize build 100_applications/20_http/http-client-db/ | kubectl apply -f -
kustomize build 100_applications/20_http/http-server-db/ | kubectl apply -f -

#### gRPC

kustomize build 100_applications/30_grpc/grpc-client/ | kubectl apply -f -
kustomize build 100_applications/30_grpc/grpc-server/ | kubectl apply -f -

#### Broker

`TODO` 80_brokers/

kustomize build 100_applications/40_broker/kafka-consumer/ | kubectl apply -f -
kustomize build 100_applications/40_broker/kafka-producer/ | kubectl apply -f -
