# Contribution Guide

## Components

TODO

## CLI

TODO

## Documentation

The documentation is built with [mkdocs] and the [mkdocs-material] theme. It uses the [mkdocs-monorepo] plugin to include
documentation from the upstream components.  
To quickly preview changes made to the files in the [docs](./docs) folder, a Docker-based script
is provided at [hacks/mkdocs](./hacks/mkdocs). Simply running the script will start the development
server at [http://localhost:8000](http://localhost:8000).

[mkdocs]: https://mkdocs.org
[mkdocs-material]: https://squidfunk.github.io/mkdocs-material
[mkdocs-monorepo]: https://backstage.github.io/mkdocs-monorepo-plugin
