package v1

import (
	"os"
	"testing"

	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/clients"
	"github.com/opentelekomcloud/gophertelekomcloud/acceptance/tools"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/hosts"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/policies"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/waf-premium/v1/rules"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
)

func TestWafPremiumBlacklistRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	optsP := policies.CreateOpts{
		Name: policyName,
	}

	t.Logf("Attempting to create WAF Premium policy: %s", policyName)
	policy, err := policies.Create(client, optsP)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	hostOpts := hosts.CreateOpts{
		Hostname: tools.RandomString("www.waf-demo.com", 3),
		Proxy:    pointerto.Bool(false),
		PolicyId: policy.ID,
		Server: []hosts.PremiumWafServer{{
			FrontProtocol: "HTTP",
			BackProtocol:  "HTTP",
			Address:       "192.168.1.110",
			Port:          80,
			Type:          "ipv4",
			VpcId:         vpcID,
		}},
		Description: "description",
	}
	t.Logf("Attempting to create WAF Premium host")
	host, err := hosts.Create(client, hostOpts)
	th.AssertNoErr(t, err)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium host: %s", host.ID)
		th.AssertNoErr(t, hosts.Delete(client, host.ID, hosts.DeleteOpts{}))
		t.Logf("Deleted WAF Premium host: %s", host.ID)
	})

	blacklistName := tools.RandomString("waf-black-", 3)
	blacklistOpts := rules.BlacklistCreateOpts{
		Name:        blacklistName,
		Description: "desc",
		Addresses:   "192.168.1.0/24",
		Action:      pointerto.Int(0),
	}
	t.Logf("Attempting to Create WAF Premium blacklist rule")
	blacklist, err := rules.CreateBlacklist(client, policy.ID, blacklistOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, blacklist.Name, blacklistName)
	th.AssertEquals(t, blacklist.Addresses, "192.168.1.0/24")
	th.AssertEquals(t, blacklist.Description, "desc")
	th.AssertEquals(t, *blacklist.Action, 0)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium blacklist rule: %s", blacklist.ID)
		th.AssertNoErr(t, rules.DeleteBlacklistRule(client, policy.ID, blacklist.ID))
		t.Logf("Deleted WAF Premium blacklist rule: %s", blacklist.ID)
	})

	t.Logf("Attempting to List WAF Premium blacklist rule")
	listAntiCrawler, err := rules.ListBlacklists(client, policy.ID, rules.ListBlacklistOpts{})
	th.AssertNoErr(t, err)
	if len(listAntiCrawler) < 1 {
		t.Fatal("empty WAF Premium blacklist rule list")
	}

	t.Logf("Attempting to Update WAF Premium blacklist rule: %s", blacklist.ID)
	updatedBl, err := rules.UpdateBlacklist(client, policy.ID, blacklist.ID, rules.UpdateBlacklistOpts{
		Name:        blacklistName + "-updated",
		Description: "updated",
		Addresses:   "10.1.100.0/24",
		Action:      pointerto.Int(2),
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium blacklist rule: %s", blacklist.ID)
	getBlacklist, err := rules.GetBlacklist(client, policy.ID, updatedBl.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getBlacklist.Addresses, "10.1.100.0/24")
	th.AssertEquals(t, getBlacklist.Description, "updated")
	th.AssertEquals(t, *getBlacklist.Action, 2)
	th.AssertEquals(t, getBlacklist.Name, blacklistName+"-updated")
}

func TestWafPremiumCcRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	ccOpts := rules.CreateCcOpts{
		Mode:        pointerto.Int(0),
		Url:         "/path",
		Description: "desc",
		Action: &rules.CcActionObject{
			Category: "captcha",
		},
		TagType:     "ip",
		LimitNum:    10,
		LimitPeriod: 60,
	}
	t.Logf("Attempting to Create WAF Premium cc rule")
	cc, err := rules.CreateCc(client, policy.ID, ccOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, cc.Mode, 0)
	th.AssertEquals(t, cc.Url, "/path")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium cc rule: %s", cc.ID)
		th.AssertNoErr(t, rules.DeleteCcRule(client, policy.ID, cc.ID))
		t.Logf("Deleted WAF Premium cc rule: %s", cc.ID)
	})

	t.Logf("Attempting to List WAF Premium cc rule")
	listCc, err := rules.ListCcs(client, policy.ID, rules.ListCcOpts{})
	th.AssertNoErr(t, err)
	if len(listCc) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium cc rule: %s", cc.ID)
	updatedCc, err := rules.UpdateCc(client, policy.ID, cc.ID, rules.CreateCcOpts{
		Mode:        pointerto.Int(0),
		Url:         "/path1",
		Description: "updated",
		Action: &rules.CcActionObject{
			Category: "log",
		},
		TagType:     "ip",
		LimitNum:    10,
		LimitPeriod: 60,
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium cc rule: %s", updatedCc.ID)
	getCc, err := rules.GetCc(client, policy.ID, updatedCc.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getCc.Url, "/path1")
	th.AssertEquals(t, getCc.Action.Category, "log")
}

func TestWafPremiumCustomRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	customOpts := rules.CreateCustomOpts{
		Time:        pointerto.Bool(false),
		Description: "desc",
		Conditions: []rules.CustomConditionsObject{{
			Category:       "url",
			LogicOperation: "contain",
			Contents:       []string{"test"},
		}},
		Action: &rules.CustomActionObject{
			Category: "block",
		},
		Priority: pointerto.Int(50),
	}
	t.Logf("Attempting to Create WAF Premium custom rule")
	custom, err := rules.CreateCustom(client, policy.ID, customOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, custom.Priority, 50)
	th.AssertEquals(t, custom.Conditions[0].Category, "url")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium custom rule: %s", custom.ID)
		th.AssertNoErr(t, rules.DeleteCustomRule(client, policy.ID, custom.ID))
		t.Logf("Deleted WAF Premium custom rule: %s", custom.ID)
	})

	t.Logf("Attempting to List WAF Premium custom rule")
	listCc, err := rules.ListCustoms(client, policy.ID, rules.ListCustomOpts{})
	th.AssertNoErr(t, err)
	if len(listCc) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium custom rule: %s", custom.ID)
	updatedCustom, err := rules.UpdateCustom(client, policy.ID, custom.ID, rules.CreateCustomOpts{
		Time:        pointerto.Bool(false),
		Description: "updated",
		Action: &rules.CustomActionObject{
			Category: "pass",
		},
		Conditions: []rules.CustomConditionsObject{{
			Category:       "url",
			LogicOperation: "contain",
			Contents:       []string{"test"},
		}},
		Priority: pointerto.Int(60),
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium custom rule: %s", updatedCustom.ID)
	getCustom, err := rules.GetCustom(client, policy.ID, updatedCustom.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getCustom.Description, "updated")
	th.AssertEquals(t, getCustom.Action.Category, "pass")
}

func TestWafPremiumAntiCrawlerRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	antiCrawlerName := tools.RandomString("waf-atc-", 3)
	antiCrawlerOpts := rules.CreateAntiCrawlerOpts{
		Url:   "/patent/id",
		Logic: 3,
		Name:  antiCrawlerName,
		Type:  "anticrawler_except_url",
	}
	t.Logf("Attempting to Create WAF Premium anti crawler rule")
	antiCrawler, err := rules.CreateAntiCrawler(client, policy.ID, antiCrawlerOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, antiCrawler.Name, antiCrawlerName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium anti crawler rule: %s", antiCrawler.ID)
		th.AssertNoErr(t, rules.DeleteAntiCrawlerRule(client, policy.ID, antiCrawler.ID))
		t.Logf("Deleted WAF Premium anti crawler rule: %s", antiCrawler.ID)
	})

	t.Logf("Attempting to List WAF Premium anti crawler rule")
	listAntiCrawler, err := rules.ListAntiCrawlers(client, policy.ID, rules.ListAntiCrawlerOpts{})
	th.AssertNoErr(t, err)
	if len(listAntiCrawler) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium anti crawler rule: %s", antiCrawler.ID)
	updatedAntiCrawler, err := rules.UpdateAntiCrawler(client, policy.ID, antiCrawler.ID, rules.UpdateAntiCrawlerOpts{
		Url:   "/patents/id",
		Logic: 4,
		Name:  antiCrawlerName + "-updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium anti crawler rule: %s", antiCrawler.ID)
	getAntiCrawler, err := rules.GetAntiCrawler(client, policy.ID, updatedAntiCrawler.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getAntiCrawler.Logic, 4)
	th.AssertEquals(t, getAntiCrawler.Url, "/patents/id")
	th.AssertEquals(t, getAntiCrawler.Name, antiCrawlerName+"-updated")
}

func TestWafPremiumDataMaskingRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	privacyOpts := rules.CreatePrivacyOpts{
		Url:         "/login",
		Category:    "params",
		Name:        "password",
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium privacy rule")
	privacy, err := rules.CreatePrivacy(client, policy.ID, privacyOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, privacy.Category, "params")
	th.AssertEquals(t, privacy.Name, "password")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium privacy rule: %s", privacy.ID)
		th.AssertNoErr(t, rules.DeletePrivacyRule(client, policy.ID, privacy.ID))
		t.Logf("Deleted WAF Premium privacy rule: %s", privacy.ID)
	})

	t.Logf("Attempting to List WAF Premium privacy rule")
	listCc, err := rules.ListPrivacy(client, policy.ID, rules.ListPrivacyOpts{})
	th.AssertNoErr(t, err)
	if len(listCc) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium privacy rule: %s", policy.ID)
	updatedPrivacy, err := rules.UpdatePrivacy(client, policy.ID, privacy.ID, rules.UpdatePrivacyOpts{
		Url:         "/path1",
		Category:    "header",
		Name:        "token",
		Description: "updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium privacy rule: %s", updatedPrivacy.ID)
	getPrivacy, err := rules.GetPrivacy(client, policy.ID, updatedPrivacy.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getPrivacy.Name, "token")
	th.AssertEquals(t, getPrivacy.Description, "updated")
	th.AssertEquals(t, getPrivacy.Category, "header")
}

func TestWafPremiumKnownAttackRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})

	kasOpts := rules.CreateKnownAttackSourceOpts{
		BlockTime:   pointerto.Int(300),
		Category:    "long_ip_block",
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium known attack source rule")
	kas, err := rules.CreateKnownAttackSource(client, policy.ID, kasOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, kas.Category, "long_ip_block")
	th.AssertEquals(t, kas.Description, "desc")
	th.AssertEquals(t, kas.BlockTime, 300)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium known attack source rule: %s", kas.ID)
		th.AssertNoErr(t, rules.DeleteKnownAttackSourceRule(client, policy.ID, kas.ID))
		t.Logf("Deleted WAF Premium known attack source rule: %s", kas.ID)
	})

	t.Logf("Attempting to List WAF Premium known attack source rule")
	listKas, err := rules.ListKnownAttackSource(client, policy.ID, rules.ListKnownAttackSourceOpts{})
	th.AssertNoErr(t, err)
	if len(listKas) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium known attack source rule: %s", policy.ID)
	updatedPrivacy, err := rules.UpdateKnownAttackSource(client, policy.ID, kas.ID, rules.UpdateKnownAttackSourceOpts{
		Description: "updated",
		BlockTime:   pointerto.Int(1200),
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium known attack source rule: %s", updatedPrivacy.ID)
	getKas, err := rules.GetKnownAttackSource(client, policy.ID, updatedPrivacy.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getKas.Description, "updated")
	th.AssertEquals(t, getKas.BlockTime, 1200)
}

func TestWafPremiumWebTamperRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})
	atOpts := rules.CreateAntiTamperOpts{
		Hostname:    "www.domain.com",
		Url:         "/login",
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium anti tamper rule")
	at, err := rules.CreateAntiTamper(client, policy.ID, atOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, at.Hostname, "www.domain.com")
	th.AssertEquals(t, at.Description, "desc")
	th.AssertEquals(t, at.Url, "/login")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium anti tamper rule: %s", at.ID)
		th.AssertNoErr(t, rules.DeleteAntiTamperRule(client, policy.ID, at.ID))
		t.Logf("Deleted WAF Premium anti tamper rule: %s", at.ID)
	})

	t.Logf("Attempting to List WAF Premium anti tamper rule")
	listAt, err := rules.ListAntiTamper(client, policy.ID, rules.ListAntiTamperOpts{})
	th.AssertNoErr(t, err)
	if len(listAt) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium anti tamper rule: %s", policy.ID)
	_, err = rules.UpdateAntiTamperCache(client, policy.ID, at.ID)
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium anti tamper rule: %s", at.ID)
	getAt, err := rules.GetAntiTamper(client, policy.ID, at.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getAt.Description, "desc")
	th.AssertEquals(t, *getAt.Status, 1)
}

func TestWafPremiumInformationLeakageProtectionRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})
	ilpOpts := rules.CreateAntiLeakageOpts{
		Url:         "/attack",
		Category:    "sensitive",
		Contents:    []string{"id_card"},
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium information leakage protection rule")
	ilp, err := rules.CreateAntiLeakage(client, policy.ID, ilpOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ilp.Contents[0], "id_card")
	th.AssertEquals(t, ilp.Description, "desc")
	th.AssertEquals(t, ilp.Url, "/attack")
	th.AssertEquals(t, ilp.Category, "sensitive")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium information leakage protection rule: %s", ilp.ID)
		th.AssertNoErr(t, rules.DeleteAntiLeakageRule(client, policy.ID, ilp.ID))
		t.Logf("Deleted WAF Premium information leakage protection rule: %s", ilp.ID)
	})

	t.Logf("Attempting to List WAF Premium information leakage protection rule")
	listIlp, err := rules.ListAntiLeakage(client, policy.ID, rules.ListAntiLeakageOpts{})
	th.AssertNoErr(t, err)
	if len(listIlp) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium information leakage protection rule: %s", policy.ID)
	ilpUpdated, err := rules.UpdateAntiLeakage(client, policy.ID, ilp.ID, rules.UpdateAntiLeakageOpts{
		Url:         "/pass",
		Category:    "sensitive",
		Contents:    []string{"id_card"},
		Description: "updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium information leakage protection rule: %s", ilp.ID)
	getIlp, err := rules.GetAntiLeakage(client, policy.ID, ilpUpdated.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getIlp.Description, "updated")
	th.AssertEquals(t, getIlp.Url, "/pass")
	th.AssertEquals(t, getIlp.Contents[0], "id_card")
}

func TestWafPremiumAlarmMaskingRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})
	ignoreOpts := rules.CreateIgnoreOpts{
		Domains: []string{"www.example.com"},
		Conditions: []rules.IgnoreCondition{{
			Category:       "url",
			Contents:       []string{"/login"},
			LogicOperation: "equal",
		}},
		Mode:        1,
		Rule:        "all",
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium alarm masking rule")
	ignore, err := rules.CreateIgnore(client, policy.ID, ignoreOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, ignore.Description, "desc")

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium alarm masking rule: %s", ignore.ID)
		th.AssertNoErr(t, rules.DeleteIgnoreRule(client, policy.ID, ignore.ID))
		t.Logf("Deleted WAF Premium information leakage protection rule: %s", ignore.ID)
	})

	t.Logf("Attempting to List WAF Premium alarm masking rule")
	listIgnore, err := rules.ListIgnore(client, policy.ID, rules.ListIgnoreOpts{})
	th.AssertNoErr(t, err)
	if len(listIgnore) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium alarm masking rule: %s", policy.ID)
	ignoreUpdated, err := rules.UpdateIgnore(client, policy.ID, ignore.ID, rules.CreateIgnoreOpts{
		Domains: []string{"www.example1.com"},
		Conditions: []rules.IgnoreCondition{{
			Category:       "ip",
			Contents:       []string{"192.168.1.1"},
			LogicOperation: "equal",
		}},
		Mode:        1,
		Rule:        "all",
		Description: "updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium alarm masking rule: %s", ignore.ID)
	getIgnore, err := rules.GetIgnore(client, policy.ID, ignoreUpdated.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getIgnore.Description, "updated")
}

