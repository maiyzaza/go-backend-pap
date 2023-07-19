package test

import (
	"PattayaAvenueProperty/service"
	"testing"
)

type PersonController struct {
	personService service.PersonService
}

func NewPersonController(personService service.PersonService) PersonController {
	return PersonController{personService: personService}
}

func TestGetProfiles(t *testing.T) {
	got := NewPersonController(service.NewPersonService()).GetProfiles()
	want := `{
				"place_id": 1,
				"place_name": "Grand Avenue",
				"Buildings": [
					{
						"building_id": 1,
						"building_name": "A",
						"floors": [
							{
								"floor_id": 1,
								"floor_number": "G"
							},
							{
								"floor_id": 2,
								"floor_number": "1"
							},
							{
								"floor_id": 3,
								"floor_number": "2"
							},
							{
								"floor_id": 4,
								"floor_number": "3"
							}
						]
					}
				]
			},
			{
				"place_id": 1,
				"place_name": "Grand Avenue",
				"Buildings": [
					{
						"building_id": 1,
						"building_name": "A",
						"floors": [
							{
								"floor_id": 1,
								"floor_number": "G"
							},
							{
								"floor_id": 2,
								"floor_number": "1"
							},
							{
								"floor_id": 3,
								"floor_number": "2"
							},
							{
								"floor_id": 4,
								"floor_number": "3"
							}
						]
					},
					{
						"building_id": 2,
						"building_name": "B",
						"floors": [
							{
								"floor_id": 5,
								"floor_number": "1"
							},
							{
								"floor_id": 6,
								"floor_number": "2"
							},
							{
								"floor_id": 7,
								"floor_number": "3"
							},
							{
								"floor_id": 8,
								"floor_number": "4"
							}
						]
					}
				]
			},
			{
				"place_id": 2,
				"place_name": "Centric Sea",
				"Buildings": [
					{
						"building_id": 3,
						"building_name": "A",
						"floors": null
					}
				]
			},
			{
				"place_id": 2,
				"place_name": "Centric Sea",
				"Buildings": [
					{
						"building_id": 3,
						"building_name": "A",
						"floors": null
					},
					{
						"building_id": 4,
						"building_name": "B",
						"floors": null
					}
				]
			}
		],
		"error": null
	}`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
