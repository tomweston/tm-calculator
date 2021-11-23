# tm-calculator

A Golang API built with Mux Router and Prometheus Metrics ready to deploy to Kubernetes.

[API V1 Postman Collection](v1_postman_collection.json)

## Example Calls

- GET defaults count of 10 random numbers: `curl http://localhost:5555/api/v1/random`
- GET 10 random numbers: `curl http://localhost:5555/api/v1/random?num=100`
- GET - Add 20 to 10: `curl http://localhost:9000/api/v1/add?num1=20&num2=10`
- GET - Subtract 5 from 10: `curl http://localhost:5555/api/v1/subtract?num1=10&num2=5`
- GET - Divide 10 by 2: `curl http://localhost:5555/api/v1/division?num1=10&num2=2`

## Health

- GET - Readiness: `curl http://localhost:5555/readiness`
- GET - Liveness: `curl http://localhost:5555/liveness`

  
## Deploying the tm-calculator API to Kubernetes

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
kubectl --namespace=tm-calculator port-forward svc/tm-calculator-service 5555:5555
```