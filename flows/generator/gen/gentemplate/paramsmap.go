package gentemplate

func (tp *Params) Map() map[string]interface{} {
	m := map[string]interface{}{}

	m["custom"] = tp.CustomKeys
	m["answers"] = tp.Answers
	m["developer"] = tp.DeveloperInfo
	m["moduleInfo"] = tp.ModuleInfo
	m["var"] = tp.Variables

	for k, v := range tp.CommonInfo {
		m[k] = v
	}

	for k, v := range tp.ProjectInfo {
		m[k] = v
	}

	return m
}
