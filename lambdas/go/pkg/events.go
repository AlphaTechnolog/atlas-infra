package pkg

import (
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// UnmarshalHTTPEventBody Unmarshal but for the `e.Body` being `e` an `APIGatewayV2HTTPRequest`. This one decodes
// the request as a base64 if the `e.IsBase64Encoded` is `true`.
func UnmarshalHTTPEventBody[T any](e events.APIGatewayV2HTTPRequest, out *T) error {
	var err error
	bodyBytes := []byte(e.Body)
	if e.IsBase64Encoded {
		bodyBytes, err = base64.StdEncoding.DecodeString(e.Body)
		if err != nil {
			return err
		}
	}

	return json.Unmarshal(bodyBytes, out)
}

// SendHTTPResponse responds in the `events.APIGatewayV2HTTPResponse` format an arbitrary response map with status.
func SendHTTPResponse[T any](status int, body T) events.APIGatewayV2HTTPResponse {
	compressedBody, _ := json.Marshal(body)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       string(compressedBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
