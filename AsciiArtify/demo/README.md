# Concept: Choosing a Local Kubernetes Tool for the AsciiArtify PoC Startup
## Introduction
AsciiArtify is a startup developing an ML product that converts images into ASCII art.
For local Kubernetes development and testing, the team is evaluating three tools:

- ***minikube** — a full-featured single-node Kubernetes cluster.
- ***kind (Kubernetes IN Docker)*** — clusters running inside Docker containers.
- ***k3d** — a lightweight wrapper around k3s (a minimal Kubernetes distribution) running in Docker.

## Characteristics

| Feature          | minikube                        | kind                       | k3d                                |
| ---------------- | ------------------------------- | -------------------------- | ---------------------------------- |
| OS Support       | Windows, macOS, Linux           | Windows, macOS, Linux      | Windows, macOS, Linux              |
| Architecture     | VM/Container                    | Docker                     | Docker (based on k3s)              |
| Automation       | Yes (addons, profiles)          | Yes (YAML cluster configs) | Yes (CLI, YAML, multi-node)        |
| Monitoring / GUI | Yes (dashboard, metrics-server) | No (manual setup)          | Partial (via k3s + external tools) |
| Podman Support   | Yes (experimental)              | Partial (via nerdctl)      | No (Docker only)                   |

## Pros and Cons
### 🟢 minikube

**Pros:**
- Easy to install
- Built-in GUI and dashboard support
- Podman support

**Cons:**
- Slower startup
- Higher resource usage

### 🟡 kind
**Pros:**
- Lightweight and fast
- Great CI/CD integration
- Simple cluster configuration

**Cons:**
- No built-in monitoring
- Limited Podman support

### 🔵 k3d
**Pros:**
- Extremely fast startup
- Multi-node support
- Based on k3s (lightweight Kubernetes)

**Cons:**
- Requires Docker
- No GUI by default

## Demo: Deploying a Hello World App on k3d
```bash
# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# Create a cluster
k3d cluster create asciiartify

# Apply the deployment manifest
kubectl apply -f demo/hello-world-deployment.yaml

# Verify
kubectl get pods
kubectl port-forward svc/hello-world 8080:80
```