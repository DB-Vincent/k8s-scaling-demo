#  k8s-scaling-demo

k8s-scaling-demo is a simple web application that visualizes Kubernetes autoscaling.

## Installation

### Using Kubernetes manifest

Deploy the latest Kubernetes manifest to your cluster using `kubectl apply`.

```bash
kubectl apply -f https://raw.githubusercontent.com/DB-Vincent/k8s-scaling-demo/refs/heads/main/deployment.yaml
```

### Using Kustomize

Deploy the latest Kubernetes manifest to your cluster using `kubectl apply`.

```bash
git clone https://github.com/DB-Vincent/k8s-scaling-demo.git
cd k8s-scaling-demo
kubectl apply -k kustomize/base
```

## Development

Feel free to adjust the application to your own liking.

### Frontend

The frontend has been built using Angular and sits in the `frontend/` directory.

1. Install the necessary packages: `npm install`
2. Run the development server: `npm run build`

### Backend

The backend application uses the client-go package to communicate with the Kubernetes cluster and is located in the `backend/` directory. Since we're not running the application in a cluster while developping, we just need to specify two environment variables:

- `POD_NAME`: set this to a pod that is running in your current Kubernetes context
- `POD_NAMESPACE`: set this to the namespace of the pod you selected for the previous variable

Once that is done;

1. Install the dependencies: `go mod tidy`
2. Start the application: `go run cmd/server/main.go`
3. Visit the API at http://localhost:8080/api or the frontend at http://localhost:8080

## Releases

The project uses semantic versioning for releases. Docker images are automatically built and published to GitHub Container Registry (ghcr.io).

### Creating a new release

1. Tag your commit with a version number:
```bash
git tag v1.0.0  # Replace with your version
git push origin v1.0.0
```

2. This will trigger the CI pipeline to:
  - Build the Docker image
  - Tag it with both the version number (e.g., `1.0.0`) and `latest`
  - Push it to `ghcr.io/DB-Vincent/k8s-scaling-demo`

You can then use the versioned image in your deployments:
```yaml
image: ghcr.io/db-vincent/k8s-scaling-demo:1.0.0
```

Or use the `latest` tag to always get the most recent version:
```yaml
image: ghcr.io/db-vincent/k8s-scaling-demo:latest
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.
