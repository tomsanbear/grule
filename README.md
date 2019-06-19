[![Build Status](https://travis-ci.org/tomsanbear/grule.svg?branch=master)](https://travis-ci.org/tomsanbear/grule) [![codecov](https://codecov.io/gh/tomsanbear/grule/branch/master/graph/badge.svg)](https://codecov.io/gh/tomsanbear/grule) [![GitHub version](https://badge.fury.io/gh/tomsanbear%2Fgrule.svg)](https://badge.fury.io/gh/tomsanbear%2Fgrule) ![GitHub issues](https://img.shields.io/github/issues/tomsanbear/grule.svg) ![GitHub pull requests](https://img.shields.io/github/issues-pr/tomsanbear/grule.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/tomsanbear/grule)](https://goreportcard.com/report/github.com/tomsanbear/grule)

# Grule - A Golang Rules Engine

Go implementation of a generic rules engine. Accepts, and compiles from a DSL language, or can be modified at runtime to provide dynamic rule changes.

## About the library

This library was created as an easy to integrate, customize and extend way to implement a generic rules engine into your app. Some general concepts are outlined below:

### Major Components

#### Conditions

Conditions are the major logical component of the system, and allow you to extend this as far as you want, and suit your needs. Basic operators are included in this library for convenience.

#### Facts

Facts are *things* that are statements about the rules system you are building. If you are building a policy system, this would be akin to the list of employees that are allowed to access a certain system.

Some other possible properties are the following:

- CPU percent limits
- A list of domain names to match against

#### Events

Events are *things* that happen within the system, they have their own properties associated with them, maybe it lines up with your Facts scheme, but it doesn't need to.

- A request to access a resource (employee name, badge id, source IP)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

- Go 12+, with modules enabled

### Installing

Just add this lib to your go.mod

## Running the tests

For basic unit test coverage:
```make test```

For benchmarks:
```make bench```

## Built With

TBD

## Contributing

Please read [CONTRIBUTING.md](https://github.com/tomsanbear/grule/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

- **Thomas Santerre** - *Initial work* - [tomsanbear](https://github.com/tomsanbear)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## Acknowledgments

- TBD
