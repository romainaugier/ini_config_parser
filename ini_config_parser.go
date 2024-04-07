package ini_config_parser

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type IniConfig struct {
	Path    string
	Content map[string]string
}

const (
	IniTokenType_Section = iota
	IniTokenType_Key     = iota
	IniTokenType_Value   = iota
	IniTokenType_Comment = iota
)

type IniToken struct {
	Data string
	Type int
}

var is_alnum = regexp.MustCompile(`[a-zA-Z0-9]`).MatchString

func ini_lex(file_content string) ([]IniToken, bool) {
	tokens := []IniToken{}

	s := file_content

	for i := 0; i < len(s); i++ {

		if s[i] == '[' {
			start_index := i + 1
			size := 0

			for s[i] != ']' {
				i++
				size++

				if i == len(s)-1 {
					fmt.Printf("Error during ini config lexing\n")
					return tokens, false
				}
			}

			tokens = append(tokens, IniToken{file_content[start_index : start_index+size-1], IniTokenType_Section})

			i++
		} else if is_alnum(string(s[i])) {
			key_start := i
			key_size := 0

			for s[i] != '=' {
				i++

				if s[i] != ' ' {
					key_size++
				}

				if i == len(s)-1 {
					fmt.Printf("Error during ini config lexing\n")
					return tokens, false
				}
			}

			tokens = append(tokens, IniToken{file_content[key_start : key_start+key_size], IniTokenType_Key})

			i++

			for i < len(s) && s[i] == ' ' {
				i++
			}

			value_start := i
			value_size := 0

			for i < len(s) && s[i] != '\n' && s[i] != '\r' && s[i] != '\t' {
				i++
				value_size++
			}

			tokens = append(tokens, IniToken{file_content[value_start : value_start+value_size], IniTokenType_Value})
		} else if s[i] == ';' {
			for s[i] != '\n' && s[i] != '\r' && s[i] != '\t' && i < len(s) {
				i++
			}
		}
	}

	return tokens, true
}

func ini_parse(tokens []IniToken) map[string]string {
	contents := make(map[string]string)

	current_section := "Default"

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == IniTokenType_Section {
			current_section = tokens[i].Data
		} else if tokens[i].Type == IniTokenType_Key {
			if (i+1) < len(tokens) && tokens[i+1].Type == IniTokenType_Value {
				key := fmt.Sprintf("%s___%s", current_section, tokens[i].Data)

				contents[key] = tokens[i+1].Data

				i++
			}
		}
	}

	return contents
}

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func IniConfigParse(file_path string) *IniConfig {
	if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("[Error] : IniConfigParse can't find ini config file path: \"%s\"\n", file_path)
		return nil
	}

	data, error := os.ReadFile(file_path)

	check_error(error)

	tokens, err := ini_lex(string(data))

	if !err {
		fmt.Print("[Error] : IniConfigParse error during config lexing\n")
		return nil
	}

	contents := ini_parse(tokens)

	config := IniConfig{file_path, contents}

	return &config
}

// Returns a ini config value, as string
func IniConfigGet(config *IniConfig, config_section string, config_key string) string {
	key := fmt.Sprintf("%s___%s", config_section, config_key)

	data := config.Content[key]

	if strings.Contains(data, "\"") {
		data = strings.ReplaceAll(data, "\"", "")
	}

	return data
}

// Returns a ini config value as int. If the value cannot be converted to int, returns default value
func IniConfigGetInt(config *IniConfig, config_section string, config_key string, default_value int) int {
	string_value := IniConfigGet(config, config_section, config_key)

	value, err := strconv.Atoi(string_value)

	if err != nil {
		return default_value
	}

	return value
}

// Returns a ini config value as bool. If the value cannot be converted to bool, returns false
func IniConfigGetBool(config *IniConfig, config_section string, config_key string, default_value bool) bool {
	string_value := IniConfigGet(config, config_section, config_key)

	value, err := strconv.ParseBool(string_value)

	if err != nil {
		return default_value
	}

	return value
}
