# 🚦 Popular Rate-Limiting Patterns

The goal is implement several rate limiter patterns and test it just for educational propourses.

## 🛠️ Tech Stack

- 🔷 Go. The patterns are implemented in Go.
- 📊 K6. I used K6 to generate the required load.
- 🪣 Prometheus. The metrics are collected by Prometheus.
  - `localhost:9090` is the address of the Prometheus server.
- 📊 Grafana. The metrics are visualized by Grafana.
  - `localhost:3000` is the address of the Grafana server.
- 🗄️ Redis. The leaky bucket is implemented using Redis.
  - `localhost:6379` is the address of the Redis server.

## 🚀 Run

Docker compose will start the rate limiter, the requester and the Prometheus and Grafana servers.

```bash
docker compose up
```

## 🔄 Manual Requester

The requester is a simple HTTP server that will send requests to the rate limiter.

```bash
# example
make requester n=1000 h=localhost:3010
```

Parameters:

- `-n`: number of requests to send
- `-h`: host to send the requests to

## 💧 Leaky Bucket

![](https://cdn-images-1.medium.com/max/2160/1*UioRG8-qID51i0rEOPVh-w.gif)

**How It Works**: Imagine a bucket with a small hole at the bottom. Requests (water) are added to the bucket and processed at a steady "drip" rate, preventing sudden floods.

**Use Cases:** Ideal for smoothing traffic flow, such as in streaming services or payment processing, where a predictable output is critical.

**Example:** A video streaming platform regulates API calls to its content delivery network, ensuring consistent playback quality.

**Drawback:** Not suitable for handling sudden bursts, like flash sales or promotional campaigns.

## 📚 Resources

- https://foojay.io/today/rate-limiting-with-redis-an-essential-guide/
- https://medium.com/@mrandiiw/how-to-set-up-prometheus-grafana-for-golang-app-monitoring-2bd5e5c3d23e
