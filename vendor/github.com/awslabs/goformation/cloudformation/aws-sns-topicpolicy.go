package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSSNSTopicPolicy AWS CloudFormation Resource (AWS::SNS::TopicPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type AWSSNSTopicPolicy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument,omitempty"`

	// Topics AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
	Topics []*Value `json:"Topics,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSTopicPolicy) AWSCloudFormationType() string {
	return "AWS::SNS::TopicPolicy"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSSNSTopicPolicy) MarshalJSON() ([]byte, error) {
	type Properties AWSSNSTopicPolicy
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSSNSTopicPolicy) UnmarshalJSON(b []byte) error {
	type Properties AWSSNSTopicPolicy
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSSNSTopicPolicy(*res.Properties)
	}

	return nil
}

// GetAllAWSSNSTopicPolicyResources retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
func (t *Template) GetAllAWSSNSTopicPolicyResources() map[string]AWSSNSTopicPolicy {
	results := map[string]AWSSNSTopicPolicy{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSSNSTopicPolicy:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SNS::TopicPolicy" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSSNSTopicPolicy{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSSNSTopicPolicyWithName retrieves all AWSSNSTopicPolicy items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSNSTopicPolicyWithName(name string) (AWSSNSTopicPolicy, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSSNSTopicPolicy:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SNS::TopicPolicy" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSSNSTopicPolicy{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSSNSTopicPolicy{}, errors.New("resource not found")
}
