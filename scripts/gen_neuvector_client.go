package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
	"github.com/jessevdk/go-flags"
)

func main() {
	log.SetOutput(os.Stdout)

	err := genNeuVectorClient()
	if err != nil {
		log.Fatalln(err)
	}
}

func genNeuVectorClient() error {
	tags := []string{
		"Authentication",
		"Scan",
		"EULA",
		"System",
		"Response Rule",
	}

	models := []string{
		"RESTAuthData",
		"RESTAuthPassword",
		"RESTAuthToken",
		"RESTToken",
		"RESTTokenData",
		"RESTError",
		"RESTScanConfig",
		"RESTScanConfigData",
		"RESTScanReport",
		"RESTScanReportData",
		"RESTScanSummary",
		"RESTVulnerability",
		"RESTScanPlatformSummary",
		"RESTRegistryImageSummary",
		"RESTRegistryImageSummaryData",
		"RESTScanLayersReport",
		"RESTScanLayersReportData",
		"RESTScanImageSummary",
		"RESTScanImageSummaryData",
		"RESTScanPlatformSummary",
		"RESTScanPlatformSummaryData",
		"RESTRegistrySummary",
		"RESTRegistrySummaryData",
		"RESTScanStatus",
		"RESTScanStatusData",
		"RESTScanSchedule",
		"RESTScanBrief",
		"RESTScanLayer",
		"RESTScanModule",
		"RESTJfrogXray",
		"RESTJfrogXrayConfig",
		"RESTGCRKey",
		"RESTGCRKeyConfig",
		"RESTAWSAccountKey",
		"RESTAWSAccountKeyConfig",
		"RESTScanner",
		"RESTScannerData",
		"RESTRegistrySummaryListData",
		"RESTRegistryConfig",
		"RESTRegistryConfigData",
		"RESTScanPackageReqData",
		"RESTScanAppPackage",
		"RESTScanPkgReport",
		"RESTScanPkgReportData",
		"RESTScanRepoReq",
		"RESTScanRepoReqData",
		"RESTScanMeta",
		"RESTScanRepoReport",
		"RESTScanRepoReportData",
		// EULA
		"RESTEULA",
		"RESTEULAData",
		// System
		"RESTLicenseInfo",
		"RESTLicenseKey",
		"RESTLicenseShow",
		"RESTLicenseRequest",
		"RESTLicenseRequestData",
		"RESTSystemConfig",
		"RESTSystemConfigData",
		"RESTLicenseShowData",
		"RESTProxy",
		"RESTSystemSummary",
		"RESTSystemSummaryData",
		"RESTSystemRequest",
		"RESTSystemRequestData",
		"RESTSystemConfigConfig",
		"RESTSystemConfigConfigData",
		"RESTUnquarReq",
		// Response rule
		"RESTResponseRuleData",
		"RESTResponseRule",
		"RESTCLUSEventCondition",
		"RESTResponseRule",
		"RESTResponseRuleInsert",
		"RESTResponseRuleActionData",
		"RESTResponseRule",
		"RESTResponseRulesData",
		"RESTResponseRuleConfig",
		"RESTResponseRuleConfigData",
	}

	args := []string{
		"--spec=./schema/nv-api-3.2.2-custom.yml",
		"--config-file=./schema/swagger_client.yml",
		"--target=.",
		"--model-package=models",
		"--client-package=client",
		"--additional-initialism=REST",
	}

	for _, tag := range tags {
		args = append(args, fmt.Sprintf("--tags=\"%s\"", tag))
	}

	for _, model := range models {
		args = append(args, fmt.Sprintf("--model=\"%s\"", model))
	}

	c := &generate.Client{}

	_, err := flags.ParseArgs(c, args)
	if err != nil {
		return err
	}

	err = c.Execute([]string{})
	if err != nil {
		return err
	}

	return nil
}
