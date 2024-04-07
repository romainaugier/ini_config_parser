package ini_config_parser

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	path, err := os.Getwd()

	if err != nil {
		t.Fatalf("Error caught during testing, cannot find current directory: %v", err)
	}

	config_path := fmt.Sprintf("%s/test_data/test.ini", path)

	if _, err := os.Stat(config_path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Cannot find ini config test file")
	}

	config := IniConfigParse(config_path)

	if config == nil {
		t.Fatalf("Error during ini config parsing, check the log for more information")
	}

	if string_key := IniConfigGet(config, "Global", "STRING_KEY"); string_key != "STRING_VALUE" {
		t.Fatalf("Wrong value for Global/STRING_KEY: %s", string_key)
	}

	if int_key := IniConfigGetInt(config, "Global", "INT_KEY", 0); int_key != 1 {
		t.Fatalf("Wrong value for Global/INT_KEY: %d", int_key)
	}

	if bool_key := IniConfigGetBool(config, "Global", "BOOL_KEY", false); bool_key != true {
		t.Fatalf("Wrong value for Global/BOOL_KEY: %t", bool_key)
	}

	if string_key_lc := IniConfigGet(config, "Local", "string_key_lc"); string_key_lc != "string value" {
		t.Fatalf("Wrong value for Local/string_key_lc: %s", string_key_lc)
	}

	if int_key := IniConfigGetInt(config, "Local", "int_key_lc", 0); int_key != 10 {
		t.Fatalf("Wrong value for Local/int_key_lc: %d", int_key)
	}

	if bool_key := IniConfigGetBool(config, "Local", "bool_key_lc", true); bool_key != false {
		t.Fatalf("Wrong value for Local/bool_key_lc: %t", bool_key)
	}
}
