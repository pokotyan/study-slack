variable "name" {}
variable "vpc_id" {}
variable "port" {}

variable "cidr_blocks" {
  type = list(string)
}

# セキュリティグループのガワ作成
resource "aws_security_group" "default" {
  name   = var.name
  vpc_id = var.vpc_id
}

# インバウンドルール。
resource "aws_security_group_rule" "ingress" {
  type              = "ingress"
  from_port         = var.port
  to_port           = var.port
  protocol          = "tcp"
  cidr_blocks       = var.cidr_blocks
  security_group_id = aws_security_group.default.id
}

# アウトバウンドルール。全ての通信を許可
resource "aws_security_group_rule" "egress" {
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.default.id
}

output "security_group_id" {
  value = aws_security_group.default.id
}

# このモジュールの使い方
# • name - セキュリティグループの名前
# • vpc_id - VPC の ID
# • port - 通信を許可するポート番号
# • cidr_blocks - 通信を許可する CIDR ブロック
#
# module "connpass-map_sg" {
#   source = "./modules/security_group"
#   name   = "module-sg"
#   vpc_id = aws_vpc.connpass-map.id
#   port   = 80
#   cidr_blocks = ["0.0.0.0/0"]
# }
