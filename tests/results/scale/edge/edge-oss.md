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

## Test TestScale_Listeners

### Reloads

- Total: 127
- Total Errors: 0
- Average Time: 147ms
- Reload distribution:
	- 500ms: 127
	- 1000ms: 127
	- 5000ms: 127
	- 10000ms: 127
	- 30000ms: 127
	- +Infms: 127

### Event Batch Processing

- Total: 386
- Average Time: 126ms
- Event Batch Processing distribution:
	- 500ms: 349
	- 1000ms: 386
	- 5000ms: 386
	- 10000ms: 386
	- 30000ms: 386
	- +Infms: 386

### Errors

- NGF errors: 0
- NGF container restarts: 0
- NGINX errors: 0
- NGINX container restarts: 0

### Graphs and Logs

See [output directory](./TestScale_Listeners) for more details.
The logs are attached only if there are errors.

## Test TestScale_HTTPSListeners

### Reloads

- Total: 128
- Total Errors: 0
- Average Time: 174ms
- Reload distribution:
	- 500ms: 128
	- 1000ms: 128
	- 5000ms: 128
	- 10000ms: 128
	- 30000ms: 128
	- +Infms: 128

### Event Batch Processing

- Total: 451
- Average Time: 122ms
- Event Batch Processing distribution:
	- 500ms: 407
	- 1000ms: 449
	- 5000ms: 451
	- 10000ms: 451
	- 30000ms: 451
	- +Infms: 451

### Errors

- NGF errors: 0
- NGF container restarts: 0
- NGINX errors: 0
- NGINX container restarts: 0

### Graphs and Logs

See [output directory](./TestScale_HTTPSListeners) for more details.
The logs are attached only if there are errors.

## Test TestScale_HTTPRoutes

### Reloads

- Total: 1001
- Total Errors: 0
- Average Time: 383ms
- Reload distribution:
	- 500ms: 721
	- 1000ms: 1001
	- 5000ms: 1001
	- 10000ms: 1001
	- 30000ms: 1001
	- +Infms: 1001

### Event Batch Processing

- Total: 1008
- Average Time: 432ms
- Event Batch Processing distribution:
	- 500ms: 628
	- 1000ms: 1008
	- 5000ms: 1008
	- 10000ms: 1008
	- 30000ms: 1008
	- +Infms: 1008

### Errors

- NGF errors: 0
- NGF container restarts: 0
- NGINX errors: 0
- NGINX container restarts: 0

### Graphs and Logs

See [output directory](./TestScale_HTTPRoutes) for more details.
The logs are attached only if there are errors.

## Test TestScale_UpstreamServers

### Reloads

- Total: 60
- Total Errors: 0
- Average Time: 126ms
- Reload distribution:
	- 500ms: 60
	- 1000ms: 60
	- 5000ms: 60
	- 10000ms: 60
	- 30000ms: 60
	- +Infms: 60

### Event Batch Processing

- Total: 62
- Average Time: 237ms
- Event Batch Processing distribution:
	- 500ms: 61
	- 1000ms: 62
	- 5000ms: 62
	- 10000ms: 62
	- 30000ms: 62
	- +Infms: 62

### Errors

- NGF errors: 0
- NGF container restarts: 0
- NGINX errors: 0
- NGINX container restarts: 0

### Graphs and Logs

See [output directory](./TestScale_UpstreamServers) for more details.
The logs are attached only if there are errors.

## Test TestScale_HTTPMatches

```text
Requests      [total, rate, throughput]         30000, 1000.04, 1000.01
Duration      [total, attack, wait]             30s, 29.999s, 886.653µs
Latencies     [min, mean, 50, 90, 95, 99, max]  596.482µs, 893.186µs, 861.306µs, 990.751µs, 1.041ms, 1.178ms, 22.076ms
Bytes In      [total, mean]                     4860000, 162.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:30000  
Error Set:
```
```text
Requests      [total, rate, throughput]         30000, 1000.04, 1000.01
Duration      [total, attack, wait]             30s, 29.999s, 887.065µs
Latencies     [min, mean, 50, 90, 95, 99, max]  692.077µs, 930.425µs, 915.478µs, 1.048ms, 1.098ms, 1.208ms, 11.806ms
Bytes In      [total, mean]                     4860000, 162.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:30000  
Error Set:
```
