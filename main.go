package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	v1alpha1 "github.com/project-copacetic/copacetic/pkg/types/v1alpha1"
)

type FakeParser struct{}

// parseFakeReport parses a fake report from a file
func parseFakeReport(file string) (*FakeReport, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var fake FakeReport
	if err = json.Unmarshal(data, &fake); err != nil {
		return nil, err
	}

	return &fake, nil
}

func newFakeParser() *FakeParser {
	return &FakeParser{}
}

func (k *FakeParser) parse(file string) (*v1alpha1.UpdateManifest, error) {
	// Parse the fake report
	report, err := parseFakeReport(file)
	if err != nil {
		return nil, err
	}

	// Create the standardized report
	updates := v1alpha1.UpdateManifest{
		APIVersion: v1alpha1.APIVersion,
		Metadata: v1alpha1.Metadata{
			OS: v1alpha1.OS{
				Type: report.OSType,
				Version: report.OSVersion,
			},
			Config: v1alpha1.Config{
				Arch: report.Arch,
			},
		},
	}

	// Convert the fake report to the standardized report
	for i := range report.Metadata.Basic_plan_component_vulnerability_fixes {
		pkgs := &report.Metadata.Basic_plan_component_vulnerability_fixes[i]
		if pkgs.Target_component_purl != "" {
			var splitCurrentComponentPurl []string
			var currentPackageWithoutArch string
			var currentPackageWithVersion []string
			var installedVersion string

			var splitTargetComponentPurl []string
			var packageNameWithVersion string
			var splitPackageNameWithVersion []string
			var targetPackageName string
			var fixedVersion string

			if strings.Contains(pkgs.Current_component_purl, "github") {
				splitCurrentComponentPurl = strings.Split(pkgs.Current_component_purl, "/")
				currentPackageWithoutArch = strings.Split(splitCurrentComponentPurl[3], "?")[0]
				currentPackageWithVersion = strings.Split(currentPackageWithoutArch, "@")
				installedVersion = currentPackageWithVersion[1]

				splitTargetComponentPurl = strings.Split(pkgs.Target_component_purl, "/")
				packageNameWithVersion = splitTargetComponentPurl[3]
				splitPackageNameWithVersion = strings.Split(packageNameWithVersion, "@")
				targetPackageName = splitPackageNameWithVersion[0]
				fixedVersion = splitPackageNameWithVersion[1]
			} else {
				splitCurrentComponentPurl = strings.Split(pkgs.Current_component_purl, "/")
				currentPackageWithoutArch = strings.Split(splitCurrentComponentPurl[2], "?")[0]
				currentPackageWithVersion = strings.Split(currentPackageWithoutArch, "@")
				installedVersion = currentPackageWithVersion[1]

				splitTargetComponentPurl = strings.Split(pkgs.Target_component_purl, "/")
				packageNameWithVersion = splitTargetComponentPurl[2]
				splitPackageNameWithVersion = strings.Split(packageNameWithVersion, "@")
				targetPackageName = splitPackageNameWithVersion[0]
				fixedVersion = splitPackageNameWithVersion[1]
			}
			
			updates.Updates = append(updates.Updates, v1alpha1.UpdatePackage{
				Name: targetPackageName,
				InstalledVersion: installedVersion,
				FixedVersion: fixedVersion,
				VulnerabilityID: pkgs.Vulnerability_id,
			})
		}
	}
	return &updates, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <image report>\n", os.Args[0])
		os.Exit(1)
	}

	// Initialize the parser
	fakeParser := newFakeParser()

	// Get the image report from command line
	imageReport := os.Args[1]

	report, err := fakeParser.parse(imageReport)
	if err != nil {
		fmt.Printf("error parsing report: %v\n", err)
		os.Exit(1)
	}

	// Serialize the standardized report and print it to stdout
	reportBytes, err := json.Marshal(report)
	if err != nil {
		fmt.Printf("Error serializing report: %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(reportBytes)
}
