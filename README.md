# Popular Rate-Limiting Patterns

The goal is implement several rate limiter patterns and test it just for educational propourses.

## Requester

The requester is a simple HTTP server that will send requests to the rate limiter.

```bash
go run requester/cmd/main.go
```

Parameters:

- `-n`: number of requests to send
- `-r`: rate of requests to send
- `-h`: host to send the requests to

## Leaky Bucket

!https://cdn-images-1.medium.com/max/2160/1*UioRG8-qID51i0rEOPVh-w.gif

**How It Works**: Imagine a bucket with a small hole at the bottom. Requests (water) are added to the bucket and processed at a steady “drip” rate, preventing sudden floods.

**Use Cases:** Ideal for smoothing traffic flow, such as in streaming services or payment processing, where a predictable output is critical.

**Example:** A video streaming platform regulates API calls to its content delivery network, ensuring consistent playback quality.

**Drawback:** Not suitable for handling sudden bursts, like flash sales or promotional campaigns.

## Resources

- https://foojay.io/today/rate-limiting-with-redis-an-essential-guide/
