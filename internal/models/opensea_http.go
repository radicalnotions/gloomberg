package models

type AssetPaymentToken struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	Decimals int     `json:"decimals"`
	EthPrice float64 `json:"eth_price"`
	UsdPrice float64 `json:"usd_price"`
}

type PrimaryAssetContract struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"`
	CreatedDate                 string `json:"created_date"`
	Name                        string `json:"name"`
	NftVersion                  string `json:"nft_version"`
	OpenseaVersion              any    `json:"opensea_version"`
	Owner                       any    `json:"owner"`
	SchemaName                  string `json:"schema_name"`
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
	Description                 string `json:"description"`
	ExternalLink                string `json:"external_link"`
	ImageURL                    string `json:"image_url"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	PayoutAddress               any    `json:"payout_address"`
}

type CollectionStats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             float64 `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            float64 `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}

type AssetCollection struct {
	Editors                 []string               `json:"editors"`
	PaymentTokens           []AssetPaymentToken    `json:"payment_tokens"`
	PrimaryAssetContracts   []PrimaryAssetContract `json:"primary_asset_contracts"`
	Traits                  struct{}               `json:"traits"`
	Stats                   CollectionStats        `json:"stats"`
	BannerImageURL          string                 `json:"banner_image_url"`
	ChatURL                 any                    `json:"chat_url"`
	CreatedDate             string                 `json:"created_date"`
	DefaultToFiat           bool                   `json:"default_to_fiat"`
	Description             string                 `json:"description"`
	DevBuyerFeeBasisPoints  string                 `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string                 `json:"dev_seller_fee_basis_points"`
	DiscordURL              string                 `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string `json:"external_url"`
	Featured                    bool   `json:"featured"`
	FeaturedImageURL            any    `json:"featured_image_url"`
	Hidden                      bool   `json:"hidden"`
	SafelistRequestStatus       string `json:"safelist_request_status"`
	ImageURL                    string `json:"image_url"`
	IsSubjectToWhitelist        bool   `json:"is_subject_to_whitelist"`
	LargeImageURL               any    `json:"large_image_url"`
	MediumUsername              any    `json:"medium_username"`
	Name                        string `json:"name"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               any    `json:"payout_address"`
	RequireEmail                bool   `json:"require_email"`
	ShortDescription            any    `json:"short_description"`
	Slug                        string `json:"slug"`
	TelegramURL                 any    `json:"telegram_url"`
	TwitterUsername             any    `json:"twitter_username"`
	InstagramUsername           any    `json:"instagram_username"`
	WikiURL                     any    `json:"wiki_url"`
	OwnedAssetCount             int    `json:"owned_asset_count"`
}

type AssetContract struct {
	Collection                  OSCollection `json:"collection"`
	Address                     string       `json:"address"`
	AssetContractType           string       `json:"asset_contract_type"`
	CreatedDate                 string       `json:"created_date"`
	Name                        string       `json:"name"`
	NftVersion                  string       `json:"nft_version"`
	OpenseaVersion              any          `json:"opensea_version"`
	Owner                       int          `json:"owner"`
	SchemaName                  string       `json:"schema_name"`
	Symbol                      string       `json:"symbol"`
	TotalSupply                 any          `json:"total_supply"`
	Description                 string       `json:"description"`
	ExternalLink                string       `json:"external_link"`
	ImageURL                    string       `json:"image_url"`
	DefaultToFiat               bool         `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int          `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int          `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool         `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int          `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int          `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int          `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int          `json:"seller_fee_basis_points"`
	PayoutAddress               string       `json:"payout_address"`
}

type OSCollection struct {
	BannerImageURL          string `json:"banner_image_url"`
	ChatURL                 any    `json:"chat_url"`
	CreatedDate             string `json:"created_date"`
	DefaultToFiat           bool   `json:"default_to_fiat"`
	Description             string `json:"description"`
	DevBuyerFeeBasisPoints  string `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
	DiscordURL              string `json:"discord_url"`
	DisplayData             struct {
		CardDisplayStyle string `json:"card_display_style"`
	} `json:"display_data"`
	ExternalURL                 string `json:"external_url"`
	Featured                    bool   `json:"featured"`
	FeaturedImageURL            any    `json:"featured_image_url"`
	Hidden                      bool   `json:"hidden"`
	SafelistRequestStatus       string `json:"safelist_request_status"`
	ImageURL                    string `json:"image_url"`
	IsSubjectToWhitelist        bool   `json:"is_subject_to_whitelist"`
	LargeImageURL               any    `json:"large_image_url"`
	MediumUsername              any    `json:"medium_username"`
	Name                        string `json:"name"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address"`
	RequireEmail                bool   `json:"require_email"`
	ShortDescription            any    `json:"short_description"`
	Slug                        string `json:"slug"`
	TelegramURL                 any    `json:"telegram_url"`
	TwitterUsername             any    `json:"twitter_username"`
	InstagramUsername           any    `json:"instagram_username"`
	WikiURL                     any    `json:"wiki_url"`
}
