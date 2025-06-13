package utility

import (
	"metaLand/utility/net"
	"metaLand/utility/sequence"
	"strconv"
	"strings"
)

// Snowflake init logic, have to first check the ip
// use the database id as the

var Sequence sequence.Senquence
var AwsFileSequence sequence.Senquence

func initSequence() (err error) {
	machineIP := net.GetDomianIP()
	machineSignature := strings.Replace(machineIP, ".", "", 4)
	machineID, err := strconv.ParseInt(machineSignature, 10, 64)
	machineID %= 32
	if err != nil {
		return
	}
	// Create snowflake sequences
	Sequence = sequence.NewSnowflake(uint64(8000), uint64(machineID))
	AwsFileSequence = sequence.NewSnowflake(uint64(8000), uint64(machineID))
	return
}
