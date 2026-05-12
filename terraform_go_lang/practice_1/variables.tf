variable "secret_path" {
  description = "Path to the secret file for Google Cloud credentials"
  type        = string
  default     = "../credential/secret.json"
}

variable "name_count" { 
  description = "Names of instances" 
  type        = list
  default     = ["server1",  "server2" , "server3"]
  