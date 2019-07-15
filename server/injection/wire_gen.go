// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injection

import (
	"outstagram/server/controllers/authcontroller"
	"outstagram/server/controllers/cmtablecontroller"
	"outstagram/server/controllers/flcontroller"
	"outstagram/server/controllers/imgcontroller"
	"outstagram/server/controllers/mecontroller"
	"outstagram/server/controllers/postcontroller"
	"outstagram/server/controllers/rctcontroller"
	"outstagram/server/controllers/storycontroller"
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
)

// Injectors from main.go:

func InitializeUserController() (*usercontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	storyBoardRepo := storybrepo.New(gormDB)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	storyBoardService := storybservice.New(storyBoardRepo, userService, reactableService)
	postRepo := postrepo.New(gormDB)
	commentableService := cmtableservice.New(commentableRepo, reactableService)
	postService := postservice.New(postRepo, userService, reactableService, commentableService)
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	postImageRepo := postimgrepo.New(gormDB)
	postImageService := postimgservice.New(postImageRepo, reactableService, commentableService)
	viewableRepo := vwablerepo.New(gormDB)
	viewableService := vwableservice.New(viewableRepo)
	controller := usercontroller.New(userService, storyBoardService, postService, imageService, postImageService, viewableService)
	return controller, nil
}

func InitializeAuthController() (*authcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	notifBoardRepo := notifbrepo.New(gormDB)
	notifBoardService := notifbservice.New(notifBoardRepo)
	storyBoardRepo := storybrepo.New(gormDB)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	storyBoardService := storybservice.New(storyBoardRepo, userService, reactableService)
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	controller := authcontroller.New(userService, notifBoardService, storyBoardService, imageService)
	return controller, nil
}

func InitializePostController() (*postcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	postRepo := postrepo.New(gormDB)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	commentableService := cmtableservice.New(commentableRepo, reactableService)
	postService := postservice.New(postRepo, userService, reactableService, commentableService)
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	postImageRepo := postimgrepo.New(gormDB)
	postImageService := postimgservice.New(postImageRepo, reactableService, commentableService)
	viewableRepo := vwablerepo.New(gormDB)
	viewableService := vwableservice.New(viewableRepo)
	controller := postcontroller.New(postService, imageService, postImageService, viewableService)
	return controller, nil
}

func InitializeReactController() (*rctcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	reactRepo := rctrepo.New(gormDB)
	reactService := rctservice.New(reactRepo)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	controller := rctcontroller.New(reactService, reactableService, userService)
	return controller, nil
}

func InitializeCommentableController() (*cmtablecontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	commentableService := cmtableservice.New(commentableRepo, reactableService)
	commentRepo := cmtrepo.New(gormDB)
	commentService := cmtservice.New(commentRepo, reactableService)
	replyRepo := replyrepo.New(gormDB)
	replyService := replyservice.New(replyRepo)
	controller := cmtablecontroller.New(commentableService, commentService, userService, reactableService, replyService)
	return controller, nil
}

func InitializeMeController() (*mecontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	postRepo := postrepo.New(gormDB)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	commentableService := cmtableservice.New(commentableRepo, reactableService)
	postService := postservice.New(postRepo, userService, reactableService, commentableService)
	storyBoardRepo := storybrepo.New(gormDB)
	storyBoardService := storybservice.New(storyBoardRepo, userService, reactableService)
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	controller := mecontroller.New(userService, postService, storyBoardService, imageService)
	return controller, nil
}

func InitializeFollowController() (*flcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	controller := flcontroller.New(userService)
	return controller, nil
}

func InitializeStoryController() (*storycontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	viewableRepo := vwablerepo.New(gormDB)
	viewableService := vwableservice.New(viewableRepo)
	storyBoardRepo := storybrepo.New(gormDB)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	commentableRepo := cmtablerepo.New(gormDB)
	reactableRepo := rctablerepo.New(gormDB, commentableRepo)
	reactableService := rctableservice.New(reactableRepo, userService)
	storyBoardService := storybservice.New(storyBoardRepo, userService, reactableService)
	controller := storycontroller.New(imageService, viewableService, storyBoardService, userService)
	return controller, nil
}

func InitializeImageController() (*imgcontroller.Controller, error) {
	gormDB, err := db.New()
	if err != nil {
		return nil, err
	}
	imageRepo := imgrepo.New(gormDB)
	imageService := imgservice.New(imageRepo)
	userRepo := userrepo.New(gormDB)
	userService := userservice.New(userRepo)
	controller := imgcontroller.New(imageService, userService)
	return controller, nil
}
