module "sqs" {
  source = "../../module/sqs"

  tags = {
    Environment = "local"
    Project     = "myproject"
  }
}

output "queue_url" {
  value = module.sqs.queue_url
}
