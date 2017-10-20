package shopifyapi

//Does not cover a query separate by , like this /admin/orders/#{id}.json?fields=id,line_items,name,total_price

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

type ShopifyCount struct {
	Count int `json:"count"`
}

type ShopifyCollects struct {
	Collects []Collect `json:"collects"`
}

type ShopifyCollect struct {
	Collect Collect `json:"collect"`
}

type Collect struct {
	ID           int       `json:"id"`
	CollectionID int       `json:"collection_id"`
	ProductID    int       `json:"product_id"`
	Featured     bool      `json:"featured"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Position     int       `json:"position"`
	SortValue    string    `json:"sort_value"`
}

type ShopifyCustomCollections struct {
	CustomCollections []CustomerCollect `json:"custom_collections"`
}

type ShopifyCustomCollection struct {
	CustomCollection CustomerCollect `json:"custom_collection"`
}

type CustomerCollect struct {
	ID             int         `json:"id"`
	Handle         string      `json:"handle"`
	Title          string      `json:"title"`
	UpdatedAt      time.Time   `json:"updated_at"`
	BodyHTML       string      `json:"body_html"`
	PublishedAt    time.Time   `json:"published_at"`
	SortOrder      string      `json:"sort_order"`
	Image          Image       `json:"image"`
	Published      bool        `json:"published"`
	PublishedScope string      `json:"published_scope"`
	TemplateSuffix interface{} `json:"template_suffix"`
}

type ShopifySmartCollections struct {
	SmartCollections []SmartCollect `json:"smart_collections"`
}

type ShopifySmartCollection struct {
	SmartCollection SmartCollect `json:"smart_collection"`
}

type SmartCollect struct {
	ID             int       `json:"id"`
	Handle         string    `json:"handle"`
	BodyHTML       string    `json:"body_html"`
	Image          Image     `json:"image"`
	PublishedAt    time.Time `json:"published_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PublishedScope string    `json:"published_scope"`
	Rules          []struct {
		Column    string `json:"column"`
		Relation  string `json:"relation"`
		Condition string `json:"condition"`
	} `json:"rules"`
	Disjunctive    bool        `json:"disjunctive"`
	SortOrder      string      `json:"sort_order"`
	TemplateSuffix interface{} `json:"template_suffix"`
	Title          string      `json:"title"`
}

type Products struct {
	Products []Produce `json:"products"`
}

type Product struct {
	Product Produce `json:"product"`
}

