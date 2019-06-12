//+build wireinject

package configs

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/cmtablecontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/rctcontroller"
	"outstagram/server/controllers/usercontroller"
	"outstagram/server/db"
	"outstagram/server/repos/cmtablerepo"
	"outstagram/server/repos/cmtrepo"
	"outstagram/server/repos/imgrepo"
	"outstagram/server/repos/notifbrepo"
	"outstagram/server/repos/postimgrepo"
	"outstagram/server/repos/postrepo"
	"outstagram/server/repos/rctablerepo"
	"outstagram/server/repos/rctrepo"
	"outstagram/server/repos/replyrepo"
	"outstagram/server/repos/storybrepo"
	"outstagram/server/repos/userrepo"
	"outstagram/server/repos/vwablerepo"
	"outstagram/server/services/cmtableservice"
	"outstagram/server/services/cmtservice"
	"outstagram/server/services/imgservice"
	"outstagram/server/services/notifbservice"
	"outstagram/server/services/postimgservice"
	"outstagram/server/services/postservice"
	"outstagram/server/services/rctableservice"
	"outstagram/server/services/rctservice"
	"outstagram/server/services/replyservice"
	"outstagram/server/services/storybservice"
	"outstagram/server/services/userservice"
	"outstagram/server/services/vwableservice"

	"github.com/google/wire"
)

func InitializeUserController() (*usercontroller.Controller, error) {
	wire.Build(
		usercontroller.New,

		userservice.New,
		userrepo.New,

		db.New)
	return &usercontroller.Controller{}, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	wire.Build(
		authcontroller.New,

		userservice.New,
		userrepo.New,

		notifbservice.New,
		notifbrepo.New,

		storybservice.New,
		storybrepo.New,

		db.New)
	return &authcontroller.Controller{}, nil
}

func InitializePostController() (*postcontroller.Controller, error) {
	wire.Build(
		postcontroller.New,

		vwablerepo.New,
		vwableservice.New,

		userservice.New
		,

		postservice.New
		,

		postimgservice.New
		,

		imgservice.New,
		imgrepo.New,

		cmtableservice.New,
		cmtablerepo.New,

		cmtservice.New,
		cmtrepo.New,

		replyservice.New,
		replyrepo.New,

		rctableservice.New,
		rctablerepo.New,

		db.New)
	return &postcontroller.Controller{}, nil
}

func InitializeReactController() (*rctcontroller.Controller, error) {
	wire.Build(
		rctcontroller.New,

		rctservice.New,
		rctrepo.New,

		db.New)

	return &rctcontroller.Controller{}, nil
}

func InitializeCommentableController() (*cmtablecontroller.Controller, error) {
	wire.Build(
		cmtablecontroller.New,

		cmtableservice.New,
		cmtablerepo.New,

		cmtservice.New,
		cmtrepo.New,

		userservice.New,
		userrepo.New,

		rctableservice.New,
		rctablerepo.New,

		replyservice.New,
		replyrepo.New,

		db.New)

	return &cmtablecontroller.Controller{}, nil
}
