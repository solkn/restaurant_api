package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/solkn/restaurant_api/comment"
	"github.com/solkn/restaurant_api/entity"
	"github.com/solkn/restaurant_api/utils"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	commentService comment.CommentService
}
func NewCommentHandler(commentServ comment.CommentService)*CommentHandler{
	return &CommentHandler{commentService: commentServ}
}

func(ch *CommentHandler)GetComment(w http.ResponseWriter, r *http.Request){
	params:= mux.Vars(r)

	id,err := strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)

		return
	}
	comment,errs:= ch.commentService.Comment(uint(id))

	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)

		return
	}
	output,err:= json.MarshalIndent(comment,"","\t\t")

	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return
}
func(ch *CommentHandler)GetComments(w http.ResponseWriter,r *http.Request){

	comments,errs:= ch.commentService.Comments()

	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err:=json.MarshalIndent(comments,"","\t\t")
	if err!= nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write(output)
	return

}
func(ch *CommentHandler)PostComment(w http.ResponseWriter,r *http.Request){
  body:=utils.BodyParser(r)

   comment := entity.Comment{}

  err:= json.Unmarshal(body,&comment)
  if err != nil{
  	w.Header().Set("content-type","application/json")
  	http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
	  return
  }
  StoredComment,errs:= ch.commentService.StoreComment(&comment)

  if len(errs)>0{

  	   w.Header().Set("content-type","application/json")
  	   http.Error(w,http.StatusText(http.StatusNotFound),404)
	  return
  }
  output,err:= json.MarshalIndent(&StoredComment,"","\t\t")

  if err != nil{

	  w.Header().Set("content-type","application/json")
	  http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
	  return
  }

  w.Header().Set("content-type","application/json")
  w.Write(output)



}

func(ch *CommentHandler)UpdateComment(w http.ResponseWriter,r *http.Request){

	params:= mux.Vars(r)

	id,err:= strconv.Atoi(params["id"])

	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	comment,errs:= ch.commentService.Comment(uint(id))

	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	body:=utils.BodyParser(r)

	err = json.Unmarshal(body,&comment)

	if err!= nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	comment,errs = ch.commentService.UpdateComment(comment)
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	output,err := json.MarshalIndent(comment,"","\t\t")

	w.Header().Set("content-type","application/json")
	w.Write(output)

	return


}

func(ch *CommentHandler)DeleteComment(w http.ResponseWriter,r *http.Request){
	params:= mux.Vars(r)

	id,err:=strconv.Atoi(params["id"])
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusUnprocessableEntity),http.StatusUnprocessableEntity)
		return
	}
	_,errs:= ch.commentService.DeleteComment(uint(id))

	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}

}