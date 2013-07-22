goconfig
========

Goconfig is a configuration file parser for the Go Programming Language. It is based on [goconf](https://code.google.com/p/goconf/).

Goconfig has a few new features:

 - It gives more detailed errors
 - There is no need to specify a section name in the config file ("default" is assumed as the section name)
 - It can now read from a byte slice or io.Reader
 - It can now write to a byte slice or io.Writer
 - Returns default value on error or missing option

NOTE: All section names and options are case insensitive. All values are case sensitive.

Example 1
---------
Config file

    host = something.com
    port = 443
    active = true
    compression = off

Code

    c, err := conf.ReadConfigFile("something.config")
    c.String("host", "") // return something.com
    c.Int("port", 443) // return 443
    c.Bool("active", true) // return true
    c.Bool("compression", false) // return false

Example 2
---------
Config file

    [default]
    host = something.com
    port = 443
    active = true
    compression = off

    [service-1]
    compression = on

    [service-2]
    port = 444

Code

    c, err := conf.ReadConfigFile("something.config")
    c.GetBool("default", "compression") // returns false
    c.GetBool("service-1", "compression") // returns true
    c.GetBool("service-2", "compression") // returns GetError
