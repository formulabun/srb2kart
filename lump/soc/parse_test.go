package soc

import "testing"

func TestParseSoc(t *testing.T) {
	testData := []struct {
		Data   string
		Blocks Soc
	}{

		// Empty
		{
			`
    `,
			Soc{},
		},

		// Level 1 from maps.kart
		{
			`
Level 1
LevelName = Green Hills
SubTitle = Ring Cup 1
TypeOfLevel = Race
NumLaps = 3
SkyNum = 88
Music = KMAP01
RecordAttack = true
NoVisitNeeded = true
      `,
			Soc{
				Block{
					Header{
						"LEVEL", "1",
					},
					map[string]string{
						"LEVELNAME":     "Green Hills",
						"SUBTITLE":      "Ring Cup 1",
						"TYPEOFLEVEL":   "Race",
						"NUMLAPS":       "3",
						"SKYNUM":        "88",
						"MUSIC":         "KMAP01",
						"RECORDATTACK":  "true",
						"NOVISITNEEDED": "true",
					},
				},
			},
		},

		// Free slots
		{
			`
# Free Slots
FREESLOT
MT_SOR_TREE
MT_SOR_TREEB
MT_SOR_ROCKET
MT_SOR_LEAVE
MT_SOR_SIGN
MT_SOR_BOOST
S_SOR_TREE
S_SOR_TREEB
S_SOR_ROCKET_ANIMATE
S_SOR_ROCKET
S_SOR_LEAVE
S_SOR_SIGN
S_SOR_BOOST
SPR_SOR_
SPR_SORS
SPR_SORT
`,
			Soc{
				Block{
					Header{
						"FREESLOT", "",
					},
					map[string]string{
						"MT_SOR_TREE":          "",
						"MT_SOR_TREEB":         "",
						"MT_SOR_ROCKET":        "",
						"MT_SOR_LEAVE":         "",
						"MT_SOR_SIGN":          "",
						"MT_SOR_BOOST":         "",
						"S_SOR_TREE":           "",
						"S_SOR_TREEB":          "",
						"S_SOR_ROCKET_ANIMATE": "",
						"S_SOR_ROCKET":         "",
						"S_SOR_LEAVE":          "",
						"S_SOR_SIGN":           "",
						"S_SOR_BOOST":          "",
						"SPR_SOR_":             "",
						"SPR_SORS":             "",
						"SPR_SORT":             "",
					},
				},
			},
		},

		// multi block soc
		{
			`
Level 10
LevelName = Grand Metropolis
NoZone = true
SubTitle = Sneaker Cup FINAL
TypeOfLevel = Race
NumLaps = 2
SkyNum = 70
Music = KMAP10
RecordAttack = true
NoVisitNeeded = true
NextLevel = EVALUATION

#######################
# MAP11-15: Water Cup #
#######################

Level 11
LevelName = Sunbeam Paradise
SubTitle = Water Cup 1
TypeOfLevel = Race
NumLaps = 3
SkyNum = 2
Music = KMAP11
RecordAttack = true
NoVisitNeeded = true
      `,
			Soc{
				Block{
					Header{"LEVEL", "10"},
					map[string]string{
						"LEVELNAME":     "Grand Metropolis",
						"NOZONE":        "true",
						"SUBTITLE":      "Sneaker Cup FINAL",
						"TYPEOFLEVEL":   "Race",
						"NUMLAPS":       "2",
						"SKYNUM":        "70",
						"MUSIC":         "KMAP10",
						"RECORDATTACK":  "true",
						"NOVISITNEEDED": "true",
						"NEXTLEVEL":     "EVALUATION",
					},
				},
				Block{
					Header{"LEVEL", "11"},
					map[string]string{
						"LEVELNAME":     "Sunbeam Paradise",
						"SUBTITLE":      "Water Cup 1",
						"TYPEOFLEVEL":   "Race",
						"NUMLAPS":       "3",
						"SKYNUM":        "2",
						"MUSIC":         "KMAP11",
						"RECORDATTACK":  "true",
						"NOVISITNEEDED": "true",
					},
				},
			},
		},
	}

	for _, d := range testData {
		soc := ParseSoc([]byte(d.Data))
		if len(soc) != len(d.Blocks) {
			t.Fatalf("len(ParseSoc(%s)) = %d but got %d instead", d.Data, len(d.Blocks), len(soc))
		}

		for iBlock, block := range d.Blocks {
			if block.Header.Type != soc[iBlock].Header.Type {
				t.Fatalf("ParseSoc(%s)[%d].Header.Type = '%s' but got '%s' instead", d.Data, iBlock, block.Header.Type, soc[iBlock].Header.Type)
			}
			if block.Header.Name != soc[iBlock].Header.Name {
				t.Fatalf("ParseSoc(%s)[%d].Header.Name = '%s' but got '%s' instead", d.Data, iBlock, block.Header.Name, soc[iBlock].Header.Name)
			}
			for kProp, vProp := range block.Properties {
				if vProp != soc[iBlock].Properties[kProp] {
					t.Fatalf("ParseSoc(%s)[%d].Properties[%s] = '%s' but got '%s' instead", d.Data, iBlock, kProp, vProp, soc[iBlock].Properties[vProp])
				}
			}
		}
	}
}
