package resource

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Resource struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Name is a name of the resource. The name contains a maximum of 32 characters, including only letters, numbers, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`
	// Type is a resource type. Can be: archive, file, jar
	Type string `json:"type" required:"true"`
	// Location is an OBS path for storing the resource file. When type is set to jar, location is the path for storing the main JAR package. The path contains a maximum of 1,023 characters. For example, obs://myBucket/test.jar
	Location string `json:"location" required:"true"`
	// JAR package and properties file that the main JAR package depends on. The description contains a maximum of 10,240 characters. If this parameter and the dependFiles parameter are both available, this parameter is preferentially parsed.
	DependPackages []*DependPackage `json:"dependPackages,omitempty"`
	// DependFiles is a JAR package and properties file that the main JAR package depends on. The description contains a maximum of 10,240 characters.
	DependFiles []string `json:"dependFiles,omitempty"`
	// Description of the resource. The description contains a maximum of 255 characters.
	Desc string `json:"desc,omitempty"`
	// Directory for storing the resource. Access the DataArts Studio console and choose Data Development. The default directory is the root directory.
	Directory string `json:"directory,omitempty"`
}

type DependPackage struct {
	// Type is a file type.
	Type string `json:"type,omitempty"`
	// Location is a file path.
	Location string `json:"location,omitempty"`
}

// Create is used to create a resource. Types of nodes, including DLI Spark, MRS Spark, and MRS MapReduce, can reference files such as JAR and properties through resources.
// Send request POST /v1/{project_id}/resources
func Create(client *golangsdk.ServiceClient, opts Resource) (*CreateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
	}

	if opts.Workspace != "" {
		reqOpts.MoreHeaders[HeaderWorkspace] = opts.Workspace
	}

	raw, err := client.Post(client.ServiceURL(resourcesEndpoint), b, nil, reqOpts)
	if err != nil {
		return nil, err
	}

	var resp CreateResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err
}

type CreateResp struct {
	ResourceId string `json:"resourceId"`
}
