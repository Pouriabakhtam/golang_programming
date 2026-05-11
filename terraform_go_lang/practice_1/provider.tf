provider "google" {
project = "core-5gc"
region  = "us-central1"
credentials = file(var.secret_path)
}