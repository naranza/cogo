package cogo_test

import (
  "os"
  "fmt"
  "testing"
  "gitlab.com/naranza/cogo"
)

type Config struct {
  Port  int
  Debug bool
  Rate  float64
  Name  string
  ScriptFilePerms os.FileMode
}

func TestLoadConfig_Ok(t *testing.T) {
  
  var cfg Config

  err := cogo.LoadConfig("ok.cogo", &cfg)
  if err != nil {
    t.Fatalf("LoadConfig failed: %v", err)
  }

  // Adjust these assertions to what you expect from ok.cogo content
  if cfg.Port != 8080 {
    t.Errorf("expected Port=8080, got %d", cfg.Port)
  }
  if !cfg.Debug {
    t.Errorf("expected Debug=true, got false")
  }
  if cfg.Rate != 0.75 {
    t.Errorf("expected Rate=0.75, got %v", cfg.Rate)
  }
  if cfg.Name != "example" {
    t.Errorf("expected Name='example', got %q", cfg.Name)
  }
  if cfg.ScriptFilePerms != os.FileMode(0755) {
    t.Errorf("expected ScriptFilePerms=0755, got %o", cfg.ScriptFilePerms)
  }
}

func TestLoadConfig_Less3Parts(t *testing.T) {
  
  var cfg Config

  wie := "Invalid config line (less than 3 parts): Port int8080"
  wig := cogo.LoadConfig("less_3_parts.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}

func TestLoadConfig_CantSet(t *testing.T) {

  var cfg Config

  wie := "cannot set field: Ports"
  wig := cogo.LoadConfig("cantset.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}

func TestLoadConfig_InvalidInt(t *testing.T) {
  
  var cfg Config

  wie := "Invalid int value for key Port"
  wig := cogo.LoadConfig("invalid_int.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}

func TestLoadConfig_InvalidBool(t *testing.T) {
  
  var cfg Config

  wie := "Invalid bool value for key Debug"
  wig := cogo.LoadConfig("invalid_bool.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}


func TestLoadConfig_InvalidFloat(t *testing.T) {
  
  var cfg Config

  wie := "Invalid float value for key Rate"
  wig := cogo.LoadConfig("invalid_float.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}

func TestLoadConfig_InvalidFilemode(t *testing.T) {
  
  var cfg Config

  wie := "Invalid filemode value for key ScriptFilePerms"
  wig := cogo.LoadConfig("invalid_filemode.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }
  fmt.Println(wig)
}

func TestLoadConfig_TypeFail(t *testing.T) {
  
  var cfg Config

  wie := "unknown config type: strings"
  wig := cogo.LoadConfig("type_fail.cogo", &cfg).Error()
  if wie != wig {
    t.Fatalf("Expected '%v', got '%v'", wie, wig)
  }

}
