[network]
  port = ${listen_port}
  host = "${private_ip}"

[discovery]
  minimum_master_nodes = ${minimum_masters}
  ping_unicast_hosts = [ ${private_ips} ]

[transport]
  port = 9300

${ "${backup_config}" == "s3" ? <<EOT
[s3]
  [s3.client.default]
    endpoint = "${endpoint}"

EOT 
: "${backup_config}" == "efs" ? <<EOT
[path]      
  repo = "${nfs_mount_path}/opensearch"
EOT 
: "" }

[tls]
  rootCA = """${opensearch_root_ca}"""
  admin_cert = """${opensearch_public_key}"""
  admin_key = """${opensearch_private_key}"""
  ssl_cert = """${opensearch_public_key}"""
  ssl_key = """${opensearch_private_key}"""
