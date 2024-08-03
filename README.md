# fashionmart

This repository contains the backend API for a single-brand e-commerce store, rewritten from NestJS to Go. The API offers comprehensive functionality to manage users, products, and orders.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Zeo-dev3/
   ```

2. Run docker-compose:
   ```bash
   docker-compose up -d
   ```

## Usage

### Running the API

After following the setup instructions, the API will be available at `http://localhost:3000`.

## Dependencies

- [Go Fiber](https://gofiber.io) - Web framework inspired by Express.js
- [GORM](https://gorm.io) - ORM library for Golang
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Package for password hashing

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any features, bug fixes, or enhancements.

1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`.
3. Make your changes and commit them: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-branch`.
5. Submit a pull request.

## License

This project is licensed under the MIT License.