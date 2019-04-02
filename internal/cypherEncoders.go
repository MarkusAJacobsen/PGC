package internal

import "pgc/internal/pkg"

const CreatePlantCypher = "MERGE (p:Plant { name: $name, latinName: $latinName }) RETURN p.name"
const CreatePlantFamilyCypher = "MERGE (f:Family { name: $name }) return f.name"
const LinkPlantAndFamilyCypher = "MATCH (p:Plant {name: $name}) MATCH (f:Family {name: $family}) MERGE (p)-[:IS_IN]->(f) RETURN p.name"
const GetAllPlants = "MATCH (p:Plant) RETURN p"

func CreatePlant(p pkg.Plant) map[string]interface{} {
	return map[string]interface{}{
		"name":      p.Name,
		"latinName": p.LatinName,
	}
}

func CreateFamily(p pkg.Plant) map[string]interface{} {
	return map[string]interface{}{
		"name": p.Family,
	}
}

func CreatePlantRelation(p pkg.Plant) map[string]interface{} {
	return map[string]interface{}{
		"name":      p.Name,
		"latinName": p.LatinName,
		"family":    p.Family,
	}
}
