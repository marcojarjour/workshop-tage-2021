
# TODO LIST

## Infrastructure

- [x] fix loki dashboards

- [x] fix grafana-self dashboard

- [x] test postgres

- [x] test kafka

- [x] import dashboards
	- [x] strimzi
	- [x] postgresql

- [ ] introduce gitops `WIP`
	- [ ] argocd
	- [ ] argocd apps
	- [ ] kustomize in all folders

- [ ] fix postgresql grafana dashboard

### Best effort

- [ ] create simple dashboards for custom apps
	- [ ] standadlone
	- [ ] http
	- [ ] grpc
	- [ ] kafka

- [ ] upgrades
	- [ ] prometheus
	- [ ] kube-state-metrics
	- [ ] node-exporter
	- [ ] grafana
	- [ ] loki
	- [ ] vector
	- [ ] jaeger
	- [ ] strimzi

## Applications

- [x] fix kafka apps

- [x] test http-db apps

- [x] test kafka apps

### Best effort

- [ ] enable metrics on standalone app

- [ ] check kubernetes probes
	- [ ] standadlone
	- [ ] http
	- [ ] grpc
	- [ ] kafka
