-- <-- this begins a comment meant for humans
local payload = {
  -- let's see some common datatypes
  nullz = null,

  true_bool = true,
  false_bool = false,

  integer = 123,
  float = 123.123,

  string = "value",

  -- since string starts with a comment sign (#)
  -- it has to be disambiguated by wrapping it in quotes
  string_quoted = "#value",

  -- this is a key (nested_map) whose value is a map
  -- with multiple key value pairs
  nested_map = {
    key1 = "value1",
    key2 = "value2",
  },

  -- this is a list with two items
  list = {"item1", "item2"},

  -- lists can carry any values including maps
  list_with_a_map = {
    {
      key1 = "value1",
      key2 = "value2",
    },
    "item2"
  },
}

return payload