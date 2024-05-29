package input

import (
    "testing"
)

func TestNewPromptTemplate(t *testing.T) {
    validTemplateString := "Hello, {{.Name}}!"
    invalidTemplateString := "Hello, {{.Name}"

    tests := []struct {
        name        string
        templateStr string
        shouldError bool
    }{
        {"Valid Template", validTemplateString, false},
        {"Invalid Template", invalidTemplateString, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            pt, err := NewPromptTemplate(tt.templateStr)
            if tt.shouldError {
                if err == nil {
                    t.Errorf("expected error, got nil")
                }
                return
            }
            if err != nil {
                t.Errorf("did not expect error, got %v", err)
            }
            if pt.Template == nil {
                t.Errorf("expected template to be initialized, got nil")
            }
        })
    }
}