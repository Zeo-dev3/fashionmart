## FashionMart: Go Backend API for E-commerce

This repository contains the backend API for a single-brand e-commerce store, rewritten from NestJS to Go. The API offers comprehensive functionality to manage users, products, and orders.

## Table of Contents\*\*

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Tech Stack](#tech-stack)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before diving in, ensure you have the following tools installed on your system:

- **Go** (version 1.16 or later): [https://go.dev/doc/install](https://go.dev/doc/install)
- **Docker**: [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)
- **Docker Compose**: [https://docs.docker.com/compose/](https://docs.docker.com/compose/)

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Zeo-dev3/fashionmart.git
   ```

2. **Fire Up Docker Compose:**

   ```bash
   docker-compose up -d
   ```

This will start all the necessary services in the background, and your API will be up and running shortly.

## Usage

The magic happens at `http://localhost:3000`! Once the API is running, you can access it using this URL to interact with your e-commerce store's functionalities.

## Tech Stack

This project leverages a powerful combination of technologies to deliver a robust e-commerce backend:

- Go Fiber: A high-performance web framework inspired by Express.js for a smooth development experience. [https://github.com/gofiber/fiber](https://github.com/gofiber/fiber)
- GORM: An Object-Relational Mapper (ORM) that simplifies database interactions in Go. [https://gorm.io/](https://gorm.io/)
- bcrypt: A built-in Go package for securely hashing passwords, ensuring user data protection. [https://github.com/golang/crypto/blob/master/bcrypt/bcrypt.go](https://github.com/golang/crypto/blob/master/bcrypt/bcrypt.go)

## Swagger Documentation

For detailed API documentation, Swagger UI is available at the `/swagger` endpoint once the server is running. You can explore all available endpoints, request parameters, and response schemas interactively.

To access Swagger UI:

1. Start the server using Docker Compose as described in the [Installation](#installation) section.
2. Open your browser and navigate to `http://localhost:3000/swagger`.

Here, you can view and test the API endpoints and their functionalities directly from the browser.

## Contributing

We welcome your contributions to make FashionMart even better! Feel free to fork the repository, make your changes, and submit a pull request. Here's a quick guide:

1. Fork the repository.
2. Create a new branch for your changes (e.g., `git checkout -b improve-readme`).
3. Implement your modifications and commit them with a descriptive message (e.g., `git commit -m "Enhanced readme for clarity"`).
4. Push your branch to your fork (e.g., `git push origin improve-readme`).
5. Submit a pull request to the main repository.

## License

This project is distributed under the permissive MIT License, providing you with the freedom to use, modify, and distribute it as per your needs.
