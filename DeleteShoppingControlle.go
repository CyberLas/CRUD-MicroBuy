//DeleteShoppingController - Borrar una compra
func DeleteShoppingController(w http.ResponseWriter, r http.Request) {
	defer r.Body.Close()
	var shoppingID model.ShoppingID
	if err = json.NewDecoder(r.Body).Decode(&shoppingID); err != nil {
		respondWithError(w, http.StatusBadRequest, Invalid request)
		return
	}
	if err = connection.Delete(shoppingID.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{result success})
}