variable "secret_path" {
  description = "Path to the secret file for Google Cloud credentials"
  type        = string
  default     = "../credential/secret.json"
}

variable "name_count" { 
  description = "Names of instances" 
  type        = list
  default     = ["server1", "server2"]
}

variable "machine_type" {
  type = map(string)
  default = {
    "dev" = "e2-micro"
    "prod" = "e2-medium"
  }
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}

# variable production_machine_type {
#   description = "Machine type for production environment"
#   type        = string
#   default     = "e2-medium"
# }

# variable development_machine_type {
#   description = "Machine type for development environment"
#   type        = string
#   default     = "e2-micro"
# }
