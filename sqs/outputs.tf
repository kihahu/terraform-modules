output "this_sqs_queue_id" {
  description = "The URL for the created Amazon SQS queue"
  value       = module.sqs.this_sqs_queue_id
}

output "this_sqs_queue_arn" {
  description = "The ARN of the SQS queue"
  value       = module.sqs.this_sqs_queue_arn
}

output "this_sqs_queue_name" {
  description = "The SQS queue name"
  value       = module.sqs.this_sqs_queue_name
}
