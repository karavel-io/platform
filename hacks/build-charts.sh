#!/usr/bin/env bash

rm -rf dist && mkdir -p dist
helm package platform/charts/* -d dist && helm repo index dist --url https://charts.mikamai.com/karavel
