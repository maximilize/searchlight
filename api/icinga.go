package api

import (
	"fmt"

	"github.com/appscode/log"
	"github.com/appscode/searchlight/data"
)

type CheckPod string

const (
	CheckPodInfluxQuery CheckPod = "influx_query"
	CheckPodStatus      CheckPod = "pod_status"
	CheckPodVolume      CheckPod = "pod_volume"
	CheckPodExec        CheckPod = "pod_exec"
)

func (c CheckPod) IsValid() bool {
	_, ok := PodCommands[c]
	return ok
}

func ParseCheckPod(s string) (*CheckPod, error) {
	cmd := CheckPod(s)
	if _, ok := PodCommands[cmd]; !ok {
		return nil, fmt.Errorf("Invalid pod check command %s", s)
	}
	return &cmd, nil
}

type CheckNode string

const (
	CheckNodeInfluxQuery CheckNode = "influx_query"
	CheckNodeVolume      CheckNode = "node_volume"
	CheckNodeStatus      CheckNode = "node_status"
)

func (c CheckNode) IsValid() bool {
	_, ok := NodeCommands[c]
	return ok
}

func ParseCheckNode(s string) (*CheckNode, error) {
	cmd := CheckNode(s)
	if _, ok := NodeCommands[cmd]; !ok {
		return nil, fmt.Errorf("Invalid node check command %s", s)
	}
	return &cmd, nil
}

type CheckCluster string

const (
	CheckComponentStatus CheckCluster = "component_status"
	CheckJsonPath        CheckCluster = "json_path"
	CheckNodeExists      CheckCluster = "node_exists"
	CheckPodExists       CheckCluster = "pod_exists"
	CheckEvent           CheckCluster = "event"
	CheckCACert          CheckCluster = "ca_cert"

	CheckHttp        CheckCluster = "any_http"
	CheckHelloIcinga CheckCluster = "hello"
	//CheckDIG         CheckCluster = "dig"
	//CheckDNS         CheckCluster = "dns"
	CheckDummy CheckCluster = "dummy"
	//CheckICMP        CheckCluster = "icmp"
)

func (c CheckCluster) IsValid() bool {
	_, ok := ClusterCommands[c]
	return ok
}

func ParseCheckCluster(s string) (*CheckCluster, error) {
	cmd := CheckCluster(s)
	if _, ok := ClusterCommands[cmd]; !ok {
		return nil, fmt.Errorf("Invalid cluster check command %s", s)
	}
	return &cmd, nil
}

type IcingaCommand struct {
	Name   string
	Vars   map[string]data.CommandVar
	States []string
}

var (
	PodCommands     map[CheckPod]IcingaCommand
	NodeCommands    map[CheckNode]IcingaCommand
	ClusterCommands map[CheckCluster]IcingaCommand
)

func init() {
	PodCommands = map[CheckPod]IcingaCommand{}
	NodeCommands = map[CheckNode]IcingaCommand{}
	ClusterCommands = map[CheckCluster]IcingaCommand{}

	cmdList, err := data.LoadIcingaData()
	if err != nil {
		log.Fatal(err)
	}
	for _, cmd := range cmdList.Command {
		vars := make(map[string]data.CommandVar)
		for _, v := range cmd.Vars {
			vars[v.Name] = v
		}
		c := IcingaCommand{
			Name:   cmd.Name,
			Vars:   vars,
			States: cmd.States,
		}
		if c.Name == string(CheckPodInfluxQuery) ||
			c.Name == string(CheckPodStatus) ||
			c.Name == string(CheckPodVolume) ||
			c.Name == string(CheckPodExec) {
			PodCommands[CheckPod(c.Name)] = c
		}
		if c.Name == string(CheckNodeInfluxQuery) ||
			c.Name == string(CheckNodeVolume) ||
			c.Name == string(CheckNodeStatus) {
			NodeCommands[CheckNode(c.Name)] = c
		}
		if c.Name == string(CheckHttp) ||
			c.Name == string(CheckComponentStatus) ||
			c.Name == string(CheckJsonPath) ||
			c.Name == string(CheckNodeExists) ||
			c.Name == string(CheckPodExists) ||
			c.Name == string(CheckEvent) ||
			c.Name == string(CheckCACert) ||
			c.Name == string(CheckHelloIcinga) ||
			// c.Name == string(CheckDIG) ||
			// c.Name == string(CheckDNS) ||
			// c.Name == string(CheckICMP) ||
			c.Name == string(CheckDummy) {
			ClusterCommands[CheckCluster(c.Name)] = c
		}
	}
}
