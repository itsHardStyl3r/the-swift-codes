# The Swift Codes (backend)

This repository contains backend for The Swift Codes.
Written in Go. Uses [MariaDB](https://github.com/mariadb) for database and [GORM](https://gorm.io/) for operations on
it.
Environmental variables are implemented with [godotenv](https://github.com/joho/godotenv).

<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="32" height="32"/> <img src="https://www.vectorlogo.zone/logos/mariadb/mariadb-icon.svg" alt="mariadb" width="32" height="32"/>

## Before you start

Make sure to create .env in the root directory of the application file with contents as follows:

```
csvDataPath = <path to csv file to read>
dbUser =
dbPassword =
dbAddr = address:port
dbDatabase =
```

Make sure to provide database, as only the tables will be created. As of right now, only MariaDB is supported.

## Current endpoints

### 1. GET: /v1/swift-codes/{swift-code}

Retrieves details of a single SWIFT code whether for a headquarters or
branches.

- Response structure if {swift-code} points to headquarters:

```
{
  "address": string,
  "bankName": string,
  "countryISO2": string,
  "countryName": string,
  "isHeadquarter": bool,
  "swiftCode": string,
  "branches": [
    {
      "address": string,
      "bankName": string,
      "countryISO2": string,
      "isHeadquarter": bool,
      "swiftCode": string
    },
    {
      "address": string,
      "bankName": string,
      "countryISO2": string,
      "isHeadquarter": bool,
      "swiftCode": string
    },
    ...
  ]
}
```

- Response structure if {swift-code} points to a branch:

```
{
  "address": string,
  "bankName": string,
  "countryISO2": string,
  "countryName": string,
  "isHeadquarter": bool,
  "swiftCode": string
}
```

## Known issues (or out of scope)

- Current implementation does not support any authentication. ⚠️
- Code performs automatic migrations on the database. This is not an issue per se, but should not be enabled in
  production environment.
