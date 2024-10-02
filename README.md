#  k8s-scaling-demo

k8s-scaling-demo is a simple web application that visualizes Kubernetes autoscaling.

## Installation

Deploy the latest Kubernetes manifest to your cluster using `kubectl apply`.

```bash
kubectl apply -f https://raw.githubusercontent.com/DB-Vincent/k8s-scaling-demo/refs/heads/main/deployment.yaml
```

## Development

Feel free to adjust the application to your own liking.

### Frontend

The frontend has been built using Angular and sits in the `frontend/` directory.

1. Install the necessary packages: `npm install`
2. Run the development server: `npm run start`
3. Visit the application at http://localhost:4200

### Backend

The backend application uses the client-go package to communicate with the Kubernetes cluster. Since we're not running the application in a cluster while developping, we just need to specify two environment variables:

- `POD_NAME`: set this to a pod that is running in your current Kubernetes context
- `POD_NAMESPACE`: set this to the namespace of the pod you selected for the previous variable

Once that is done;

1. Install the dependencies: `go mod tidy`
2. Start the application: `go run main.go`
3. Visit the API at http://localhost:8080/api

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
