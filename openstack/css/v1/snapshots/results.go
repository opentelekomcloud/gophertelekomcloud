package snapshots

import (
	"time"
)

// Snapshot contains all the information associated with a Cluster Snapshot.
type Snapshot struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"backupType"`
	Method        string `json:"backupMethod"`
	Description   string `json:"description"`
	ClusterID     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	Indices       string `json:"indices"`
	TotalShards   int    `json:"totalShards"`
	FailedShards  int    `json:"failedShards"`
	KeepDays      int    `json:"backupKeepDay"`
	Period        string `json:"backupPeriod"`
	Bucket        string `json:"bucketName"`
	Version       string `json:"version"`
	Status        string `json:"status"`
	RestoreStatus string `json:"restoreStatus"`
	// type of the data search engine
	DataStore DataStore `json:"datastore"`
	// the information about times
	ExpectedStartTime time.Time `json:"-"`
	StartTime         time.Time `json:"-"`
	EndTime           time.Time `json:"-"`
	Created           string    `json:"created"`
	Updated           string    `json:"updated"`
}

type DataStore struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}
