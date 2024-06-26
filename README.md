<div align='center' style="text-align: center;">

<h1 style="border:0;margin:1rem">Coffeeshop Backend</h1>

Backend for CoffeeShop

<hr>
<br>

</div>

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Postman Collection](#postman-collection)
- [Resources](#resources)
- [Contributors](#contributors)
- [Related Project](#related-projects)
- [License](#license)
- [Suggestion](#suggestion)

## Overview

CoffeeShop.

## Features

- CRUD

## Technologies Used

- [Golang](https://nodejs.org/en/docs)
- [Go Gin](https://expressjs.com/)
- etc.

## Getting Started

### Installation

1. Clone this repo

   ```bash
   git clone https://github.com/ninja1cak/coffeshop-be
   ```

2. Enter the directory

   ```bash
   cd coffeshop-be
   ```

3. Install all dependencies

   ```bash
   go get .
   ```

4. Create .env file

   ```env
    APP_ENV=dev
   PORT=8081
   ```

   DB_HOST=""
   DB_NAME=""
   DB_USER=""
   DB_PASS=""
   DB_PORT=5432

   <!-- secret key JWT -->

   JWT_KEYS="randomkey"

   <!-- cloudinary -->

   CD_NAME="yourclooudinaryname"
   CD_KEY="yourcdkey"
   CD_SECRET="yourcdsecret"

````

5. Start the local server

```bash
go run cmd/main.go
````

## Postman Collection

You can download in <a href='#'> Here </a>

## Contributors

Currently, there are no contributors to this project. If you would like to contribute, you can submit a pull request.

## Suggestion

If you find bugs / find better ways / suggestions you can pull request.

### build program docker

1. build file go
   go build -o "./build/coffeeshopbe.exe" ./cmd/main.go

2. build docker image

<!-- docker build -t nama_image:tag lokasi_dockerfile -->
<!-- docker build -t nama_user_dockerhub/coffeeshopbe:tag .(titik = folder saat ini, ketika di terminal) -->

docker build -t zikrigusli/coffeeshopbe:1 .

## run and create container from image

docker compose up -d

## delete container

docker compose down
