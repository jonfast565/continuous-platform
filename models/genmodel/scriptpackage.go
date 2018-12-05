package genmodel

import (
	"../../stringutil"
	"../jenkinsmodel"
	"github.com/ahmetb/go-linq"
)

type ScriptPackage struct {
	Scripts []ScriptKeyValuePair
}

func (sp ScriptPackage) GetScriptContentsByKey(key jenkinsmodel.JenkinsJobKey) *string {
	for _, script := range sp.Scripts {
		myKeys := script.GetJenkinsKeyList()
		myKeys.SanitizeKeyList()
		result := linq.From(myKeys).FirstWithT(func(myKey jenkinsmodel.JenkinsJobKey) bool {
			return stringutil.StringArrayCompare(key.Keys, myKey.Keys) &&
				myKey.Type == jenkinsmodel.PipelineJob
		})

		if result != nil {
			return &script.Value
		}
	}
	return nil
}
