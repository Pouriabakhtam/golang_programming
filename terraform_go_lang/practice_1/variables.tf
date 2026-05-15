variable "secret_path" {
  description = "Path to the secret file for Google Cloud credentials"
  type        = string
  default     = "../credential/secret.json"
}

variable "name_count" { 
  description = "Names of instances" 
  type        = list
  default     = ["server1",  "server2" , "server3"]
}

variable "machine_type" {
  type = map(string)
  default = {
    "dev" = "e2-micro"
    "prod" = "e2-medium"
  }

}