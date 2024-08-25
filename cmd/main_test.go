package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Strawhub/go-api/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックデータとモック関数の定義
type MockModel struct {
	mock.Mock
}

func (m *MockModel) GetAll() []model.Question {
	args := m.Called()
	return args.Get(0).([]model.Question)
}

func (m *MockModel) GetBy(tag, num string) []model.Question {
	args := m.Called(tag, num)
	return args.Get(0).([]model.Question)
}

func TestGetAll(t *testing.T) {
	// モックの設定
	mockModel := new(MockModel)
	mockModel.On("GetAll").Return([]model.Question{
		{ID: 1, Year: 2024, Genre: "Science", Question: "Goとは何ですか？", Answer: "プログラミング言語です。", Commentary: "Goの基本的な説明です。"},
		{ID: 2, Year: 2024, Genre: "Technology", Question: "Ginアプリケーションのテスト方法は？", Answer: "テスト用のフレームワークを使用します。", Commentary: "Ginフレームワークのテスト方法です。"},
	})

	// Ginのルーターをセットアップ
	r := gin.Default()

	// モックを使用するように設定
	r.GET("/questions", func(c *gin.Context) {
		questions := mockModel.GetAll()
		c.JSON(http.StatusOK, questions)
	})

	// テストリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/questions", nil)
	w := httptest.NewRecorder()

	// リクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードをアサート
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディをアサート
	expected := `[{"ID":1,"Year":2024,"Genre":"Science","Question":"Goとは何ですか？","Answer":"プログラミング言語です。","Commentary":"Goの基本的な説明です。"},{"ID":2,"Year":2024,"Genre":"Technology","Question":"Ginアプリケーションのテスト方法は？","Answer":"テスト用のフレームワークを使用します。","Commentary":"Ginフレームワークのテスト方法です。"}]`
	assert.JSONEq(t, expected, w.Body.String())
}

func TestGetBy(t *testing.T) {
	// モックの設定
	mockModel := new(MockModel)
	mockModel.On("GetBy", "genre", "Science").Return([]model.Question{
		{ID: 1, Year: 2024, Genre: "Science", Question: "Goとは何ですか？", Answer: "プログラミング言語です。", Commentary: "Goの基本的な説明です。"},
	})

	// Ginのルーターをセットアップ
	r := gin.Default()

	// モックを使用するように設定
	r.GET("/:tag/:num", func(c *gin.Context) {
		tag := c.Param("tag")
		num := c.Param("num")
		questions := mockModel.GetBy(tag, num)
		c.JSON(http.StatusOK, questions)
	})

	// テストリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/genre/Science", nil)
	w := httptest.NewRecorder()

	// リクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードをアサート
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディをアサート
	expected := `[{"ID":1,"Year":2024,"Genre":"Science","Question":"Goとは何ですか？","Answer":"プログラミング言語です。","Commentary":"Goの基本的な説明です。"}]`
	assert.JSONEq(t, expected, w.Body.String())
}
