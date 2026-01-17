api_addr         = "http://127.0.0.1:8200"
plugin_directory = "./build/plugins"

storage "inmem" {}

listener "tcp" {
  address = "127.0.0.1:8200"
  tls_disable = "true"
}
