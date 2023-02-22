# go-gin-demo

A sample web application written in Go for [Tanzu Application Platform](https://docs.vmware.com/en/VMware-Tanzu-Application-Platform/index.html).

## Prerequisites

* Installed the `tanzu` CLI
* Installed the Tanzu Application Platform
* Set up a developer namespace

## How to deploy

Using `config/workload.yaml`, simply deploy it using:

```sh
tanzu apps workload apply -f config/workload.yaml -n DEVELOPER-NAMESPACE
```

## Troubleshooting

In case the version(s) of Go Buildpack is not compatible with the version specified in `go.mod`, please update the Go version with `go mod edit` command.

```sh
go mod edit -go=1.19
```
