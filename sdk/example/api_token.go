package example

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func APIToken() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.APIToken,
		DocsURL:       sdk.URL("http://example.com/docs/api_token"),
		ManagementURL: sdk.URL("http://dashboard.example.com/user/security/tokens"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.AccountID,
				MarkdownDescription: "The Example API account ID.",
				Secret:              false,
				Composition: &schema.ValueComposition{
					Length: 12,
					Charset: schema.Charset{
						Digits: true,
					},
				},
			},
			{
				Name:                fieldname.Token,
				MarkdownDescription: "The API token used to authenticate to the Example API.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 40,
					Prefix: "tkn_",
					Charset: schema.Charset{
						Uppercase: true,
						Digits:    true,
					},
				},
			},
		},
		Provisioner: provision.EnvVars(map[string]string{
			fieldname.AccountID: "EXAMPLE_ACCOUNT_ID",
			fieldname.Token:     "EXAMPLE_API_TOKEN",
		}),
		Importer: importer.TryAll(
			importer.TryEnvVarPair(map[string]string{
				fieldname.AccountID: "EXAMPLE_ACCOUNT_ID",
				fieldname.Token:     "EXAMPLE_API_TOKEN",
			}),
			importer.TryEnvVarPair(map[string]string{
				fieldname.AccountID: "EXAMPLE_ACCOUNT_ID",
				fieldname.Token:     "EXAMPLE_TOKEN",
			}),
		),
	}
}
