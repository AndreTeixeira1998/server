package tag

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/traggo/server/generated/gqlmodel"
	"github.com/traggo/server/model"
)

// CreateTag creates a tag.
func (r *ResolverForTag) CreateTag(ctx context.Context, key string, color string, typeArg gqlmodel.TagDefinitionType) (*gqlmodel.TagDefinition, error) {
	definition := &model.TagDefinition{
		Key:   key,
		Color: color,
		Type:  model.TagDefinitionType(typeArg),
	}

	if !r.DB.Find(definition).RecordNotFound() {
		return nil, fmt.Errorf("tag with key '%s' does already exist", definition.Key)
	}

	create := r.DB.Create(&definition)
	gqlTag := &gqlmodel.TagDefinition{}
	copier.Copy(gqlTag, definition)
	return gqlTag, create.Error
}
