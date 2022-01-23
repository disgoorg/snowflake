package snowflake

import (
	"os"
	"strconv"
	"time"
)

var Epoch int64 = 1420070400000

// NewSnowflake returns a new Snowflake based on the given time.Time
//goland:noinspection GoUnusedExportedFunction
func NewSnowflake(timestamp time.Time) Snowflake {
	return Snowflake(strconv.FormatInt(((timestamp.UnixNano()/1_000_000)-Epoch)<<22, 10))
}

// GetSnowflakeEnv returns a new Snowflake from an environment variable
//goland:noinspection GoUnusedExportedFunction
func GetSnowflakeEnv(key string) Snowflake {
	return Snowflake(os.Getenv(key))
}

// Snowflake is a general utility type around discord's IDs
type Snowflake string

// String returns the string representation of the Snowflake
func (s Snowflake) String() string {
	return string(s)
}

// Int64 returns the int64 representation of the Snowflake
func (s Snowflake) Int64() int64 {
	snowflake, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return snowflake
}

// Deconstruct returns DeconstructedSnowflake (https://discord.com/developers/docs/reference#snowflakes-snowflake-id-format-structure-left-to-right)
func (s Snowflake) Deconstruct() DeconstructedSnowflake {
	snowflake := s.Int64()
	return DeconstructedSnowflake{
		Time:      time.Unix(0, ((snowflake>>22)+Epoch)*1_000_000),
		WorkerID:  (snowflake & 0x3E0000) >> 17,
		ProcessID: (snowflake & 0x1F000) >> 12,
		Increment: snowflake & 0xFFF,
	}
}

// Time returns the time.Time when the snowflake was created
func (s Snowflake) Time() time.Time {
	return s.Deconstruct().Time
}

// DeconstructedSnowflake contains the properties used by Discord for each CommandID
type DeconstructedSnowflake struct {
	Time      time.Time
	WorkerID  int64
	ProcessID int64
	Increment int64
}
