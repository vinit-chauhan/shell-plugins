package jira

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "jira",
		Platform: schema.PlatformInfo{
			Name:     "Jira",
			Homepage: sdk.URL("https://jira.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			APIToken(),
		},
		Executables: []schema.Executable{
			JiraCLI(),
		},
	}
}
