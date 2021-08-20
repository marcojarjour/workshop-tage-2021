
# TODO LIST

## Infrastructure

- [x] fix loki dashboards

- [x] fix grafana-self dashboard

- [x] test postgres

- [x] test kafka

- [ ] introduce argocd

- [ ] import dashboards
	- [ ] strimzi
	- [ ] postgresql

### Best effort

- [ ] create simple dashboards for custom apps
	- [ ] standadlone
	- [ ] http
	- [ ] http-db
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

## Applications

- [x] fix kafka apps

- [x] test http-db apps

- [x] test kafka apps

- [ ] fix http-client error `"error": "Get products failed: json: cannot unmarshal string into Go struct field Product.id of type int"`

### Best effort

- [ ] enable metrics on standalone app

- [ ] check kubernetes probes
	- [ ] standadlone
	- [ ] http
	- [ ] http-db
	- [ ] grpc
	- [ ] kafka
