# Results

## Test environment

NGINX Plus: false

GKE Cluster:

- Node count: 12
- k8s version: v1.29.5-gke.1091002
- vCPUs per node: 16
- RAM per node: 65855012Ki
- Max pods per node: 110
- Zone: us-west1-b
- Instance Type: n2d-standard-16

## Test: Send http /coffee traffic

```text
Requests      [total, rate, throughput]         6000, 100.02, 100.01
Duration      [total, attack, wait]             59.991s, 59.99s, 986.387µs
Latencies     [min, mean, 50, 90, 95, 99, max]  723.73µs, 1.102ms, 1.037ms, 1.193ms, 1.256ms, 1.839ms, 13.076ms
Bytes In      [total, mean]                     960000, 160.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:6000  
Error Set:
```

![http.png](http.png)

## Test: Send https /tea traffic

```text
Requests      [total, rate, throughput]         6000, 100.02, 100.01
Duration      [total, attack, wait]             59.995s, 59.99s, 4.904ms
Latencies     [min, mean, 50, 90, 95, 99, max]  606.976µs, 1.042ms, 1.016ms, 1.151ms, 1.203ms, 1.458ms, 13.052ms
Bytes In      [total, mean]                     922049, 153.67
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:6000  
Error Set:
```

![https.png](https.png)
