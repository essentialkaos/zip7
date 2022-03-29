<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zip7.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/zip7.v1"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://kaos.sh/w/zip7/ci"><img src="https://kaos.sh/w/zip7/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/zip7"><img src="https://kaos.sh/r/zip7.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/zip7"><img src="https://codebeat.co/badges/11fb655d-8da8-4694-a32b-b95ff9eed602" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/zip7/codeql"><img src="https://kaos.sh/w/zip7/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#compatibility-and-os-support">Compatibility and OS support</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`zip7` package provides methods for working with 7z archives (`p7zip` wrapper).

### Installation

Make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get github.com/essentialkaos/zip7
```

If you want to update `zip7` to latest stable release, do:

```
go get -u github.com/essentialkaos/zip7
```

### Compatibility and OS support

|      Version |     1.x |
|--------------|---------|
|  `p7zip 9.x` | Partial |
| `p7zip 15.x` |    Full |
| `p7zip 16.x` |    Full |

| OS       | Support            |
|----------|--------------------|
| Linux    | :heavy_check_mark: |
| Mac OS X | :heavy_check_mark: |
| FreeBSD  | :heavy_check_mark: |
| Windows  | :x:                |

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/zip7/ci.svg?branch=master)](https://kaos.sh/w/zip7/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/zip7/ci.svg?branch=develop)](https://kaos.sh/w/zip7/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
