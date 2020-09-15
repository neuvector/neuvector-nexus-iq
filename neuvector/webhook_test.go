package neuvector

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	registryScanReportWebhookBody = `{
  "text": "*Audit: CRITICAL level*\n_Registry.Scan.Report: high 30 medium 34_\n>>> notification=audit,name=Registry.Scan.Report,level=Critical,reported_timestamp=1593010960,reported_at=2020-06-24T15:02:40Z,cluster_name=cluster.local,response_rule_id=7,host_id=,host_name=,enforcer_id=,enforcer_name=,image_id=a3590c0e9ff9eca5b4863e13b4271435476104701740b3f7470dffdd481b0a60,registry=https://registry.hub.docker.com/,repository=library/debian,tag=jessie-20200607,base_os=debian:8,high_vul_cnt=30,medium_vul_cnt=34,high_vuls=[CVE-2017-16997 CVE-2014-9761 CVE-2017-1000408 CVE-2018-1000001 CVE-2018-6485 CVE-2019-9169 CVE-2017-1000408 CVE-2019-9169 CVE-2017-16997 CVE-2018-6485 CVE-2014-9761 CVE-2018-1000001 CVE-2019-9169 CVE-2017-16997 CVE-2017-1000408 CVE-2014-9761 CVE-2018-1000001 CVE-2018-6485 CVE-2020-10543 CVE-2018-6797 CVE-2020-10878 CVE-2017-12424 CVE-2017-12424 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779],medium_vuls=[CVE-2018-12886 CVE-2015-5276 CVE-2018-12886 CVE-2018-12886 CVE-2015-5276 CVE-2015-5276 CVE-2018-12886 CVE-2016-10739 CVE-2017-12132 CVE-2017-1000409 CVE-2020-1751 CVE-2017-12133 CVE-2016-10739 CVE-2017-12132 CVE-2017-12133 CVE-2020-1751 CVE-2017-1000409 CVE-2017-12133 CVE-2020-1751 CVE-2017-1000409 CVE-2016-10739 CVE-2017-12132 CVE-2017-10790 CVE-2017-7244 CVE-2015-3217 CVE-2017-7186 CVE-2020-12723 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011],cvedb_version=1.924,message=,user=,error=,aggregation_from=0,count=0,platform=,platform_version=",
  "username": "NeuVector - cluster.local"
}`

	containerScanReportWebhookBody = `{
  "text": "*Audit: CRITICAL level*\n_Container.Scan.Report: high 30 medium 33_\n>>> notification=audit,name=Container.Scan.Report,level=Critical,reported_timestamp=1593423866,reported_at=2020-06-29T09:44:26Z,cluster_name=cluster.local,response_rule_id=8,host_id=docker-desktop:TKIG:AHYH:HP2Y:PGG3:LM5P:LMA3:J4E4:S66R:FX2E:ZAOZ:FZCZ:GJGK,host_name=docker-desktop,enforcer_id=5c203bb7c1dbbf95a2925cc8392cecbbc37c246ea9eab206eeea816e5fd1b0db,enforcer_name=neuvector-nexus-iq-integration_neuvector_1,workload_id=2e51c3c15775aeec19071bbe4a8a562e61faf1c4ee8058905709b2eb88480157,workload_name=recursing_goldstine,workload_image=debian:jessie-20200607-slim,workload_service=debian,base_os=debian:8,high_vul_cnt=30,medium_vul_cnt=33,high_vuls=[CVE-2018-1000001 CVE-2018-6485 CVE-2017-1000408 CVE-2017-16997 CVE-2019-9169 CVE-2014-9761 CVE-2014-9761 CVE-2018-6485 CVE-2017-16997 CVE-2019-9169 CVE-2018-1000001 CVE-2017-1000408 CVE-2017-1000408 CVE-2018-1000001 CVE-2019-9169 CVE-2018-6485 CVE-2017-16997 CVE-2014-9761 CVE-2018-6797 CVE-2020-10878 CVE-2020-10543 CVE-2017-12424 CVE-2017-12424 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779 CVE-2016-2779],medium_vuls=[CVE-2018-12886 CVE-2015-5276 CVE-2018-12886 CVE-2015-5276 CVE-2018-12886 CVE-2018-12886 CVE-2015-5276 CVE-2017-12132 CVE-2017-1000409 CVE-2017-12133 CVE-2020-1751 CVE-2016-10739 CVE-2017-1000409 CVE-2017-12132 CVE-2017-12133 CVE-2016-10739 CVE-2020-1751 CVE-2017-12132 CVE-2020-1751 CVE-2016-10739 CVE-2017-1000409 CVE-2017-12133 CVE-2015-3217 CVE-2017-7244 CVE-2017-7186 CVE-2020-12723 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011 CVE-2016-5011],cvedb_version=1.924,message=,user=,error=,aggregation_from=0,count=0,platform=,platform_version=",
  "username": "NeuVector - cluster.local"
}`
)

func createWebhookRequest(body string) (*WebhookRequest, error) {
	var wr WebhookRequest
	err := json.Unmarshal([]byte(body), &wr)
	if err != nil {
		return nil, err

	}
	return &wr, nil
}

func createRegistryScanReportWebhookRequest() (*WebhookRequest, error) {
	return createWebhookRequest(registryScanReportWebhookBody)
}

func createContainerScanReportWebhookRequest() (*WebhookRequest, error) {
	return createWebhookRequest(containerScanReportWebhookBody)
}

func TestWebhookRequest_ToWebhookData(t *testing.T) {
	wr, err := createRegistryScanReportWebhookRequest()
	if err != nil {
		t.Error(err)
	}

	wd, err := wr.ToWebhookData()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Registry.Scan.Report", wd.Event)
	assert.Equal(t, "Critical", wd.Level)

	assert.Len(t, wd.Values, 28)
	assert.Contains(t, wd.Values, "name")
	assert.Contains(t, wd.Values, "level")
	assert.Contains(t, wd.Values, "image_id")
}

func TestWebhookData_ToWebhookRegistryScanReportData(t *testing.T) {
	wr, err := createRegistryScanReportWebhookRequest()
	if err != nil {
		t.Error(err)
	}

	wd, err := wr.ToWebhookData()
	if err != nil {
		t.Error(err)
	}

	rsrd, err := wd.ToWebhookRegistryScanReportData()
	if err != nil {
		t.Error(err)
	}

	expRsrd := &WebhookRegistryScanReportData{
		Registry: "https://registry.hub.docker.com/",
		ImageId:  "a3590c0e9ff9eca5b4863e13b4271435476104701740b3f7470dffdd481b0a60",
	}

	assert.EqualValues(t, expRsrd, rsrd)
}

func TestWebhookData_ToWebhookContainerScanReportData(t *testing.T) {
	wr, err := createContainerScanReportWebhookRequest()
	if err != nil {
		t.Error(err)
	}

	wd, err := wr.ToWebhookData()
	if err != nil {
		t.Error(err)
	}

	csrd, err := wd.ToWebhookContainerScanReportData()
	if err != nil {
		t.Error(err)
	}

	expCsrd := &WebhookContainerScanReportData{
		WorkloadId:      "2e51c3c15775aeec19071bbe4a8a562e61faf1c4ee8058905709b2eb88480157",
		WorkloadName:    "recursing_goldstine",
		WorkloadImage:   "debian:jessie-20200607-slim",
		WorkloadService: "debian",
	}

	assert.EqualValues(t, expCsrd, csrd)
}
