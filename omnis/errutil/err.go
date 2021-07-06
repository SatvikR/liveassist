// Copyright (c) 2021 Satvik Reddy
package errutil

import "github.com/gin-gonic/gin"

// Represents a JSON field with an error
type ErrorField struct {
	Field string
	Err   error
}

// CreateErrJSON creates a error json with an errors array mapping fields to
// their corresponding errors. This makes the frontend error handling much easier
//
// ex.
// 	errutil.CreateErrJSON(
// 		[]errutil.ErrorField{
// 			{Field: "password", Err: "incorrect password"},
// 		},
// 	),
// 	// returns
// 	{
// 		"errors": [
// 			{
// 				"password": "incorrect password"
// 			}
// 		]
// 	}
func CreateErrJSON(errs []ErrorField) gin.H {
	errList := make([]gin.H, len(errs))
	for i, v := range errs {
		errList[i] = gin.H{
			v.Field: v.Err.Error(),
		}
	}
	return gin.H{
		"errors": errList,
	}
}
