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

## Running tests

In root directory run:

```
cd test
go test -v
```

### Example run

<details>
<summary>Show example run...</summary>

```
PS E:\the-swift-codes\backend\test> go test -v
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
- using env:   export GIN_MODE=release
- using code:  gin.SetMode(gin.ReleaseMode)

=== RUN   TestAPI
2025/03/30 21:53:02 INFO Current database status:
2025/03/30 21:53:02 INFO - countries: 4 entries
2025/03/30 21:53:02 INFO - banks: 3 entries
2025/03/30 21:53:02 INFO - bics: 4 entries
=== RUN   TestAPI/TestByCountryISO2
[GIN] 2025/03/30 - 21:53:02 | 200 |            0s |                 | GET      "/v1/swift-codes/country/BG"
=== RUN   TestAPI/TestByCountryISO2OnInvalid
[GIN] 2025/03/30 - 21:53:02 | 400 |            0s |                 | GET      "/v1/swift-codes/country/A"
[GIN] 2025/03/30 - 21:53:02 | 400 |            0s |                 | GET      "/v1/swift-codes/country/AAA"
=== RUN   TestAPI/TestByCountryISO2OnNonExistent

2025/03/30 21:53:02 E:/the-swift-codes/backend/cmd/api/getByCountryISO2.go:27 record not found
[0.000ms] [rows:0] SELECT * FROM `countries` WHERE iso2 = "EU" ORDER BY `countries`.`id` LIMIT 1
[GIN] 2025/03/30 - 21:53:02 | 404 |            0s |                 | GET      "/v1/swift-codes/country/EU"
=== RUN   TestAPI/TestByCountryISO2WithSwifts
[GIN] 2025/03/30 - 21:53:02 | 200 |            0s |                 | GET      "/v1/swift-codes/country/PL"
=== RUN   TestAPI/TestBySwiftCodeBranches
[GIN] 2025/03/30 - 21:53:02 | 200 |            0s |                 | GET      "/v1/swift-codes/LITWLTDEADD"
=== RUN   TestAPI/TestBySwiftCodeHeadquarters
[GIN] 2025/03/30 - 21:53:02 | 200 |       551.9µs |                 | GET      "/v1/swift-codes/POLSPLAWXXX"
=== RUN   TestAPI/TestBySwiftCodeOnNonExistent

2025/03/30 21:53:02 E:/the-swift-codes/backend/cmd/api/getBySwiftCode.go:27 record not found
[0.000ms] [rows:0] SELECT `bics`.`id`,`bics`.`country_id`,`bics`.`bic`,`bics`.`code_type`,`bics`.`bank_id`,`bics`.`address`,`bics`.`town`,`bics`.`time_zone`,`bics`.`location_code`,`bics`.`branch`,`Bank`.`id` AS `Bank__id`,`Bank`.`name` AS `Bank__name`,`Bank`.`bank_code` AS `Bank__bank_code`,`Country`.`id` AS `Country__id`,`Country`.`name` AS `Country__name`,`Country`.`iso2` AS `Country__iso2` FROM `bics` LEFT JOIN `banks` `Bank` ON `bics`.`bank_id` = `Bank`.`id` LEFT JOIN `countries` `Country` ON `bics`.`country_id` = `Country`.`id` WHERE bic = "ABCDEFGHIJK" ORDER BY `bics`.`id` LIMIT 1
[GIN] 2025/03/30 - 21:53:02 | 404 |            0s |                 | GET      "/v1/swift-codes/ABCDEFGHIJK"
=== RUN   TestAPI/TestDeleteSwiftDeleteBySwiftCode
[GIN] 2025/03/30 - 21:53:02 | 200 |            0s |                 | DELETE   "/v1/swift-codes/LITWLTDEXXX"
=== RUN   TestAPI/TestDeleteSwiftOnInvalid
[GIN] 2025/03/30 - 21:53:02 | 400 |            0s |                 | DELETE   "/v1/swift-codes/1234"
=== RUN   TestAPI/TestDeleteSwiftOnNonExistent
[GIN] 2025/03/30 - 21:53:02 | 404 |            0s |                 | DELETE   "/v1/swift-codes/XXXXXXXXXXX"
=== RUN   TestAPI/TestPostSwiftCode

2025/03/30 21:53:02 E:/the-swift-codes/backend/cmd/api/postSwiftCode.go:49 record not found
[0.000ms] [rows:0] SELECT * FROM `bics` WHERE bic = "POLSAWAWXXX" ORDER BY `bics`.`id` LIMIT 1
[GIN] 2025/03/30 - 21:53:02 | 200 |       524.7µs |                 | POST     "/v1/swift-codes"
=== RUN   TestAPI/TestPostSwiftCodeInvalidBody
2025/03/30 21:53:02 DEBU Invalid JSON body: Key: 'PostSwiftRequest.Address' Error:Field validation for 'Address' failed on the 'required' tag
Key: 'PostSwiftRequest.BankName' Error:Field validation for 'BankName' failed on the 'required' tag
Key: 'PostSwiftRequest.CountryISO2' Error:Field validation for 'CountryISO2' failed on the 'required' tag
Key: 'PostSwiftRequest.CountryName' Error:Field validation for 'CountryName' failed on the 'required' tag
Key: 'PostSwiftRequest.SwiftCode' Error:Field validation for 'SwiftCode' failed on the 'required' tag.
[GIN] 2025/03/30 - 21:53:02 | 400 |       519.5µs |                 | POST     "/v1/swift-codes"
=== RUN   TestAPI/TestPostSwiftCodeInvalidCountry

2025/03/30 21:53:02 E:/the-swift-codes/backend/cmd/api/postSwiftCode.go:49 record not found
[0.000ms] [rows:0] SELECT * FROM `bics` WHERE bic = "POLSDDAWXXX" ORDER BY `bics`.`id` LIMIT 1

2025/03/30 21:53:02 E:/the-swift-codes/backend/cmd/api/postSwiftCode.go:57 record not found
[0.000ms] [rows:0] SELECT * FROM `countries` WHERE iso2 = "PL" AND name = "POLAN" ORDER BY `countries`.`id` LIMIT 1
[GIN] 2025/03/30 - 21:53:02 | 404 |            0s |                 | POST     "/v1/swift-codes"
=== RUN   TestAPI/TestPostSwiftCodeInvalidSwift
[GIN] 2025/03/30 - 21:53:02 | 400 |            0s |                 | POST     "/v1/swift-codes"
2025/03/30 21:53:02 INFO Current database status:
2025/03/30 21:53:02 INFO - countries: 4 entries
2025/03/30 21:53:02 INFO - banks: 3 entries
2025/03/30 21:53:02 INFO - bics: 4 entries
--- PASS: TestAPI (0.02s)
--- PASS: TestAPI/TestByCountryISO2 (0.00s)
--- PASS: TestAPI/TestByCountryISO2OnInvalid (0.00s)
--- PASS: TestAPI/TestByCountryISO2OnNonExistent (0.00s)
--- PASS: TestAPI/TestByCountryISO2WithSwifts (0.00s)
--- PASS: TestAPI/TestBySwiftCodeBranches (0.00s)
--- PASS: TestAPI/TestBySwiftCodeHeadquarters (0.00s)
--- PASS: TestAPI/TestBySwiftCodeOnNonExistent (0.00s)
--- PASS: TestAPI/TestDeleteSwiftDeleteBySwiftCode (0.00s)
--- PASS: TestAPI/TestDeleteSwiftOnInvalid (0.00s)
--- PASS: TestAPI/TestDeleteSwiftOnNonExistent (0.00s)
--- PASS: TestAPI/TestPostSwiftCode (0.00s)
--- PASS: TestAPI/TestPostSwiftCodeInvalidBody (0.00s)
--- PASS: TestAPI/TestPostSwiftCodeInvalidCountry (0.00s)
--- PASS: TestAPI/TestPostSwiftCodeInvalidSwift (0.00s)
PASS
ok      github.com/itsHardStyl3r/the-swift-codes/test   0.104s
```

</details>

## Known issues (or out of scope)

- Current implementation does not support any authentication. ⚠️
- Code performs automatic migrations on the database. This is not an issue per se, but should not be enabled in
  production environment.
- Ignoring the fact whether the codes should be marked BIC8 or BIC11.
- Not all necessary fields (e.g. town, timezone) are filled due to lack of information or the way structures are
  required to work.
- Graceful shutdown.
