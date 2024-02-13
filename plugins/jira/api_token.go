package jira

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIToken() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIToken,
		DocsURL:       sdk.URL("https://jira.com/docs/api_token"),               // TODO: Replace with actual URL
		ManagementURL: sdk.URL("https://console.jira.com/user/security/tokens"), // TODO: Replace with actual URL
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Credential,
				MarkdownDescription: "Token used to authenticate to Jira.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 192,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
						Symbols:   true,
					},
				},
			},
			{
				Name:                fieldname.Email,
				MarkdownDescription: "Email used to authenticate to Jira.",
				Secret:              false,
			},
		},
		DefaultProvisioner: jiraProvisioner{},
		Importer: importer.TryAll(
			importer.TryEnvVarPair(defaultEnvVarMapping),
		)}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"JIRA_TOKEN": fieldname.Credential,
	"JIRA_EMAIL": fieldname.Email,
}

type jiraProvisioner struct{}

func (j jiraProvisioner) Description() string {
	return "Jira cli token provisioner"
}

func (j jiraProvisioner) Provision(_ context.Context, input sdk.ProvisionInput, output *sdk.ProvisionOutput) {

	output.AddArgs("--user", input.ItemFields[fieldname.Email]+":"+input.ItemFields[fieldname.Credential])
}

func (j jiraProvisioner) Deprovision(_ context.Context, _ sdk.DeprovisionInput, _ *sdk.DeprovisionOutput) {
	// No-op
}
