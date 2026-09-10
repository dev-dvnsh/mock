# mock

A small Go CLI for generating fake, realistic data from a schema you define on the command line.

Useful for testing APIs, databases, seed data, or whenever you need some data quickly.

## Installation

```bash
go install github.com/dev-dvnsh/mock@latest
```

## Usage

```bash
mock "field1, field2, field3" count [--format json|csv] [--output filename]
```

Supported fields:

```text
name
email
age
phone
city
country
uuid
date
bool
int
password
```

JSON is the default output format. Use `--format csv` for CSV output, or `--output` to write the result to a file.

## Examples

Generate 5 users as JSON:

```bash
mock "name, email, age" 5
```

Generate 10 records as CSV:

```bash
mock "name, phone, city, country" 10 --format csv
```

Generate 20 records and save them to a file:

```bash
mock "uuid, name, email, password" 20 --output users.json
```

## License

MIT
