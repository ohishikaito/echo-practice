package repository

// type (
// 	testHelper struct {
// 		T    *testing.T
// 		Ctrl *gomock.Controller
// 		sut  userRepo
// 		database.DB
// 	}
// )

// func createDataAccessorTestHelper(t *testing.T) *testHelper {
// 	ctrl := gomock.NewController(t)
// 	// DBのテーブル中身全部消す
// 	// insertする
// 	// テスト終了したらrollback!?
// 	// →依存関係あるテーブルあったら消せないからどうしようかな〜
// 	// テスト用DBを新しく立ち上げる？
// 	// ↓
// 	// sqlx で createDB して fileを読み込んでmigrationしてsetup?

// 	dba := config.dbProvider.GetDBAccessor()
// 	ctx := shared.NewMockAuthContext(ctrl)
// 	clock := shared.NewMockClock(ctrl)
// 	// sut := NewDataAccessor(ctx, clock, dba)
// 	sut := NewUserRepo(dba)

// 	return &testHelper{
// 		TestHelper: softiatest.TestHelper{T: t, Ctrl: ctrl},
// 		sut:        sut,
// 		dba:        dba,
// 		clock:      clock,
// 		ctx:        ctx,
// 	}
// }
