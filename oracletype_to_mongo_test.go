package oracletype_to_mongo

import (
	"testing"
)

func TestXlat(t *testing.T) {
	// Define test cases
	testCases := []struct {
		oracleType    string
		MongoExpected string
	}{
		{"NUMBER[19,5]", "decimal"},
		{"VARCHAR2", "string"},
		{"NUMBER[1,0]", "bool"},
		{"NUMBER[2,0]", "int"},
		{"NUMBER[2,2]", "decimal"},
		{"NUMBER[19,18]", "decimal"},
		{"NUMBER[38,19]", "decimal"},
		{"NUMBER[0,0]", "int"},
		{"FLOAT", "float"},
		{"NUMBER[5,10]", "unknown"},
		{"DATE", "date"},
		{"TIMESTAMP(4)", "timestamp"},
		{"INT64", "long"},
	}

	type result_t struct {
		mongoType string
		err       error
	}

	// Run test cases
	for _, tc := range testCases {
		var r result_t
		r.mongoType, r.err = Xlat(tc.oracleType)

		// log.Print(r)

		if r.mongoType != tc.MongoExpected {
			t.Errorf("Oracle2MongoXlat(%s) = (%s, %v), expected %s", tc.oracleType, r.mongoType, r.err, tc.MongoExpected)
		}

	}
}
