
# Cleanup

## With GitOps

`TODO`

## Without GitOps

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
