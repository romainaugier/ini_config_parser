# ini_config_parser
Simple and easy to use .ini config parser written in Go

![BuildTest](https://github.com/romainaugier/ini_config_parser/actions/workflows/build_and_test.yml/badge.svg)

API:
```go
func main() {
    config_path := "/path/to/config.ini"

    // Parses the configuration
    config := IniConfigParse(config_path)

    // Config will be nil if there was an error during parsing
    if config == nil {
        t.Fatalf("Error during ini config parsing, check the log for more information")
    }

    // Gets a key value as string. By default, all values are stored as strings. If the key is not present,
    // returns an empty one
    string_key := IniConfigGet(config, "Global", "STRING_KEY")

    // Gets a key value as int. If there was an error during the conversion string->int, returns the
    // default value
    int_key := IniConfigGetInt(config, "Global", "INT_KEY", 0)

    // Gets a key value as bool. If there was an error during the conversion string->bool, returns the
    // default value
    bool_key := IniConfigGetBool(config, "Global", "BOOL_KEY", false)
}
```