package internal

import "pgc/internal/pkg"

const CreatePlantCypher = "CREATE (p:Plant { name: $name, latinName: $latinName }) RETURN p.name"
const CreatePlantFamilyCypher = "CREATE (f:Family { name: $name }) return f.name"
const LinkPlantAndFamilyCypher = "MATCH (p:Plant {name: $name}) MATCH (f:Family {name: $family}) CREATE (p)-[:IS_IN]->(f) return p.name"

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
