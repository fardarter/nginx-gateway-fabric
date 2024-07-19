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

## Test1: Running latte path based routing

```text
Requests      [total, rate, throughput]         29999, 1000.00, 999.97
Duration      [total, attack, wait]             30s, 29.999s, 919.496µs
Latencies     [min, mean, 50, 90, 95, 99, max]  581.765µs, 846.843µs, 813.639µs, 928.064µs, 972.795µs, 1.152ms, 29.428ms
Bytes In      [total, mean]                     4769841, 159.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:29999  
Error Set:
```

## Test2: Running coffee header based routing

```text
Requests      [total, rate, throughput]         30000, 1000.04, 1000.01
Duration      [total, attack, wait]             30s, 29.999s, 839.899µs
Latencies     [min, mean, 50, 90, 95, 99, max]  597.349µs, 874.558µs, 849.664µs, 1.002ms, 1.082ms, 1.331ms, 7.203ms
Bytes In      [total, mean]                     4800000, 160.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:30000  
Error Set:
```

## Test3: Running coffee query based routing

```text
Requests      [total, rate, throughput]         29999, 1000.00, 999.97
Duration      [total, attack, wait]             30s, 29.999s, 926.302µs
Latencies     [min, mean, 50, 90, 95, 99, max]  622.494µs, 886.352µs, 869.509µs, 975.016µs, 1.021ms, 1.147ms, 22.066ms
Bytes In      [total, mean]                     5039832, 168.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:29999  
Error Set:
```

## Test4: Running tea GET method based routing

```text
Requests      [total, rate, throughput]         30000, 1000.03, 1000.01
Duration      [total, attack, wait]             30s, 29.999s, 754.279µs
Latencies     [min, mean, 50, 90, 95, 99, max]  614.59µs, 854.062µs, 838.681µs, 971.694µs, 1.029ms, 1.147ms, 7.002ms
Bytes In      [total, mean]                     4710000, 157.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:30000  
Error Set:
```

## Test5: Running tea POST method based routing

```text
Requests      [total, rate, throughput]         30000, 1000.01, 999.98
Duration      [total, attack, wait]             30.001s, 30s, 773.301µs
Latencies     [min, mean, 50, 90, 95, 99, max]  624.169µs, 860µs, 847.581µs, 962.369µs, 1.004ms, 1.131ms, 11.857ms
Bytes In      [total, mean]                     4710000, 157.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:30000  
Error Set:
```
