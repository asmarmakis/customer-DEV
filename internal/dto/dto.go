package dto

import "time"

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"user123"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
	RoleID   string `json:"role_id" example:"01HXYZ123456789ABCDEF"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"user123"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginResponse represents login response with token
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  User   `json:"user"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid request"`
}

// CustomerRequest represents customer creation/update request
type CustomerRequest struct {
	Name string `json:"name" binding:"required" example:"PT Teknologi Maju"`
	/* Email       string  `json:"email" example:"info@teknologimaju.com"` */
	/* Phone       string  `json:"phone" example:"021-12345678"`
	Website     string  `json:"website" example:"https://teknologimaju.com"` */
	/* Description string  `json:"description" example:"Perusahaan teknologi informasi"` */
	Status           string  `json:"status" example:"Active"`
	Category         string  `json:"category" example:"Technology"`
	Rating           float64 `json:"rating" example:"4.5"`
	AverageCost      float64 `json:"average_cost" example:"50000000"`
	AccountManagerID *string `json:"account_manager_id" example:"AM001"`
}

// CustomersResponse represents customers list response
type CustomersResponse struct {
	Customers []Customer `json:"customers"`
	Stats     Stats      `json:"stats"`
}

// CustomerListResponse represents simplified customer response for lists
type CustomerListResponse struct {
	ID          string  `json:"id" example:"01HXYZ123456789ABCDEF"`
	Name        string  `json:"name" example:"PT Teknologi Maju"`
	BrandName   string  `json:"brand_name" example:"TechMaju"`
	Code        string  `json:"code" example:"TM001"`
	Logo        string  `json:"logo" example:"uploads/logos/logo_1.png"`
	Status      string  `json:"status" example:"Active"`
	Category    string  `json:"category" example:"Technology"`
	Rating      float64 `json:"rating" example:"4.5"`
	AverageCost float64 `json:"average_cost" example:"50000000"`
	LogoSmall   string  `json:"logo_small" example:"uploads/logos_small/logo_small_1.png"`
	CreatedAt   string  `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt   string  `json:"updated_at" example:"2024-01-15T08:00:00Z"`
	ManagerName *string `json:"manager_name" example:"John Doe"`
}

// Stats represents customer statistics
type Stats struct {
	TotalCustomers   int64   `json:"total_customers" example:"100"`
	NewCustomers     int64   `json:"new_customers" example:"10"`
	AvgCost          float64 `json:"avg_cost" example:"45000000"`
	BlockedCustomers int64   `json:"blocked_customers" example:"5"`
}

// User represents user data in responses
type User struct {
	ID       uint   `json:"id" example:"1"`
	Username string `json:"username" example:"user123"`
	Email    string `json:"email" example:"user@example.com"`
	RoleID   string `json:"role_id" example:"1"`
}

// Customer represents customer data
type Customer struct {
	ID        uint   `json:"id" example:"1"`
	Name      string `json:"name" example:"PT Teknologi Maju"`
	BrandName string `json:"brand_name" example:"TechMaju"`
	Code      string `json:"code" example:"TM001"`
	/* Email       string  `json:"email" example:"info@teknologimaju.com"`
	Phone       string  `json:"phone" example:"021-12345678"`
	Website     string  `json:"website" example:"https://teknologimaju.com"` */
	/* Description string  `json:"description" example:"Perusahaan teknologi informasi"` */
	Logo             string                `json:"logo" example:"uploads/logos/logo_1.png"`
	LogoSmall        string                `json:"logo_small" example:"uploads/logos_small/logo_small_1.png"` // Field baru untuk logo kecil
	Status           string                `json:"status" example:"Active"`
	Category         string                `json:"category" example:"Technology"`
	Rating           float64               `json:"rating" example:"4.5"`
	AverageCost      float64               `json:"average_cost" example:"50000000"`
	AccountManagerID *string               `json:"account_manager_id" example:"AM001"`
	AccountManager   *AccountManagerDetail `json:"account_manager,omitempty"`
}

// CustomerResponse represents customer with simplified relations
type CustomerResponse struct {
	ID               string                `json:"id" example:"1"`
	Name             string                `json:"name" example:"PT Teknologi Maju"`
	BrandName        string                `json:"brand_name" example:"TechMaju"`
	Code             string                `json:"code" example:"TM001"`
	AccountManagerID *string               `json:"account_manager_id" example:"AM001"`
	AccountManager   *AccountManagerDetail `json:"account_manager,omitempty"`
	/* Email            string            `json:"email" example:"info@teknologimaju.com"`
	Phone            string            `json:"phone" example:"021-12345678"`
	Website          string            `json:"website" example:"https://teknologimaju.com"`
	Description      string            `json:"description" example:"Perusahaan teknologi informasi"` */
	Logo        string            `json:"logo" example:"uploads/logos/logo_1.png"`
	LogoSmall   string            `json:"logo_small" example:"uploads/logos_small/logo_small_1.png"`
	Status      string            `json:"status" example:"Active"`
	Category    string            `json:"category" example:"Technology"`
	Rating      float64           `json:"rating" example:"4.5"`
	AverageCost float64           `json:"average_cost" example:"50000000"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Addresses   []AddressResponse `json:"addresses,omitempty"`
	Contacts    []ContactResponse `json:"contacts,omitempty"`
	Others      []OtherResponse   `json:"others,omitempty"`
}

