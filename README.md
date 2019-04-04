# go-cli

Makes a request to SWAPI, parses the JSON and prints output. Currently.

## Installation

```
go install go-cli
```

## Running in dev

```
go run main.go request -v -i 2
```

### Flags

#### Version

```
go-cli version
```

#### Request

```
go-cli request
```

`-i <int>` - ID to pass to request

`-v` - print response