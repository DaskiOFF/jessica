package gen_test

import (
	"testing"

	"github.com/daskioff/jessica/flows/generator/gen"
)

func Test_NewParams_WithAllTypesSequentialOrder(t *testing.T) {
	args := []string{"usecase", "App", "--notest", "--nomock", "ck1:cv1", "ck2:cv2"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "App" {
		t.Error("Expected ModuleName == App, got ", p.ModuleName)
	}

	if p.NeedGenerateTests != false {
		t.Error("Expected NeedGenerateTests == false, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != false {
		t.Error("Expected NeedGenerateMock == false, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 2 {
		t.Error("Expected 2 custom keys, got ", len(p.CustomKeys))
	}

	if p.CustomKeys["ck1"] != "cv1" {
		t.Error("Expected value 'cv1' for key 'ck1', got ", p.CustomKeys["ck1"])
	}

	if p.CustomKeys["ck2"] != "cv2" {
		t.Error("Expected value 'cv2' for key 'ck2', got ", p.CustomKeys["ck2"])
	}
}

func Test_NewParams_WithAllTypesNonSequentialOrder(t *testing.T) {
	args := []string{"usecase", "App", "ck2:cv2", "--nomock", "ck1:cv1", "--notest"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "App" {
		t.Error("Expected ModuleName == App, got ", p.ModuleName)
	}

	if p.NeedGenerateTests != false {
		t.Error("Expected NeedGenerateTests == false, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != false {
		t.Error("Expected NeedGenerateMock == false, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 2 {
		t.Error("Expected 2 custom keys, got ", len(p.CustomKeys))
	}

	if p.CustomKeys["ck1"] != "cv1" {
		t.Error("Expected value 'cv1' for key 'ck1', got ", p.CustomKeys["ck1"])
	}

	if p.CustomKeys["ck2"] != "cv2" {
		t.Error("Expected value 'cv2' for key 'ck2', got ", p.CustomKeys["ck2"])
	}
}

func Test_NewParams_WithoutModuleName(t *testing.T) {
	args := []string{"usecase", "--notest", "--nomock", "ck1:cv1", "ck2:cv2"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "" {
		t.Error("Expected ModuleName == '', got ", p.ModuleName)
	}

	if p.NeedGenerateTests != false {
		t.Error("Expected NeedGenerateTests == false, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != false {
		t.Error("Expected NeedGenerateMock == false, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 2 {
		t.Error("Expected 2 custom keys, got ", len(p.CustomKeys))
	}

	if p.CustomKeys["ck1"] != "cv1" {
		t.Error("Expected value 'cv1' for key 'ck1', got ", p.CustomKeys["ck1"])
	}

	if p.CustomKeys["ck2"] != "cv2" {
		t.Error("Expected value 'cv2' for key 'ck2', got ", p.CustomKeys["ck2"])
	}
}

func Test_NewParams_WithoutNoTest(t *testing.T) {
	args := []string{"usecase", "App", "--nomock", "ck1:cv1", "ck2:cv2"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "App" {
		t.Error("Expected ModuleName == 'App', got ", p.ModuleName)
	}

	if p.NeedGenerateTests != true {
		t.Error("Expected NeedGenerateTests == true, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != false {
		t.Error("Expected NeedGenerateMock == false, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 2 {
		t.Error("Expected 2 custom keys, got ", len(p.CustomKeys))
	}

	if p.CustomKeys["ck1"] != "cv1" {
		t.Error("Expected value 'cv1' for key 'ck1', got ", p.CustomKeys["ck1"])
	}

	if p.CustomKeys["ck2"] != "cv2" {
		t.Error("Expected value 'cv2' for key 'ck2', got ", p.CustomKeys["ck2"])
	}
}

func Test_NewParams_WithoutNoMock(t *testing.T) {
	args := []string{"usecase", "App", "--notest", "ck1:cv1", "ck2:cv2"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "App" {
		t.Error("Expected ModuleName == 'App', got ", p.ModuleName)
	}

	if p.NeedGenerateTests != false {
		t.Error("Expected NeedGenerateTests == false, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != true {
		t.Error("Expected NeedGenerateMock == true, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 2 {
		t.Error("Expected 2 custom keys, got ", len(p.CustomKeys))
	}

	if p.CustomKeys["ck1"] != "cv1" {
		t.Error("Expected value 'cv1' for key 'ck1', got ", p.CustomKeys["ck1"])
	}

	if p.CustomKeys["ck2"] != "cv2" {
		t.Error("Expected value 'cv2' for key 'ck2', got ", p.CustomKeys["ck2"])
	}
}

func Test_NewParams_WithoutCustomKeys(t *testing.T) {
	args := []string{"usecase", "App", "--notest", "--nomock"}

	p := gen.NewParams(args)

	if p.TemplateName != "usecase" {
		t.Error("Expected TemplateName == usecase, got ", p.TemplateName)
	}

	if p.ModuleName != "App" {
		t.Error("Expected ModuleName == 'App', got ", p.ModuleName)
	}

	if p.NeedGenerateTests != false {
		t.Error("Expected NeedGenerateTests == false, got ", p.NeedGenerateTests)
	}

	if p.NeedGenerateMock != false {
		t.Error("Expected NeedGenerateMock == false, got ", p.NeedGenerateTests)
	}

	if len(p.CustomKeys) != 0 {
		t.Error("Expected 0 custom keys, got ", len(p.CustomKeys))
	}
}
