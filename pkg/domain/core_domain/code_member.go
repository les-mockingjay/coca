package core_domain

type CodeMember struct {
	ID            string
	AliasPackage  string
	Name          string
	Type          string
	Structures    []CodeDataStruct
	FunctionNodes []CodeFunction
	Namespace     []string
	FileID        string
	DataStructID  string
	Position      CodePosition
}

func NewCodeMember() *CodeMember {
	return &CodeMember{}
}

func (c *CodeMember) BuildMemberId() {
	IsDefaultFunction := c.Name == "default"
	if IsDefaultFunction {
		for _, function := range c.FunctionNodes {
			c.ID = c.AliasPackage + "::" + function.Name
		}
	} else {
		packageName := c.FileID
		if c.FileID != c.AliasPackage {
			packageName = c.FileID + "|" + c.AliasPackage
		}
		c.ID = packageName + ":" + c.DataStructID
	}
}
