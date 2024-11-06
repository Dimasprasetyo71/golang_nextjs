# Gin Web Framework

# Project Folder Structure

This document provides an overview of the project folder structure golang .

```sh 
📦config
 ┣ 📂fakers
 ┃ ┣ 📜addres.go
 ┃ ┣ 📜product.fakers.go
 ┃ ┗ 📜user.fakers.go
 ┣ 📂seeders
 ┃ ┗ 📜seders.go
 ┗ 📜database.go
 📦deployment
 ┗ 📜docker-compose.yml
 📦controllers
 ┣ 📜addresController.go
 ┣ 📜CategoryController.go
 ┣ 📜ProductControllers.go
 ┣ 📜ProductImageControllers.go
 ┗ 📜UserControllers.go
 📦repositories
 ┣ 📜addres_repository.go
 ┣ 📜categoy.repository.go
 ┣ 📜productimage.go
 ┣ 📜product_repository.go
 ┗ 📜user_repository.go
 📦models
 ┣ 📜addres.go
 ┣ 📜category.go
 ┣ 📜payment.go
 ┣ 📜product.go
 ┣ 📜productimages.go
 ┗ 📜user.go
 📦docs
 ┣ 📜product.yaml
 ┗ 📜user.yaml
 📦dto
 ┣ 📜addres_dto.go
 ┣ 📜category_dto.go
 ┣ 📜midtrans.dto.go
 ┣ 📜product_dto.go
 ┣ 📜product_image_dto.go
 ┗ 📜user_dto.go
 📦services
 ┣ 📜category_service.go
 ┣ 📜product_service.go
 ┗ 📜user_service.go
 📦middleware
 ┣ 📜auth.go
 ┣ 📜error.handling.go
 ┗ 📜rate.limiting.go
 📦validations
 ┣ 📜addres_validation.go
 ┣ 📜category_validation.go
 ┣ 📜productimage_validation.go
 ┣ 📜product_validation.go
 ┗ 📜user_validation.go
 📦utils
 ┣ 📜generate.slug.go
 ┣ 📜jwt.token.go
 ┣ 📜logger.go
 ┗ 📜random.go
 📦helpers
 ┣ 📜formatter.go
 ┣ 📜json.go
 ┣ 📜parser.go
 ┗ 📜string.go
 📦testing
 ┣ 📜address_testing.go
 ┣ 📜category_testing.go
 ┣ 📜test.go
 ┗ 📜user_testing.go
 📦routes
 ┗ 📜route.go
```

## Root Folders and Files

- `.github/`: Contains GitHub configuration files such as workflows for CI/CD pipelines.
- `.vercel/`: Configuration files for Vercel deployment.
- `.vscode/`: Settings and configurations for VS Code workspace.
- `cmd/`: Contains the entry point for the application, typically a main file for launching services.
- `config/`: Configuration files and settings (e.g., database configuration, application settings).
- `controllers/`: Controller files that handle requests and responses.
- `deployment/`: Docker, Kubernetes, and other deployment-related files.
- `docs/`: Documentation files, possibly including API documentation.
- `dto/`: Data Transfer Objects, used for structuring data that’s sent or received.
- `go/`: A general folder for Go-related dependencies or tools.
- `helpers/`: Helper functions or utilities used across the application.
- `middleware/`: Middleware functions for request processing (e.g., authentication, logging).
- `models/`: Models representing the structure of database tables or domain objects.
- `repositories/`: Repository layer for data access logic.
- `routes/`: Routing files that define the API endpoints and link them to controllers.
- `services/`: Service layer containing business logic.
- `testing/`: Contains testing files and related configurations.
- `utils/`: Utility functions or classes used throughout the project.
- `validations/`: Validation rules and schemas for input data validation.

## Root Files

- `.env`: Environment variables for sensitive configurations.
- `.gitignore`: Specifies files and folders to be ignored by Git.
- `code_of_conduct.md`: Project code of conduct.
- `CONTRIBUTING.md`: Guidelines for contributing to the project.
- `Dockerfile`: Docker configuration for containerizing the application.
- `go.mod`: Module file listing dependencies.
- `go.sum`: Sum file for dependency verification.
- `LICENSE`: Project license file.
- `LICENSE.txt`: Text-based license file.
- `main.go`: Main entry point for the Go application.
- `makefile`: A makefile for build automation.
- `README.md`: Project overview, setup instructions, and documentation.
- `SECURITY.md`: Security policy or guidelines for reporting vulnerabilities.
- `swagger.yaml`: Swagger configuration file for API documentation.

