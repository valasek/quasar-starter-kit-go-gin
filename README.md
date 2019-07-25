Go / Gin / Quasar Framework Starter Kit
==================

## Features

&nbsp; &nbsp; ✓ Allows you to start server on HTTPS / HTTP

&nbsp; &nbsp; ✓ Includes file download example

&nbsp; &nbsp; ✓ Includes file upload example

DB support can be added via go/db or e.g. via gorm- [Gorm](https://github.com/jinzhu/gorm)

# Standing on the shoulders of giants

[Go](https://golang.org/), [Gin web framework](https://github.com/gin-gonic), [Vue](https://vuejs.org/), [Quasar](https://quasar.dev/)

## Go Backend

- [Gin web framework](https://github.com/gin-gonic)
- [Logrus](https://github.com/sirupsen/logrus), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), [lumberjackrus](https://github.com/orandin/lumberjackrus)

## JS Frontend

- [Vue.js](https://vuejs.org/) spa client with webpack
- [Quasar framework](https://quasar.dev/)
- [Axios](https://github.com/axios/axios)

## Prerequisites

* OS X, Windows or Linux
* [Node.js](https://nodejs.org) v6 or newer
* [Go](https://golang.org/)
* [Visual Studio Code](https://code.visualstudio.com/) or your prefered IDE.

# Getting Started

**Step 1**. Clone the latest version of **Go / Gin / Quasar Framework Starter Kit** on your local machine by running:

```shell
$ git clone -o quasar-starter-kit-go-gin -b master --single-branch https://github.com/valasek/quasar-starter-kit-go-gin.git MyApp
$ cd MyApp
```

## Client

**Step 2**. Install project dependencies listed in [`client/package.json`](client/package.json) files: 

```shell
$ cd client
```

```shell
$ npm install                   # Install Node.js dependencies
```

or using Yarn:

```shell
$ yarn install                   # Install Node.js dependencies
```

**Step 3**. Launch your web app:

```shell
$ quasar dev                      # Compile and lanch the app, the same as running: npm dev or yarn dev
```

The app should become available at [http://localhost:8080/](http://localhost:8080/). See [`client/package.json`](client/package.json) for other available commands such as `quasar build` etc.

## Server

**Step 4**. Launch your server: 

```shell
$ cd ../server
$ go run quasar-starter-kit-go-gin.go server    # Will install required Go packages and run the server in dev mode
```

The server will listen on [http://localhost:3000/](http://localhost:3000/).

## How to Update

You can always fetch and merge the latest changes from this (upstream) repo back into your project by running:

```shell
$ git checkout master
$ git fetch quasar-starter-kit-go-gin
$ git merge quasar-starter-kit-go-gin/master 
```

## How to Contribute

Anyone and everyone is welcome to contribute. The best way to start is by checking our [open issues](https://github.com/valasek/quasar-starter-kit-go-gin/issues), [submit a new issues](https://github.com/valasek/quasar-starter-kit-go-gin/issues/new?labels=bug) or [feature request](https://github.com/valasek/quasar-starter-kit-go-gin/issues/new?labels=enhancement), participate in discussions, upvote or downvote the issues you like or dislike, send pull requests.

## Get in Touch

[Twitter](https://twitter.com/valasek), [Web](https://twitter.com/valasek)

## License

Copyright © 2019-present [Stanislav Valasek](https://valasek.wordpress.com). This source code is licensed under the MIT

---
Made by Stanislav Valasek [@valasek](https://twitter.com/valasek)
