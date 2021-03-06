package internal

import (
	"pgc/internal/pkg"
)

// Plant
const CreatePlantCypher = "MERGE (p:Plant { id: $id, name: $name, latinName: $latinName, barcode: $barcode, category: $category }) RETURN p.name"
const CreatePlantFamilyCypher = "MERGE (f:Family { name: $name }) RETURN f.name"
const LinkPlantAndFamilyCypher = "MATCH (p:Plant {name: $name}) MATCH (f:Family {name: $family}) MERGE (p)-[:IS_IN]->(f) RETURN p.name"
const GetAllPlantsCypher = "MATCH (p:Plant) RETURN p"
const GetPlantCypher = "MATCH (p:Plant {id: $pId}) RETURN p"
const GetPlantBarcodeCypher = "MATCH (p:Plant {barcode: $barcode}) RETURN p"
const DeletePlantCypher = "MATCH (p:Plant {id: $pId}) DETACH DELETE p"

// User
const CreateUserCypher = "MERGE (u:User {idToken: $idToken}) ON MATCH SET u.origin = $origin, u.email = $email ON CREATE SET u.name = $name, u.origin = $origin, u.email = $email RETURN u.idToken"
const CreateAreaCypher = "MERGE (a:Area {area: $area}) RETURN a.area"
const LinkUserAndAreaCypher = "MATCH (u:User {idToken: $idToken}) MATCH (a:Area {area: $area}) MERGE (u)-[:LIVES]->(a) RETURN u.idToken"
const GetUserCypher = "MATCH (u:User {idToken: $idToken}) RETURN u"
const DeleteUserCypher = "MATCH (u:User {idToken: $idToken}) DETACH DELETE u"

// Project
const CreateProjectCypher = "MERGE (pr:Project {id: $id, name: $name, startDate: $startDate, status: $status, climate: $climate}) RETURN pr.id"
const GetProjectsCypher = "MATCH (u:User {idToken: $idToken})-[:HAS_PROJECT]->(pr:Project) RETURN pr"
const GetProjectCypher = "MATCH (u:User {idToken: $idToken})-[:HAS_PROJECT]->(pr:Project {id: $pId}) RETURN pr"
const LinkProjectCypher = "MATCH (pr:Project {id: $prId}) MATCH (u:User {idToken: $idToken}) MATCH (p:Plant {id: $pId}) MERGE (u)-[:HAS_PROJECT]->(pr) MERGE (pr)-[:IS_PLANT]->(p) RETURN pr.id"
const DeleteProjectCypher = "MATCH (pr:Project {id: $id}) DETACH DELETE pr"

// Guide
const CreateGuideCypher = "MERGE (g:Guide {id: $id, title: $title, chapterTitles: $chapterTitles}) return g.id"
const CreateStageCypher = "MERGE (s:Stage {id: $id, title: $title, text: $text, images: $images}) WITH s MATCH (g:Guide {id: $gId}) MERGE (g)-[:CONTAINS_STAGE {pageNr: $pageNr, chapterNr: $chapterNr, filter: $filter}]->(s) RETURN g.id"
const GetGuideCypher = "MATCH (g:Guide {id: $id})-[c:CONTAINS_STAGE]->(s:Stage) RETURN g.id AS id, g.title AS title, g.chapterTitles AS chapterTitles, collect(s) AS stages, collect(c) AS containsStageRel"
const DeleteGuideCypher = "MATCH (g:Guide {id: $id}) DETACH DELETE g"
const DeleteOrphanedStages = "MATCH (s:Stage) WHERE NOT (s)--() delete (s)"
const LinkPlantToGuideCypher = "MATCH (g:Guide {id: $gId}) MATCH (p:Plant {id: $pId}) MERGE (p)-[:HAS_GUIDE]->(g) RETURN p.id"

func CreatePlant(p pkg.Plant) map[string]interface{} {
	return map[string]interface{}{
		"id":        p.Id,
		"name":      p.Name,
		"latinName": p.LatinName,
		"barcode":   p.Barcode,
		"category":  p.Category,
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

func CreateProject(pr pkg.Project) map[string]interface{} {
	return map[string]interface{}{
		"id":        pr.Id,
		"name":      pr.Name,
		"startDate": pr.StartDate,
		"status":    pr.Status,
		"climate":   pr.Climate,
	}
}

func CreateProjectRelation(pl pkg.ProjectLink) map[string]interface{} {
	return map[string]interface{}{
		"prId":    pl.Project.Id,
		"idToken": pl.UserId,
		"pId":     pl.PlantId,
	}
}

func CreateGuide(g pkg.Guide) map[string]interface{} {
	return map[string]interface{}{
		"id":            g.Id,
		"title":         g.Title,
		"chapterTitles": g.ChapterTitles,
	}
}

func CreateStage(s pkg.Stage) map[string]interface{} {
	return map[string]interface{}{
		"id":        s.Id,
		"title":     s.Title,
		"pageNr":    s.PageNr,
		"chapterNr": s.ChapterNr,
		"filter":    s.Filter,
		"text":      s.Text,
		"images":    s.Images,
	}
}
