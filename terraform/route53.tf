data "aws_route53_zone" "connpass-map" {
  name = "connpass.net"
}

# albへのエイリアスレコード
resource "aws_route53_record" "connpass-map" {
  zone_id = data.aws_route53_zone.connpass-map.zone_id
  name    = data.aws_route53_zone.connpass-map.name
  type    = "A"

  alias {
    name                   = aws_lb.connpass-map.dns_name
    zone_id                = aws_lb.connpass-map.zone_id
    evaluate_target_health = true
  }
}

# sslのdns検証用のレコード
resource "aws_route53_record" "connpass-map_certificate" {
  name    = aws_acm_certificate.connpass-map.domain_validation_options[0].resource_record_name
  type    = aws_acm_certificate.connpass-map.domain_validation_options[0].resource_record_type
  records = [aws_acm_certificate.connpass-map.domain_validation_options[0].resource_record_value]
  zone_id = data.aws_route53_zone.connpass-map.id
  ttl     = 60
}
