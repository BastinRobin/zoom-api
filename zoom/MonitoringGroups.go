package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type ListMonitoringGroups struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Sites         Sites
}
type MonitoringGroups []MonitoringGroup
type MonitoringGroup struct {
	ID                    string   `json:"id"`
	MonitorMembersCount   int      `json:"monitor_members_count"`
	MonitoredMembersCount int      `json:"monitored_members_count"`
	MonitoringPrivileges  []string `json:"monitoring_privileges"`
	Name                  string   `json:"name"`
	Prompt                bool     `json:"prompt"`
}
type Monitorings []Monitoring
type Monitoring struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MonitoringGroupID struct {
	ID                    string   `json:"id"`
	MonitorMembersCount   int      `json:"monitor_members_count"`
	MonitoredMembersCount int      `json:"monitored_members_count"`
	MonitoringPrivileges  []string `json:"monitoring_privileges"`
	Name                  string   `json:"name"`
	Prompt                bool     `json:"prompt"`
	Monitors              Monitors
}
type Monitors []Monitor
type Monitor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//return the list of monitoring groups
func (z *Zoom) GetListMonitoringGroups() (ListMonitoringGroups, error) {
	response, err := z.Request("/phone/monitoring_groups", "GET")
	if err != nil {
		log.Println(err)
		return ListMonitoringGroups{}, err
	}

	var listmonitoringroups ListMonitoringGroups

	// Unmarshal the response into ListMonitoringGroups struct
	err = json.Unmarshal(response, &listmonitoringroups)
	if err != nil {
		log.Println(err)
		return ListMonitoringGroups{}, err
	}
	//return  listmonitoringroupsstruct
	return listmonitoringroups, nil

}

// return monitoring group by ID
func (z *Zoom) GetMonitoringGroupId(monitoringGroupId string) (MonitoringGroupID, error) {
	// Create the url for getting group
	endpoint := fmt.Sprintf("/phone/monitoring_groups/%v", monitoringGroupId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return MonitoringGroupID{}, err
	}

	// Umarshal the response into MonitoringGroupID struct
	var monitoringgroupid MonitoringGroupID
	err = json.Unmarshal(response, &monitoringgroupid)
	if err != nil {
		log.Println(err)
		return MonitoringGroupID{}, err
	}
	// return monitoringgroupid struct
	return monitoringgroupid, nil
}
