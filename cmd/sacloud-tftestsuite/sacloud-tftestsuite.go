package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"golang.org/x/xerrors"
)

type Resource struct {
	Resource []SakuracloudServer `hcl:"resource,block"`
	Remain   hcl.Body            `hcl:",remain"`
}

type SakuracloudServer struct {
	ResourceType string   `hcl:"resource_type,label"`
	ResourceName string   `hcl:"resource_name,label"`
	Commitment   string   `hcl:"commitment,optional"`
	Name         string   `hcl:"name,optional"`
	Core         int      `hcl:"core,optional"`
	Memory       int      `hcl:"memory,optional"`
	Remain       hcl.Body `hcl:",remain"`
}

var ServerPlan = make(map[int][]int, 0) // [Core][](RAM/1024)

var Results []string

func main() {
	if err := sacloudTFTestSuite(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func testTF(tfPath string) error {
	b, err := ioutil.ReadFile(tfPath)
	if err != nil {
		return xerrors.Errorf(": %v", err)
	}
	parser := hclparse.NewParser()
	file, diagnostics := parser.ParseHCL(b, "test")

	if diagnostics.HasErrors() {
		panic(diagnostics)
	}

	var example Resource
	diags := gohcl.DecodeBody(file.Body, nil, &example)
	if diags.HasErrors() {
		panic(diags)
	}
	for _, server := range example.Resource {
		if server.ResourceType != "sakuracloud_server" {
			continue
		}
		if server.Core == 0 {
			server.Core = 1
		}
		if server.Memory == 0 {
			server.Memory = 1
		}
		if server.Commitment == "" {
			server.Commitment = "standard"
		}
		if !contains(ServerPlan[server.Core], server.Memory) {
			Results = append(Results, fmt.Sprintf("%s: not serviced in %dCore/%dGB\n", server.ResourceName, server.Core, server.Memory))
		}
	}

	if len(Results) > 0 {
		for _, v := range Results {
			fmt.Fprintf(os.Stderr, v)
		}
		return xerrors.New("test failed")
	}
	return nil
}

func sacloudTFTestSuite(args []string) error {
	if len(args) < 2 {
		return xerrors.New("file name argument is required")
	}

	zone := os.Getenv("SAKURACLOUD_ZONE")
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	caller := sacloud.NewClient(token, secret)
	serverPlanOp := sacloud.NewServerPlanOp(caller)
	condition := &sacloud.FindCondition{
		Filter: search.Filter{
			search.Key("Availability"): search.AndEqual("available"),
			search.Key("Commitment"):   search.AndEqual("standard"),
		},
	}
	searched, err := serverPlanOp.Find(context.Background(), zone, condition)
	if err != nil {
		return xerrors.Errorf(": %v", err)
	}
	for _, v := range searched.ServerPlans {
		ServerPlan[v.CPU] = append(ServerPlan[v.CPU], v.MemoryMB/1024)
	}

	return testTF(args[1])
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
