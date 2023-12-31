package data

const (
	ImageRequest = 0
	HostRequest  = 1
)

type Report struct {
	RequestType int
	ServerUrl   string

	General         *GeneralType
	Sensitive       *SensitiveType
	Malware         *MalwareType
	Vulnerabilities *VulnerabilitiesType
	ScanHistory     *ScanHistoryType

	BenchResults *BenchResultsType
}

func (report *Report) GetImageAssurancePolicies() (map[string]bool, map[string][]CheckPerformedType) {
	result := make(map[string]bool)
	checks := make(map[string][]CheckPerformedType)

	for _, policy := range report.General.AssuranceResults.ChecksPerformed {
		if _, ok := result[policy.PolicyName]; !ok {
			result[policy.PolicyName] = false
		}
		if policy.Failed {
			result[policy.PolicyName] = true
		}
		checks[policy.PolicyName] = append(checks[policy.PolicyName], policy)
	}
	return result, checks
}
