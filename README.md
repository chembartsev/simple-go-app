# simple-go-app

### A very simple Go app for [K3S Example Project](https://github.com/madlopt/k3s-example-project ) 

The logic is exactly the same as [here](https://github.com/madlopt/simple-nodejs-app ). So, it just providing the `/metrics` endpoint with Prometheus metrics and sending "I'm alive" log.

You can configure it via the following environment variables:

          - name: PROJECT_NAME
            value: "go-test-app"
          - name: CONSOLE_LOG_TIMEOUT
            value: "5000"
          - name: METRICS_ROUTE
            value: "/metrics"
          - name: METRICS_PORT
            value: "8080"
          - name: METRICS_COMPONENT
            value: "applications"
          - name: METRICS_SOURCE
            value: "go-test-app"
          - name: METRICS_NAMESPACE
            value: "default"

The app is just sending `PROJECT_NAME is alive` message to the console every `CONSOLE_LOG_TIMEOUT` milliseconds and exposing the metrics endpoint on the `METRICS_ROUTE` route and `METRICS_PORT` port.

### Metrics & Prom-Client Memory Leak

There are a lot of default metrics available, from the [client_golang](https://github.com/prometheus/client_golang/) library. And the one is custom one ``http_requests_counter`` which is a counter of all the requests to the application just as an example of custom metric you can add.

### DockerHub Image

The image I'm using for [K3S Example Project](https://github.com/madlopt/k3s-example-project ) is [there](https://hub.docker.com/r/madloptus/goapp) you can pull it by running `docker pull madloptus/goapp`

### Running the app locally

Not sure even if someone will need it, but if you want to run the app locally, there is the docket-compose file for that. Just run `docker-compose up` and the app will be available on the `http://localhost:8080`. 

