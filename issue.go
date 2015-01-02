package main

var hlpIssue = &Command{
	UsageLine: "issue",
	Short:     "information on reporting of issues",
	Long: `
In case you find this program misbehaving, acting strange or want to request a feature,
please file an issue at

    https://github.com/coduno/cli/issues

In order to make resolution easier please make sure your description contains, but is not limited to,
the following points:

  1. Expected behaviour or output
  2. Actual behaviour or output
  3. Steps to reproduce the relevant behaviour or output
  4. Version information as printed by 'coduno version'

The more concise any issue report is, the faster it will be resolved.
	`,
}
