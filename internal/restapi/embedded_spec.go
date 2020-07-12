// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "api",
    "version": "1.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/users/{user_id}": {
      "get": {
        "summary": "ユーザーの詳細情報を取得する",
        "operationId": "get-users-user_id",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "user": {
                  "$ref": "../models/user.v1.yaml"
                }
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "user_id",
          "in": "path",
          "required": true
        }
      ]
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "api",
    "version": "1.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/users/{user_id}": {
      "get": {
        "summary": "ユーザーの詳細情報を取得する",
        "operationId": "get-users-user_id",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "user": {
                  "$ref": "#/definitions/userV1"
                }
              }
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "user_id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "userV1": {
      "type": "object",
      "title": "user",
      "required": [
        "id",
        "display_name",
        "iconImagePath",
        "backgroundImagePath",
        "created_at",
        "updated_at"
      ],
      "properties": {
        "backgroundImagePath": {
          "type": "string"
        },
        "created_at": {
          "type": "number"
        },
        "deleted_at": {
          "type": "number"
        },
        "display_name": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "iconImagePath": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "profile": {
          "type": "string"
        },
        "updated_at": {
          "type": "number"
        }
      }
    }
  }
}`))
}
