package helloworld

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/cfn"
	"log"
)

func EchoResource(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	v, _ := event.ResourceProperties["Echo"].(string)

	data = map[string]interface{}{
		"Response": v,
		"Echo": v,
	}

	log.Printf("data is %v", data)

	switch event.RequestType {
	case cfn.RequestCreate:
	case cfn.RequestUpdate:
	case cfn.RequestDelete:
	default:
		return "", nil, fmt.Errorf("unsupported request type: %s", event.RequestType)
	}

	return
}
