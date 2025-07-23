package integration

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestIntegrationAuth(t *testing.T) {
	config := setupTestConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool := setupTestDB(t, config.Database, ctx)
	t.Cleanup(func() {
		pool.Close()
	})

	server := setupTestServer(t, pool, *config)
	t.Cleanup(func() {
		server.Close()
	})
	t.Run("successful registration", func(t *testing.T) {
		_, err := pool.Exec(context.Background(), `TRUNCATE TABLE users RESTART IDENTITY CASCADE`)
		assert.NoError(t, err)
		resp, err := http.Post(server.URL+"/auth/register", "application/json", strings.NewReader(`{"email":"a.starikova@innopolis.university","password":"lexandrinnn_t", "role_id": 1}`))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var count int
		err = pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email = 'a.starikova@innopolis.university'").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})
	t.Run("invalid mail", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/auth/register", "application/json", strings.NewReader(`{"email":"feduardo2006@gmail.com","password":"1234567890", "role_id": 1}`))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var count int
		err = pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email = 'feduardo2006@gmail.com'").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})
	t.Run("pass too short", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/auth/register", "application/json", strings.NewReader(`{"email":"f.markin@innopolis.university","password":"bb4", "role_id": 1}`))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var count int
		err = pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email = 'f.markin@innopolis.university'").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})
	t.Run("pass too long", func(t *testing.T) {
		resp, err := http.Post(server.URL+"/auth/register", "application/json", strings.NewReader(`{"email":"m.kazakov@innopolis.university","password":"qwertyuiop[]asdfghjkl;'zxcvbnm,./1234567890", "role_id": 1}`))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		var count int
		err = pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email = 'm.kazakov@innopolis.university'").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 0, count)

	})
	t.Run("user exist", func(t *testing.T) {
		_, err := pool.Exec(context.Background(), `TRUNCATE TABLE users RESTART IDENTITY CASCADE`)
		assert.NoError(t, err)
		passhash, _ := bcrypt.GenerateFromPassword([]byte("AtYourDisposal"), config.Sequrity.Bcrypt.Cost)
		id := uuid.New().String()
		query := `INSERT INTO users (id, email, password_hash, role_id) VALUES ($1, $2, $3,$4)`
		_, err = pool.Exec(context.Background(), query, id, "p.balandin@innopolis.university", passhash, 1)
		assert.NoError(t, err)
		resp, _ := http.Post(server.URL+"/auth/register", "application/json", strings.NewReader(`{"email":"p.balandin@innopolis.university","password":"AtYourDisposal", "role_id":1}`))
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		var count int
		err = pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email = 'p.balandin@innopolis.university'").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})
	t.Run("succesful login", func(t *testing.T) {
		_, err := pool.Exec(context.Background(), `TRUNCATE TABLE users RESTART IDENTITY CASCADE`)
		assert.NoError(t, err)
		passhash, _ := bcrypt.GenerateFromPassword([]byte("ApproachTomorrow"), config.Sequrity.Bcrypt.Cost)
		id := uuid.New().String()
		query := `INSERT INTO users (id, email, password_hash, role_id) VALUES ($1, $2, $3,$4)`
		_, err = pool.Exec(context.Background(), query, id, "i.vavilov@innopolis.university", passhash, 1)
		assert.NoError(t, err)
		resp, err := http.Post(server.URL+"/auth/login", "application/json", strings.NewReader(`{"email":"i.vavilov@innopolis.university","password":"ApproachTomorrow"}`))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotEmpty(t, resp.Cookies(), "Should set tokens cookie")
	})
}
