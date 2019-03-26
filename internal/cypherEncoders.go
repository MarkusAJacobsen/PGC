package internal

import "pgc/internal/pkg"

const CreatePlantCypher = "CREATE (p:Plant { name: $name, latinName: $latinName }) RETURN p.name"

func CreatePlant(p pkg.Plant) (obj map[string]interface{}) {
	return map[string]interface{}{
		"name":      p.Name,
		"latinName": p.LatinName,
		"family":    p.Family,
	}
}
