package user

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)


func setupTestDB(t *testing.T)(*sql.DB, func()){
	t.Helper();

	pool, err := dockertest.NewPool(``);

	require.NoError(t, err);

	resource, err := pool.Run(`postgres`, `15-alpine`, []string{
		"POSTGRES_USER=test",
		"POSTGRES_PASSWORD=test",
		"POSTGRES_DB=testdb",
	});

	require.NoError(t, err);
	resource.Expire(120);

	var db *sql.DB;
	require.NoError(t, pool.Retry(func() error{
		var err error;
		dsn := fmt.Sprintf("postgres://test:test@localhost:%s/testdb?sslmode=disable", resource.GetPort("5432/tcp"))
		db, err = sql.Open("postgres", dsn);

		if err != nil {
			return err
		}

		return db.Ping();
	}))

	cleanup := func(){
		if db != nil{
			db.Close();
		}

		if err := pool.Purge(resource); err != nil{
			fmt.Println(`Failed to purge resource`, err.Error());
		}
	}

	return db, cleanup;
}


func TestUserRepo_Integration(t *testing.T){

	db, cleanup := setupTestDB(t);
	defer cleanup();

	repo := NewRepo(db);

	require.NoError(t, repo.CreateTable());

	id, err := repo.InsertUser(`Alice`);

	require.NoError(t, err);
	require.NotZero(t, id);


	user, err := repo.GetUser(id);

	require.NoError(t, err);
	require.Equal(t, `Alice`, user.Name);

}
