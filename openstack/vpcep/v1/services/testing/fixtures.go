package testing

const (
	createRequest = `
{
   "port_id":"4189d3c2-8882-4871-a3c2-d380272eed88",
   "vpc_id":"4189d3c2-8882-4871-a3c2-d380272eed80",
   "approval_enabled":false,
   "service_type":"interface",
   "server_type":"VM",
   "ports":
  [
    {
      "client_port":8080,
      "server_port":90,
      "protocol":"TCP"
    },
    {
      "client_port":8081,
      "server_port":80,
      "protocol":"TCP"
    }
  ]
}
`
	createResponse = `
{
    "id":"4189d3c2-8882-4871-a3c2-d380272eed83",
    "port_id":"4189d3c2-8882-4871-a3c2-d380272eed88",
    "vpc_id":"4189d3c2-8882-4871-a3c2-d380272eed80",
    "pool_id":"5289d3c2-8882-4871-a3c2-d380272eed80",
    "status":"available",
    "approval_enabled":false,
    "service_name":"test123",
    "service_type":"interface",
    "server_type":"VM",
    "project_id":"6e9dfd51d1124e8d8498dce894923a0d",
    "created_at":"2018-01-30T07:42:01.174",
    "ports":
              [
                {
                    "client_port":8080,
                    "server_port":90,
                    "protocol":"TCP"
                },
                {
                    "client_port":8081,
                    "server_port":80,
                    "protocol":"TCP"
                }
              ]
}
`
	updateRequest = `
{
   "approval_enabled":true,
   "service_name":"test",
   "ports":[
             {
                "client_port":8081,
                "server_port":22,
                "protocol":"TCP"
             },
             {
                "client_port":8082,
                "server_port":23,
                "protocol":"UDP"
             }
           ]
}
`

	updateResponse = `
{
    "id":"4189d3c2-8882-4871-a3c2-d380272eed83",
    "port_id":"4189d3c2-8882-4871-a3c2-d380272eed88",
    "vpc_id":"4189d3c2-8882-4871-a3c2-d380272eed80",
    "pool_id":"5289d3c2-8882-4871-a3c2-d380272eed80",
    "status":"available",
    "approval_enabled":true,
    "service_name":"test123",
    "service_type":"interface",
    "server_type":"VM",
    "project_id":"6e9dfd51d1124e8d8498dce894923a0d",
    "created_at":"2018-01-30T07:42:01.174",
    "ports":[
             {
                "client_port":8081,
                "server_port":22,
                "protocol":"TCP"
             },
             {
                "client_port":8082,
                "server_port":23,
                "protocol":"UDP"
             }
           ]
}
`
	listResponse = `
{
   "endpoint_services":[
         {
           "id":"4189d3c2-8882-4871-a3c2-d380272eed83",
           "port_id":"4189d3c2-8882-4871-a3c2-d380272eed88",
           "vpc_id":"4189d3c2-8882-4871-a3c2-d380272eed80",
           "status":"available",
           "approval_enabled":false,
           "service_name":"test123",
           "server_type":"VM",
           "service_type":"interface",
           "ports":[
                {
                  "client_port":8080,
                  "server_port":90,
                  "protocol":"TCP"
                },
                {
                  "client_port":8081,
                  "server_port":80,
                  "protocol":"TCP"
                }
             ],
           "project_id":"6e9dfd51d1124e8d8498dce894923a0d",
           "created_at":"2018-01-30T07:42:01.174",
           "update_at":"2018-01-30T07:42:01.174"
         }
     ],
   "total_count":100
}
`
)
