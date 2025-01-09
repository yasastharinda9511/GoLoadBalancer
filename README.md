# Go Load Balancer

This repository contains a Go-based Load Balancer that distributes incoming requests among multiple backend servers. The Load Balancer supports features such as round-robin and least-connections algorithms, and is designed for high performance and scalability.

## Features

- **Load Balancing Algorithms**: Supports Round-Robin, Least Connections, and more.
- **Health Checks**: Periodic health checks to ensure backend server availability.
- **Concurrency**: Handles multiple simultaneous requests efficiently.
- **Logging**: Detailed logs for debugging and monitoring.

---

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- Git

---

## Clone the Repository

```bash
$ git clone https://github.com/yasastharinda9511/GoLoadBalancer.git
$ cd GoLoadBalancer
```

---

## Build the Load Balancer

### Steps to Build:
1. Navigate to the project directory:

   ```bash
   $ cd GoLoadBalancer
   ```

2. Build the project:

   ```bash
   $ go build -o bin/loadbalancer ./cmd
   ```

   This will generate an executable file named `loadbalancer` inside the `bin` directory.

---

## Run the Load Balancer

### Using Command Line:

1. Start the Load Balancer with a configuration file:

   ```bash
   $ ./bin/loadbalancer
   ```

2. Sample configuration file:

   ```yaml
   server:
     base_port: 3333
     server_count: 2

   rules:
     - id: "rule1"
       path_rule:
         path: "/api"
         type: "PREFIX"
       header_rules:
         - key: "Yasas"
           value: "tharinda"
       pool:
         load_balancer: "WEIGHTEDLOADBALANCER"
         backends:
           - url: "http://localhost:3000"
             weight: 1
           - url: "http://localhost:3001"
             weight: 1
           - url: "http://localhost:3002"
             weight: 1
     - id: "rule2"
       path_rule:
         path: "/pet"
         type: "PREFIX"
       pool:
         load_balancer: "WEIGHTEDLOADBALANCER"
         backends:
           - url: "http://localhost:3005"
             weight: 1
           - url: "http://localhost:3006"
             weight: 1
           - url: "http://localhost:3007"
             weight: 1
   ```

## Testing

You can test the Load Balancer by sending HTTP requests using tools like `curl`:

```bash
$ curl -H "Yasas:tharinda" http://127.0.0.1:3333
$ curl -H "Yasas:tharinda" http://127.0.0.1:3334
```

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes.
4. Open a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contact

For questions or support, feel free to open an issue or contact the repository owner.

