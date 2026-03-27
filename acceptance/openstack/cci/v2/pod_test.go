package v2

import (
	"strconv"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	ns "github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/namespace"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/cci/v2/pod"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestPodLifecycle(t *testing.T) {
	t.Skip("Tenant not whitelisted to run CCI")
	client, err := clients.NewCCIClient()
	th.AssertNoErr(t, err)

	nsName := "cci-namespace-" + strconv.Itoa(tools.RandomInt(1, 1000))
	podName := "test-cm-" + strconv.Itoa(tools.RandomInt(1, 1000))

	createNsOpts := ns.CreateOpts{
		APIVersion: "cci/v2",
		Kind:       "Namespace",
		Metadata: ns.Metadata{
			Name: nsName,
		},
	}

	t.Logf("Attempting to create namespace")
	_, err = ns.Create(client, createNsOpts)
	th.AssertNoErr(t, err)

	err = waitForStatusActive(client, 600, nsName)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete namespace")
		nsDeleteOpts := ns.DeleteOpts{
			Name: nsName,
		}
		_, err = ns.Delete(client, nsDeleteOpts)
		th.AssertNoErr(t, err)
		err = waitForStatusDeleted(client, 500, nsName)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to create pod")
	createOpts := pod.CreateOpts{
		Pod: pod.Pod{
			APIVersion: "cci/v2",
			Kind:       "Pod",
			Metadata: pod.ObjectMeta{
				Name: podName,
				Annotations: map[string]string{
					"resource.cci.io/pod-size-specs": "2.00_4.0",
				},
			},
			Spec: pod.PodSpec{
				Containers: []pod.Container{
					{
						Env: []pod.EnvVar{
							{
								Name:  "ENV1",
								Value: "false",
							},
						},
						Image:           "nginx:latest",
						ImagePullPolicy: "IfNotPresent",
						Name:            "deploy-example",
						Resources: pod.ResourceRequirements{
							Limits: map[string]string{
								"cpu":    "500m",
								"memory": "1Gi",
							},
							Requests: map[string]string{
								"cpu":    "500m",
								"memory": "1Gi",
							},
						},
						TerminationMessagePolicy: "File",
						TerminationMessagePath:   "/dev/termination-log",
					},
				},
				DNSPolicy: "Default",
				ImagePullSecrets: []pod.LocalObjectReference{
					{
						Name: "image-pull-secret",
					},
				},
				RestartPolicy:                 "Always",
				TerminationGracePeriodSeconds: pointerto.Int64(30),
			},
		},
	}

	newPod, err := pod.Create(client, nsName, createOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, podName, newPod.Metadata.Name)

	t.Cleanup(func() {
		t.Logf("Attempting to delete pod")
		deleteOpts := pod.DeleteOpts{
			NameSpace: nsName,
			PodName:   podName,
		}
		_, err := pod.Delete(client, deleteOpts)
		th.AssertNoErr(t, err)
	})

	t.Logf("Attempting to get pod")
	getPod, err := pod.Get(client, nsName, podName)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, createOpts.Spec.Containers[0].Env[0].Name, getPod.Spec.Containers[0].Env[0].Name)
	th.AssertEquals(t, createOpts.Spec.Containers[0].Image, getPod.Spec.Containers[0].Image)

	t.Logf("Attempting to list pods")
	pods, err := pod.List(client, nsName, pod.ListOpts{})
	th.AssertNoErr(t, err)
	found := false
	for _, p := range pods {
		if p.Metadata.Name == podName {
			found = true
			break
		}
	}
	th.AssertEquals(t, true, found)

	t.Logf("Attempting to update pod")
	latestPod, err := pod.Get(client, nsName, podName)
	th.AssertNoErr(t, err)

	updatedLabels := latestPod.Metadata.Labels
	if updatedLabels == nil {
		updatedLabels = map[string]string{}
	}
	updatedLabels["updated"] = "true"

	updateOpts := pod.UpdateOpts{
		Pod: pod.Pod{
			APIVersion: "cci/v2",
			Kind:       "Pod",
			Metadata: pod.ObjectMeta{
				Name:            podName,
				Annotations:     latestPod.Metadata.Annotations,
				Labels:          updatedLabels,
				ResourceVersion: latestPod.Metadata.ResourceVersion,
			},
			Spec: latestPod.Spec,
		},
	}

	updatedPod, err := pod.Update(client, nsName, podName, updateOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, "true", updatedPod.Metadata.Labels["updated"])

	t.Run("Patch", func(t *testing.T) {
		t.Skip("PATCH endpoint not authorized for current IAM user (APIGW.0302)")
		patchBody := map[string]interface{}{
			"metadata": map[string]interface{}{
				"labels": map[string]string{
					"patched": "true",
				},
			},
		}
		patchOpts := pod.PatchOpts{
			Namespace: nsName,
			Name:      podName,
		}
		patchedPod, err := pod.Patch(client, patchOpts, patchBody, "application/merge-patch+json")
		th.AssertNoErr(t, err)
		th.AssertEquals(t, "true", patchedPod.Metadata.Labels["patched"])
	})

	t.Run("Log", func(t *testing.T) {
		t.Skip("Log endpoint not published in the environment")
		logOpts := pod.LogOpts{
			Namespace: nsName,
			Name:      podName,
			TailLines: pointerto.Int(10),
		}
		_, err := pod.Log(client, logOpts)
		th.AssertNoErr(t, err)
	})

	t.Run("ConnectPost", func(t *testing.T) {
		t.Skip("Exec endpoint not authorized for current IAM user (APIGW.0302)")
		connectPostOpts := pod.ConnectPostOpts{
			Command:   "ls",
			Container: "deploy-example",
			Stdout:    pointerto.Bool(true),
			Stderr:    pointerto.Bool(true),
		}
		err := pod.ConnectPost(client, nsName, podName, connectPostOpts)
		th.AssertNoErr(t, err)
	})

	t.Run("ConnectGet", func(t *testing.T) {
		t.Skip("Exec endpoint not authorized for current IAM user (APIGW.0302)")
		connectGetOpts := pod.ConnectGetOpts{
			Command:   "ls",
			Container: "deploy-example",
			Stdout:    pointerto.Bool(true),
			Stderr:    pointerto.Bool(true),
		}
		err := pod.ConnectGet(client, nsName, podName, connectGetOpts)
		th.AssertNoErr(t, err)
	})
}
