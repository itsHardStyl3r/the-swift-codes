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
httpListenOn = 0.0.0.0:8080
```

Make sure to provide database, as only the tables will be created. As of right now, only MariaDB is supported.

## Current endpoints

All responses and requests should be in `content-type: application/json; charset=utf-8`.

### 1. GET: /v1/swift-codes/{swift-code}

Retrieves details of a single SWIFT code whether for a headquarters or branches.

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

### 2. GET: /v1/swift-codes/country/{countryISO2code}

Returns all SWIFT codes with details for a specific country (both headquarters and branches).

- Response structure:

```
{
  "countryISO2": string,
  "countryName": string,
  "swiftCodes": [
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
    }, ...
  ]
}
```

### 3. POST: /v1/swift-codes

Adds new SWIFT code entries to the database for a specific country.

- Request structure:

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

- Response structure:

```
{
  "message": string
}
```

### 4. DELETE: /v1/swift-codes/{swift-code}

Deletes swift-code data if swiftCode matches the one in the database.

- Response structure:

```
{
  "message": string
}
```

## Known issues (or out of scope)

- Current implementation does not support any authentication. ⚠️
- Code performs automatic migrations on the database. This is not an issue per se, but should not be enabled in
  production environment.
- Ignoring the fact whether the codes should be marked BIC8 or BIC11.
- Not all necessary fields (e.g. town, timezone) are filled due to lack of information or the way structures are
  required to work.
- Graceful shutdown.
