package bookcontroller

import (
	"final-project/entities"
	"final-project/libraries"
	"final-project/models"
	"net/http"
	"strconv"
	"text/template"
)

var validation = libraries.NewValidation()

var bookModel = models.NewBookModel()

func Index(response http.ResponseWriter, request *http.Request) {
	books, _ := bookModel.FindAll()

	data := map[string]interface{}{
		"books": books,
	}

	temp, err := template.ParseFiles("views/book/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/book/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var book entities.Book
		book.Title = request.Form.Get("title")
		book.Author = request.Form.Get("author")
		book.Publisher = request.Form.Get("publisher")
		book.ISBN = request.Form.Get("isbn")
		book.Year = request.Form.Get("year")
		book.Category = request.Form.Get("category")

		data := make(map[string]interface{})

		vErrors := validation.Struct(book)

		if vErrors != nil {
			data["book"] = book
			data["validation"] = vErrors
		} else {
			data["message"] = "Data is successfully stored"
			bookModel.Create(book)
		}

		temp, _ := template.ParseFiles("views/book/add.html")
		temp.Execute(response, data)
	}
}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var book entities.Book
		bookModel.Find(id, &book)

		data := map[string]interface{}{
			"book": book,
		}

		temp, err := template.ParseFiles("views/book/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)
	} else if request.Method == http.MethodPost {
		request.ParseForm()

		var book entities.Book
		book.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		book.Title = request.Form.Get("title")
		book.Author = request.Form.Get("author")
		book.Publisher = request.Form.Get("publisher")
		book.ISBN = request.Form.Get("isbn")
		book.Year = request.Form.Get("year")
		book.Category = request.Form.Get("category")

		data := make(map[string]interface{})

		vErrors := validation.Struct(book)

		if vErrors != nil {
			data["book"] = book
			data["validation"] = vErrors
		} else {
			data["message"] = "Data is successfully edited"
			bookModel.Update(book)
		}

		temp, _ := template.ParseFiles("views/book/edit.html")
		temp.Execute(response, data)
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	bookModel.Delete(id)

	http.Redirect(response, request, "/book", http.StatusSeeOther)
}
