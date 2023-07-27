package golangsdk

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"
)

// RFC3339Milli describes a common time format used by some API responses.
const RFC3339Milli = "2006-01-02T15:04:05.999999Z"

type JSONRFC3339Milli time.Time

func (jt *JSONRFC3339Milli) UnmarshalJSON(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	t, err := time.Parse(RFC3339Milli, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339Milli(t)
	return nil
}

const RFC3339MilliNoZ = "2006-01-02T15:04:05.999999"

type JSONRFC3339MilliNoZ time.Time

func (jt *JSONRFC3339MilliNoZ) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339MilliNoZ, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339MilliNoZ(t)
	return nil
}

type JSONRFC1123 time.Time

func (jt *JSONRFC1123) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(time.RFC1123, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC1123(t)
	return nil
}

type JSONUnix time.Time

func (jt *JSONUnix) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	unix, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	TempT = time.Unix(unix, 0)
	*jt = JSONUnix(TempT)
	return nil
}

// RFC3339NoZ is the time format used in Heat (Orchestration).
const RFC3339NoZ = "2006-01-02T15:04:05"

type JSONRFC3339NoZ time.Time

func (jt *JSONRFC3339NoZ) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339NoZ, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339NoZ(t)
	return nil
}

// RFC3339ZNoT is the time format used in Zun (Containers Service).
const RFC3339ZNoT = "2006-01-02 15:04:05-07:00"

type JSONRFC3339ZNoT time.Time

func (jt *JSONRFC3339ZNoT) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339ZNoT, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339ZNoT(t)
	return nil
}

// RFC3339ZNoTNoZ is another time format used in Zun (Containers Service).
const RFC3339ZNoTNoZ = "2006-01-02 15:04:05"

type JSONRFC3339ZNoTNoZ time.Time

func (jt *JSONRFC3339ZNoTNoZ) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse(RFC3339ZNoTNoZ, s)
	if err != nil {
		return err
	}
	*jt = JSONRFC3339ZNoTNoZ(t)
	return nil
}
