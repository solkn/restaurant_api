package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solkn/restaurant_api/entity"
	"github.com/solkn/restaurant_api/order"
	"github.com/solkn/restaurant_api/utils"
	"net/http"
	"strconv"
)

type  OrderHandler struct{
	orderService order.OrderService
}
func NewOrderHandler(orderServ order.OrderService)*OrderHandler{
	return &OrderHandler{orderService: orderServ}
}
func(oh *OrderHandler)GetOrder(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	order,errs:= oh.orderService.Order(uint(id))
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(order,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(oh *OrderHandler)GetOrders(w http.ResponseWriter,r *http.Request){
	order,errs:= oh.orderService.Orders()
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(order,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}
func(oh *OrderHandler)GetCustomerOrders(w http.ResponseWriter,r *http.Request){
	customerOrder:=entity.User{}
	order,errs:= oh.orderService.CustomerOrders(&customerOrder)
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(order,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}


func(oh *OrderHandler)PostOrder(w http.ResponseWriter,r *http.Request){
	body:= utils.BodyParser(r)

	order:= entity.Order{}
	err:= json.Unmarshal(body,&order)
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	storedOrder,errs:= oh.orderService.StoreOrder(&order)
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(storedOrder,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func(oh *OrderHandler)UpdateOrder(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)
	id,err:= strconv.Atoi(params["id"])

	order,errs:=oh.orderService.Order(uint(id))
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	body:= utils.BodyParser(r)

	err = json.Unmarshal(body,&order)
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	order,errs = oh.orderService.UpdateOrder(order)
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(order,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}

func (oh *OrderHandler)DeleteHandler(w http.ResponseWriter,r *http.Request){
	params:=mux.Vars(r)
	id,err:=strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	_,errs:= oh.orderService.DeleteOrder(uint(id))
	if len(errs)>0 {
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
}