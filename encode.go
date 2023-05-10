package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"io"
)

func Encode(r io.Reader, w io.Writer) error {
	normalJSONs, err := readDynamoDBJSONLines(r)
	if err != nil {
		return fmt.Errorf("read stdin: %w", err)
	}

	encoder := dynamodbattribute.NewEncoder()

	for _, normalJSON := range normalJSONs {
		dynamodbJSONValue, err := encoder.Encode(normalJSON)
		if err != nil {
			return fmt.Errorf("parse json: %w", err)
		}

		if err := writeNormalJSON(w, dynamodbJSONValue); err != nil {
			return fmt.Errorf("write normal json: %w", err)
		}
	}

	return nil
}

func readDynamoDBJSONLines(r io.Reader) ([]interface{}, error) {
	decoder := json.NewDecoder(r)

	normalJSONs := []interface{}{}

	for {
		var normalJSON interface{}

		err := decoder.Decode(&normalJSON)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("read dynamodb json: %w", err)
		}

		normalJSONs = append(normalJSONs, normalJSON)
	}

	return normalJSONs, nil
}

func writeNormalJSON(w io.Writer, value *dynamodb.AttributeValue) error {
	encoder := json.NewEncoder(w)

	if v := value.NULL; v != nil {
		if err := encoder.Encode(value); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.BOOL; v != nil {
		if err := encoder.Encode(value); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.N; v != nil {
		if err := encoder.Encode(value); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.S; v != nil {
		if err := encoder.Encode(value); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.B; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.BS; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.NS; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.SS; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.L; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	if v := value.M; v != nil {
		if err := encoder.Encode(v); err != nil {
			return fmt.Errorf("write stdout: %w", err)
		}
		return nil
	}

	return nil
}
