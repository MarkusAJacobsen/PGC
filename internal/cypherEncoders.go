package internal

import "pgc/internal/pkg"

const CreatePlantCypher = "MERGE (p:Plant { name: $name, latinName: $latinName }) RETURN p.name"
const CreatePlantFamilyCypher = "MERGE (f:Family { name: $name }) RETURN f.name"
const LinkPlantAndFamilyCypher = "MATCH (p:Plant {name: $name}) MATCH (f:Family {name: $family}) MERGE (p)-[:IS_IN]->(f) RETURN p.name"
const GetAllPlantsCypher = "MATCH (p:Plant) RETURN p"

const CreateUserCypher = "MERGE (u:User {idToken: $idToken}) ON MATCH SET u.origin = $origin, u.email = $email ON CREATE SET u.name = $name, u.origin = $origin, u.email = $email RETURN u.idToken"
const CreateAreaCypher = "MERGE (a:Area {area: $area}) RETURN a.area"
const LinkUserAndAreaCypher = "MATCH (u:User {idToken: $idToken}) MATCH (a:Area {area: $area}) MERGE (u)-[:LIVES]->(a) RETURN u.idToken"

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

func CreateUser(u pkg.User) map[string]interface{} {
	return map[string]interface{}{
		"idToken": u.IdToken,
		"name":    u.Name,
		"origin":  u.Origin,
		"email":   u.Email,
	}
}

func CreateArea(u pkg.User) map[string]interface{} {
	return map[string]interface{}{
		"area": u.Area,
	}
}

func CreateUserAreaRelation(u pkg.User) map[string]interface{} {
	return map[string]interface{}{
		"idToken": u.IdToken,
		"area":    u.Area,
	}
}
