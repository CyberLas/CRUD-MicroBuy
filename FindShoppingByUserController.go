// FindShoppingByUserController - Encuentra una compra por su ID
func FindShoppingByUserController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, err := strconv.Atoi(params["id_user"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	shoppings, err := connection.FindByUser(idUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error")
		return
	}
	respondWithJSON(w, http.StatusOK, shoppings)
}