type Produce struct {
	BodyHTML       string      `json:"body_html"`
	CreatedAt      time.Time   `json:"created_at"`
	Handle         string      `json:"handle"`
	ID             int         `json:"id"`
	Image          Image       `json:"image"`
	Images         []Image     `json:"images"`
	Options        []Option    `json:"options"`
	ProductType    string      `json:"product_type"`
	PublishedAt    time.Time   `json:"published_at"`
	PublishedScope string      `json:"published_scope"`
	Tags           string      `json:"tags"`
	TemplateSuffix interface{} `json:"template_suffix"`
	Title          string      `json:"title"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Variants       Variants    `json:"variants"`
	Vendor         string      `json:"vendor"`
}

type Images struct {
	Images []ProductImage `json:"Images"`
}

type Image struct {
	Image []ProductImage `json:"Image"`
}

type ProductImage struct {
	CreatedAt  time.Time `json:"created_at"`
	Height     int       `json:"height"`
	ID         int       `json:"id"`
	Position   int       `json:"position"`
	ProductID  int       `json:"product_id"`
	Src        string    `json:"src"`
	UpdatedAt  string    `json:"updated_at"`
	VariantIds []int     `json:"variant_ids"`
	Width      int64     `json:"width"`
}

type Variants struct {
	Variants []ProductVariant `json:"variants"`
}
type Variant struct {
	Variant ProductVariant `json:"variant"`
}

type ProductVariant struct {
	Barcode                     string      `json:"barcode"`
	CompareAtPrice              interface{} `json:"compare_at_price"`
	CreatedAt                   time.Time   `json:"created_at"`
	FulfillmentService          string      `json:"fulfillment_service"`
	Grams                       int         `json:"grams"`
	ID                          int         `json:"id"`
	ImageID                     int         `json:"image_id"`
	InventoryManagement         string      `json:"inventory_management"`
	InventoryPolicy             string      `json:"inventory_policy"`
	InventoryQuantity           int         `json:"inventory_quantity"`
	InventoryQuantityAdjustment int         `json:"inventory_quantity_adjustment"`
	OldInventoryQuantity        int         `json:"old_inventory_quantity"`
	Option1                     string      `json:"option1"`
	Option2                     string      `json:"option2"`
	Option3                     string      `json:"option3"`
	Position                    int         `json:"position"`
	Price                       string      `json:"price"`
	ProductID                   int         `json:"product_id"`
	RequiresShipping            bool        `json:"requires_shipping"`
	Sku                         string      `json:"sku"`
	Taxable                     time.Time   `json:"taxable"`
	Title                       string      `json:"title"`
	UpdatedAt                   string      `json:"updated_at"`
	Weight                      float64     `json:"weight"`
	WeightUnit                  string      `json:"weight_unit"`
}

type Orders struct {
	Orders []ProductOrder `json:"orders"`
}

type Order struct {
	Order ProductOrder `json:"order"`
}

type ProductOrder struct {
	AppID                int       `json:"app_id"`
	BillingAddress       Address   `json:"billing_address"`
	BrowserIP            string    `json:"browser_ip"`
	BuyerAcceptMarketing bool      `json:"buyer_accepts_marketing"`
	CancelReason         string    `json:"cancel_reason"`
	CancelledAt          time.Time `json:"cancelled_at"`
	CartToken            string    `json:"cart_token"`
	ClientDetails        struct {
		AcceptLanguage interface{} `json:"accept_language"`
		BrowserIP      string      `json:"browser_ip"`
		BrowserHeight  interface{} `json:"browser_height"`
		BrowserWidth   interface{} `json:"browser_width"`
		SessionHash    interface{} `json:"session_hash"`
		UserAgent      interface{} `json:"user_agent"`
	} `json:"client_details"`
	ClosedAt       time.Time `json:"closed_at"`
	CreatedAt      time.Time `json:"created_at"`
	Currency       string    `json:"currency"`
	Customer       Customer  `json:"customer"`
	CustomerLocale string    `json:"customer_locale"`
	DiscountCodes  struct {
		Amount string `json:"amount"`
		Code   string `json:"code"`
		Type   string `json:"type"`
	} `json:"discount_codes"`
	Email           string `json:"email"`
	FinancialStatus string `json:"financial_status"`
	Fulfillments    struct {
		CreatedAt       time.Time `json:"created_at"`
		ID              int       `json:"id"`
		LineItems       string    `json:"line_items"`
		OrderID         int       `json:"order_id"`
		Receipt         string    `json:"receipt"`
		Status          string    `json:"status"`
		TrackingCompany string    `json:"tracking_company"`
		UpdatedAt       time.Time `json:"updated_at"`
	} `json:"fulfillments"`
	FulfillmentsStatus  string            `json:"fulfillment_status"`
	Tags                string            `json:"tags"`
	ID                  int               `json:"id"`
	LineItems           []LineItem        `json:"line_items"`
	LocationID          int               `json:"location_id"`
	Name                string            `json:"name"`
	Note                string            `json:"note"`
	NoteAttributes      map[string]string `json:"note_attributes"`
	Number              int               `json:"number"`
	OrderNumber         int               `json:"order_number"`
	PaymentGatewayNames map[string]string `json:"payment_gateway_names"`
	Phone               string            `json:"phone"`
	ProcessedAt         time.Time         `json:"processed_at"`
	ProcessingMethod    string            `json:"processing_method"`
	ReferringSite       string            `json:"referring_site"`
	Refunds             []Refund          `json:"refunds"`
	ShippingAddress     Address           `json:"shipping_address"`
	ShippingLines       struct {
		Code                          string    `json:"code"`
		Price                         string    `json:"price"`
		Source                        string    `json:"source"`
		Title                         string    `json:"title"`
		TaxLines                      []TaxLine `json:"tax_lines"`
		CarrierIdentifier             string    `json:"carrier_identifier"`
		RequestedFulfillmentServiceID string    `json:"requested_fulfillment_service_id"`
	}
	SourceName          string    `json:"source_name"`
	SubtotalPrice       float64   `json:"subtotal_price"`
	TaxLines            []TaxLine `json:"tax_lines"`
	TaxIncluded         bool      `json:"taxes_included"`
	Token               string    `json:"token"`
	TotalDiscounts      string    `json:"total_discounts"`
	TotalLineItemsPrice string    `json:"total_line_items_price"`
	TotalPrice          string    `json:"total_price"`
	TotalTax            string    `json:"total_tax"`
	TotalWeight         int       `json:"total_weight"`
	UpdatedAt           time.Time `json:"updated_at"`
	UserID              int       `json:"user_id"`
	OrderStatusURL      string    `json:"order_status_url"`
}

type Address struct {
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Company      string `json:"company"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	FirstName    string `json:"first_name"`
	ID           int    `json:"id"`
	LastName     string `json:"last_name"`
	Latitude     string `json:"latitude"`
	Longtitude   string `json:"longtitude"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	ProvinceCode string `json:"province_code"`
	Province     string `json:"province"`
	Zip          string `json:"zip"`
	Default      bool   `json:"default"`
}

type Customer struct {
	AcceptMarketing bool        `json:"accepts_marketing"`
	CreatedAt       time.Time   `json:"created_at"`
	Email           string      `json:"email"`
	Phone           string      `json:"phone"`
	FirstName       string      `json:"first_name"`
	ID              int         `json:"id"`
	LastName        string      `json:"last_name"`
	Note            interface{} `json:"note"`
	OrdersCount     int         `json:"orders_count"`
	State           string      `json:"state"`
	TotalSpent      string      `json:"total_spent"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Tags            string      `json:"tags"`
}

