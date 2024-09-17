package validators

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func NotNullObject() validator.Object {
	return notNullObjectNestedAttributeValidator{}
}

type notNullObjectNestedAttributeValidator struct{}

func (notNullObjectNestedAttributeValidator) Description(context.Context) string {
	return "value defaults to state value, if set"
}

func (m notNullObjectNestedAttributeValidator) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m notNullObjectNestedAttributeValidator) ValidateObject(ctx context.Context, request validator.ObjectRequest, response *validator.ObjectResponse) {
	if !request.ConfigValue.IsNull() {
		return
	}

	response.Diagnostics.AddAttributeError(
		request.Path,
		"Missing Attribute Value",
		request.Path.String()+": "+m.Description(ctx),
	)
}
