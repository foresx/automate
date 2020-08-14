package relaxting

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/chef/automate/components/compliance-service/ingest/ingestic/mappings"
  "time"
  "strings"
)


func TestGetEsIndex(t *testing.T) {
  sum := mappings.IndexNameSum
  rep := mappings.IndexNameRep

  filters := make(map[string][]string)
  // Used by calls like ListReportIds to get all report IDs from the beginning of time
  index, err :=	GetEsIndex(filters, false, true)
  assert.NoError(t, err)
  assert.True(t, strings.HasPrefix(index, rep + "-2017*,"+rep+"-2018*,"+rep+"-2019*,"+rep+"-2020"))

  filters["profile_id"] = []string{"prof1", "prof2"}
  filters["end_time"] = []string{"2019-01-05T23:59:59Z", "2019-01-06T23:59:59Z"} // 2019-01-06 should be ignored

  index, err =	GetEsIndex(filters, true, false)
  assert.Equal(t, sum + "-2019.01.05*", index)
  assert.NoError(t, err)

  index, err =	GetEsIndex(filters, true, true)
  assert.Equal(t, sum+"-2017*,"+sum+"-2018*,"+sum+"-2019.01.01*,"+sum+"-2019.01.02*,"+sum+"-2019.01.03*,"+sum+"-2019.01.04*,"+sum+"-2019.01.05*", index)
  assert.NoError(t, err)

  index, err =	GetEsIndex(filters, false, false)
  assert.Equal(t, rep + "-2019.01.05*", index)
  assert.NoError(t, err)

  index, err =	GetEsIndex(filters, false, true)
  assert.Equal(t, rep+"-2017*,"+rep+"-2018*,"+rep+"-2019.01.01*,"+rep+"-2019.01.02*,"+rep+"-2019.01.03*,"+rep+"-2019.01.04*,"+rep+"-2019.01.05*", index)
  assert.NoError(t, err)

  filters["start_time"] = []string{"2100-01-12T23:59:59Z"}
  index, err =	GetEsIndex(filters, false, true)
  assert.EqualError(t, err, "bad date range. end date must be greater than or equal to start date")

  filters["start_time"] = []string{"2019-01-02T23:59:59Z", "2019-01-03T23:59:59Z"} // 2019-01-03 should be ignored
  index, err =	GetEsIndex(filters, false, true)
  assert.Equal(t, rep+"-2019.01.02*,"+rep+"-2019.01.03*,"+rep+"-2019.01.04*,"+rep+"-2019.01.05*", index)
  assert.NoError(t, err)

  // Last 24h test when no start_time and end_time filters are provided
  filters = make(map[string][]string)
  index, err =	GetEsIndex(filters, true, false)
  today := time.Now().UTC().Format("2006.01.02")
  yesterday := time.Now().Add(-24 * time.Hour).UTC().Format("2006.01.02")
  assert.Equal(t, sum+"-"+yesterday+"*,"+sum+"-"+today+"*", index)
  assert.NoError(t, err)
}