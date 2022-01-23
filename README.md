[![Go Reference](https://pkg.go.dev/badge/github.com/DisgoOrg/snowflake.svg)](https://pkg.go.dev/github.com/DisgoOrg/disgo)
[![Go Version](https://img.shields.io/github/go-mod/go-version/DisgoOrg/snowflake)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/DisgoOrg/disgo/blob/master/LICENSE)
[![Disgo Version](https://img.shields.io/github/v/tag/DisgoOrg/snowflake?label=release)](https://github.com/DisgoOrg/snowflake/releases/latest)
[![Disgo Discord](https://discord.com/api/guilds/817327181659111454/widget.png)](https://discord.gg/TewhTfDpvW)

# snowflake

snowflake is a golang library for parsing [snowflake IDs](https://docs.snowflake.com) from discord.
This package provides a custom `snowflake` type which has various utility methods for parsing snowflake IDs.

### Installing

```sh
go get github.com/DisgoOrg/snowflake
```

## Usage

```go

id := Snowflake("123456789012345678")

// deconstructs the snowflake ID into its components timestamp, worker ID, process ID, and increment
id.Deconstruct()

// time.Time when the snowflake was generated
id.Time()

// returns the string representation of the snowflake ID
id.String()

// returns the int64 representation of the snowflake ID
id.Int64()

// returns a new snowflake with worker ID, process ID, and increment set to 0
// this can be used for various pagination requests to the discord api
id = NewSnowflake(time.Now())

// returns a snowflake from an environment variable
id = NewSnowflakeEnv("guild_id")
```

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/
Org/snowflake/blob/master/LICENSE). See LICENSE for more information.
