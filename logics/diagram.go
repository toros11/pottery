package logics

import (
	"fmt"
	"github.com/jinzhu/gorm"
	loamLogics "github.com/qb0C80aE/loam/logics"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
)

const diagramImageRoot string = "/ui/files/images/diagram"

var physicalNodeIconPaths = map[int]string{
	1: fmt.Sprintf("%s/%s", diagramImageRoot, "l2switch.png"),
	2: fmt.Sprintf("%s/%s", diagramImageRoot, "l3switch.png"),
	3: fmt.Sprintf("%s/%s", diagramImageRoot, "firewall.png"),
	4: fmt.Sprintf("%s/%s", diagramImageRoot, "router.png"),
	5: fmt.Sprintf("%s/%s", diagramImageRoot, "loadbalancer.png"),
	6: fmt.Sprintf("%s/%s", diagramImageRoot, "server.png"),
	7: fmt.Sprintf("%s/%s", diagramImageRoot, "network.png"),
}

var virtualNodeIconPaths = map[int]string{
	1: fmt.Sprintf("%s/%s", diagramImageRoot, "l2switch_v.png"),
	2: fmt.Sprintf("%s/%s", diagramImageRoot, "l3switch_v.png"),
	3: fmt.Sprintf("%s/%s", diagramImageRoot, "firewall_v.png"),
	4: fmt.Sprintf("%s/%s", diagramImageRoot, "router_v.png"),
	5: fmt.Sprintf("%s/%s", diagramImageRoot, "loadbalancer_v.png"),
	6: fmt.Sprintf("%s/%s", diagramImageRoot, "server_v.png"),
	7: fmt.Sprintf("%s/%s", diagramImageRoot, "network.png"),
}

var segmentIconPath = fmt.Sprintf("%s/%s", diagramImageRoot, "segment.png")

type PhysicalDiagramLogic struct {
}

type LogicalDiagramLogic struct {
}

func (_ *PhysicalDiagramLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {
	diagram := &models.Diagram{}

	nodes := []*loamModels.Node{}
	if err := db.Preload("Ports").Select(queryFields).Find(&nodes).Error; err != nil {
		return nil, err
	}

	nodeMap := make(map[int]*loamModels.Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	ports := []*loamModels.Port{}
	if err := db.Select(queryFields).Find(&ports).Error; err != nil {
		return nil, err
	}

	portMap := make(map[int]*loamModels.Port)
	for _, port := range ports {
		portMap[port.ID] = port
	}

	for _, node := range nodes {
		var iconPathMap map[int]string = nil
		if node.NodePvID == 1 {
			iconPathMap = physicalNodeIconPaths
		} else {
			iconPathMap = virtualNodeIconPaths
		}
		diagramNode := &models.DiagramNode{
			Name: node.Name,
			Icon: iconPathMap[node.NodeTypeID],
		}
		diagram.Nodes = append(diagram.Nodes, diagramNode)
	}

	registerdPortMap := make(map[int]int)
	for _, port := range ports {
		_, exists := registerdPortMap[int(port.DestinationPortID.Int64)]
		if (port.DestinationPortID.Valid) && (!exists) {
			sourceNode := nodeMap[port.NodeID]
			destinationPort := portMap[int(port.DestinationPortID.Int64)]
			destinationNode := nodeMap[destinationPort.NodeID]

			diagramInterface := &models.DiagramInterface{
				Source: port.Name,
				Target: destinationPort.Name,
			}
			diagramMeta := &models.DiagramMeta{
				Interface: diagramInterface,
			}
			diagramLink := &models.DiagramLink{
				Source: sourceNode.Name,
				Target: destinationNode.Name,
				Meta:   diagramMeta,
			}

			diagram.Links = append(diagram.Links, diagramLink)

			registerdPortMap[port.ID] = port.ID
		}
	}

	return diagram, nil
}

func (_ *PhysicalDiagramLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {
	return nil, nil
}

func (_ *PhysicalDiagramLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {
	return nil, nil
}

func (_ *PhysicalDiagramLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (_ *PhysicalDiagramLogic) Delete(db *gorm.DB, id string) error {
	return nil
}

func (_ *PhysicalDiagramLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *PhysicalDiagramLogic) Options(db *gorm.DB) error {
	return nil
}

func (_ *LogicalDiagramLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {
	nodePvs := []*loamModels.NodePv{}
	if err := db.Select(queryFields).Find(&nodePvs).Error; err != nil {
		return nil, err
	}

	nodeTypes := []*loamModels.NodeType{}
	if err := db.Select(queryFields).Find(&nodeTypes).Error; err != nil {
		return nil, err
	}

	nodes := []*loamModels.Node{}
	if err := db.Preload("Ports").Select(queryFields).Find(&nodes).Error; err != nil {
		return nil, err
	}

	ports := []*loamModels.Port{}
	if err := db.Select(queryFields).Find(&ports).Error; err != nil {
		return nil, err
	}

	nodeMap := make(map[int]*loamModels.Node)
	portMap := make(map[int]*loamModels.Port)
	consumedPortMap := make(map[int]*loamModels.Port)

	for _, node := range nodes {
		nodeMap[node.ID] = node
	}
	for _, port := range ports {
		portMap[port.ID] = port
	}

	segments := loamLogics.GenerateSegments(nodeMap, portMap, consumedPortMap)

	diagram := &models.Diagram{}

	for _, node := range nodes {
		if node.NodeTypeID != 1 {
			var iconPathMap map[int]string = nil
			if node.NodePvID == 1 {
				iconPathMap = physicalNodeIconPaths
			} else {
				iconPathMap = virtualNodeIconPaths
			}
			diagramNode := &models.DiagramNode{
				node.Name,
				iconPathMap[node.NodeTypeID],
			}
			diagram.Nodes = append(diagram.Nodes, diagramNode)
		}
	}

	for i, segment := range segments {

		diagramNode := &models.DiagramNode{
			fmt.Sprintf("[%d]%s", i, segment.Cidr),
			segmentIconPath,
		}
		diagram.Nodes = append(diagram.Nodes, diagramNode)

		for _, port := range segment.Ports {
			diagramInterface := &models.DiagramInterface{
				Source: "",
				Target: fmt.Sprintf("%s[%s](%s/%d)",
					port.Name,
					port.MacAddress.String,
					port.Ipv4Address.String,
					port.Ipv4Prefix.Int64,
				),
			}
			diagramMeta := &models.DiagramMeta{
				Interface: diagramInterface,
			}
			diagramLink := &models.DiagramLink{
				Source: fmt.Sprintf("[%d]%s", i, segment.Cidr),
				Target: nodeMap[port.NodeID].Name,
				Meta:   diagramMeta,
			}
			diagram.Links = append(diagram.Links, diagramLink)
		}

	}

	return diagram, nil
}

func (_ *LogicalDiagramLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {
	return nil, nil
}

func (_ *LogicalDiagramLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {
	return nil, nil
}

func (_ *LogicalDiagramLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {
	return nil, nil
}

func (_ *LogicalDiagramLogic) Delete(db *gorm.DB, id string) error {
	return nil
}

func (_ *LogicalDiagramLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *LogicalDiagramLogic) Options(db *gorm.DB) error {
	return nil
}

var PhysicalDiagramLogicInstance = &PhysicalDiagramLogic{}
var LogicalDiagramLogicInstance = &LogicalDiagramLogic{}

func init() {
}
