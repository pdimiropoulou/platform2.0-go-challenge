package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"platform2.0-go-challenge/controllers"
)

var DB *gorm.DB
var err error

// func TestGetAssets(t *testing.T) {
// 	controller := controllers.Controller{}
// 	//cannot read .env
// 	DB, err = gorm.Open(postgres.Open("postgres://baprmzsc:O4Fw-Oe0oNgarBmobmVmmB8vt1cuG9lj@tai.db.elephantsql.com/baprmzsc"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

// 	req, err := http.NewRequest("GET", "/api/assets/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(controller.GetUserAssets(DB))
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Check the response body is the expected
// 	expected := `{"ID": 1, "Description": "Movies Chart Q12021", "AssetTypeID": 1, "UserID": 1, "Favourite": true, "ExternalID": 1}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

func TestGetRequest(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/api/assets/1", nil)
	w := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"mystring": "abcd",
	}

	// CHANGE THIS LINE!!!
	r = mux.SetURLVars(r, vars)
	DB, err = gorm.Open(postgres.Open("postgres://baprmzsc:O4Fw-Oe0oNgarBmobmVmmB8vt1cuG9lj@tai.db.elephantsql.com/baprmzsc"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	controller := controllers.Controller{}
	controller.GetUserAssets(DB)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte("abcd"), w.Body.Bytes())
}
