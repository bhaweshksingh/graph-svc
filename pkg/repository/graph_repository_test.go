package repository

//
//func setUp() (*gorm.DB, context.Context) {
//	ctx := context.Background()
//	dbConn.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.EventSnapshot{})
//	dbConn.WithContext(ctx).Where("user_id = ?", userId).Delete(&model.ChartData{})
//	return dbConn, ctx
//}
//func TestGormgraphRepository_GetAnswer(t *testing.T) {
//	dbConn, ctx := setUp()
//	expectedEvent := model.EventSnapshot{"name", "john", userId}
//	dbConn.WithContext(ctx).Create(expectedEvent)
//	repository := NewgraphRepository(dbConn)
//
//	eventSnapshot, err := repository.GetAnswer(ctx, &dto.EventQuery{"name", userId})
//
//	assertions.So(err, assertions.ShouldBeNil)
//	assertions.ShouldEqual(eventSnapshot, expectedEvent)
//	checkForEmptyHistory(dbConn, ctx)
//}
//
//func TestGormgraphRepository_GetAnswer_fails(t *testing.T) {
//	dbConn, ctx := setUp()
//	repository := NewgraphRepository(dbConn)
//
//	eventSnapshot, err := repository.GetAnswer(ctx, &dto.EventQuery{"name", userId})
//
//	assertions.So(eventSnapshot, assertions.ShouldBeNil)
//	assertions.ShouldContain(err.Error(), "record not found")
//}
//
