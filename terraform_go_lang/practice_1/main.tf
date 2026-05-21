resource "google_compute_instance" "core_5g_instance" {
    count        = length(var.name_count)
    name         = "core-5g-instance-${count.index}"
    # here we are using conditional expression to determine the machine type based on the environment variable
    machine_type = var.environment == "prod" ? var.machine_type["prod"] : var.machine_type["dev"]
    # machine_type = var.environment == "prod" ? var.production_machine_type : var.development_machine_type
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
    # depends_on = [ google_compute_instance.core_lte_instance ]
}

# resource "google_compute_instance" "core_lte_instance" {
#     count = length((var.name_count))
#     name         = "core-lte-instance-${count.index}"
#     machine_type = var.machine_type["prod"]
#     zone         = "us-central1-a"
#     boot_disk {
#         initialize_params {
#             image = "ubuntu-os-cloud/ubuntu-2204-lts"
#         }
#     }
#     network_interface {
#         network = "default"
#         access_config {
#         }
#     }
#     service_account {
#         scopes = ["userinfo-email", "compute-ro", "storage-ro"]
#     }
# }

output "instance_name" {
    value = google_compute_instance.core_5g_instance.*.name
}
output "instance_zone" { value = google_compute_instance.core_5g_instance.*.zone }
output "instance_machine_type" { value = google_compute_instance.core_5g_instance.*.machine_type }
output "instance_id" { value = join(",", google_compute_instance.core_5g_instance.*.id) }