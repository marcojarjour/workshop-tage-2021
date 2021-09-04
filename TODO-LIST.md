
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
	- [x] kustomize in all folders

- [ ] fix all grafana dashboard
	- [ ] postgresql
	- [ ] others?

- [ ] create dashboard for all custom apps
	- [ ] standadlone
	- [ ] http
	- [ ] grpc
	- [ ] kafka

### Best effort

- [ ] upgrades
	- [ ] prometheus
	- [ ] kube-state-metrics
	- [ ] node-exporter
	- [ ] grafana
	- [ ] loki
	- [ ] vector
	- [ ] jaeger
	- [ ] strimzi

---

## Applications

- [x] fix kafka apps

- [x] test http-db apps

- [x] test kafka apps

- [x] enable metrics on standalone app

- [x] migrate to go 1.17

- [x] introduce custom metrics

- [x] update env-vars
