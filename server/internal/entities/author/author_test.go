package author_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/bensaufley/aud-it/internal/db"
	"github.com/bensaufley/aud-it/internal/entities/author"
	"github.com/bensaufley/aud-it/internal/entities/base"

	"github.com/stretchr/testify/assert"
)

var ulidRegExp = regexp.MustCompile(`^[0123456789ABCDEFGHJKMNPQRSTVWXYZ]{26}$`)

func TestCreate(t *testing.T) {
	var database *sql.DB
	testCases := []struct {
		it string

		setup func() error

		authors []*author.Author

		assert func(*testing.T, error)
	}{
		{
			it: "adds new authors",

			authors: []*author.Author{
				author.New("Anders", "Charlie Jane"),
				author.New("Gailey", "Sarah"),
				author.New("Unknown"),
			},

			assert: func(test *testing.T, err error) {
				if !assert.NoError(test, err) {
					return
				}
				rows, err := database.Query(`SELECT ulid, first_name, last_name FROM authors WHERE last_name IN ('Anders', 'Gailey', 'Unknown')`)
				if !assert.NoError(test, err) {
					return
				}
				count := 0
				for rows.Next() {
					a := author.New()
					if err := rows.Scan(&a.ULID, &a.FirstName, &a.LastName); err != nil {
						assert.FailNowf(test, err.Error(), "error scanning row")
					}

					count += 1
					assert.Regexp(test, ulidRegExp, a.ULID)
					switch a.LastName {
					case "Anders":
						assert.Equal(test, base.StrPointer("Charlie Jane"), a.FirstName)
					case "Gailey":
						assert.Equal(test, base.StrPointer("Sarah"), a.FirstName)
					case "Unknown":
						assert.Nil(test, a.FirstName)
					default:
						assert.FailNowf(test, "unexpected last name value", "value was %s", a.LastName)
					}
				}
				assert.Equal(test, 3, count)
			},
		},
	}

	dir := t.TempDir()
	dbCfg := &db.Config{
		DBPath:         dir + "/test.db",
		MigrationsPath: "../../../migrations",
	}
	var err error
	database, err = dbCfg.Get()
	if err != nil {
		assert.FailNowf(t, err.Error(), "could not get database")
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc

			g.Run(testCase.it, func(test *testing.T) {
				if testCase.setup != nil {
					if err := testCase.setup(); err != nil {
						assert.FailNowf(test, err.Error(), "could not run test setup")
					}
				}

				tx, err := database.BeginTx(context.Background(), &sql.TxOptions{})
				if err != nil {
					assert.FailNowf(test, err.Error(), "could not begin transaction")
				}
				err = author.Create(context.Background(), tx, testCase.authors...)
				if err != nil {
					tx.Rollback()
					testCase.assert(test, err)
				}

				err = tx.Commit()
				testCase.assert(test, err)
			})
		}
	})
}
