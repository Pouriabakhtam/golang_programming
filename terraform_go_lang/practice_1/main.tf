resource "google_compute_instance" "core_5g_instance" {
    count        = $lenght(var.name_count)
    name         = "core-5g-instance"
    machine_type = "e2-micro"
    zone         = "us-central1-a"
    boot_disk {
        initialize_params {
            image = "ubuntu-os-cloud/ubuntu-2204-lts"
        }
    }
    network_interface {
        network = "default"
        access_config {
        }
    }
    service_account {
        scopes = ["userinfo-email", "compute-ro", "storage-ro"]
    }
}

output "instance_name" {
    value = google_compute_instance.core_5g_instance.name
}
output "instance_zone" { value = google_compute_instance.core_5g_instance.zone }
output "instance_machine_type" { value = google_compute_instance.core_5g_instance.machine_type }