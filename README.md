# wee-web-test
Wee Web Test Docker

[![CodeQL](https://github.com/wee-ops/wee-web-test/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/wee-ops/wee-web-test/actions/workflows/codeql-analysis.yml)
[![Go](https://github.com/wee-ops/wee-web-test/actions/workflows/go.yml/badge.svg)](https://github.com/wee-ops/wee-web-test/actions/workflows/go.yml)
[![release](https://github.com/wee-ops/wee-web-test/actions/workflows/release.yaml/badge.svg)](https://github.com/wee-ops/wee-web-test/actions/workflows/release.yaml)

Web server for testing

## Docker usage

```bash
docker pull ghcr.io/wee-ops/wee-web-test:latest
```
## Environment vars

* HTTP_PORT=8080 (default)

## Endpoints

### / : Home

```html
<h1>Wee Web Test</h1>
<h2>Version development</h2>
<h3>Built @ now</h3>
<a href="../info">info</a>
```

### /info : Extra information

```json
{
  "BuildTime": "now",
  "BuildVersion": "development",
  "HostName": "c1e836be01d5",
  "RunningBy": "8.60959565s",
  "StartTime": "2022-03-08 03:32:48.737862704 +0000 UTC m=+0.000455213"
}
```
