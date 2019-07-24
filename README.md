Go / Gin / Quasar Framework Starter Kit
==================

## Features

&nbsp; &nbsp; ✓ Component-based front-end development via [Webpack](https://webpack.github.io/) and [React](https://facebook.github.io/react) (see [`webpack.config.js`](webpack.config.js))  
&nbsp; &nbsp; ✓ Static type checking with [TypeScript](https://www.typescriptlang.org)  
&nbsp; &nbsp; ✓ Application state management via [Redux](http://redux.js.org/)  
&nbsp; &nbsp; ✓ Universal cross-stack routing and navigation [`history`](https://github.com/ReactJSTraining/history) (see [`client/routes.tsx`](client/routes.tsx))  
&nbsp; &nbsp; ✓ Hot Module Replacement ([HMR](https://webpack.github.io/docs/hot-module-replacement.html)) /w [React Hot Loader](http://gaearon.github.io/react-hot-loader/)  
&nbsp; &nbsp; ✓ Lightweight build automation with plain JavaScript (see [`run.js`](run.js))  
&nbsp; &nbsp; ✓ Cross-device testing with [Browsersync](https://browsersync.io/)

## Styling
The project is framework agnostic so you can easily add your preferred styling framework. 
   
&nbsp; &nbsp; [Bootstrap 4 components](https://reactstrap.github.io/)   
&nbsp; &nbsp; [Ant design components](https://ant.design/)   
&nbsp; &nbsp; [Material UI components](https://material-ui.com/)   
   
Or extend the loaders in webpack.config.js to compile your own [SASS](https://github.com/webpack-contrib/sass-loader) or [LESS](https://github.com/webpack-contrib/less-loader) styles.

Another option instead of css could be to use css-in-js. Your can see a list of frameworks at MicheleBertoli's great repo [here](https://github.com/MicheleBertoli/css-in-js). [Emotion](https://github.com/emotion-js/emotion) has worked in a few test projects but feedback on a good library that plays well with typescript are appreciated. 

## Prerequisites

* OS X, Windows or Linux
* [Node.js](https://nodejs.org) v6 or newer
* [.NET Core](https://www.microsoft.com/net/core) and [.NET Core SDK](https://www.microsoft.com/net/core)
* [Visual Studio Code](https://code.visualstudio.com/) or your prefered IDE.

### Getting Started

**Step 1**. Clone the latest version of **Go / Gin / Quasar Framework Starter Kit** on your local machine by running:

```shell
$ git clone -o aspnet-starter-kit -b master --single-branch \
      https://github.com/valasek/quasar-starter-kit-go-gin.git MyApp
$ cd MyApp
```

**Step 2**. Install project dependencies listed in
[`client/package.json`](client/package.json) files: 

```shell
$ npm install                   # Install both Node.js and .NET Core dependencies
```

or using Yarn:

```shell
$ yarn install                   # Install both Node.js and .NET Core dependencies
```

**Step 3**. Finally, launch your web app:

```shell
$ quasar dev                      # Compile and lanch the app, the same as running: npm dev or yarn dev
```

The app should become available at [http://localhost:8080/](http://localhost:8080/).
See [`client/package.js`](client/package.js) for other available commands such as `node run build` etc.
You can also run your app in a release (production) mode by running `node run --release`, or without
Hot Module Replacement (HMR) by running `node run --no-hmr`.

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
