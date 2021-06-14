package pattern

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"strings"

	"git.canopsis.net/canopsis/go-engines/lib/canopsis/types"
	mgobson "github.com/globalsign/mgo/bson"
	mongobson "go.mongodb.org/mongo-driver/bson"
)

// InfoRegexMatches is a type that contains the values of the sub-expressions
// of regular expressions for each of the fields of an Info that contain
// strings.
type InfoRegexMatches struct {
	Name        RegexMatches
	Description RegexMatches
	Value       RegexMatches
}

// InfoFields is a type representing a pattern that can be applied to a piece
// of information on an entity.
// The fields are not defined directly in the Info struct to make the
// unmarshalling easier.
type InfoFields struct {
	Name        StringPattern `bson:"name"`
	Description StringPattern `bson:"description"`
	Value       StringPattern `bson:"value"`

	// When unmarshalling a BSON document, the fields of this document that are
	// not defined in this struct are added to UnexpectedFields.
	UnexpectedFields map[string]interface{} `bson:",inline"`
}

// InfoPattern is a type representing a pattern that can be applied to a piece
// of information on an entity.
type InfoPattern struct {
	ShouldNotBeSet bool

	InfoFields
}

// AsMongoQuery returns a mongodb filter from the InfoPattern for mgo-driver
func (p InfoPattern) AsMongoQuery() mgobson.M {
	var query mgobson.M
	query = make(mgobson.M)
	if p.ShouldNotBeSet {
		return nil
	}

	if !p.Name.Empty() {
		query["name"] = p.Name.AsMongoQuery()
	}

	if !p.Description.Empty() {
		query["description"] = p.Description.AsMongoQuery()
	}

	if !p.Value.Empty() {
		query["value"] = p.Value.AsMongoQuery()
	}

	return query
}

// AsMongoQuery returns a mongodb filter from the InfoPattern for mongo-driver
func (p InfoPattern) AsMongoDriverQuery() mongobson.M {
	var query mongobson.M
	query = make(mongobson.M)
	if p.ShouldNotBeSet {
		return nil
	}

	if !p.Name.Empty() {
		query["name"] = p.Name.AsMongoDriverQuery()
	}

	if !p.Description.Empty() {
		query["description"] = p.Description.AsMongoDriverQuery()
	}

	if !p.Value.Empty() {
		query["value"] = p.Value.AsMongoDriverQuery()
	}

	return query
}

// Matches returns true if a piece of information is matched by a pattern. If
// the pattern contains regular expressions with sub-expressions, the values
// of the sub-expressions are written in the matches argument.
func (p InfoPattern) Matches(info types.Info, isSet bool, matches *InfoRegexMatches) bool {
	if p.ShouldNotBeSet {
		return !isSet
	}

	return isSet &&
		p.Name.Matches(info.Name, &matches.Name) &&
		p.Description.Matches(info.Description, &matches.Description) &&
		p.Value.Matches(info.Value, &matches.Value)
}

// SetBSON unmarshals a BSON value into an InfoPattern.
func (p *InfoPattern) SetBSON(raw mgobson.Raw) error {
	switch raw.Kind {
	case mgobson.ElementNil:
		// The BSON value is null. The field should not be set.
		p.ShouldNotBeSet = true
		return nil

	default:
		err := raw.Unmarshal(&p.InfoFields)
		if err != nil {
			return err
		}

		if len(p.UnexpectedFields) != 0 {
			unexpectedFieldNames := make([]string, 0, len(p.UnexpectedFields))
			for key := range p.UnexpectedFields {
				unexpectedFieldNames = append(unexpectedFieldNames, key)
			}
			return fmt.Errorf("Unexpected pattern fields: %s", strings.Join(unexpectedFieldNames, ", "))
		}

		return nil
	}
}

func (p InfoPattern) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if p.ShouldNotBeSet {
		return bsontype.Null, []byte{}, nil
	}

	resultBson := mongobson.M{}

	if p.Name.RegexMatch.Set || p.Name.Equal.Set {
		bsonFieldName, err := GetFieldBsonName(p, "Name", "name")
		if err != nil {
			return bsontype.Undefined, nil, err
		}

		resultBson[bsonFieldName] = p.Name
	}

	if p.Description.RegexMatch.Set || p.Description.Equal.Set {
		bsonFieldName, err := GetFieldBsonName(p, "Description", "description")
		if err != nil {
			return bsontype.Undefined, nil, err
		}

		resultBson[bsonFieldName] = p.Description
	}

	if p.Value.RegexMatch.Set || p.Value.Equal.Set {
		bsonFieldName, err := GetFieldBsonName(p, "Value", "value")
		if err != nil {
			return bsontype.Undefined, nil, err
		}

		resultBson[bsonFieldName] = p.Value
	}

	if len(resultBson) > 0 {
		return mongobson.MarshalValue(resultBson)
	}

	return bsontype.Undefined, nil, nil
}

func (p *InfoPattern) UnmarshalBSONValue(valueType bsontype.Type, b []byte) error {
	switch valueType {
	case bsontype.Null:
		// The BSON value is null. The field should not be set.
		p.ShouldNotBeSet = true
	default:
		err := mongobson.Unmarshal(b, &p.InfoFields)
		if err != nil {
			return err
		}

		if len(p.UnexpectedFields) != 0 {
			unexpectedFieldNames := make([]string, 0, len(p.UnexpectedFields))
			for key := range p.UnexpectedFields {
				unexpectedFieldNames = append(unexpectedFieldNames, key)
			}
			return UnexpectedFieldsError{
				Err: fmt.Errorf("Unexpected pattern fields: %s", strings.Join(unexpectedFieldNames, ", ")),
			}
		}
	}

	return nil
}