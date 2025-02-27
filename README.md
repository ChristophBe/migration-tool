# Migration Tool

This migration tool should support the migration process between versions of applications.
With the migration tool steps that are based of bash scripts can be defined and orchestrated. 

# Usage

## Run with Docker
```bash
docker run -v ./examples/migrations:/migrations ghcr.io/christophbe/migration-tool:latest run
```


## Commands
### Run Migrations
Run all migration steps, that were not executed before.
```bash
migration-tool run -folder migrations
```

If any migration file has changed, execution will be aborted to prevent inconsistency.

### Recalculate Hashes
Recalculate the hashes of migration scripts and update `migrations.yaml`.
```bash
migration-tool recalculate-hashes -folder migrations
```

### Verify Migrations
Check if any migration files have changed based on stored hashes.
```bash
migration-tool verify -folder migrations
```

## Migrations File

The `migrations.yaml` file should be structured as follows:
```yaml
migrations:
  - filename: 001_first_step.sh
    description: |
      Execute step one of the migrations.
    hash: <auto-generated>
  - filename: 002_second_step.sh
    description: |
      Execute step two of the migrations.
    hash: <auto-generated>
```

# Development

## Setup
1. Clone this Repository
2. Install needed development tooling.
   - [Golang CLI](https://go.dev/dl) Version 1.24 or later
   - (optional) [mockery](https://vektra.github.io/mockery/latest/installation/) for generation of mocks

## Run and Build 
### Run from Code 
```bash
go run  github.com/ChristophBe/migration-tool/cmd/migration-tool
```

### Build executable
```bash
go build  github.com/ChristophBe/migration-tool/cmd/migration-tool
```

## Run Tests
```bash
go test ./...
```

## Generate Code
For testing mock implementations of interfaces are used. To generate the mocks [mockery](https://vektra.github.io/mockery/latest/) is used.

To run regenerate the generated files run the following command:
```bash
go generate
```