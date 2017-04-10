package models

type Diagram struct {
	Nodes []*DiagramNode `json:"nodes"`
	Links []*DiagramLink `json:"links"`
}

type DiagramNode struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type DiagramInterface struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type DiagramMeta struct {
	Interface *DiagramInterface `json:"interface"`
}

type DiagramLink struct {
	Source string       `json:"source"`
	Target string       `json:"target"`
	Meta   *DiagramMeta `json:"meta"`
}

func NewDiagramModel() *Diagram {
	return &Diagram{}
}

func NewDiagramNodeModel() *DiagramNode {
	return &DiagramNode{}
}

func NewDiagramInterfaceModel() *DiagramInterface {
	return &DiagramInterface{}
}

func NewDiagramMetaModel() *DiagramMeta {
	return &DiagramMeta{}
}

func NewDiagramLinkModel() *DiagramLink {
	return &DiagramLink{}
}

var sharedDiagramModel = NewDiagramModel()
var sharedDiagramNodeModel = NewDiagramNodeModel()
var sharedDiagramInterfaceModel = NewDiagramInterfaceModel()
var sharedDiagramMetaModel = NewDiagramMetaModel()
var sharedDiagramLinkModel = NewDiagramLinkModel()

func SharedDiagramModel() *Diagram {
	return sharedDiagramModel
}

func SharedDiagramNodeModel() *DiagramNode {
	return sharedDiagramNodeModel
}

func SharedDiagramInterfaceModel() *DiagramInterface {
	return sharedDiagramInterfaceModel
}

func SharedDiagramMetaModel() *DiagramMeta {
	return sharedDiagramMetaModel
}

func SharedDiagramLinkModel() *DiagramLink {
	return sharedDiagramLinkModel
}

func init() {
}