type AddressResponse struct {
	ID      uint   `json:"id" example:"1"`
	Name    string `json:"name" example:"Head Office"`
	Address string `json:"address" example:"Jl. Sudirman No. 123, Jakarta Selatan"`
	IsMain  bool   `json:"isMain" example:"true"`
	Active  bool   `json:"active" example:"true"`
}

type ContactResponse struct {
	ID          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"Budi Santoso"`
	Birthdate   string `json:"birthdate" example:"1985-03-15"`
	JobPosition string `json:"jobPosition" example:"CEO"`
	Email       string `json:"email" example:"budi@digiinno.com"`
	Phone       string `json:"phone" example:"021-5551234"`
	Mobile      string `json:"mobile" example:"0812-3456-7890"`
	IsMain      bool   `json:"isMain" example:"true"`
	Active      bool   `json:"active" example:"true"`
}

type OtherResponse struct {
	ID     uint   `json:"id" example:"1"`
	Key    string `json:"key" example:"company_size"`
	Value  string `json:"value" example:"50-100 employees"`
	Active bool   `json:"active" example:"true"`
}

// AccountManagerDetail represents account manager details in responses
type AccountManagerDetail struct {
	ID          string `json:"id" example:"AM001"`
	ManagerName string `json:"manager_name" example:"John Doe"`
}

// CreateAccountManagerRequest represents account manager creation request
type CreateAccountManagerRequest struct {
	ManagerName string `json:"manager_name" binding:"required" example:"John Doe"`
	Email       string `json:"email" binding:"required,email" example:"john@example.com"`
	Phone       string `json:"phone" example:"+62812345678"`
}

// UpdateAccountManagerRequest represents account manager update request
type UpdateAccountManagerRequest struct {
	ManagerName *string `json:"manager_name" example:"Updated Name"`
	Email       *string `json:"email" example:"updated@example.com"`
	Phone       *string `json:"phone" example:"+62812345679"`
	IsActive    *bool   `json:"is_active" example:"true"`
}

// AccountManagerResponse represents account manager response
type AccountManagerResponse struct {
	ID          string `json:"id" example:"AM001"`
	ManagerName string `json:"manager_name" example:"John Doe"`
	Email       string `json:"email" example:"john@example.com"`
	Phone       string `json:"phone" example:"+62812345678"`
	IsActive    bool   `json:"is_active" example:"true"`
	CreatedAt   string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt   string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

// CreateCustomerRequest represents comprehensive customer creation request
type CreateCustomerRequest struct {
	ID               string                   `json:"id"`
	ParentID         string                   `json:"parent_id"` // <- UBAH dari *uint ke *string
	Name             *string                  `json:"name" binding:"required"`
	BrandName        *string                  `json:"brandName"`
	Code             *string                  `json:"code"`
	AccountManagerID *string                  `json:"account_manager_id"`
	ManagerName      *string                  `json:"manager_name"` // Field baru untuk pemilihan berdasarkan nama
	Logo             string                   `json:"logo"`
	LogoSmall        string                   `json:"logoSmall"`
	StatusName       string                   `json:"status_name"`
	Addresses        []CreateAddressRequest   `json:"addresses,omitempty"`
	Socials          []CreateSocialRequest    `json:"socials,omitempty"`
	Contacts         []CreateContactRequest   `json:"contacts,omitempty"`
	Structures       []CreateStructureRequest `json:"structures,omitempty"`
	Groups           CreateGroupsRequest      `json:"groups,omitempty"` // Ubah dari []CreateGroupsRequest ke CreateGroupsRequest
	Others           []CreateOtherRequest     `json:"others,omitempty"`
}

// CreateAddressRequest represents address creation in customer request
type CreateAddressRequest struct {
	// CustomerID uint   `json:"customer_id" binding:"required"` // Hapus field ini
	Name    string `json:"name" binding:"required" example:"Head Office"`
	Address string `json:"address" binding:"required" example:"Jl. Sudirman No. 123, Jakarta Selatan"`
	IsMain  bool   `json:"isMain" example:"true"`
	Active  bool   `json:"active" example:"true"`
}

// CreateSocialRequest represents social media creation in customer request
type CreateSocialRequest struct {
	// Name     string `json:"name" binding:"required" example:"Instagram"` // Hapus field ini karena duplikat dengan Platform
	Platform string `json:"platform" binding:"required" example:"Instagram"`
	Handle   string `json:"handle" binding:"required" example:"@digiinno_id"`
	Active   bool   `json:"active" example:"true"`
}

// CreateContactRequest represents contact creation in customer request
type CreateContactRequest struct {
	// CustomerID  uint   `json:"customer_id" binding:"required"` // Hapus field ini
	Name        string `json:"name" binding:"required" example:"Budi Santoso"`
	Birthdate   string `json:"birthdate" example:"1985-03-15"`
	JobPosition string `json:"jobPosition" example:"CEO"`
	Email       string `json:"email" example:"budi@digiinno.com"`
	Phone       string `json:"phone" example:"021-5551234"`
	Mobile      string `json:"mobile" example:"0812-3456-7890"`
	IsMain      bool   `json:"isMain" example:"true"`
	Active      bool   `json:"active" example:"true"`
}

// CreateStructureRequest represents structure creation in customer request
type CreateStructureRequest struct {
	// CustomerID uint    `json:"customer_id" binding:"required"` // Hapus field ini
	TempKey   string  `json:"tempKey" example:"1"`
	ParentKey *string `json:"parentKey" example:"null"`
	Name      string  `json:"name" binding:"required" example:"Board of Directors"`
	Level     int     `json:"level" binding:"required" example:"1"`
	Address   string  `json:"address" example:"Jakarta"`
	Active    bool    `json:"active" example:"true"`
}

// CreateGroupsRequest represents groups assignment in customer request
type CreateGroupsRequest struct {
	IndustryID        string `json:"industryId" example:"1"` // Perbaiki nama field
	IndustryActive    bool   `json:"industryActive" example:"true"`
	ParentGroupID     string `json:"parentGroupId" example:"2"` // Perbaiki nama field
	ParentGroupActive bool   `json:"parentGroupActive" example:"true"`
}

// CreateOtherRequest represents other attributes in customer request
type CreateOtherRequest struct {
	// CustomerID uint    `json:"customer_id" binding:"required"` // Hapus field ini
	Key    string  `json:"key" binding:"required" example:"company_size"`
	Value  *string `json:"value" example:"50-100 employees"`
	Active bool    `json:"active" example:"true"`
}

// CreateGroupConfigRequest represents group config creation request
type CreateGroupConfigRequest struct {
	Name   string `json:"name" binding:"required" example:"Sales Activity"`
	Field  string `json:"field" binding:"required" example:"id"`
	Active bool   `json:"active" example:"true"`
}

// GroupConfigResponse represents group config response without is_deleted field
type GroupConfigResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Field     string    `json:"field"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// CreateGroupConfigDetailRequest represents group config detail creation request
type CreateGroupConfigDetailRequest struct {
	Name string `json:"name" binding:"required" example:"Regular Meeting"`
	Icon string `json:"icon" binding:"required" example:"bla.icon"`
}

// GroupConfigDetailResponse represents group config detail response with minimal fields
type GroupConfigDetailResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// CreateActivityRequest represents activity creation request
type CreateActivityRequest struct {
	CustomerID   uint   `json:"customer_id" binding:"required" example:"1"`
	Title        string `json:"title" binding:"required" example:"Client Meeting"`
	Type         string `json:"type" binding:"required" example:"Meeting"`
	Agenda       string `json:"agenda" example:"Discuss project requirements"`
	StartTime    string `json:"start_time" binding:"required" example:"2024-01-15T10:00:00Z"`
	EndTime      string `json:"end_time" binding:"required" example:"2024-01-15T12:00:00Z"`
	LocationName string `json:"location_name" example:"Conference Room A"`
	Status       string `json:"status" example:"Scheduled"`
	// Hapus field Lat dan Lng yang masih ada
}

// UpdateActivityRequest represents activity update request
type UpdateActivityRequest struct {
	Title        *string `json:"title" example:"Updated Meeting"`
	Type         *string `json:"type" example:"Meeting"`
	Agenda       *string `json:"agenda" example:"Updated agenda"`
	StartTime    *string `json:"start_time" example:"2024-01-15T10:00:00Z"`
	EndTime      *string `json:"end_time" example:"2024-01-15T12:00:00Z"`
	LocationName *string `json:"location_name" example:"Conference Room B"`
	Status       *string `json:"status" example:"Completed"`
}

// ActivityResponse represents activity response
type ActivityResponse struct {
	ID           string `json:"id"`
	CustomerID   uint   `json:"customer_id" example:"1"`
	Title        string `json:"title" example:"Client Meeting"`
	Type         string `json:"type" example:"Meeting"`
	Agenda       string `json:"agenda" example:"Discuss project requirements"`
	StartTime    string `json:"start_time" example:"2024-01-15T10:00:00Z"`
	EndTime      string `json:"end_time" example:"2024-01-15T12:00:00Z"`
	LocationName string `json:"location_name" example:"Conference Room A"`
	Status       string `json:"status" example:"Scheduled"`
	CreatedBy    uint   `json:"created_by" example:"1"`
	CreatedAt    string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt    string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

// ActivitiesResponse represents activities list response
type ActivitiesResponse struct {
	Activities []ActivityResponse `json:"activities"`
	Total      int64              `json:"total" example:"10"`
}

// ActivityAttendeeRequest represents activity attendee request
type ActivityAttendeeRequest struct {
	UserIDs []uint `json:"user_ids" binding:"required" example:"[1,2,3]"`
}

// ActivityCheckinRequest represents activity check-in request
type ActivityCheckinRequest struct {
	// Bisa ditambahkan field jika diperlukan, misalnya:
	// Notes string `json:"notes" example:"Arrived on time"`
	// Location string `json:"location" example:"Conference Room A"`
}

// Invoice DTOs
type InvoiceResponse struct {
	ID            string            `json:"id" example:"01HXYZ123456789ABCDEF"`
	CustomerID    uint              `json:"customer_id" example:"1"`
	ProjectID     string            `json:"project_id" example:"01HXYZ123456789ABCDEF"`
	InvoiceNumber string            `json:"invoice_number" example:"INV-2024-001"`
	Amount        float64           `json:"amount" example:"1000000"`
	IssuedDate    time.Time         `json:"issued_date" example:"2024-01-15T00:00:00Z"`
	DueDate       time.Time         `json:"due_date" example:"2024-02-15T00:00:00Z"`
	PaidAmount    float64           `json:"paid_amount" example:"500000"`
	Balance       float64           `json:"balance" example:"500000"`
	Status        string            `json:"status" example:"partial"`
	Customer      *CustomerResponse `json:"customer,omitempty"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type CreateInvoiceRequest struct {
	CustomerID    uint      `json:"customer_id" binding:"required" example:"1"`
	ProjectID     string    `json:"project_id" example:"01HXYZ123456789ABCDEF"`
	InvoiceNumber string    `json:"invoice_number" binding:"required" example:"INV-2024-001"`
	Amount        float64   `json:"amount" binding:"required" example:"1000000"`
	IssuedDate    time.Time `json:"issued_date" binding:"required" example:"2024-01-15T00:00:00Z"`
	DueDate       time.Time `json:"due_date" binding:"required" example:"2024-02-15T00:00:00Z"`
}

type UpdateInvoiceRequest struct {
	CustomerID    *uint      `json:"customer_id" example:"1"`
	ProjectID     *string    `json:"project_id" example:"01HXYZ123456789ABCDEF"`
	InvoiceNumber *string    `json:"invoice_number" example:"INV-2024-001"`
	Amount        *float64   `json:"amount" example:"1000000"`
	IssuedDate    *time.Time `json:"issued_date" example:"2024-01-15T00:00:00Z"`
	DueDate       *time.Time `json:"due_date" example:"2024-02-15T00:00:00Z"`
	PaidAmount    *float64   `json:"paid_amount" example:"500000"`
}

// Payment DTOs
type PaymentResponse struct {
	ID        string           `json:"id" example:"01HXYZ123456789ABCDEF"`
	InvoiceID uint             `json:"invoice_id" example:"1"`
	Amount    float64          `json:"amount" example:"500000"`
	PaidAt    time.Time        `json:"paid_at" example:"2024-01-20T10:00:00Z"`
	Invoice   *InvoiceResponse `json:"invoice,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type CreatePaymentRequest struct {
	InvoiceID uint      `json:"invoice_id" binding:"required" example:"1"`
	Amount    float64   `json:"amount" binding:"required" example:"500000"`
	PaidAt    time.Time `json:"paid_at" binding:"required" example:"2024-01-20T10:00:00Z"`
}

type UpdatePaymentRequest struct {
	InvoiceID *uint      `json:"invoice_id" example:"1"`
	Amount    *float64   `json:"amount" example:"500000"`
	PaidAt    *time.Time `json:"paid_at" example:"2024-01-20T10:00:00Z"`
}

// Status DTOs - removed duplicate CreateStatusRequest

// CreateAssessmentRequest represents assessment creation request
type CreateAssessmentRequest struct {
	Name   string `json:"name" binding:"required" example:"Assessment Name"`
	RoleID string `json:"role_id" binding:"required" example:"01HXYZ123456789ABCDEF"`
}

type UpdateAssessmentRequest struct {
	Name     *string `json:"name" example:"Updated Assessment"`
	RoleID   *string `json:"role_id" example:"01HXYZ123456789ABCDEF"`
	IsActive *bool   `json:"is_active" example:"true"`
}

// AssessmentResponse represents assessment response
type AssessmentResponse struct {
	ID        string `json:"id" example:"01HXYZ123456789ABCDEF"`
	Name      string `json:"name" example:"Assessment Name"`
	RoleID    string `json:"role_id" example:"01HXYZ123456789ABCDEF"`
	RoleName  string `json:"role_name" example:"Admin"`
	IsActive  bool   `json:"is_active" example:"true"`
	CreatedAt string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

// AccountManagerListResponse represents simplified account manager response for lists
type AccountManagerListResponse struct {
	ID          string `json:"id" example:"AM001"`
	ManagerName string `json:"manager_name" example:"John Doe"`
	CreatedAt   string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt   string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

// Event DTOs
type CreateEventRequest struct {
	CustomerID   string `json:"customer_id" binding:"required" example:"01HXYZ123456789ABCDEF"`
	Title        string `json:"title" binding:"required" example:"Product Launch Event"`
	Description  string `json:"description" example:"Launch event for new product line"`
	EventType    string `json:"event_type" binding:"required" example:"Launch"`
	StartDate    string `json:"start_date" binding:"required" example:"2024-03-15T09:00:00Z"`
	EndDate      string `json:"end_date" binding:"required" example:"2024-03-15T17:00:00Z"`
	Location     string `json:"location" example:"Convention Center"`
	Agenda       string `json:"agenda" example:"Product presentation and networking"`
	Status       string `json:"status" example:"Planned"`
	MaxAttendees int    `json:"max_attendees" example:"100"`
}

type UpdateEventRequest struct {
	Title        *string `json:"title" example:"Updated Event Title"`
	Description  *string `json:"description" example:"Updated description"`
	EventType    *string `json:"event_type" example:"Conference"`
	StartDate    *string `json:"start_date" example:"2024-03-15T09:00:00Z"`
	EndDate      *string `json:"end_date" example:"2024-03-15T17:00:00Z"`
	Location     *string `json:"location" example:"Updated Location"`
	Agenda       *string `json:"agenda" example:"Updated agenda"`
	Status       *string `json:"status" example:"Confirmed"`
	MaxAttendees *int    `json:"max_attendees" example:"150"`
}

type EventResponse struct {
	ID           string `json:"id" example:"01HXYZ123456789ABCDEF"`
	CustomerID   string `json:"customer_id" example:"01HXYZ123456789ABCDEF"`
	Title        string `json:"title" example:"Product Launch Event"`
	Description  string `json:"description" example:"Launch event for new product line"`
	EventType    string `json:"event_type" example:"Launch"`
	StartDate    string `json:"start_date" example:"2024-03-15T09:00:00Z"`
	EndDate      string `json:"end_date" example:"2024-03-15T17:00:00Z"`
	Location     string `json:"location" example:"Convention Center"`
	Agenda       string `json:"agenda" example:"Product presentation and networking"`
	Status       string `json:"status" example:"Planned"`
	MaxAttendees int    `json:"max_attendees" example:"100"`
	CreatedAt    string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt    string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

type AssessmentDetail struct {
	ID           string `json:"id" example:"01HXYZ123456789ABCDEF"`
	AssessmentID string `json:"assessment_id" example:"01HXYZ123456789ABCDEF"`
	Name         string `json:"name" example:"Assessment Detail Item"`
	IsActive     bool   `json:"is_active" example:"true"`
	CreatedAt    string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt    string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

type CreateAssessmentDetailRequest struct {
	AssessmentID string `json:"assessment_id" binding:"required" example:"01HXYZ123456789ABCDEF"`
	Name         string `json:"name" binding:"required" example:"Assessment Detail Item"`
}

type UpdateAssessmentDetailRequest struct {
	AssessmentID *string `json:"assessment_id" example:"01HXYZ123456789ABCDEF"`
	Name         *string `json:"name" example:"Updated Detail Item"`
	IsActive     *bool   `json:"is_active" example:"true"`
}

// CreateStatusRequest represents status creation request
type CreateStatusRequest struct {
	StatusName string `json:"status_name" binding:"required" example:"Active"`
}

type UpdateStatusRequest struct {
	StatusName string `json:"status_name" binding:"required" example:"Active"`
}

// StatusResponse represents status response
type StatusResponse struct {
	ID         string `json:"id" example:"01HXYZ123456789ABCDEF"`
	StatusName string `json:"status_name" example:"Active"`
	CreatedAt  string `json:"created_at" example:"2024-01-15T08:00:00Z"`
	UpdatedAt  string `json:"updated_at" example:"2024-01-15T08:00:00Z"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
