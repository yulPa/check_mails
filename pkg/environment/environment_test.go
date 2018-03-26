package environment

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateANewEnvironment(t *testing.T) {
	data := []byte(`
    {
			"name": "an_environment",
      "ips": [
        "192.168.0.1",
        "192.168.0.2",
        "192.168.0.3"
      ],
      "abuse": "abuse@domain.tld",
      "open": false,
			"options": {
				"quota": {
					"tenlastminutes": 150,
					"sixtylastminutes": 200,
					"lastday": 1000,
					"lastweek": 3000,
					"lastmonth": 10000
				}
			}
    }
    `)
	env, err := NewEnvironment(data)
	assert.Equal(t, 3000, env.Options.Quota.LastWeek)
	assert.Equal(t, "an_environment", env.Name)
	assert.Nil(t, err)
}

func TestCreateNewDefaultEnvironment(t *testing.T) {
	env := NewDefaultEnvironment(
		"a_default_environment",
		[]string{"192.168.0.1", "192.168.0.2"},
		"abuse@domain.tld",
		false,
	)
	assert.Equal(t, "a_default_environment", env.Name)
}

func TestCreateAnArrayofEnvironment(t *testing.T) {
	data := []byte(`
		[
		  {
		    "name": "environment1",
		    "ips": [
		      "192.168.0.1",
		      "192.168.0.2",
		      "192.168.0.3"
		    ],
		    "abuse": "abuse@domain.tld",
		    "open": false,
		    "options": {
		      "quota": {
		        "tenlastminutes": 150,
		        "sixtylastminutes": 200,
		        "lastday": 1000,
		        "lastweek": 3000,
		        "lastmonth": 10000
		      }
		    },
		    "entity": "an_entity"
		  },
		  {
		    "name": "environment2",
		    "ips": [
		      "192.168.0.1",
		      "192.168.0.2",
		      "192.168.0.3"
		    ],
		    "abuse": "abuse2@domain.tld",
		    "open": true,
		    "options": {
		      "quota": {
		        "tenlastminutes": 150,
		        "sixtylastminutes": 200,
		        "lastday": 1000,
		        "lastweek": 1234,
		        "lastmonth": 10000
		      }
		    },
		    "entity": "an_entity"
		  }
		]
		`)
	envs, _ := NewEnvironments(data)
	assert.Equal(t, 1234, envs[1].Options.Quota.LastWeek)
}
