# go-prometheus-service

1. mac vim  /opt/homebrew/etc/prometheus.yml

2. global:
scrape_interval: 15s

scrape_configs:
- job_name: "prometheus"
  static_configs:
    - targets: ["localhost:9090"]
- job_name: "faucet"
  static_configs:
    - targets: ["127.0.0.1:9185"]
- job_name: "go-prometheus-service"
  static_configs:
    - targets: ["127.0.0.1:9011"]
- job_name: "node3"
  static_configs:
    - targets: ["127.0.0.1:52319"]
- job_name: "node4"
  static_configs:
    - targets: ["127.0.0.1:52333"]