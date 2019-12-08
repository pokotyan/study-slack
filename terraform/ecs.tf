resource "aws_ecs_cluster" "connpass-map" {
  name = "connpass-map"
}

# resource "aws_ecs_task_definition" "connpass-map" {
#   family                   = "connpass-map"
#   cpu                      = "256"
#   memory                   = "512"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["FARGATE"]
#   container_definitions    = file("./container_definitions.json")
# }

# resource "aws_ecs_service" "connpass-map" {
#   name    = "connpass-map"
#   cluster = aws_ecs_cluster.connpass-map.arn
#   task_definition                   = aws_ecs_task_definition.connpass-map.arn
#   desired_count                     = 2
#   launch_type                       = "FARGATE"
#   platform_version                  = "1.3.0"
#   health_check_grace_period_seconds = 60

#   network_configuration {
#     assign_public_ip = false
#     security_groups  = [module.ecs_service_sg.security_group_id]

#     subnets = [
#       aws_subnet.private_0.id,
#       aws_subnet.private_1.id,
#     ]
#   }

#   load_balancer {
#     target_group_arn = aws_lb_target_group.connpass-map.arn
#     container_name   = "connpass-map"
#     container_port   = 80
#   }

#   lifecycle {
#     ignore_changes = [task_definition]
#   }
# }

module "ecs_service_sg" {
  source      = "./modules/security_group"
  name        = "ecs_service_sg"
  vpc_id      = aws_vpc.connpass-map.id
  port        = 80
  cidr_blocks = [aws_vpc.connpass-map.cidr_block]
}
