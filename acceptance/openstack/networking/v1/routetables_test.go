package v1

import (
	"log"
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ecs/v1/cloudservers"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/routetables"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/subnets"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v1/vpcs"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/networking/v2/peerings"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestRouteTablesList(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	listOpts := routetables.ListOpts{}
	routeTablesList, err := routetables.List(client, listOpts)
	th.AssertNoErr(t, err)

	for _, rtb := range routeTablesList {
		tools.PrintResource(t, rtb)
	}
}

func TestRouteTablesLifecycle(t *testing.T) {
	client, err := clients.NewNetworkV1Client()
	th.AssertNoErr(t, err)

	clientV2, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create first vpc
	createOpts1 := vpcs.CreateOpts{
		Name: tools.RandomString("acc-vpc-1-", 3),
		CIDR: "192.168.0.0/16",
	}

	t.Logf("Attempting to create first vpc: %s", createOpts1.Name)

	vpc1, err := vpcs.Create(client, createOpts1).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created first vpc: %s", vpc1.ID)

	createSubnetOpts1 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-1-1-", 3),
		Description: "some description 1",
		CIDR:        "192.168.0.0/24",
		GatewayIP:   "192.168.0.1",
		VpcID:       vpc1.ID,
	}
	t.Logf("Attempting to create first subnet: %s", createSubnetOpts1.Name)

	subnet1, err := subnets.Create(client, createSubnetOpts1).Extract()
	th.AssertNoErr(t, err)

	createSubnetOpts2 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-2-1-", 3),
		Description: "some description 2",
		CIDR:        "192.168.10.0/24",
		GatewayIP:   "192.168.10.1",
		VpcID:       vpc1.ID,
	}
	t.Logf("Attempting to create second subnet: %s", createSubnetOpts2.Name)

	subnet2, err := subnets.Create(client, createSubnetOpts2).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		deleteSubnet(t, client, subnet1.VpcID, subnet1.ID)
		deleteSubnet(t, client, subnet2.VpcID, subnet2.ID)
		deleteVpc(t, client, vpc1.ID)
	})

	// Create second vpc
	createOpts2 := vpcs.CreateOpts{
		Name: tools.RandomString("acc-vpc-2-", 3),
		CIDR: "172.16.0.0/16",
	}

	t.Logf("Attempting to create first vpc: %s", createOpts2.Name)

	vpc2, err := vpcs.Create(client, createOpts2).Extract()
	th.AssertNoErr(t, err)
	t.Logf("Created second vpc: %s", vpc2.ID)

	createSubnetOpts3 := subnets.CreateOpts{
		Name:        tools.RandomString("acc-subnet-1-2-", 3),
		Description: "some description 3",
		CIDR:        "172.16.10.0/24",
		GatewayIP:   "172.16.10.1",
		VpcID:       vpc2.ID,
	}
	t.Logf("Attempting to create second subnet: %s", createSubnetOpts3.Name)

	subnet3, err := subnets.Create(client, createSubnetOpts3).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		deleteSubnet(t, client, subnet3.VpcID, subnet3.ID)
		deleteVpc(t, client, vpc2.ID)
	})

	createOpts := peerings.CreateOpts{
		Name: tools.RandomString("acc-peering-", 3),
		RequestVpcInfo: peerings.VpcInfo{
			VpcId: vpc1.ID,
		},
		AcceptVpcInfo: peerings.VpcInfo{
			VpcId: vpc2.ID,
		},
	}
	peering, err := peerings.Create(clientV2, createOpts).Extract()
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete vpc peering connection: %s", peering.ID)
		err := peerings.Delete(clientV2, peering.ID).ExtractErr()
		if err != nil {
			t.Fatalf("Error deleting vpc peering connection: %v", err)
		}
	})

	createRtbOpts := routetables.CreateOpts{
		Name:        tools.RandomString("acc-rtb-", 3),
		Description: "route table",
		VpcID:       vpc1.ID,
	}
	rtb, err := routetables.Create(client, createRtbOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		err = routetables.Delete(client, rtb.ID)
		th.AssertNoErr(t, err)
	})

	getRtb, err := routetables.Get(client, rtb.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getRtb.Description, "route table")

	t.Logf("Attempting to associate subnets to vpc route table: %s", rtb.ID)
	actionOpts := routetables.ActionOpts{
		Subnets: routetables.ActionSubnetsOpts{
			Associate: []string{subnet1.ID, subnet2.ID},
		},
	}
	associate, err := routetables.Action(client, rtb.ID, actionOpts)
	th.AssertNoErr(t, err)
	if len(associate.Subnets) < 1 {
		t.Fatalf("Number of associated subnet lower than expected")
	}

	log.Printf("Attempting to update VPC route table with peering route: %#v", rtb.ID)
	routesOpts := map[string][]routetables.RouteOpts{}
	updateOpts := routetables.UpdateOpts{
		Name:        "First",
		Description: pointerto.String("description"),
	}
	addRouteOpts := []routetables.RouteOpts{
		{
			Destination: "172.16.0.0/16",
			Type:        "peering",
			NextHop:     peering.ID,
			Description: pointerto.String("route 1"),
		},
	}
	routesOpts["add"] = addRouteOpts
	updateOpts.Routes = routesOpts
	err = routetables.Update(client, rtb.ID, updateOpts)
	th.AssertNoErr(t, err)

	// ECS
	imageId := os.Getenv("OS_IMAGE_ID")
	flavorId := os.Getenv("OS_FLAVOR_ID")
	az := os.Getenv("OS_AVAILABILITY_ZONE")
	if imageId != "" && flavorId != "" && az != "" {
		clientCompute, err := clients.NewComputeV1Client()
		th.AssertNoErr(t, err)
		createEcsOpts := cloudservers.CreateOpts{
			ImageRef:  imageId,
			FlavorRef: flavorId,
			Name:      tools.RandomString("acc-ecs-", 3),
			VpcId:     vpc1.ID,
			Nics: []cloudservers.Nic{
				{
					SubnetId: subnet1.NetworkID,
				},
			},
			RootVolume: cloudservers.RootVolume{
				VolumeType: "SSD",
			},
			DataVolumes: []cloudservers.DataVolume{
				{
					VolumeType: "SSD",
					Size:       40,
				},
			},
			AvailabilityZone: az,
		}
		ecs := openstack.CreateCloudServer(t, clientCompute, createEcsOpts)
		t.Cleanup(func() {
			openstack.DeleteCloudServer(t, clientCompute, ecs.ID)
		})

		log.Printf("Attempting to update VPC route table with ecs route: %#v", rtb.ID)
		routesEcsOpts := map[string][]routetables.RouteOpts{}
		updateEcsOpts := routetables.UpdateOpts{
			Name:        "First",
			Description: pointerto.String("description"),
		}
		addRouteEcsOpts := []routetables.RouteOpts{
			{
				Destination: "192.168.0.0/24",
				Type:        "ecs",
				NextHop:     ecs.ID,
				Description: pointerto.String("route 2"),
			},
		}
		routesEcsOpts["mod"] = addRouteEcsOpts
		updateEcsOpts.Routes = routesEcsOpts
		err = routetables.Update(client, rtb.ID, updateEcsOpts)
		th.AssertNoErr(t, err)
	}
	getChRtb, err := routetables.Get(client, rtb.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getChRtb.Description, "description")

	t.Logf("Attempting to disassociate subnets to vpc route table: %s", rtb.ID)
	actionDisOpts := routetables.ActionOpts{
		Subnets: routetables.ActionSubnetsOpts{
			Disassociate: []string{subnet1.ID, subnet2.ID},
		},
	}
	disassociate, err := routetables.Action(client, rtb.ID, actionDisOpts)
	th.AssertNoErr(t, err)
	if len(disassociate.Subnets) > 1 {
		t.Fatalf("Number of associated subnet higher than expected")
	}
}
