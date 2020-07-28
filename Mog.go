// Insert - Insert a Shopping
func Insert(shopping model.Shopping) error {
	session, err := mgo.DialWithInfo(INFO)
	defer session.Close()

	shopping.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(shopping)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}