type LineItem struct {
	FulfillableQuantity int         `json:"fulfillable_quantity"`
	FulfillmentService  string      `json:"fulfillment_service"`
	FulfillmentStatus   string      `json:"fulfillment_status"`
	Grams               int         `json:"grams"`
	ID                  int         `json:"id"`
	Price               string      `json:"price"`
	ProductID           int         `json:"product_id"`
	Quantity            int         `json:"quantity"`
	RequiresShipping    bool        `json:"requires_shipping"`
	Sku                 string      `json:"sku"`
	Title               string      `json:"title"`
	VariantID           int         `json:"variant_id"`
	VariantTitle        string      `json:"variant_title"`
	Vendor              string      `json:"vendor"`
	Name                string      `json:"name"`
	GiftCard            bool        `json:"gift_card"`
	Properties          interface{} `json:"properties"`
	Taxable             bool        `json:"taxable"`
	TaxLines            []TaxLine   `json:"tax_lines"`
	TotalDiscount       string      `json:"total_discount"`
}

type Refund struct {
	CreatedAt       time.Time        `json:"created_at"`
	ProcessedAt     time.Time        `json:"processed_at"`
	ID              int              `json:"id"`
	Note            string           `json:"note"`
	RefundLineItems []RefundLineItem `json:"refund_line_items"`
	Restock         bool             `json:"restock"`
	Transactions    []Transaction    `json:"transactions"`
	UserID          int              `json:"user_id"`
}

