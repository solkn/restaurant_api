
package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/solkn/restaurant_api/authentication"
	"github.com/solkn/restaurant_api/entity"
	"github.com/solkn/restaurant_api/user"
	"github.com/solkn/restaurant_api/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userServices user.UserService) *UserHandler {
	return &UserHandler{userService: userServices}
}
func (uh *UserHandler) Authenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authentication.TokenValid(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
func (uh *UserHandler) Authorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userX entity.User
		uid, err := authentication.ExtractTokenID(r)
		userX.ID = uid
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, errs := uh.userService.User(uid)
		fmt.Println(user.Roles)
		if len(errs)>0 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if strings.ToUpper(user.Roles[0].Name) != strings.ToUpper("ADMIN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		next(w, r)
	}

}
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	users, errs := uh.userService.User(uint(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uh *UserHandler) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return

	}
	userRoles, errs := uh.userService.UserRoles(&user1)
	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(userRoles, "", "\t\t")

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uh *UserHandler) IsEmailExists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["email"]

	ok := uh.userService.EmailExists(email)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uh *UserHandler) IsPhoneExists(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	phone := params["phone"]

	ok := uh.userService.PhoneExists(phone)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, errs := uh.userService.Users()
	if errs != nil {
		fmt.Println("err1")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		fmt.Println("err2")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(output)
	if err != nil {
		fmt.Println("err3")
		fmt.Println(err.Error())
	}
	return

}
func (uh *UserHandler) GetUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request) {

	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	user2, errs := uh.userService.UserByEmail(user1.Email)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(*user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return

}
func (uh *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		print(err.Error())
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user1.Password), 12)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user1.Password = string(hashedPassword)
	user2, errs := uh.userService.StoreUser(&user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	output, err := json.MarshalIndent(user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--logging---")
	body := utils.BodyParser(r)
	var user1 entity.User

	var token string
	var expiry int64

	err := json.Unmarshal(body, &user1)
	if err != nil {
		fmt.Println("error 263")
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Println(user1)
	userFromDatabase, errs := uh.userService.UserByEmail(user1.Email)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	err = bcrypt.CompareHashAndPassword([]byte(user1.Password), []byte(userFromDatabase.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	token, err = authentication.CreateToken(userFromDatabase.ID)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	expiry = time.Now().Add(time.Minute * 30).Unix()

	output, err := json.MarshalIndent(userFromDatabase, "", "\t\t")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	expiryString := strconv.Itoa(int(expiry))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("token", token)
	w.Header().Add("expiry_date", expiryString)
	_, _ = w.Write(output)
	return
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var token string
	var expiry int64
	body := utils.BodyParser(r)
	var user1 entity.User
	err := json.Unmarshal(body, &user1)
	if err != nil {
		print(err.Error())
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user1.Password), 12)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user1.Password = string(hashedPassword)
	user2, errs := uh.userService.StoreUser(&user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}

	output, err := json.MarshalIndent(user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	token, err = authentication.CreateToken(user2.ID)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	expiry = time.Now().Add(time.Minute * 30).Unix()
	expiryString := strconv.Itoa(int(expiry))
	w.Header().Add("token", token)
	w.Header().Add("token", expiryString)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uh *UserHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user1, errs := uh.userService.User(uint(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength

	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	_ = json.Unmarshal(body, &user1)
	user1, errs = uh.userService.UpdateUser(user1)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(user1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uh.userService.DeleteUser(uint(id))

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
