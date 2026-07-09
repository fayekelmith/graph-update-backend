package db

import "github.com/fayekelmith/graph-update-backend/models"

var bounties = map[string]models.Bounty{
	"1": {ID: "1", Title: "Fix login bug", Amount: 500, Status: "open"},
}

var people = map[string]models.Person{
	"1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
}

func AllBounties() []models.Bounty {
	out := []models.Bounty{}
	for _, b := range bounties {
		out = append(out, b)
	}
	return out
}

func FindBounty(id string) (models.Bounty, bool) { b, ok := bounties[id]; return b, ok }
func SaveBounty(b models.Bounty)                 { bounties[b.ID] = b }
func RemoveBounty(id string)                     { delete(bounties, id) }

func AllPeople() []models.Person {
	out := []models.Person{}
	for _, p := range people {
		out = append(out, p)
	}
	return out
}

func FindPerson(id string) (models.Person, bool) { p, ok := people[id]; return p, ok }
