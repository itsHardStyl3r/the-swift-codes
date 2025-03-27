# The Swift Codes project

This repository contains both backend and frontend, although frontend was out of the scope of this project.

<img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="32" height="32"/> <img src="https://www.vectorlogo.zone/logos/mariadb/mariadb-icon.svg" alt="mariadb" width="32" height="32"/> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/typescript/typescript-plain.svg" alt="typescript" width="32" height="32"/> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/vitejs/vitejs-original.svg" alt="vitejs" width="32" height="32"/> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/pnpm/pnpm-original.svg" alt="pnpm" width="32" height="32"/>

## Requirements

Backend is written in Go and uses mariadb. Frontend uses typescript and pnpm, so naturally you'd need node.js.

### Installing Go

Follow detailed instructions on [golang's website](https://go.dev/doc/install) for all supported platforms.

- Windows (10 or newer)

Alternatively to the aforementioned instructions, you can install go using [scoop.sh](https://scoop.sh/).
To install scoop, in powershell run:

```
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

and then

```
scoop install go
```

and go should be installed and added to your path.

- Linux (apt package manager)

Update your package lists using

```
sudo apt update
```

and install `golang-go` with

```
apt install -y golang-go
```

and go should be installed on your system.

### Installing node.js and pnpm

You can follow detailed instructions on [nodejs.org webiste](https://nodejs.org/en/download), especially for Linux.

- Windows

Again, you can use [scoop.sh](https://scoop.sh/) to install both. Head above to see its installation details and once
complete, run

```
scoop install nodejs-lts pnpm
```

and you're done. Alternatively, use winget
as [explained here](https://nodejs.org/en/download/package-manager/all#windows-1).

### Installing mariadb

- Windows

Use scoop to install mariadb, run

```
scoop install mariadb
```

or follow instructions [on mariadb website](https://mariadb.com/kb/en/installing-mariadb-msi-packages-on-windows/).

- Linux (apt package manager)

Run

```
sudo apt update
sudo apt install -y mariadb-server
```

and then

```
sudo mysql_secure_installation
```

to configure the database.

## Installation

Clone the repository in a desired place using:

```
git clone https://github.com/itsHardStyl3r/the-swift-codes
```

### Installing backend

Change directory to `./backend` and run

```
go mod tidy
go mod verify
```

### Installing frontend

Change directory to `./frontend` and run

```
pnpm install
```

## Running the applications

### Running backend

```
go run cmd/main.go
```

### Running frontend

```
pnpm run dev
```

## Running tests (backend only)

See [backend README](backend/README.md).