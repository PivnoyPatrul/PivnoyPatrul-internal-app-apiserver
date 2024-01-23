package apiserver
type Config struct {
  BindAddr string 'toml:"bind_addr"'
  }
func Newconfig() Config {
  retunr &Config {
   
  BindAddr:":8080,
  }
  }
