web-shipping
============

This image runs a web service in 8080 port, used for show microservice concept. It comes from rawmind/alpine-base.

## Build

```
docker build -t rawmind/web-shipping:<version> .
```

## Versions

- `0.1` [(Dockerfile)](https://github.com/rawmind0/web-shipping/blob/0.1/Dockerfile)


## Usage

```
docker run rawmind/web-shipping:<version> 
```

## Example

See [rancher-example][rancher-example], rancher catalog package that runs web-monolith in a cattle environment.

## Microservice uri

- / web-shipping app

[rancher-example]: https://github.com/rawmind0/web-shipping/tree/master/rancher