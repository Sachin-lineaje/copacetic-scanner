// Type definitions for fake scanner report
package main

// FakeReport contains OS, Arch, and Package information
type FakeReport struct {
	OSType    string
	OSVersion string
	Arch      string
	Metadata  Metadata `json:"meta_data"`
}

type Metadata struct {
	Basic_plan_component_vulnerability_fixes []FakePackage `json:"basic_plan_component_vulnerability_fixes"`
}

// FakePackage contains package and vulnerability information
type FakePackage struct {
	Current_package_id         string `json:"current_package_id"`
	Current_sbom_id            string `json:"current_sbom_id"`
	Target_hub_package_id      string `json:"target_hub_package_id"`
	Target_hub_sbom_id         string `json:"target_hub_sbom_id"`
	Current_component_purl     string `json:"current_component_purl"`
	Target_component_purl      string `json:"target_component_purl"`
	Fixed_vuln                 int 	  `json:"fixed_vuln"`
	Vulnerability_id           string `json:"vulnerability_id"`
}