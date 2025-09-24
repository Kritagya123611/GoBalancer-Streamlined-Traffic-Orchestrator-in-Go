##  GoBalancer: Streamlined Traffic Orchestrator

GoBalancer is a lightweight, scalable HTTP load balancer built in Go, designed to intelligently route traffic to multiple backend servers while ensuring reliability through health checks and observability via metrics.

It leverages Go’s concurrency, Redis for state management, and Docker for cloud-native deployment — making it an excellent project for demonstrating modern backend engineering skills.

## ✨ Features

Round-Robin Load Balancing – Distributes incoming HTTP requests across multiple backends.

Health Checks – Pings /health endpoints to route traffic only to healthy servers.

Redis Integration – Persists backend health status & request metrics.

Observability – Built-in /metrics endpoint for monitoring & Prometheus integration (future).

Cloud-Native Deployment – Fully Dockerized with docker-compose.

Concurrency – Non-blocking health checks & proxying using Go’s goroutines.

Extensible Design – Easy to add rate limiting, advanced routing, or service discovery.

## 💡 Why GoBalancer?

Unlike typical CRUD APIs, GoBalancer is a systems-level project that showcases:

Systems Programming – Reverse proxying and HTTP routing with net/http & httputil.

Scalability – Distributed state management using Redis.

Modern Cloud Practices – Dockerized, observable, and ready for Prometheus/Grafana.

Concurrency Mastery – Parallel health checks and request handling using goroutines.

A great portfolio piece that shows recruiters you can build real-world infrastructure like what powers AWS, Netflix, and Google Cloud.

## 🛠️ Tech Stack

Go – Core language (net/http, httputil).

Redis – Health & metrics persistence.

Docker & Docker Compose – Containerization & local orchestration.

(Future) Prometheus + Grafana – Advanced monitoring.

(Future) Gorilla WebSocket – Real-time health updates.

- Health Checker runs in background goroutines
- Status & metrics stored in Redis

## ⚙️ Installation & Setup
1️⃣ Prerequisites

Go
 ≥ 1.21

Docker
 & Docker Compose

Redis
 (local or containerized)

Git

2️⃣ Clone the Repo
git clone https://github.com/yourusername/gobalancer.git
cd gobalancer

3️⃣ Install Dependencies
go mod tidy

4️⃣ Run with Docker Compose (Recommended)
docker-compose up --build


This will spin up:

Go load balancer at localhost:8080

Backends at :8081, :8082

Redis for health & metrics

5️⃣ Run Locally (Without Docker)

Start Redis (e.g., redis-server or hosted):

go run backend/backend.go --port 8081
go run backend/backend.go --port 8082
go run cmd/gobalancer/main.go

6️⃣ Test the Load Balancer
curl http://localhost:8080
 → "Hello from Backend X"

curl http://localhost:8080/metrics
 → {"total_requests": 5}

## 🚀 Future Enhancements

Rate Limiting – Token bucket or leaky bucket.

Prometheus Metrics – Deeper observability.

Dynamic Configuration – Add/remove backends via HTTP API.

WebSocket Dashboard – Real-time health view.

Kubernetes Integration – Service discovery & scaling.

## 🧪 Challenges Overcome

Efficient goroutine-based concurrency for health checks.

Redis state management for distributed scaling.

Reverse proxying with httputil.ReverseProxy.

Seamless containerization for cloud-native environments.

## 🏅 Why This Project Stands Out

✅ Systems-level (not just CRUD)
✅ Modern cloud-native practices
✅ Scalable design with Redis
✅ Resume-worthy & recruiter-friendly

⚡ Perfect for anyone aiming for backend/SRE/devops/system design roles.

🤝 Contributing

Fork the repo

Create a branch: git checkout -b feature/awesome-feature

Commit: git commit -m "Add awesome feature"

Push: git push origin feature/awesome-feature

Open a Pull Request 🚀

📜 License

This project is licensed under the MIT License.

📬 Contact

Built with ❤️ by Kritagya
📧 Email: jhakritagya45@gmail.com
🐙 GitHub Issues for questions or feedback
