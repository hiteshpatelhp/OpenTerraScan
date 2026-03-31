## 🎉 Community Revival

> **OpenTerraScan** is the community-maintained fork of the project that Tenable decided to lock away behind a paywall because apparently "open source" was just a phase they were going through.
> Big thanks to Tenable for open-sourcing it long enough for us to fork it — truly a gift that keeps on giving, unlike their support contracts.
> This project is alive, maintained, and free — unlike Tenable's conscience when they archived it.

## OpenTerraScan

[![License: Apache 2.0](https://img.shields.io/badge/license-Apache%202-blue)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](../../pulls)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](code_of_conduct.md)

## Introduction

OpenTerraScan is a static code analyzer for Infrastructure as Code. OpenTerraScan allows you to:

- Seamlessly scan infrastructure as code for misconfigurations.
- Monitor provisioned cloud infrastructure for configuration changes that introduce posture drift, and enables reverting to a secure posture.
- Detect security vulnerabilities and compliance violations.
- Mitigate risks before provisioning cloud native infrastructure.
- Offers flexibility to run locally or integrate with your CI\CD.

## Key features
* 500+ Policies for security best practices
* Scanning of [Terraform](https://runopenterrascan.io/docs/usage/command_line_mode/#scanning-current-directory-containing-terraform-files-for-aws-resources) (HCL2)
* Scanning of AWS CloudFormation Templates (CFT)
* Scanning of Azure Resource Manager (ARM)
* Scanning of [Kubernetes](https://runopenterrascan.io/docs/usage/command_line_mode/#scanning-for-a-specific-iac-provider) (JSON/YAML), [Helm](https://runopenterrascan.io/docs/usage/command_line_mode/#scanning-a-helm-chart) v3, and [Kustomize](https://runopenterrascan.io/docs/usage/command_line_mode/#scanning-a-kustomize-chart)
* Scanning of [Dockerfiles](https://runopenterrascan.io/docs/usage/command_line_mode/#scanning-a-dockerfile)
* Support for [AWS](https://runopenterrascan.io/docs/policies/aws/), [Azure](https://runopenterrascan.io/docs/policies/azure/), [GCP](https://runopenterrascan.io/docs/policies/gcp/), [Kubernetes](https://runopenterrascan.io/docs/policies/k8s/), [Dockerfile](https://runopenterrascan.io/docs/policies/docker/), and [GitHub](https://runopenterrascan.io/docs/policies/github/)
* Integrates with docker image vulnerability scanning for AWS, Azure, GCP, Harbor container registries.

## Quick Start

1. [Install](#install)
2. [Scan](#scan)
3. [Integrate](#integrate)

### Step 1: Install

> **Note:** Packaged releases (AUR, Homebrew, Docker) are not yet available for OpenTerraScan. For now, build from source (see [Building OpenTerraScan](#building-openterrascan)) or download a binary from the [releases](../../releases) page once available.

#### Install as a native executable

```sh
  curl -L "$(curl -s https://api.github.com/repos/OpenTerraScan/OpenTerraScan/releases/latest | grep -o -E "https://.+?_Linux_x86_64.tar.gz")" > openterrascan.tar.gz
  tar -xf openterrascan.tar.gz openterrascan && rm openterrascan.tar.gz
  sudo install openterrascan /usr/local/bin && rm openterrascan
  openterrascan
```

### Step 2: Scan
To scan your code for security issues you can run the following (defaults to scanning Terraform).

```sh
$ openterrascan scan
```
**Note**: OpenTerraScan will exit with an error code if any errors or violations are found during a scan.

#### List of possible Exit Codes
| Scenario      | Exit Code |
| ----------- | ----------- |
| scan summary has errors and violations | 5 |
| scan summary has errors but no violations | 4 |
| scan summary has violations but no errors | 3 |
| scan summary has no violations or errors | 0 |
| scan command errors out due to invalid inputs | 1 |

### Step 3: Integrate with CI\CD

OpenTerraScan can be integrated into CI/CD pipelines to enforce security best practices.
Please refer to our [documentation to integrate with your pipeline](https://runopenterrascan.io/docs/integrations/).

## OpenTerraScan Commands
You can use the `openterrascan` command with the following options:

```sh
$ openterrascan
OpenTerraScan

Usage:
  openterrascan [command]

Available Commands:
  help        Help about any command
  init        Initialize OpenTerraScan
  scan        Detect compliance and security violations across Infrastructure as Code.
  server      Run OpenTerraScan as an API server
  version     OpenTerraScan version

Flags:
  -c, --config-path string   config file path
  -h, --help                 help for openterrascan
  -l, --log-level string     log level (debug, info, warn, error, panic, fatal) (default "info")
  -x, --log-type string      log output type (console, json) (default "console")
  -o, --output string        output type (human, json, yaml, xml) (default "Human")

Use "openterrascan [command] --help" for more information about a command.
```

## Policies
OpenTerraScan policies are written using the [Rego policy language](https://www.openpolicyagent.org/docs/latest/policy-language/). Every rego includes a JSON "rule" file which defines metadata for the policy.
By default, OpenTerraScan downloads policies from OpenTerraScan repositories while scanning for the first time. However, if you want to download the latest policies, you need to run the Initialization process. See [Usage](https://runopenterrascan.io/docs/usage/command_line_mode/) for information about the Initialization process.

Note: The scan command will implicitly run the initialization process if there are no policies found.

## Docker Image Vulnerabilities
You can use the `--find-vuln` flag to collect vulnerabilities as reported in its registry as part of OpenTerraScan's output. Currently OpenTerraScan supports Elastic Container Registry (ECR), Azure Container Registry, Google Container Registry, and Google Artifact Registry.

The `--find-vuln` flag can be used when scanning IaC files as follows:

```
$ openterrascan scan -i <IaC provider> --find-vuln
```

## Customizing scans

By default, OpenTerraScan scans your entire configuration against all policies. However, OpenTerraScan supports granular configuration of policies and resources.

For now, some quick tips:

- [Exclude a particular policy for a specific resource.](#How_to_exclude_a_policy_while_scanning_a_resource)
- [Manually configure policies to be suppressed or applied globally from a scan across all resources or, for just a particular resource.](#_How_to_include_or_exclude_specific_policies_or_resources_from_being_scanned)

### How to exclude a policy while scanning a resource

You can configure OpenTerraScan to skip a particular policy (rule) while scanning a resource. Follow these steps depending on your platform:

#### Terraform
Use Terraform scripts to configure OpenTerraScan to skip rules by inserting a comment with the phrase `"ts:skip=<RULENAME><SKIP_REASON>"`. The comment should be included inside the resource as shown in the example below.

![tf](docs/img/tf_skip_rule.png)

#### Kubernetes
In Kubernetes yamls, you can configure OpenTerraScan to skip policies by adding an annotation as seen in the snippet below.

![k8s](docs/img/skiprules.png)

### How to include or exclude specific policies or resources from being scanned

Use the OpenTerraScan config file to manually select the policies which should be included or excluded from the entire scan. This is suitable for edge use cases.
Use the "in-file" suppression option to specify resources that should be excluded from being tested against selected policies. This ensures that the policies are skipped only for particular resources, rather than all of the resources.

![config](https://user-images.githubusercontent.com/74685902/105115887-83e2f380-5a7e-11eb-82b8-a1d18c83a405.png)

### Sample scan output

OpenTerraScan's default output is a list of violations present in the scanned IaC. A sample output:

![Screenshot 2021-01-19 at 10 52 47 PM](https://user-images.githubusercontent.com/74685902/105115731-32d2ff80-5a7e-11eb-93b0-2f0620eb1295.png)

## Building OpenTerraScan
OpenTerraScan can be built locally. This is helpful if you want to be on the latest version or when developing OpenTerraScan. [gcc](https://gcc.gnu.org/install/) and [Go](https://go.dev/doc/install) 1.19 or above are required.

```sh
$ git clone git@github.com:OpenTerraScan/OpenTerraScan.git
$ cd OpenTerraScan
$ make build
$ ./bin/openterrascan
```

### To build your own docker, refer to this example (Alpine Linux):
```
FROM golang:alpine AS build-env

RUN apk add --update git

RUN git clone https://github.com/OpenTerraScan/OpenTerraScan && cd OpenTerraScan \
  && CGO_ENABLED=0 GO111MODULE=on go build -o /go/bin/openterrascan cmd/openterrascan/main.go

```

## Developing OpenTerraScan
To learn more about developing and contributing to OpenTerraScan, refer to the [contributing guide](CONTRIBUTING.md).

## Code of Conduct
We believe having an open and inclusive community benefits all of us. Please note that this project is released with a [Contributor Code of Conduct](code_of_conduct.md). By participating in this project you agree to abide by its terms.

## License

OpenTerraScan is licensed under the [Apache 2.0 License](LICENSE).
