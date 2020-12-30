# Logging components

Karavel provides a robust logging stack based on open source projects released by [Grafana Labs].
The core of the stack revolves around [Grafana], the well-known dashboard and analytics tool, as
it aggregates all the observability components offered by the Karavel platform.

The log aggregation and query pipeline is provided by [Loki], a very lightweight and powerful 
server developed by Grafana. Logs are collected by a [Promtail] agent installed on each Kubernetes node
and are then accessible and queryable via the [Explorer tab] in Grafana.

Other logging solutions will be available as alternatives in the future.

## Components

- [Loki]  
  Logging stack based on Loki and Promtail. Requires `mikamai.karavel.grafana` to be installed as well  
  KBT Import:
  ```yaml
  import_role:
    name: mikamai.karavel.logging_loki
  ```
  [config](../variables.md#loki)
  

[Grafana Labs]: https://grafana.com/oss/
[Grafana]: ./grafana.md
[Loki]: https://grafana.com/oss/loki
[Promtail]: https://grafana.com/docs/loki/latest/clients/promtail/
[Explorer tab]: https://grafana.com/docs/grafana/latest/explore/
