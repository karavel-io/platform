# Grafana component

Karavel provides an instance of [Grafana] managed by the [grafana-operator] that is used as the central
observability center for the entire cluster. Grafana aggregates logs, metrics and tracing information so that
they can be easily queried and correlated.

The [grafana-operator] allows to manage Grafana instances, dashboards and datasources as Kubernetes resources.
This allows users to add dashboards to their application deployment manifests and have them automatically
provisioned on Grafana. Check the [official documentation](https://github.com/integr8ly/grafana-operator/tree/master/documentation)
for more information.

## Components

- [Grafana]  
  Installs the [grafana-operator] and a cluster-wide Grafana instance.  
  KBT Import:
  ```yaml
  import_role:
    name: mikamai.karavel.grafana
  ```
  [config](../variables.md#grafana)
  

[Grafana]: https://grafana.com/oss/grafana
[grafana-operator]: https://github.com/integr8ly/grafana-operator
[Loki]: https://grafana.com/oss/loki
[Promtail]: https://grafana.com/docs/loki/latest/clients/promtail/
[Explorer tab]: https://grafana.com/docs/grafana/latest/explore/
