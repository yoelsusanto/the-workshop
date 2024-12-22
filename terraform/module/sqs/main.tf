resource "aws_sqs_queue" "main_queue" {
  name                      = "main"
  visibility_timeout_seconds = 60

  tags = var.tags
}
