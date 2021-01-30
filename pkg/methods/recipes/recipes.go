package recipes

import (
	"encoding/json"
	"net/http"

	"github.com/algocook/proto/recipes"
	"github.com/gorilla/mux"
)

// GetOne метод который возвращает информацию о рецепте
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	client, err := recipes.NewClient("5301")
	if err != nil {
		json.NewEncoder(w).Encode(recipes.RecipeResponse{
			Error: &recipes.Error{
				ErrorCode: 1,
				ErrorStr:  err.Error(),
			},
		})
		return
	}

	recipe := client.GetRecipe(&recipes.Recipe_ID{
		RecipeId: params["id"],
		OwnerId:  1,
	})
	json.NewEncoder(w).Encode(recipe)
}

// PostOne метод загрузки нового рецепта
func PostOne(w http.ResponseWriter, r *http.Request) {
	return
}

// DeleteOne удаляет из бд
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	return
}

// Search метод поиска по рецептам
func Search(w http.ResponseWriter, r *http.Request) {
	return
}
