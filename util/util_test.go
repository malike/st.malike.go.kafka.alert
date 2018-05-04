package util

import (
	"fmt"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	NewConfiguration()
	LogLevel = "INFO"
	if AppConfiguration.SMTPConfig.EmailSender == "" {
		t.Errorf("Required configuration not loaded ")
		t.FailNow()
	}
	t.Log("Email Sender available as " + AppConfiguration.SMTPConfig.EmailSender)
}

func TestLoadConfigurationWithTemplates(t *testing.T) {
	NewConfiguration()
	if len(AppConfiguration.Templates) == 0 {
		t.Error("Required configuration not loaded. No Templates found ")
		t.FailNow()
	}
	t.Log(fmt.Printf("templates found '%d' ", len(AppConfiguration.Templates)))
	for k, v := range AppConfiguration.Templates {
		t.Log("Template ID '" + k + "' ==> " + v)
	}
}

func TestConfiguration_GetTemplate(t *testing.T) {
	NewConfiguration()
	if len(AppConfiguration.Templates) == 0 {
		t.Errorf("Required configuration not loaded. No Templates found ")
		t.FailNow()
	}
	var randomTemplateID string
	var randomTemplateContent string
	for k, v := range AppConfiguration.Templates {
		randomTemplateID = k
		randomTemplateContent = v
		break
	}
	templateContent := AppConfiguration.GetTemplate(randomTemplateID)
	if templateContent != randomTemplateContent {
		t.Errorf("Could not fetch template")
		t.FailNow()
	}
	t.Log("Temaplate ID '" + randomTemplateID + "' matched it's  content '" + randomTemplateContent + "'")
}

func BenchmarkConfiguration_GetTemplate(b *testing.B) {
	NewConfiguration()
	for i := 0; i < b.N; i++ {
		var randomTemplateID string
		for k := range AppConfiguration.Templates {
			randomTemplateID = k
			break
		}
		AppConfiguration.GetTemplate(randomTemplateID)
	}
}
