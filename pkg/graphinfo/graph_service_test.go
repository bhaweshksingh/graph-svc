package graphinfo

//
//func TestGormgraphRepository_GetAnswer(t *testing.T) {
//	ctx := context.Background()
//	eventSnapshot := model.EventSnapshot{"name", "john", userId}
//	repositoryMock := mock.graphRepositoryMock{
//		GetAnswerFunc: func(ctx context.Context, eventQuery *dto.EventQuery) (*model.EventSnapshot, error) {
//			return &eventSnapshot, nil
//		}}
//
//	service := NewGraphService(&repositoryMock)
//	event := dto.EventQuery{"name", userId}
//
//	actualSnapshot, err := service.GetAnswer(ctx, &event)
//
//	assertions.So(err, assertions.ShouldBeNil)
//	assertions.ShouldEqual(actualSnapshot, eventSnapshot)
//}
//
//func TestGormgraphRepository_GetAnswer_fails(t *testing.T) {
//	ctx := context.Background()
//	mockError := errors.New("failed to get")
//	repositoryMock := mock.graphRepositoryMock{
//		GetAnswerFunc: func(ctx context.Context, eventQuery *dto.EventQuery) (*model.EventSnapshot, error) {
//			return nil, mockError
//		}}
//
//	service := NewGraphService(&repositoryMock)
//	event := dto.EventQuery{"name", userId}
//
//	actualSnapshot, err := service.GetAnswer(ctx, &event)
//
//	assertions.So(actualSnapshot, assertions.ShouldBeNil)
//	assertions.ShouldContain(err.Error(), mockError.Error())
//}
