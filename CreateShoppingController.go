// CreateShoppingController - Crear una compra
func CreateShoppingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping model.Shopping
	if err := json.NewDecoder(r.Body).Decode(&shopping); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	shopping.ID = bson.NewObjectId()
	if err := connection.Insert(shopping); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, shopping)
}