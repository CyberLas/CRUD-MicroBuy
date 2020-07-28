// FindShoppingByIDController - Encuentra una compra por su ID
func FindShoppingByIDController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shopping, err := connection.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, shopping)
}