package jira

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestAPITokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIToken().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Token: "X_6*@pb^M\H`0tc=#T0,BHSa-$Q|`M{^[D\rMvo:orttY{6rcm_vzJ_BDT}pUWM$Rlx&#34;x&#39;0N&amp;Crj&lt;9aes$lw&#39;aM:~&#34;$&amp;#z9fYgmM|{]v/`Aq$c}(pQBV7kUy]&#43;(j%Jf\s[&gt;91@J#c%87CkN1js4Pc$7IUVO`\33\9)G)I:b]b^cPNThwz*4&#39;CwB\;EXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"JIRA_TOKEN": "X_6*@pb^M\H`0tc=#T0,BHSa-$Q|`M{^[D\rMvo:orttY{6rcm_vzJ_BDT}pUWM$Rlx&#34;x&#39;0N&amp;Crj&lt;9aes$lw&#39;aM:~&#34;$&amp;#z9fYgmM|{]v/`Aq$c}(pQBV7kUy]&#43;(j%Jf\s[&gt;91@J#c%87CkN1js4Pc$7IUVO`\33\9)G)I:b]b^cPNThwz*4&#39;CwB\;EXAMPLE",
				},
			},
		},
	})
}

func TestAPITokenImporter(t *testing.T) {
	plugintest.TestImporter(t, APIToken().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"JIRA_TOKEN": "X_6*@pb^M\H`0tc=#T0,BHSa-$Q|`M{^[D\rMvo:orttY{6rcm_vzJ_BDT}pUWM$Rlx&#34;x&#39;0N&amp;Crj&lt;9aes$lw&#39;aM:~&#34;$&amp;#z9fYgmM|{]v/`Aq$c}(pQBV7kUy]&#43;(j%Jf\s[&gt;91@J#c%87CkN1js4Pc$7IUVO`\33\9)G)I:b]b^cPNThwz*4&#39;CwB\;EXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Token: "X_6*@pb^M\H`0tc=#T0,BHSa-$Q|`M{^[D\rMvo:orttY{6rcm_vzJ_BDT}pUWM$Rlx&#34;x&#39;0N&amp;Crj&lt;9aes$lw&#39;aM:~&#34;$&amp;#z9fYgmM|{]v/`Aq$c}(pQBV7kUy]&#43;(j%Jf\s[&gt;91@J#c%87CkN1js4Pc$7IUVO`\33\9)G)I:b]b^cPNThwz*4&#39;CwB\;EXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in jira/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "X_6*@pb^M\H`0tc=#T0,BHSa-$Q|`M{^[D\rMvo:orttY{6rcm_vzJ_BDT}pUWM$Rlx&#34;x&#39;0N&amp;Crj&lt;9aes$lw&#39;aM:~&#34;$&amp;#z9fYgmM|{]v/`Aq$c}(pQBV7kUy]&#43;(j%Jf\s[&gt;91@J#c%87CkN1js4Pc$7IUVO`\33\9)G)I:b]b^cPNThwz*4&#39;CwB\;EXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
