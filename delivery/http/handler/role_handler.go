package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solkn/restaurant_api/entity"
	"github.com/solkn/restaurant_api/user"
	"github.com/solkn/restaurant_api/utils"
	"net/http"
	"strconv"
)

type RoleHandler struct {
	roleService user.RoleService
}
func NewRoleHandler(roleServ user.RoleService)*RoleHandler{
	return &RoleHandler{roleService: roleServ}
}
func(rh *RoleHandler)GetRole(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	role,errs:= rh.roleService.Role(uint(id))
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(role,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(rh *RoleHandler)GetRoles(w http.ResponseWriter,r *http.Request){
	role,errs:= rh.roleService.Roles()
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(role,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(rh *RoleHandler)PostRole(w http.ResponseWriter,r *http.Request){
	body:= utils.BodyParser(r)
	role:= entity.Role{}
	err:= json.Unmarshal(body,&role)
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	storedRole,errs:= rh.roleService.StoreRole(&role)
	if len(errs)>0{
		 w.Header().Set("content-type","application/json")
		 http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(storedRole,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(rh *RoleHandler)UpdateRole(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	role,errs:= rh.roleService.Role(uint(id))
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	body:= utils.BodyParser(r)
	err = json.Unmarshal(body,&role)
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	role,errs = rh.roleService.UpdateRole(role)
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:= json.MarshalIndent(role,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(rh *RoleHandler)DeleteRole(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	_,errs:= rh.roleService.DeleteRole(uint(id))
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
}

