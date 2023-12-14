resource "google_sql_database" "database" {
  name     = "rose-auth"
  instance = google_sql_database_instance.instance.name
}

resource "google_sql_database_instance" "instance" {
  name             = "my-database-instance"
  region           = "us-central1"
  database_version = "MYSQL_8_0"
  settings {
    tier = "db-f1-micro"
  }

  deletion_protection  = "true"
}