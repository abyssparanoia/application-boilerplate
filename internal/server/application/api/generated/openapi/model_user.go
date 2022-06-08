/*
 * api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type User struct {
	Id string `json:"id"`

	DisplayName string `json:"display_name"`

	IconImagePath string `json:"iconImagePath"`

	BackgroundImagePath string `json:"backgroundImagePath"`

	Profile string `json:"profile"`

	Email string `json:"email"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	DeletedAt string `json:"deleted_at,omitempty"`
}

// AssertUserRequired checks if the required fields are not zero-ed
func AssertUserRequired(obj User) error {
	elements := map[string]interface{}{
		"id":                  obj.Id,
		"display_name":        obj.DisplayName,
		"iconImagePath":       obj.IconImagePath,
		"backgroundImagePath": obj.BackgroundImagePath,
		"profile":             obj.Profile,
		"email":               obj.Email,
		"created_at":          obj.CreatedAt,
		"updated_at":          obj.UpdatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of User (e.g. [][]User), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUser, ok := obj.(User)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserRequired(aUser)
	})
}