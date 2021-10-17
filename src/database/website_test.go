package database_test

import (
	"testing"

	"code.thetadev.de/ThetaDev/gotry/try"
	"github.com/Theta-Dev/Talon/src/database"
	"github.com/Theta-Dev/Talon/src/fixtures/testdb"
	"github.com/stretchr/testify/assert"
)

func TestWebsiteAdd(t *testing.T) {
	db := testdb.Open()

	t.Run("ok", func(t *testing.T) {
		ws := &database.Website{
			Name:        "KeepTalking",
			Path:        "Test/KeepTalking/",
			Logo:        testdb.Files[13],
			Color:       "#0000cc",
			Visibility:  database.VISIBILITY_SEARCHABLE,
			UserID:      2,
			MaxVersions: 3,
			SourceUrl:   "https://github.com/Theta-Dev/KeepTalkingBomb",
			SourceType:  "github",
		}
		try.Check(db.WebsiteAdd(ws))

		id := len(testdb.Websites) + 1
		assert.EqualValues(t, id, ws.ID)

		gotWs := try.X(db.WebsiteByID(uint(id), true)).(*database.Website)

		assert.Equal(t, "KeepTalking", gotWs.Name)
		assert.Equal(t, "Test/KeepTalking", gotWs.Path)
		assert.Equal(t, "test/keeptalking", gotWs.PathLower)
		assert.Equal(t, testdb.Files[13].Hash, gotWs.Logo.Hash)
		assert.Equal(t, "#0000cc", gotWs.Color)
		assert.Equal(t, database.VISIBILITY_SEARCHABLE, gotWs.Visibility)
		assert.Equal(t, "Zoey", gotWs.User.Name)
		assert.Equal(t, 3, gotWs.MaxVersions)
		assert.Equal(t, "https://github.com/Theta-Dev/KeepTalkingBomb", gotWs.SourceUrl)
		assert.Equal(t, "github", gotWs.SourceType)
	})

	t.Run("duplicate", func(t *testing.T) {
		ws := &database.Website{
			Name:   "KeepTalking2",
			Path:   "Test/KeepTalking/",
			UserID: 2,
		}

		err := db.WebsiteAdd(ws)
		assert.ErrorIs(t, err, database.ErrSitePathAlreadyExists)
	})

	t.Run("no_user", func(t *testing.T) {
		ws := &database.Website{
			Name: "KeepTalking3",
			Path: "KeepTalking",
		}

		err := db.WebsiteAdd(ws)
		assert.ErrorIs(t, err, database.ErrEmptyUser)
	})
}

func TestWebsiteUpdate(t *testing.T) {
	db := testdb.Open()

	ws := try.X(db.WebsiteByID(1, true)).(*database.Website)
	ws.Name = "KeepTalking"
	try.Check(db.WebsiteUpdate(ws))

	gotWs := try.X(db.WebsiteByID(1, true)).(*database.Website)
	assert.Equal(t, "KeepTalking", gotWs.Name)
}

func TestWebsiteByID(t *testing.T) {
	db := testdb.Open()

	t.Run("found", func(t *testing.T) {
		ws := try.X(db.WebsiteByID(1, false)).(*database.Website)

		assert.EqualValues(t, 1, ws.ID)
		assert.Equal(t, "ThetaDev", ws.Name)
	})

	t.Run("not_found", func(t *testing.T) {
		noWs := try.X(db.WebsiteByID(0, false)).(*database.Website)
		assert.Nil(t, noWs)
	})
}

func TestWebsiteByPath(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name   string
		path   string
		expect string
	}{
		{
			name:   "ThetaDev",
			path:   "/",
			expect: "ThetaDev",
		},
		{
			name:   "Talon",
			path:   "/tALOn/",
			expect: "Talon",
		},
		{
			name:   "Spotify-Gender-Ex",
			path:   "spotify-gender-ex",
			expect: "Spotify-Gender-Ex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := try.X(db.WebsiteByPath(tt.path, true)).(*database.Website)
			assert.Equal(t, tt.expect, ws.Name)
			assert.NotNil(t, ws.User)
			assert.NotEmpty(t, ws.Versions)
		})
	}

	t.Run("not_found", func(t *testing.T) {
		ws := try.X(db.WebsiteByPath("XYZ", false)).(*database.Website)
		assert.Nil(t, ws)
	})
}

func TestWebsitePathExists(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name   string
		path   string
		expect bool
	}{
		{
			name:   "ThetaDev",
			path:   "/",
			expect: true,
		},
		{
			name:   "Talon",
			path:   "/tALOn/",
			expect: true,
		},
		{
			name:   "XYZ",
			path:   "/XYZ",
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := try.Bool(db.WebsitePathExists(tt.path))
			assert.Equal(t, tt.expect, exists)
		})
	}
}

func TestWebsitesGet(t *testing.T) {
	db := testdb.Open()

	expectedVersionsMap := [][]int{
		{0, 1},
		{2},
		{3},
		{},
	}

	tests := []struct {
		name   string
		query  []interface{}
		expect []int
	}{
		{
			name:   "all",
			expect: []int{0, 1, 2, 3},
		},
		{
			name:   "by_name",
			query:  []interface{}{"websites.name = ?", "ThetaDev"},
			expect: []int{0},
		},
		{
			name:   "by_visibility",
			query:  []interface{}{"visibility >= ?", database.VISIBILITY_VISIBLE},
			expect: []int{0, 1},
		},
		{
			name:   "none",
			query:  []interface{}{"websites.id = 0"},
			expect: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sites := try.X(db.WebsitesGet(true, tt.query...)).([]*database.Website)

			for _, s := range sites {
				i := s.ID - 1
				expSite := testdb.Websites[tt.expect[i]]
				expVersions := expectedVersionsMap[tt.expect[i]]

				assert.Equal(t, expSite.Name, s.Name)

				for i, v := range s.Versions {
					assert.Equal(t, testdb.Versions[expVersions[i]].Name, v.Name)
				}
			}
		})
	}
}

func TestWebsitesCount(t *testing.T) {
	db := testdb.Open()

	tests := []struct {
		name   string
		query  []interface{}
		expect int
	}{
		{
			name:   "all",
			query:  []interface{}{},
			expect: 4,
		},
		{
			name:   "by_name",
			query:  []interface{}{"websites.name = ?", "ThetaDev"},
			expect: 1,
		},
		{
			name:   "by_visibility",
			query:  []interface{}{"visibility >= ?", database.VISIBILITY_VISIBLE},
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := try.Int(db.WebsitesCount(tt.query...))
			assert.Equal(t, tt.expect, count)
		})
	}
}

func TestWebsiteDeleteByID(t *testing.T) {
	db := testdb.Open()

	try.Check(db.WebsiteDeleteByID(2))

	gotSite := try.X(db.WebsiteByID(2, false)).(*database.Website)
	assert.Nil(t, gotSite)
}
