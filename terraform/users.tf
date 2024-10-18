# User john

resource "aws_iam_user" "john_user" {
  name = "john"
  tags = { 
    name = "john", 
    project = "x",
  }
}

data "aws_iam_policy_document" "john_policy" { 
  statement {
    effect    = "Allow"
    actions   = [ 
      "ec2:Describe*",
    ]
    resources = [ 
      "*",
    ]
  }
}

resource "aws_iam_user_policy" "john_user_policy" {
  name   = "john-policy"
  user   = aws_iam_user.john_user.name
  policy = data.aws_iam_policy_document.john_policy.json
}

###################################################
###################################################
# User jim

resource "aws_iam_user" "jim_user" {
  name = "jim"
  tags = { 
    name = "jim", 
    project = "banana",
  }
}

data "aws_iam_policy_document" "jim_policy" { 
  statement {
    effect    = "Allow"
    actions   = [ 
      "ec2:Describe*",
    ]
    resources = [ 
      "*",
    ]
  }
}

resource "aws_iam_user_policy" "jim_user_policy" {
  name   = "jim-policy"
  user   = aws_iam_user.jim_user.name
  policy = data.aws_iam_policy_document.jim_policy.json
}

###################################################
###################################################
# User ana

resource "aws_iam_user" "ana_user" {
  name = "ana"
  tags = { 
    name = "ana", 
    project = "bot",
  }
}

data "aws_iam_policy_document" "ana_policy" { 
  statement {
    effect    = "Allow"
    actions   = [ 
      "ec2:Describe*",
    ]
    resources = [ 
      "*",
    ]
  }
}

resource "aws_iam_user_policy" "ana_user_policy" {
  name   = "ana-policy"
  user   = aws_iam_user.ana_user.name
  policy = data.aws_iam_policy_document.ana_policy.json
}

###################################################
###################################################
