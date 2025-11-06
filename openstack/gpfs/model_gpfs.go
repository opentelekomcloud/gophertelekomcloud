package gpfs

// ListFSInput is the input parameter of ListFS function
type ListFSInput struct {
	BucketType string
}

// ListFSOutput is the result of ListFS function
type ListFSOutput struct {
	BaseModel
	Owner   Owner    `xml:"Owner"`
	Buckets []Bucket `xml:"Buckets>Bucket"`
}

// CreateFSInput is the input parameter of CreateBucket function
type CreateFSInput struct {
	BucketLocation
	FSName     string `xml:"-"`
	Redundancy string `xml:"-"`
	BucketType string `xml:"-"`
}
