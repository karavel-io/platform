resource "aws_s3_bucket" "charts" {
  bucket = var.bucket_name
  acl = "private"

}

resource "aws_cloudfront_origin_access_identity" "bucket_identity" {
  comment = "Mikamai Charts repo access identity"
}

data "aws_iam_policy_document" "s3_policy" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.charts.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.bucket_identity.iam_arn]
    }
  }
}

resource "aws_s3_bucket_policy" "example" {
  bucket = aws_s3_bucket.charts.id
  policy = data.aws_iam_policy_document.s3_policy.json
}


resource "aws_acm_certificate" "certificate" {
  provider = aws.us-east-1
  domain_name = var.public_domain
  validation_method = "DNS"
}

resource "cloudflare_record" "validations" {
  for_each = {
  for dvo in aws_acm_certificate.certificate.domain_validation_options : dvo.domain_name => {
    name = dvo.resource_record_name
    value = dvo.resource_record_value
    type = dvo.resource_record_type
  }
  }
  name = each.value.name
  type = each.value.type
  value = each.value.value
  zone_id = var.cloudflare_zone_id
}

resource "aws_cloudfront_distribution" "cdn" {
  enabled = true
  aliases = [var.public_domain]
  origin {
    domain_name = aws_s3_bucket.charts.bucket_regional_domain_name
    origin_id = var.bucket_name
    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.bucket_identity.cloudfront_access_identity_path
    }
  }
  default_cache_behavior {
    allowed_methods = [
      "GET",
      "HEAD",
      "OPTIONS"]
    cached_methods = [
      "GET",
      "HEAD"]
    target_origin_id = var.bucket_name
    viewer_protocol_policy = "redirect-to-https"
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
    min_ttl = 0
    default_ttl = 3600
    max_ttl = 86400
  }
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }
  viewer_certificate {
    acm_certificate_arn = aws_acm_certificate.certificate.arn
    ssl_support_method = "sni-only"
  }
}

resource "cloudflare_record" "repo" {
  name = var.public_domain
  type = "CNAME"
  zone_id = var.cloudflare_zone_id
  value = aws_cloudfront_distribution.cdn.domain_name
}
