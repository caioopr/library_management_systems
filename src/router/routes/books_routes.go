package routes

import (
	"api/src/controllers"
	"net/http"
)

var booksRoutes = []Route{
	{
		URI:          "/books",
		Method:       http.MethodPost,
		Function:     controllers.CreateBook,
		AuthRequired: true,
	},
	{
		URI:          "/books",
		Method:       http.MethodGet,
		Function:     controllers.GetBooks,
		AuthRequired: true,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodGet,
		Function:     controllers.GetBook,
		AuthRequired: true,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodPatch,
		Function:     controllers.UpdateBook,
		AuthRequired: true,
	},
	{
		URI:          "/books/{bookId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteBook,
		AuthRequired: true,
	},
}
