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
                        "description": "OK",
                        "schema": {
                            "type": "string"
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
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 4,
                    "example": "Jeff"
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
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
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
