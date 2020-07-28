// FindByID - ...
func FindByID(id string) (model.Shopping, error) {
	var shopping model.Shopping
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return shopping, err
	}

	session, err := mgo.DialWithInfo(INFO)
	if err != nil {
		log.Fatal(err)
		return shopping, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&shopping)
	return shopping, err
}