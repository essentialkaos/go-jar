<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-jar.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/jar.v1"><img src="https://godoc.org/pkg.re/essentialkaos/jar.v1?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/jar"><img src="https://goreportcard.com/badge/github.com/essentialkaos/jar"></a>
  <a href="https://travis-ci.com/essentialkaos/jar"><img src="https://travis-ci.com/essentialkaos/jar.svg"></a>
  <a href="https://github.com/essentialkaos/jar/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/jar/workflows/CodeQL/badge.svg" /></a>
  <a href="https://coveralls.io/github/essentialkaos/jar?branch=master"><img src="https://coveralls.io/repos/github/essentialkaos/jar/badge.svg?branch=master" alt="Coverage Status" /></a>
  <a href="https://github.com/essentialkaos/jar/actions"><img src="https://github.com/essentialkaos/jar/workflows/Go/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-jar-master"><img alt="codebeat badge" src="https://codebeat.co/badges/ca62b3ee-901f-4410-9a0d-bca2d071b95d" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg" alt="License" />
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`jar` is a very simple Go package for reading manifest data from JAR files.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.12+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/jar.v1
```

For update to the latest stable release, do:

```
go get -u pkg.re/essentialkaos/jar.v1
```

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.com/essentialkaos/jar.svg?branch=master)](https://travis-ci.com/essentialkaos/jar) |
| `develop` | [![Build Status](https://travis-ci.com/essentialkaos/jar.svg?branch=develop)](https://travis-ci.com/essentialkaos/jar) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
