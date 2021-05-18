{{/*
Expand the name of the chart.
*/}}
{{- define "velero.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "velero.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "velero.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "velero.labels" -}}
{{ include "velero.selectorLabels" . }}
karavel.io/component-version: {{ .Chart.Version }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "velero.selectorLabels" -}}
app.kubernetes.io/part-of: {{ include "velero.name" . }}
app.kubernetes.io/managed-by: karavel
karavel.io/component-name: {{ .Chart.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "velero.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "velero.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}
