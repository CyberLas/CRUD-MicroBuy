// FindByUser - ...
func FindByUser(idUser int) ([]model.Shopping, error) {
	var shoppings []model.Shopping
	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shoppings, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"user": idUser}).All(&shoppings)
	return shoppings, err
}