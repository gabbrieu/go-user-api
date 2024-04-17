variable "envfile" {
  type    = string
  default = ".env"
}

locals {
  envfile = {
    for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
    if !startswith(line, "#") && length(split("=", line)) > 1
  }
}

data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "gorm" {
  url = "postgres://${local.envfile["DATABASE_USER"]}:${local.envfile["DATABASE_PASSWORD"]}@${local.envfile["DATABASE_HOST"]}:${local.envfile["DATABASE_PORT"]}/${local.envfile["DATABASE_NAME"]}"
  dev = "docker://postgres/15/dev"
  src = data.external_schema.gorm.url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

