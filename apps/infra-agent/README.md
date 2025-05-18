## What is agent in InfraSim?
#### The agent is a standalone program that runs in EC2 instances to simulate real users
- Receives instructions (target URL, Type of request, payload, number of users)
- generate traffic to simulate user (GET/POST) etc
- Monitor results (success/ failures, latecy, errors)
- Sends data back to control plane

## Target
| Feature                        | Description                                                              |
| ------------------------------ | ------------------------------------------------------------------------ |
| 🟢 **Input Parameters**        | What to test (URL), how (GET/POST), how many (concurrency level)         |
| 🚀 **Concurrency**             | Simulate multiple users at the same time (via goroutines)                |
| 📊 **Result Tracking**         | Count of success, failure, maybe latency                                 |
| ⏱️ **Timeout Handling**        | Prevent hanging requests (via timeout)                                   |
| 🔒 **Safe Parallel Execution** | Use mutexes or atomic counters to track results across goroutines safely |
| 📤 **(Optional) Report Back**  | Send aggregated results to a control server (control plane)              |
| 📦 **Self-contained Binary**   | Can run standalone on EC2 or any machine                                 |



## Logs
### Current MVP Objective 18 May 2025
1. Accepts a URL and number of users
2. Spawns that many goroutines to make GET requests
3. Track how many succeded and how many failed
4. Prints the result

