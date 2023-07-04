package config

import "fmt"

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type LocalConfig struct {
	Type string `hcl:"type"`
	AWS  AWS    `hcl:"aws,block"`
}

func init() {
	localConfig := LocalConfig{Type: "local", AWS: AWS{
		Label: "test",
		EC2: EC2WithSchedule{
			//EC2:      EC2{Name: "test"},
			Schedule: Schedule{Time: "test"},
		}}}
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(&localConfig, f.Body())
	fmt.Printf("%s", f.Bytes())
}
