package conf

import (
	"strconv"
	"testing"
)

const confFile = `
[default]
host = example.com
port = 43
compression = on
active = false

[service-1]
port = 443
`

//url = http://%(host)s/something

type stringtest struct {
	section string
	option  string
	answer  string
}

type inttest struct {
	section string
	option  string
	answer  int
}

type booltest struct {
	section string
	option  string
	answer  bool
}

var testSet = []interface{}{
	stringtest{"", "host", "example.com"},
	inttest{"default", "port", 43},
	booltest{"default", "compression", true},
	booltest{"default", "active", false},
	inttest{"service-1", "port", 443},
	//stringtest{"service-1", "url", "http://example.com/something"},
}

var options struct {
    Host string
    Port uint
    Compression bool
}

func TestBuild(t *testing.T) {
	c, err := ReadConfigBytes([]byte(confFile))
	if err != nil {
		t.Error(err)
	}

	for _, element := range testSet {
		switch element.(type) {
		case stringtest:
			e := element.(stringtest)
			ans, err := c.GetString(e.section, e.option)
			if err != nil {
				t.Error("c.GetString(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.GetString(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer: " + ans)
			}
		case inttest:
			e := element.(inttest)
			ans, err := c.GetInt(e.section, e.option)
			if err != nil {
				t.Error("c.GetInt(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.GetInt(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer: " + strconv.Itoa(ans))
			}
		case booltest:
			e := element.(booltest)
			ans, err := c.GetBool(e.section, e.option)
			if err != nil {
				t.Error("c.GetBool(\"" + e.section + "\",\"" + e.option + "\") returned error: " + err.Error())
			} else if ans != e.answer {
				t.Error("c.GetBool(\"" + e.section + "\",\"" + e.option + "\") returned incorrect answer")
			}
		}
	}

    if 43 != c.Int("port", 123) {
        t.Error("c.Int(\"port\") did not return 43")
    }

    if 123 != c.Int("ports", 123) {
        t.Error("c.Int(\"ports\") did not return 123")
    }

    c.Struct(&options)

    if "example.com" != options.Host {
        t.Error("options.Host did not equal to 'example.com' string")
    }

    if 43 != options.Port {
        t.Error("options.Port did not equal to int 43")
    }

    if true != options.Compression {
        t.Error("options.Compression did not equal to bool true")
    }
}
