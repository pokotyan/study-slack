resource "aws_acm_certificate" "connpass-map" {
  domain_name               = data.aws_route53_zone.connpass-map.name
  subject_alternative_names = []
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

# ssl証明書の検証完了まで待機する。何かリソースが作成されるわけではない
resource "aws_acm_certificate_validation" "connpass-map" {
  certificate_arn         = aws_acm_certificate.connpass-map.arn
  validation_record_fqdns = [aws_route53_record.connpass-map_certificate.fqdn]
}
