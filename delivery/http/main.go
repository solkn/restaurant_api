package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/solkn/restaurant_api/delivery/http/handler"
	"github.com/solkn/restaurant_api/entity"

	ur "github.com/solkn/restaurant_api/user/repository"
	us "github.com/solkn/restaurant_api/user/service"

	cr "github.com/solkn/restaurant_api/comment/repository"
	cs "github.com/solkn/restaurant_api/comment/service"

	or "github.com/solkn/restaurant_api/order/repository"
	os "github.com/solkn/restaurant_api/order/service"

	mr "github.com/solkn/restaurant_api/menu/repository"
	ms "github.com/solkn/restaurant_api/menu/service"



)

func createTables(dbConn *gorm.DB) []error {
	dbConn.DropTableIfExists(&entity.Role{}, &entity.User{}, &entity.Comment{}, &entity.Order{},
		&entity.Item{},&entity.Category{},&entity.Ingredient{}).GetErrors()
	errs := dbConn.CreateTable(&entity.Role{}, &entity.User{}, &entity.Comment{},&entity.Order{},
		&entity.Item{},&entity.Category{},&entity.Ingredient{}).GetErrors()

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func main() {

	dbconn, err := gorm.Open(
		"postgres", "postgres://postgres:solomon@localhost/restaurant?sslmode=disable")
	if dbconn != nil {
		defer dbconn.Close()
	}
	if err != nil {
		panic(err)
	}

	createTables(dbconn)


	router:= mux.NewRouter()

	roleRepo:= ur.NewRoleGormRepo(dbconn)
	roleService:= us.NewRoleService(roleRepo)
	roleHandler:=handler.NewRoleHandler(roleService)


	userRepo:= ur.NewUserGormRepo(dbconn)
	userService:= us.NewUserService(userRepo)
	usersHandler:= handler.NewUserHandler(userService)

	router.HandleFunc("/api/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoles))).Methods("GET")
	//router.HandleFunc("/api/roles/{name}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByName))).Methods("GET")
	//router.HandleFunc("/api/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.GetRoleByID))).Methods("GET")
	router.HandleFunc("/api/role", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.PostRole))).Methods("POST")
	router.HandleFunc("/api/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.UpdateRole))).Methods("PUT")
	router.HandleFunc("/api/role/{id}", usersHandler.Authenticated(usersHandler.Authorized(roleHandler.DeleteRole))).Methods("DELETE")

	router.HandleFunc("/api/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/api/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/api/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/api/admin/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/api/admin/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.DeleteUser))).Methods("DELETE")
	router.HandleFunc("/api/admin/email/{email}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsEmailExists))).Methods("GET")
	router.HandleFunc("/api/admin/phone/{phone}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.IsPhoneExists))).Methods("GET")
	router.HandleFunc("/api/admin/check", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUserByUsernameAndPassword))).Methods("POST")

	router.HandleFunc("/api/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PostUser))).Methods("POST")
	router.HandleFunc("/api/user/users", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUsers))).Methods("GET")
	router.HandleFunc("/api/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.GetUser))).Methods("GET")
	router.HandleFunc("/api/user/users/{id}", usersHandler.Authenticated(usersHandler.Authorized(usersHandler.PutUser))).Methods("PUT")
	router.HandleFunc("/api/user/email/{email}", usersHandler.IsEmailExists).Methods("GET")
	router.HandleFunc("/api/user/phone/{phone}", usersHandler.IsPhoneExists).Methods("GET")
	router.HandleFunc("/api/user/check", usersHandler.GetUserByUsernameAndPassword).Methods("POST")
	router.HandleFunc("/api/user/login", usersHandler.Login).Methods("POST")
	router.HandleFunc("/api/user/signup", usersHandler.SignUp).Methods("POST")

	commentRepo:= cr.NewCommentGormRepo(dbconn)
	commentService:= cs.NewCommentService(commentRepo)
	commentHandler:= handler.NewCommentHandler(commentService)

	router.HandleFunc("/api/comment", commentHandler.GetComments).Methods("GET")
	router.HandleFunc("/api/comment/{id}", commentHandler.GetComment).Methods("GET")
	router.HandleFunc("/api/comment", usersHandler.Authenticated(usersHandler.Authorized(commentHandler.PostComment))).Methods("POST")
	router.HandleFunc("/api/comment/{id}", usersHandler.Authenticated(usersHandler.Authorized(commentHandler.UpdateComment))).Methods("PUT")
	router.HandleFunc("/api/comment/{id}", usersHandler.Authenticated(usersHandler.Authorized(commentHandler.DeleteComment))).Methods("DELETE")

	orderRepo:= or.NewOrderGormRepo(dbconn)
	orderService:= os.NewOrderService(orderRepo)
	orderHandler:= handler.NewOrderHandler(orderService)

	router.HandleFunc("/api/order", orderHandler.GetOrders).Methods("GET")
	router.HandleFunc("/api/order/{id}", orderHandler.GetOrder).Methods("GET")
	router.HandleFunc("/api/order", usersHandler.Authenticated(usersHandler.Authorized(orderHandler.PostOrder))).Methods("POST")
	router.HandleFunc("/api/order/{id}", usersHandler.Authenticated(usersHandler.Authorized(orderHandler.UpdateOrder))).Methods("PUT")
	router.HandleFunc("/api/order/{id}", usersHandler.Authenticated(usersHandler.Authorized(orderHandler.DeleteHandler))).Methods("DELETE")

	menuCategoryRepo:= mr.NewCategoryGormRepo(dbconn)
	menuCategoryService:= ms.NewCategoryService(menuCategoryRepo)
	menuItemRepo:= mr.NewItemGormRepo(dbconn)
	menuItemService:=ms.NewItemService(menuItemRepo)
	menuIngredientRepo:=mr.NewIngredientGormRepo(dbconn)
	menuIngredientService:=ms.NewIngredientService(menuIngredientRepo)
	menuHandler:= handler.NewMenuHandler(menuCategoryService,menuIngredientService,menuItemService)

	router.HandleFunc("/api/menu", menuHandler.GetMenu).Methods("GET")

	err = http.ListenAndServe("192.168.56.1:8080", router)

	if err != nil {
		panic(err)
	}
}