type RefundLineItem struct {
	ID         int      `json:"id"`
	LineItem   LineItem `json:"line_item"`
	LineItemID int      `json:"line_item_id"`
	Quantity   int      `json:"quantity"`
	Subtotal   float64  `json:"subtotal"`
	TotalTax   float64  `json:"total_tax"`
}

type Transaction struct {
	Amount         string    `json:"amount"`
	Authorization  string    `json:"authorization"`
	CreatedAt      time.Time `json:"created_at"`
	DeviceID       string    `json:"device_id"`
	Gateway        string    `json:"gateway"`
	SourceName     string    `json:"source_name"`
	PaymentDetails struct {
		AvsResultCode     string `json:"avs_result_code"`
		CreditCardBin     string `json:"credit_card_bin"`
		CreditCardCompany string `json:"credit_card_company"`
		CreditCardNumber  string `json:"credit_card_number"`
		CvvResultCode     string `json:"cvv_result_code"`
	} `json:"payment_details"`
	ID        int         `json:"id"`
	Kind      string      `json:"kind"`
	OrderID   int         `json:"order_id"`
	Receipt   interface{} `json:"receipt"`
	ErrorCode string      `json:"error_code"`
	Status    string      `json:"status"`
	Test      bool        `json:"test"`
	UserID    int         `json:"user_id"`
	Currency  string      `json:"currency"`
}

type TaxLine struct {
	Price string  `json:"price"`
	Rate  float64 `json:"rate"`
	Title string  `json:"title"`
}

type Option struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	ProductID int      `json:"product_id"`
	Values    []string `json:"values"`
}

