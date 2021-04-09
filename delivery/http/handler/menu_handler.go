package handler

import (
	"encoding/json"
	"github.com/solkn/restaurant_api/menu"
	"net/http"
)

type MenuHandler struct {
	menuCategoryService menu.CategoryService
	menuIngredientService menu.IngredientService
	menuItemService menu.ItemService


}

func NewMenuHandler(menuCategoryServ menu.CategoryService,menuIngridientServ menu.IngredientService,
	               menuItemServ menu.ItemService)*MenuHandler{

	return &MenuHandler{menuCategoryService: menuCategoryServ,menuIngredientService: menuIngridientServ,
	                      menuItemService: menuItemServ}
}

func(mh *MenuHandler)GetMenu(w http.ResponseWriter,r *http.Request){
      category,errs:= mh.menuCategoryService.Categories()
      if len(errs)>0{
      	w.Header().Set("content-type","application/json")
      	http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		  return
	  }
	  categoryOutput,err:= json.MarshalIndent(category,"","\t\t")
	  if err != nil{
	  	w.Header().Set("content-type","application/json")
	  	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		  return
	  }
	item,errs:= mh.menuItemService.Items()
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	itemOutput,err:= json.MarshalIndent(item,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	ingredient,errs:= mh.menuIngredientService.Ingredients()
	if len(errs)>0{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		return
	}
	ingredientOutput,err:= json.MarshalIndent(ingredient,"","\t\t")
	if err != nil{
		w.Header().Set("content-type","application/json")
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	  w.Header().Set("content-type","application/json")
	w.Write(categoryOutput)
	w.Write(itemOutput)
	w.Write(ingredientOutput)
	return

}