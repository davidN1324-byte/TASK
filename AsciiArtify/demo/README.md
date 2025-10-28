# 🧠 AsciiArtify — Proof of Concept (PoC)

**AsciiArtify** is a startup building an ML service that converts images into ASCII art.

This repository contains the **PoC** infrastructure and a demo of deploying a simple application on local Kubernetes.

---

## 🚀 Project Goals

- Master Kubernetes tools for local development.

- Prepare a Proof of Concept.

- Test the ML project's functionality in containers.

---

## 🧩 Technologies Used

| Component | Purpose |
|-------------|------------|
| **k3d**     | Local Kubernetes cluster (based on k3s) |
| **kubectl** | Cluster Management |
| **Docker**  | Containerization and Environments |
| **GitHub**  | Version Control and Documentation System |

---

## ⚙️ Quick Start (Hello World)

> Requires **Docker**, **kubectl**, and **k3d** installed.

```bash
# 1️⃣ Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# 2️⃣ Create a local cluster
k3d cluster create asciiartify

# 3️⃣ Apply the demo manifest
kubectl apply -f demo/hello-world-deployment.yaml

# 4️⃣ Check pods
kubectl get pods

# 5️⃣ Forward the port and open the application
kubectl port-forward svc/hello-world 8080:80