# Daily Coding Problems
Resource to hold daily coding problem solutions

# Table of Contents

* [Overview](#Overview)
* [Deployment and Configuration](#deployment-and-configuration)
    * [Analytics](#analytics)
* [Developer Details](#developer-details)
    * [Installation](#installation)
    * [File Linting](#file-linting)
    * [Testing](#testing)
    * [Gotchas](#gotchas)



# Overview
Brain teasers for keeping the mind fresh.  

# Deployment and Configuration

## Configuration Variables

| Name | Example | Description|
|------|---------|------------|

## Analytics
N/A

# Developer Details

## Installation

### Install Go
Install the Go programming language compiler (version 1.16 or later) and toolset: [golang.org](https://golang.org/doc/install)

### Clone from Github
Clone the repo:  `git clone git@github.com:pocketazn/CodingProblems.git`

### Setup

Run  `go mod vendor` and `go mod tidy` to sync dependencies

## File Linting

All major editors have Go plugins that can automate running `gofmt` and/or `gometalinter` and other tools to keep your code clean, consistent, and properly formatted.

## Testing

Testing should cover all use cases and edge cases for problems.

### Requirements
We require 100%* test coverage of this repo. You *must* strive to increase test coverage as much as possible. No exceptions.

### Running Tests

run `go test ./...` to run all tests

## Gotchas