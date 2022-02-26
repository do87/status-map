package payload

type InfraList []InfraListItem

type InfraListItem struct {
	Name        string `json:"name"`
	Href        string `json:"href"`
	Environment string `json:"environment"`
}

type StageList []StageListItem

type StageListItem struct {
	Name        string `json:"name"`
	Href        string `json:"href"`
	Environment string `json:"environment"`
}
