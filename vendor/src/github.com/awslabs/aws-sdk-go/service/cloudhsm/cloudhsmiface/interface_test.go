// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudhsmiface_test

import (
	"testing"

	"github.com/awslabs/aws-sdk-go/service/cloudhsm"
	"github.com/awslabs/aws-sdk-go/service/cloudhsm/cloudhsmiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudhsmiface.CloudHSMAPI)(nil), cloudhsm.New(nil))
}