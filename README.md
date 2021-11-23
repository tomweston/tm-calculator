# tm-calculator

A Golang API built with Mux Router and Prometheus Metrics ready to deploy to Kubernetes.

[Postman Collection](postman_collection.json)

## Deploying the tm-calculator API

To deploy it, run the following command:

```
kubectl apply -f https://raw.githubusercontent.com/tomweston/tm-calculator/master/manifest.yml
```

## Accessing the tm-calculator API

```sh
kubectl --namespace=tm-calculator port-forward svc/tm-calculator-service 5555:5555
```

## Testing

```sh
go test -v ./...
```

## Build & Run with Docker

```sh
docker build -t tm-calculator .
```
```
docker run -d -p 5555:5555 tm-calculator
```

## Local Install

```sh
go build
```
```sh
./tm-calculator
```

---

## Prometheus Metrics

[Metrics](http://127.0.0.1:5555/metrics)

### Useful Metrics

`processed_adds_total` - The total number of processed add events

`processed_subtracts_total` - The total number of processed subtraction events

`processed_division_total` - The total number of processed division events

`processed_random_total` - The total number of processed random events

---

## Kubernetes

### Deploy

**NOTE: This deploys by default the public tomweston/tm-calculator:latest image**

```sh
kubectl apply -f manifest.yml
```

### Port Forward

**NOTE: Binds to 9000 so as not to be confused with local service**

```sh
kubectl --namespace=tm-calculator port-forward svc/tm-calculator-service 9000:5555
```

### TODO: Add default return of 10 random numbers if count is not provided