type Shop struct {
	Address1                string      `json:"address1"`
	Address2                string      `json:"address2"`
	City                    string      `json:"city"`
	Country                 string      `json:"country"`
	CountryCode             string      `json:"country_code"`
	CountryName             string      `json:"country_name"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	CustomerEmail           string      `json:"customer_email"`
	Currency                string      `json:"currency"`
	Domain                  string      `json:"domain"`
	Email                   string      `json:"email"`
	GoogleAppDomain         interface{} `json:"google_apps_domain"`
	GoogleAppLoginEnabled   interface{} `json:"google_apps_login_enabled"`
	ID                      int         `json:"id"`
	LastName                float64     `json:"last_name"`
	Latitude                float64     `json:"latitude"`
	MoneyFormat             string      `json:"money_format"`
	MoneyWithCurrencyFormat string      `json:"money_with_currency_format"`
	WeightUnit              string      `json:"weight_unit"`
	MyShopifyDomain         string      `json:"myshopify_domain"`
	Name                    string      `json:"name"`
	PlanName                string      `json:"plan_name"`
	HasDiscounts            bool        `json:"has_discounts"`
	HasGiftCards            bool        `json:"has_gift_cards"`
	PlanDisplayName         string      `json:"plan_display_name"`
	PasswordEnabled         bool        `json:"password_enabled"`
	Phone                   string      `json:"phone"`
	PrimaryLocale           string      `json:"primary_locale"`
	Province                string      `json:"province"`
	ProvinceCode            string      `json:"province_code"`
	ShopOwner               string      `json:"shop_owner"`
	Source                  interface{} `json:"source"`
	ForceSSL                bool        `json:"force_ssl"`
	TaxShipping             bool        `json:"tax_shipping"`
	TaxesIncluded           bool        `json:"taxes_included"`
	CountyTaxes             bool        `json:"county_taxes"`
	TimeZone                string      `json:"timezone"`
	IanaTimeZone            string      `json:"iana_timezone"`
	Zip                     string      `json:"zip"`
	HasStoreFront           bool        `json:"has_storefront"`
	SetupRequired           bool        `json:"setup_required"`
}

type ShopifyDownloader struct {
	NetClient   *http.Client
	Domain      string
	AccessToken string
}

type iShopifyDownloader interface {
	BuildURL(path string, query map[string]string) (string, error)
	OneIDURL(path string, query map[string]string, id string) (string, error)
	TwoIDURL(path string, query map[string]string, fid, sid string) (string, error)
	GetCountCollect(query map[string]string) (ShopifyCount, error)
	GetOneCollect(query map[string]string, id string) (ShopifyCollect, error)
	GetListCollect(query map[string]string) (ShopifyCollects, error)
	GetCountCustomCollection(query map[string]string) (ShopifyCount, error)
	GetOneCustomCollection(query map[string]string, id string) (ShopifyCustomCollection, error)
	GetListCustomCollection(query map[string]string) (ShopifyCustomCollections, error)
	GetCountSmartCollection(query map[string]string) (ShopifyCount, error)
	GetOneSmartCollection(query map[string]string, id string) (ShopifySmartCollection, error)
	GetListSmartCollection(query map[string]string) (ShopifySmartCollections, error)
	GetCountProduct(query map[string]string) (ShopifyCount, error)
	GetOneProduct(query map[string]string, id string) (Product, error)
	GetListProduct(query map[string]string) (Products, error)
	GetCountImage(query map[string]string, id string) (ShopifyCount, error)
	GetOneImage(query map[string]string, fid, sid string) (Image, error)
	GetListImage(query map[string]string, id string) (Images, error)
	GetCountProductVariant(query map[string]string, id string) (ShopifyCount, error)
	GetOneProductVariant(query map[string]string, id string) (Variant, error)
	GetListProductVariant(query map[string]string, id string) (Variants, error)
	GetCountOrder(query map[string]string) (ShopifyCount, error)
	GetOneOrder(query map[string]string, id string) (Order, error)
	GetListOrder(query map[string]string) (Orders, error)
	GetStoreConfig(query map[string]string) (Shop, error)
}

// RequestAndRead perform HttpRequest and read response
func RequestAndRead(NetClient *http.Client, url string) ([]byte, error) {
	res, err := NetClient.Get(url)

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// BuildURL with no dynamics
func (sd *ShopifyDownloader) BuildURL(path string, query map[string]string) (string, error) {
	u, err := url.Parse("https://dungda-staging.myshopify.com/admin/orders/count.json?")
	if err != nil {
		fmt.Println("Cannot parse", err)
		return "", err
	}
	u.Host = sd.Domain
	u.Path = path
	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}
	u.RawQuery = params.Encode()
	nu := u.String()
	var buffer bytes.Buffer
	buffer.WriteString(nu)

	ac := ""
	p := &ac

	if len(query) == 0 {
		*p = "access_token="
	} else {
		*p = "&access_token="
	}

	buffer.WriteString(*p)
	buffer.WriteString(sd.AccessToken)
	bu := buffer.String()
	return bu, err
}

// OneIDURL build URL with one id
func (sd *ShopifyDownloader) OneIDURL(path string, query map[string]string, id string) (string, error) {
	r := mux.NewRouter()
	r.Schemes("https").Host(sd.Domain).Path(path).Name("article")
	u, err := r.Get("article").URL("id", id)

	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}

	u.RawQuery = params.Encode()
	nu := u.String()
	var buffer bytes.Buffer
	buffer.WriteString(nu)

	ac := ""
	p := &ac

	if len(query) == 0 {
		*p = "access_token="
	} else {
		*p = "&access_token="
	}

	buffer.WriteString(*p)
	buffer.WriteString(sd.AccessToken)

	bu := buffer.String()
	return bu, err
}

//TwoIDURL build URL with two id
func (sd *ShopifyDownloader) TwoIDURL(path string, query map[string]string, fid, sid string) (string, error) {
	r := mux.NewRouter()
	r.Schemes("https").Host(sd.Domain).Path(path).Name("article")
	u, err := r.Get("article").URL("fid", fid, "sid", sid)

	params := url.Values{}
	for key, value := range query {
		params.Add(key, value)
	}

	u.RawQuery = params.Encode()
	nu := u.String()
	var buffer bytes.Buffer
	buffer.WriteString(nu)

	ac := ""
	p := &ac

	if len(query) == 0 {
		*p = "access_token="
	} else {
		*p = "&access_token="
	}

	buffer.WriteString(*p)
	buffer.WriteString(sd.AccessToken)

	bu := buffer.String()
	return bu, err
}

//GetCountCollect get count collect
func (sd *ShopifyDownloader) GetCountCollect(query map[string]string) (ShopifyCount, error) {
	path := "/admin/collects/count.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneCollect with collect id
func (sd *ShopifyDownloader) GetOneCollect(query map[string]string, cid string) (ShopifyCollect, error) {
	path := "/admin/collects/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, cid)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCollect
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetListCollect get list of collect
func (sd *ShopifyDownloader) GetListCollect(query map[string]string) (ShopifyCollects, error) {
	path := "/admin/collects.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var scs ShopifyCollects
	json.Unmarshal(data, &scs)
	if err != nil {
		return scs, nil
	}
	return scs, nil
}

//GetCountCustomCollection get count of custom collection
func (sd *ShopifyDownloader) GetCountCustomCollection(query map[string]string) (ShopifyCount, error) {
	path := "/admin/custom_collections/count.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneCustomCollection get specific collection by id
func (sd *ShopifyDownloader) GetOneCustomCollection(query map[string]string, cid string) (ShopifyCustomCollection, error) {
	path := "/admin/custom_collections/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, cid)

	data, err := RequestAndRead(sd.NetClient, url)

	var scc ShopifyCustomCollection
	json.Unmarshal(data, &scc)
	if err != nil {
		return scc, nil
	}
	return scc, nil
}

//GetListCustomCollection get list of custom collection
func (sd *ShopifyDownloader) GetListCustomCollection(query map[string]string) (ShopifyCustomCollections, error) {
	path := "/admin/custom_collections.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var scc ShopifyCustomCollections
	json.Unmarshal(data, &scc)
	if err != nil {
		return scc, nil
	}
	return scc, nil
}

//GetCountSmartCollection get count of smart collection
func (sd *ShopifyDownloader) GetCountSmartCollection(query map[string]string) (ShopifyCount, error) {
	path := "/admin/smart_collections/count.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneSmartCollection get specific collection by id
func (sd *ShopifyDownloader) GetOneSmartCollection(query map[string]string, cid string) (ShopifySmartCollection, error) {
	path := "/admin/smart_collections/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, cid)

	data, err := RequestAndRead(sd.NetClient, url)

	var ssc ShopifySmartCollection
	json.Unmarshal(data, &ssc)
	if err != nil {
		return ssc, nil
	}
	return ssc, nil
}

//GetListSmartCollection get list of smart collection
func (sd *ShopifyDownloader) GetListSmartCollection(query map[string]string) (ShopifySmartCollections, error) {
	path := "/admin/smart_collections.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var ssc ShopifySmartCollections
	json.Unmarshal(data, &ssc)
	if err != nil {
		return ssc, nil
	}
	return ssc, nil
}

//GetCountProduct get count of product
func (sd *ShopifyDownloader) GetCountProduct(query map[string]string) (ShopifyCount, error) {
	path := "/admin/products/count.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneProduct get specific product by id
func (sd *ShopifyDownloader) GetOneProduct(query map[string]string, pid string) (Product, error) {
	path := "/admin/products/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, pid)

	data, err := RequestAndRead(sd.NetClient, url)

	var p Product
	json.Unmarshal(data, &p)
	if err != nil {
		return p, nil
	}
	return p, nil
}

//GetListProduct get list of product
func (sd *ShopifyDownloader) GetListProduct(query map[string]string) (Products, error) {
	path := "/admin/products.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var p Products
	json.Unmarshal(data, &p)
	if err != nil {
		return p, nil
	}
	return p, nil
}

//GetCountImage get count of image of a product id
func (sd *ShopifyDownloader) GetCountImage(query map[string]string, pid string) (ShopifyCount, error) {
	path := "/admin/products/{id:[0-9]+}/images/count.json"
	url, err := sd.OneIDURL(path, query, pid)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneImage get an specific image of a product with product id and image id
func (sd *ShopifyDownloader) GetOneImage(query map[string]string, pid, iid string) (Image, error) {
	path := "/admin/products/{fid:[0-9]+}/images/{sid:[0-9]+}.json"
	url, err := sd.TwoIDURL(path, query, pid, iid)

	data, err := RequestAndRead(sd.NetClient, url)

	var i Image
	json.Unmarshal(data, &i)
	if err != nil {
		return i, nil
	}
	return i, nil
}

//GetListImage get list of image from a product id
func (sd *ShopifyDownloader) GetListImage(query map[string]string, pid string) (Images, error) {
	path := "/admin/products/{id:[0-9]+}/images.json"
	url, err := sd.OneIDURL(path, query, pid)

	data, err := RequestAndRead(sd.NetClient, url)

	var i Images
	json.Unmarshal(data, &i)
	if err != nil {
		return i, nil
	}
	return i, nil
}

//GetCountProductVariant get count of product variant with product id
func (sd *ShopifyDownloader) GetCountProductVariant(query map[string]string, pid string) (ShopifyCount, error) {
	path := "/admin/products/{id:[0-9]+}/variants/count.json"
	url, err := sd.OneIDURL(path, query, pid)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneProductVariant get one product variant with variant id
func (sd *ShopifyDownloader) GetOneProductVariant(query map[string]string, vid string) (Variant, error) {
	path := "/admin/variants/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, vid)

	data, err := RequestAndRead(sd.NetClient, url)

	var v Variant
	json.Unmarshal(data, &v)
	if err != nil {
		return v, nil
	}
	return v, nil
}

//GetListProductVariant get list product variant with product id
func (sd *ShopifyDownloader) GetListProductVariant(query map[string]string, pid string) (Variants, error) {
	path := "/admin/products/{id:[0-9]+}/variants.json"
	url, err := sd.OneIDURL(path, query, pid)

	data, err := RequestAndRead(sd.NetClient, url)

	var v Variants
	json.Unmarshal(data, &v)
	if err != nil {
		return v, nil
	}
	return v, nil
}

//GetCountOrder get count of order
func (sd *ShopifyDownloader) GetCountOrder(query map[string]string) (ShopifyCount, error) {
	path := "/admin/orders/count.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var sc ShopifyCount
	json.Unmarshal(data, &sc)
	if err != nil {
		return sc, nil
	}
	return sc, nil
}

//GetOneOrder one order with order id
func (sd *ShopifyDownloader) GetOneOrder(query map[string]string, oid string) (Order, error) {
	path := "/admin/orders/{id:[0-9]+}.json"
	url, err := sd.OneIDURL(path, query, oid)

	data, err := RequestAndRead(sd.NetClient, url)

	var o Order
	json.Unmarshal(data, &o)
	if err != nil {
		return o, nil
	}
	return o, nil
}

//GetListOrder get list order
func (sd *ShopifyDownloader) GetListOrder(query map[string]string) (Orders, error) {
	path := "/admin/orders.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var o Orders
	json.Unmarshal(data, &o)
	if err != nil {
		return o, nil
	}
	return o, nil
}

//GetStoreConfig get store info
func (sd *ShopifyDownloader) GetStoreConfig(query map[string]string) (Shop, error) {
	path := "/admin/shop.json"
	url, err := sd.BuildURL(path, query)

	data, err := RequestAndRead(sd.NetClient, url)

	var s Shop
	json.Unmarshal(data, &s)
	if err != nil {
		return s, nil
	}
	return s, nil
}
