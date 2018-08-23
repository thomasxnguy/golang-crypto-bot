package influxdb

import (
	"log"
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
	"encoding/json"
)

// WritePoints write an entry in influxDB.
func WritePoints(clnt client.Client, bp client.BatchPoints) {
	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

// GetLastTimestamp get the last timestamp from a specific measurement.
func GetLastTimestamp(clnt client.Client, db string, precision string, measurement string, value string) (int64, error) {
	ts, err := queryDB(clnt, db, precision,
		fmt.Sprintf("Select last(%s),time from %s", value, measurement))

	if err != nil || ts == nil || ts[0].Series == nil {
		return 0, err
	}
	return ts[0].Series[0].Values[0][0].(json.Number).Int64()
}

// queryDB call InfluxDB with specific cmd.
func queryDB(clnt client.Client, db string, precision string, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:   cmd,
		Database:  db,
		Precision: precision,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
