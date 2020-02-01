package validation

// バリデーションタグ
const (
	RequiredTag = "required"
	EqFieldTag  = "eqfield"
	MinTag      = "min"
	MaxTag      = "max"
	EmailTag    = "email"
	PasswordTag = "password"
)

// バリデーションメッセージ
const (
	RequiredMessage = "入力必須です"
	EqFieldMessage  = "%sと入力が一致しません"
	MinMessage      = "%s文字以上で入力してください"
	MaxMessage      = "%s文字以下で入力してください"
	EmailMessage    = "メールアドレスの形式で入力してください"
	PasswordMessage = "パスワードの形式で入力してください"
)

// カスタムバリデーションメッセージ
const (
	CustomUniqueMessage = "すでに存在します"
)