func TestWafPremiumGeoIpRuleWorkflow(t *testing.T) {
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)

	policyName := tools.RandomString("waf-policy-", 3)
	opts := policies.CreateOpts{
		Name: policyName,
	}
	policy, err := policies.Create(client, opts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, policy.Name, policyName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium policy: %s", policy.ID)
		th.AssertNoErr(t, policies.Delete(client, policy.ID))
		t.Logf("Deleted WAF Premium policy: %s", policy.ID)
	})
	geoIpOpts := rules.CreateGeoIpOpts{
		GeoIp:       "BR",
		Action:      pointerto.Int(0),
		Name:        "Test",
		Description: "desc",
	}
	t.Logf("Attempting to Create WAF Premium geo ip rule")
	geoIP, err := rules.CreateGeoIp(client, policy.ID, geoIpOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, geoIP.Description, "desc")
	th.AssertEquals(t, geoIP.GeoIp, "BR")
	th.AssertEquals(t, geoIP.Name, "Test")
	th.AssertEquals(t, geoIP.Action, 0)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium geo ip rule: %s", geoIP.ID)
		th.AssertNoErr(t, rules.DeleteGeoIpRule(client, policy.ID, geoIP.ID))
		t.Logf("Deleted WAF Premium geo ip rule: %s", geoIP.ID)
	})

	t.Logf("Attempting to List WAF Premium geo ip rule")
	listGeo, err := rules.ListGeoIp(client, policy.ID, rules.ListGeoIpOpts{})
	th.AssertNoErr(t, err)
	if len(listGeo) < 1 {
		t.Fatal("empty WAF Premium rule list")
	}

	t.Logf("Attempting to Update WAF Premium geo ip rule: %s", policy.ID)
	geoUpdated, err := rules.UpdateGeoIp(client, policy.ID, geoIP.ID, rules.UpdateGeoIpOpts{
		GeoIp:       "DE",
		Action:      pointerto.Int(1),
		Name:        "Updated",
		Description: "updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium geo ip rule: %s", geoIP.ID)
	getGeo, err := rules.GetGeoIp(client, policy.ID, geoUpdated.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getGeo.Description, "updated")
	th.AssertEquals(t, getGeo.GeoIp, "DE")
	th.AssertEquals(t, getGeo.Action, 1)
	th.AssertEquals(t, getGeo.Name, "Updated")
}

func TestWafPremiumRefTableWorkflow(t *testing.T) {
	t.Skip("Deletion not working")
	region := os.Getenv("OS_REGION_NAME")
	vpcID := os.Getenv("OS_VPC_ID")
	if vpcID == "" && region == "" {
		t.Skip("OS_REGION_NAME, OS_VPC_ID env vars is required for this test")
	}

	client, err := clients.NewWafdV1Client()
	th.AssertNoErr(t, err)
	tableName := tools.RandomString("waf-table-", 3)
	refTableOpts := rules.CreateReferenceTableOpts{
		Name:   tableName,
		Type:   "url",
		Values: []string{"/demo"},
	}
	t.Logf("Attempting to Create WAF Premium ref table")
	table, err := rules.CreateReferenceTable(client, refTableOpts)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, table.Name, tableName)

	t.Cleanup(func() {
		t.Logf("Attempting to delete WAF Premium ref table: %s", table.ID)
		th.AssertNoErr(t, rules.DeleteReferenceTable(client, table.ID))
		t.Logf("Deleted WAF Premium ref table: %s", table.ID)
	})

	t.Logf("Attempting to List WAF Premium ref table")
	listTable, err := rules.ListReferenceTable(client, rules.ListReferenceTableOpts{})
	th.AssertNoErr(t, err)
	if len(listTable) < 1 {
		t.Fatal("empty WAF Premium ref table list")
	}

	t.Logf("Attempting to Update WAF Premium ref table: %s", table.ID)
	tableUpdated, err := rules.UpdateReferenceTable(client, table.ID, rules.UpdateReferenceTableOpts{
		Name:        tableName + "-updated",
		Type:        "url",
		Values:      []string{"/demo"},
		Description: "updated",
	})
	th.AssertNoErr(t, err)

	t.Logf("Attempting to Get WAF Premium ref table: %s", table.ID)
	getTable, err := rules.GetReferenceTable(client, tableUpdated.ID)
	th.AssertNoErr(t, err)
	th.AssertEquals(t, getTable.Description, "updated")
	th.AssertEquals(t, getTable.Name, tableName+"-updated")
}
