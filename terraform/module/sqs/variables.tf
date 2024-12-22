variable "tags" {
  description = "Tags to be applied to the SQS queue"
  type        = map(string)
  default     = {}
}
