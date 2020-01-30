package main

// @Router              /v1/user/{uid}/subscriber [GET]
// @Summary             User's subscribers
// @Description         Query user's subscribers, return page data
// @Tag                 User
// @Tag                 Subscribe
// @Param               uid path integer true false "user id" 1
// @Param               page query integer false true "page" 1
// @Accept              multipart/form-data
// @ResponseDesc 400    "request param error"
// @ResponseDesc 404    "user not found"
// @Response 200        ${resp_user}
/* @Response 400        {
							"code": 400,
							"message": "request param error"
 						} */

// @Router              /v1/user/subscribing [PUT]
// @Security            Jwt
// @Template            Auth Other
// @Summary             Subscribe user
// @Description         Subscribe someone
// @Tag                 User
// @Tag                 Subscribe
// @Param               to formData integer true false "user id"
// @Accept              multipart/form-data
// @Produce             application/json
// @ResponseDesc 400    "request param error"
// @ResponseDesc 400    "request format error"
// @ResponseDesc 404    "user not found"
// @ResponseDesc 500    "subscribe failed"
/* @Response 200        {
							"code": 200,
							"message": "success",
							"data": ${array}
 						} */
func func1() {

}

// @Router              /v1/user/subscribing [DELETE]
// @Security            Jwt
// @Template            Auth
// @Summary             Unsubscribe user
// @Description         Unsubscribe someone
// @Tag                 User
// @Param               to formData integer true false "user id"
// @ResponseDesc 400    "request param error"
// @ResponseDesc 400    "request format error"
// @ResponseDesc 404    "user not found"
// @ResponseDesc 500    "unsubscribe failed"
// @ResponseHeader 200  { "Content-Type": "application/json; charset=utf-8" }
/* @Response 200 		{
							"code": 200,
							"message": "success"
 						} */
// @ResponseHeader 400  { "Content-Type": "application/json; charset=utf-8" }
/* @Response 400 		{
							"code": 400,
							"message": "request param error"
 						} */
/* @Response 500        {
							"code": 500,
							"message": "unsubscribe failed"
 						} */
func func2() {

}
