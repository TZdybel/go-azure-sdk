package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MLAssistConfiguration = MLAssistConfigurationDisabled{}

type MLAssistConfigurationDisabled struct {

	// Fields inherited from MLAssistConfiguration
}

var _ json.Marshaler = MLAssistConfigurationDisabled{}

func (s MLAssistConfigurationDisabled) MarshalJSON() ([]byte, error) {
	type wrapper MLAssistConfigurationDisabled
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MLAssistConfigurationDisabled: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MLAssistConfigurationDisabled: %+v", err)
	}
	decoded["mlAssist"] = "Disabled"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MLAssistConfigurationDisabled: %+v", err)
	}

	return encoded, nil
}
