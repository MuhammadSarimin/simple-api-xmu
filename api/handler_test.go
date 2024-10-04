package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/api"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFindAll(t *testing.T) {

	store := &mockStore{}
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.Use(api.BasicAuth())
	api.NewHandler(v1, store)

	t.Run("Get all movies success", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/api/v1/movies", nil)

		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "200", body["response_code"])
		assert.Equal(t, "Success", body["response_message"])
	})

	t.Run("Get all movies fail", func(t *testing.T) {

		//dont have auth
		req, err := http.NewRequest("GET", "/api/v1/movies", nil)

		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "401", body["response_code"])
		assert.Equal(t, "Unauthorized", body["response_message"])
	})

	t.Run("Create movie success", func(t *testing.T) {

		payload := types.Movie{
			Title:  "Spiderman",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/movies", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Create movie failed", func(t *testing.T) {

		payload := types.Movie{
			Title:  "",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/v1/movies", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "001", body["response_code"])
		assert.Equal(t, "title is required", body["response_message"])
	})

	t.Run("Update movie success", func(t *testing.T) {

		payload := types.Movie{
			Title:  "Spiderman",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("PATCH", "/api/v1/movies/1", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Update movie failed", func(t *testing.T) {

		payload := types.Movie{
			Title:  "",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("PATCH", "/api/v1/movies/1", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 400, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "001", body["response_code"])
		assert.Equal(t, "title is required", body["response_message"])
	})

	t.Run("Find detail movie success", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/api/v1/movies/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Find detail movie failed", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/api/v1/movies/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "404", body["response_code"])
		assert.Equal(t, "movie not found", body["response_message"])
	})

	t.Run("Delete movie success", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", "/api/v1/movies/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "200", body["response_code"])
		assert.Equal(t, "Success", body["response_message"])
	})

	t.Run("Delete movie failed", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", "/api/v1/movies/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.SetBasicAuth("simple-api", "xmu")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 404, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "404", body["response_code"])
		assert.Equal(t, "movie not found", body["response_message"])
	})

	t.Run("Delete movie failed auth", func(t *testing.T) {

		req, err := http.NewRequest("DELETE", "/api/v1/movies/2", nil)
		if err != nil {
			t.Fatal(err)
		}

		// wrong password
		req.SetBasicAuth("simple-api", "xmux")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		assert.Equal(t, 401, w.Code)

		var body map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &body)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "401", body["response_code"])
		assert.Equal(t, "Unauthorized", body["response_message"])
	})
}

type mockStore struct{}

func (s *mockStore) FindAll() ([]types.Movie, error) {
	return []types.Movie{
		{
			Title:  "Spiderman",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		},
	}, nil
}

func (s *mockStore) Create(m *types.Movie) error {
	return nil
}
func (s *mockStore) FindByID(id int) (*types.Movie, error) {
	if id == 1 {
		return &types.Movie{
			Title:  "Spiderman",
			Rating: 4.5,
			Image:  "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg",
		}, nil
	}

	return nil, gorm.ErrRecordNotFound
}
func (s *mockStore) Update(m *types.Movie) error {
	return nil
}
func (s *mockStore) Delete(id int) error {

	if id == 1 {
		return nil
	}

	return gorm.ErrRecordNotFound
}
