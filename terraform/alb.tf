# # alb用のログバケット
# resource "aws_s3_bucket" "alb_log" {
#   bucket        = "alb-log-pragmatic-terraform-on-aws-by-pokotyan"
#   force_destroy = true

#   lifecycle_rule {
#     enabled = true

#     expiration {
#       days = "180"
#     }
#   }
# }

# # s3のバケットポリシー。上記で作成したバケットに対するPutObjectを許可
# resource "aws_s3_bucket_policy" "alb_log" {
#   bucket = aws_s3_bucket.alb_log.id
#   policy = data.aws_iam_policy_document.alb_log.json
# }

# data "aws_iam_policy_document" "alb_log" {
#   statement {
#     effect    = "Allow"
#     actions   = ["s3:PutObject"]
#     resources = ["arn:aws:s3:::${aws_s3_bucket.alb_log.id}/*"]

#     principals {
#       type        = "AWS"
#       identifiers = ["582318560864"]
#     }
#   }
# }

# albのセキュリティグループ定義
module "http_sg" {
  source      = "./modules/security_group"
  name        = "http-sg"
  vpc_id      = aws_vpc.connpass-map.id
  port        = 80
  cidr_blocks = ["0.0.0.0/0"]
}

module "https_sg" {
  source      = "./modules/security_group"
  name        = "https-sg"
  vpc_id      = aws_vpc.connpass-map.id
  port        = 443
  cidr_blocks = ["0.0.0.0/0"]
}

module "http_redirect_sg" {
  source      = "./modules/security_group"
  name        = "http-redirect-sg"
  vpc_id      = aws_vpc.connpass-map.id
  port        = 8080
  cidr_blocks = ["0.0.0.0/0"]
}

# albの作成
resource "aws_lb" "connpass-map" {
  name                       = "connpass-map"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = false # 削除保護

  subnets = [
    aws_subnet.public_0.id,
    aws_subnet.public_1.id,
  ]

  # access_logs {
  #   bucket  = aws_s3_bucket.alb_log.id
  #   enabled = true
  # }

  security_groups = [
    module.http_sg.security_group_id,
    module.https_sg.security_group_id,
    module.http_redirect_sg.security_group_id,
  ]
}

# httpリスナーの定義
resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.connpass-map.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "これは『HTTP』です"
      status_code  = "200"
    }
  }
}

# httpsリスナーの定義
resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.connpass-map.arn
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.connpass-map.arn
  ssl_policy        = "ELBSecurityPolicy-2016-08"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "これは『HTTPS』です"
      status_code  = "200"
    }
  }
}

# http => httpsのリダイレクトリスナーの定義
resource "aws_lb_listener" "redirect_http_to_https" {
  load_balancer_arn = aws_lb.connpass-map.arn
  port              = "8080"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

# ターゲットグループ
resource "aws_lb_target_group" "connpass-map" {
  name                 = "connpass-map"
  target_type          = "ip" # ターゲットグループにfargateを紐付ける場合はip
  vpc_id               = aws_vpc.connpass-map.id
  port                 = 80
  protocol             = "HTTP" # httpsの終端はalbなので、ターゲットグループ内はhttp
  deregistration_delay = 300

  health_check {
    path                = "/status"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }

  depends_on = [aws_lb.connpass-map] # ターゲットグループはalbが作成できてから作成
}

# ターゲットグループにリクエストをフォワードするリスナー。
# 84行目でhttpsリスナーを作成しているが、そのhttpsリスナーに以下のルールを追加。
# 「パスが /* だったらターゲットグループに転送」
# これによりhttpsリスナーのデフォルトアクションは無効化（デフォルトアクションは一番優先度が低い）される。
resource "aws_lb_listener_rule" "connpass-map" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.connpass-map.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/*"]
  }
}

output "alb_dns_name" {
  value = aws_lb.connpass-map.dns_name
}
