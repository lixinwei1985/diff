package ormdiff

import (
	diff "github.com/r3labs/diff/v2"
)

type OrmDiff struct {
	new      interface{}
	old      interface{}
	upfields map[string]interface{}
	cl       diff.Changelog
}

func (df *OrmDiff) GetUpFields() map[string]interface{} {
	return df.upfields
}
func (df *OrmDiff) Up(new interface{}) (err error) {
	df.old = df.new
	df.new = new
	df.cl, err = diff.Diff(df.old, df.new)
	if err != nil {
		df.upfields = nil
		return
	}
	df.upfields = make(map[string]interface{}, len(df.cl))
	for _, change := range df.cl {
		if change.Type == "update" || change.Type == "create" {
			///////
			df.upfields[change.Path[0]] = change.To
		}

	}
	return err
}