## Additional Files

<!-- BY DIMAS PRASETYO -->

<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)](https://github.com/gin-gonic/gin/actions?query=branch%3Amaster)
[![codecov](https://codecov.io/gh/gin-gonic/gin/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-gonic/gin)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)
[![Go Reference](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)
[![Sourcegraph](https://sourcegraph.com/github.com/gin-gonic/gin/-/badge.svg)](https://sourcegraph.com/github.com/gin-gonic/gin?badge)
[![Open Source Helpers](https://www.codetriage.com/gin-gonic/gin/badges/users.svg)](https://www.codetriage.com/gin-gonic/gin)
[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/gin-gonic/gin)](https://www.tickgit.com/browse?repo=github.com/gin-gonic/gin)

Gin is a web framework written in [Go](https://go.dev/). It features a martini-like API with performance that is up to 40 times faster thanks to [httprouter](https://github.com/julienschmidt/httprouter).
If you need performance and good productivity, you will love Gin.

**Gin's key features are:**

- authentication
- JWT TOKEN
- Login with goggle
- login with github
- docker
- Route grouping
- Error management
- Built-in rendering
- ecomerce 


## Getting started

### Prerequisites

Gin requires [Go](https://go.dev/) version [1.21](https://go.dev/doc/devel/release#go1.21.0) or above.

### Getting Gin

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/gin-gonic/gin"
```

Alternatively, use `go get`:

```sh
go get -u github.com/gin-gonic/gin
```

### how to clone you project 

```sh
     git clone https://github.com/Dimasprasetyo71/golang_nextjs?tab=security-ov-file
    #  install all 
     go get 
    # running port you golang 
    go run main.go
```

### Running Gin

A basic example:

```go
package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GOLANG_PORT")
	if port == "" {
		log.Fatal("$GOLANG_PORT must be set")
	}

	router := gin.New()
	router.Use(cors.Default())
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.RateLimitMiddleware())

    db := config.ConnectDB()
	gin.SetMode(gin.ReleaseMode)


	// Set up routes
	routes.UserRoutes(router, db)
	routes.AddressRoutes(router, db)
	routes.CategoryRoutes(router, db)
	routes.ProductImageRoutes(router, db)
	routes.ProductRoutes(router, db)

	// Set up seeders
	// seeders.DBSeed(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	router.Static("/swagger-ui", "./static/swagger-ui")
	router.GET("/swagger.yaml", func(c *gin.Context) {
    c.File("./swagger.yaml")
	})

	log.Println("Server running on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

```

To run the code, use the `go run` command, like:

```sh
$ go run main.go
```

2024/11/06 23:29:34 Server running on port localhost:3000

### See more examples

#### Quick Start

Learn and practice with the [Gin Quick Start](docs/doc.md), which includes API examples and builds tag.

#### Examples

A number of ready-to-run examples demonstrating various use cases of Gin are available in the [Gin examples](https://github.com/gin-gonic/examples) repository.

## Documentation

See the [API documentation on go.dev](https://pkg.go.dev/github.com/gin-gonic/gin).

The documentation is also available on [gin-gonic.com](https://gin-gonic.com) in several languages:

- [English](https://gin-gonic.com/docs/)
- [简体中文](https://gin-gonic.com/zh-cn/docs/)
- [繁體中文](https://gin-gonic.com/zh-tw/docs/)
- [日本語](https://gin-gonic.com/ja/docs/)
- [Español](https://gin-gonic.com/es/docs/)
- [한국어](https://gin-gonic.com/ko-kr/docs/)
- [Turkish](https://gin-gonic.com/tr/docs/)
- [Persian](https://gin-gonic.com/fa/docs/)
- [Português](https://gin-gonic.com/pt/docs/)

### Articles

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

## Benchmarks

Gin uses a custom version of [HttpRouter](https://github.com/julienschmidt/httprouter), [see all benchmarks](/BENCHMARKS.md).

| Benchmark name                 |       (1) |             (2) |          (3) |             (4) |
| ------------------------------ | ---------:| ---------------:| ------------:| ---------------:|
| BenchmarkGin_GithubAll         | **43550** | **27364 ns/op** |   **0 B/op** | **0 allocs/op** |
| BenchmarkAce_GithubAll         |     40543 |     29670 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkAero_GithubAll        |     57632 |     20648 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkBear_GithubAll        |      9234 |    216179 ns/op |   86448 B/op |   943 allocs/op |
| BenchmarkBeego_GithubAll       |      7407 |    243496 ns/op |   71456 B/op |   609 allocs/op |
| BenchmarkBone_GithubAll        |       420 |   2922835 ns/op |  720160 B/op |  8620 allocs/op |
| BenchmarkChi_GithubAll         |      7620 |    238331 ns/op |   87696 B/op |   609 allocs/op |
| BenchmarkDenco_GithubAll       |     18355 |     64494 ns/op |   20224 B/op |   167 allocs/op |
| BenchmarkEcho_GithubAll        |     31251 |     38479 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkGocraftWeb_GithubAll  |      4117 |    300062 ns/op |  131656 B/op |  1686 allocs/op |
| BenchmarkGoji_GithubAll        |      3274 |    416158 ns/op |   56112 B/op |   334 allocs/op |
| BenchmarkGojiv2_GithubAll      |      1402 |    870518 ns/op |  352720 B/op |  4321 allocs/op |
| BenchmarkGoJsonRest_GithubAll  |      2976 |    401507 ns/op |  134371 B/op |  2737 allocs/op |
| BenchmarkGoRestful_GithubAll   |       410 |   2913158 ns/op |  910144 B/op |  2938 allocs/op |
| BenchmarkGorillaMux_GithubAll  |       346 |   3384987 ns/op |  251650 B/op |  1994 allocs/op |
| BenchmarkGowwwRouter_GithubAll |     10000 |    143025 ns/op |   72144 B/op |   501 allocs/op |
| BenchmarkHttpRouter_GithubAll  |     55938 |     21360 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkHttpTreeMux_GithubAll |     10000 |    153944 ns/op |   65856 B/op |   671 allocs/op |
| BenchmarkKocha_GithubAll       |     10000 |    106315 ns/op |   23304 B/op |   843 allocs/op |
| BenchmarkLARS_GithubAll        |     47779 |     25084 ns/op |       0 B/op |     0 allocs/op |
| BenchmarkMacaron_GithubAll     |      3266 |    371907 ns/op |  149409 B/op |  1624 allocs/op |
| BenchmarkMartini_GithubAll     |       331 |   3444706 ns/op |  226551 B/op |  2325 allocs/op |
| BenchmarkPat_GithubAll         |       273 |   4381818 ns/op | 1483152 B/op | 26963 allocs/op |
| BenchmarkPossum_GithubAll      |     10000 |    164367 ns/op |   84448 B/op |   609 allocs/op |
| BenchmarkR2router_GithubAll    |     10000 |    160220 ns/op |   77328 B/op |   979 allocs/op |
| BenchmarkRivet_GithubAll       |     14625 |     82453 ns/op |   16272 B/op |   167 allocs/op |
| BenchmarkTango_GithubAll       |      6255 |    279611 ns/op |   63826 B/op |  1618 allocs/op |
| BenchmarkTigerTonic_GithubAll  |      2008 |    687874 ns/op |  193856 B/op |  4474 allocs/op |
| BenchmarkTraffic_GithubAll     |       355 |   3478508 ns/op |  820744 B/op | 14114 allocs/op |
| BenchmarkVulcan_GithubAll      |      6885 |    193333 ns/op |   19894 B/op |   609 allocs/op |

- (1): Total Repetitions achieved in constant time, higher means more confident result
- (2): Single Repetition Duration (ns/op), lower is better
- (3): Heap Memory (B/op), lower is better
- (4): Average Allocations per Repetition (allocs/op), lower is better

## Middleware

You can find many useful Gin middlewares at [gin-contrib](https://github.com/gin-contrib).

## Uses

Here are some awesome projects that are using the [Gin](https://github.com/gin-gonic/gin) web framework.

- [gorush](https://github.com/appleboy/gorush): A push notification server.
- [fnproject](https://github.com/fnproject/fn): A container native, cloud agnostic serverless platform.
- [photoprism](https://github.com/photoprism/photoprism): Personal photo management powered by Google TensorFlow.
- [lura](https://github.com/luraproject/lura): Ultra performant API Gateway with middleware.
- [picfit](https://github.com/thoas/picfit): An image resizing server.
- [dkron](https://github.com/distribworks/dkron): Distributed, fault tolerant job scheduling system.

## Contributing

Gin is the work of hundreds of contributors. We appreciate your help!

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.