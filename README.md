# RSA Key Rotator CLI 🔄

```bash              
┏┓┏┓  ┏┓┏┓╋┏┓╋┏┓
┗┫┗┛  ┛ ┗┛┗┗┻┗┗ 
 ┛     
 
 $ go-rotate --path /my/parameter/store/path --name my_key_name
 
 🔐 Generated and stored keys: 
 
    /my/parameter/store/path/my_key_name_private.pem
    /my/parameter/store/path/my_key_name_public.pem
```

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)
[![License](https://img.shields.io/github/license/GitGuardian/ggshield?color=%231B2D55&style=for-the-badge)](LICENSE)

![Build](https://github.com/kmesiab/go-key-rotator-cli/actions/workflows/go-build.yml/badge.svg)
![Lint](https://github.com/kmesiab/go-key-rotator-cli/actions/workflows/go-lint.yml/badge.svg)
![Test](https://github.com/kmesiab/go-key-rotator-cli/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/kmesiab/go-key-rotator-cli)


## Overview

RSA Key Rotator CLI is a command-line tool designed to facilitate the
secure rotation of RSA keys backed by AWS Parameter store 🔑.

Utilizing the powerful
[`go-key-rotator`](https://github.com/kmesiab/go-key-rotator)
library, this tool offers a user-friendly approach to manage RSA keys.
Ideal as a standalone tool or as a companion to the `go-key-rotator`
library, `go-rotate` streamlines the process of rotating RSA keys and
integrates seamlessly with AWS Parameter Store for safe storage and
retrieval 🛡️.

## Features

- Generate and rotate RSA keys with customizable options 🔧.
- Integration with AWS Parameter Store for secure key management 🔐.
- User-friendly command-line interface 💻.
- Suitable for standalone use or in conjunction with the `go-key-rotator`
- library 🤝.

## Installation

Install RSA Key Rotator CLI using the following Go command:

```bash
go get github.com/kmesiab/go-key-rotator-cli
```

## Usage

Once installed, you can use the `go-rotate` command to manage your
RSA keys. Common commands include:

- **Generate a new RSA key**:

```bash
go-rotate generate --size 2048
```

- **Rotate an existing RSA key**:

```bash
go-rotate rotate --key /path/to/your/key.pem
```

- **Store a key in AWS Parameter Store**:

```bash
go-rotate store --key /path/to/your/key.pem --name mykey
```

## Contributing

Contributions to RSA Key Rotator CLI are welcome! Please read our
contributing guidelines to get started 🤗.

## License

RSA Key Rotator CLI is open-source software licensed under the MIT
license 📜.
