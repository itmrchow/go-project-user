// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "HelloWorld example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "登入",
                "parameters": [
                    {
                        "description": "Login sample , account 和 email 需擇一輸入",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqdto.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回token訊息",
                        "schema": {
                            "$ref": "#/definitions/respdto.LoginResp"
                        }
                    },
                    "default": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/respdto.ApiErrorResp"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "確認服務正常",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Ping Server",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error response",
                        "schema": {
                            "$ref": "#/definitions/respdto.ApiErrorResp"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "建立用戶",
                "parameters": [
                    {
                        "description": "Create user sample",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqdto.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{userId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "查詢用戶 by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "完整更新用戶 by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Put user sample",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqdto.PutUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "刪除用戶 by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "部分更新用戶 by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patch user sample",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqdto.PatchUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "查詢用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name",
                        "name": "userName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User Phone",
                        "name": "phone",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok\" \"返回用户信息",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "err_code：10002 参数错误； err_code：10003 校验错误",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "err_code：10001 登录失败",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "reqdto.CreateUserReq": {
            "type": "object",
            "required": [
                "account",
                "email",
                "password",
                "phone",
                "userName"
            ],
            "properties": {
                "account": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "jeff7777"
                },
                "email": {
                    "type": "string",
                    "example": "jeff@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "password"
                },
                "phone": {
                    "type": "string",
                    "example": "+886955555555"
                },
                "userName": {
                    "description": "使用者名稱",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "Jeff"
                }
            }
        },
        "reqdto.LoginReq": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string",
                    "example": "jeff7777"
                },
                "email": {
                    "type": "string",
                    "example": "jeff@gmail.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 8,
                    "example": "password"
                }
            }
        },
        "reqdto.PatchUserReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "jeff7777"
                },
                "email": {
                    "type": "string",
                    "example": "jeff@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "jeffpwd"
                },
                "phone": {
                    "type": "string",
                    "example": "+886955555555"
                },
                "userName": {
                    "type": "string",
                    "example": "Jeff"
                }
            }
        },
        "reqdto.PutUserReq": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "jeff7777"
                },
                "email": {
                    "type": "string",
                    "example": "jeff@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "jeffpwd"
                },
                "phone": {
                    "type": "string",
                    "example": "+886955555555"
                },
                "userName": {
                    "type": "string",
                    "example": "Jeff"
                }
            }
        },
        "respdto.ApiErrorResp": {
            "type": "object",
            "properties": {
                "detail": {
                    "description": "問題描述",
                    "type": "string"
                },
                "title": {
                    "description": "Type     string ` + "`" + `json:\"type,omitempty\"` + "`" + `\nStatus   int    ` + "`" + `json:\"status,omitempty\"` + "`" + `   // http status",
                    "type": "string"
                }
            }
        },
        "respdto.LoginResp": {
            "type": "object",
            "properties": {
                "exp": {
                    "type": "integer",
                    "example": 1719218836
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoiamVmZjc3NzciLCJlbWFpbCI6ImplZmZAZ21haWwuY29tIiwiZXhwIjoxNzE5MjE4ODM2LCJpZCI6ImZhNzkxODE2LWRkMzUtNDJlNi1hNDc1LTAwZjg3ZDRhYzlhYSIsInBob25lIjoiKzg4Njk1NTU1NTU1NSIsInVzZXJOYW1lIjoiSmVmZiJ9.ALccbWnDW4Tg6NvIS8aCw3B96okQ3gLVqiEz3Ukq_eA"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "tags": [
        {
            "description": "User API",
            "name": "User"
        },
        {
            "description": "Example API",
            "name": "Example"
        },
        {
            "description": "Other description",
            "name": "Other"
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "User Service API",
	Description:      "User Service API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
