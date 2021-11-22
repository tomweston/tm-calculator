# tm-calculator

A Golang API built with Mux Router and Prometheus Metrics ready to deploy to Kubernetes.

[Postman Collection](postman_collection.json)

## Build & Run with Docker
```
docker build -t tm-calculator .
```
```
docker run -d -p 5555:5555 tm-calculator
```

## Local Install
```
go build
```
```
./tm-calculator
```

## Prometheus Metrics

[Metrics](http://127.0.0.1:5555/metrics)

### Useful Metrics

`processed_adds_total`

`processed_subtracts_total`

`processed_division_total`

`processed_random_total`

## Kubernetes

### Deploy

**NOTE: This deploys by default the public tomweston/tm-calculator:latest image**

```
kubectl apply -f manifest.yml
```

### Port Forward

**NOTE: Binds to 9000 so as not to be confused with local service**

```
kubectl --namespace=tm-calculator port-forward svc/tm-calculator-service 9000:5555
```
