package pg

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		LeaveType:            LeaveType{},
		EmployeeLeave:        EmployeeLeave{},
		LeaveApplication:     LeaveApplication{},
		LeaveApplicationTest: LeaveApplicationTest{},
	}
}

type Models struct {
	LeaveType            LeaveType
	EmployeeLeave        EmployeeLeave
	LeaveApplication     LeaveApplication
	LeaveApplicationTest LeaveApplicationTest
}

type LeaveType struct {
	ID                     int       `json:"id"`
	LeaveType              string    `json:"leave_type"`
	Unpaid                 bool      `json:"unpaid"`
	Limit                  int       `json:"limit"`
	EntitlementCalculation float32   `json:"entitlement_calculation"`
	Gender                 int       `json:"gender"`
	AttachmentMandatory    bool      `json:"attachment_mandatory"`
	EncashmentLeave        bool      `json:"encashment_leave"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	CreatedBy              string    `json:"created_by"`
	UpdatedBy              string    `json:"updated_by"`
}

type EmployeeLeave struct {
	ID            int       `json:"id"`
	LeaveTypeID   int       `json:"leave_type_id"`
	LeaveType     string    `json:"leave_type"`
	Entitled      float32   `json:"entitled"`
	Taken         float32   `json:"taken"`
	CreditExpired float32   `json:"credit_expired"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	EmployeeID    int       `json:"employee_id"`
}

type LeaveApplication struct {
	ID           int       `json:"id"`
	LeaveTypeID  int       `json:"leave_type_id"`
	EmployeeID   int       `json:"employee_id"`
	LeaveType    string    `json:"leave_type"`
	ApplyReason  string    `json:"apply_reason"`
	Approved     bool      `json:"approved"`
	Cancelled    bool      `json:"cancelled"`
	RejectReason string    `json:"reject_reason"`
	CancelReason string    `json:"cancel_reason"`
	CreatedAt    time.Time `json:"created_at"`
	ApprovedAt   time.Time `json:"approved_at"`
	CreatedBy    string    `json:"created_by"`
	ApprovedBy   string    `json:"approved_by"`
	LeaveDate    time.Time `json:"leave_date"`
	Filename     string    `json:"attachments"`
}

type LeaveApplicationTest struct {
	Email     string `json:"email"`
	LeaveType int    `json: leave_type`
	LeaveDate string `json:"leave_date"`
	HalfDay   bool   `json:"half_day"`
	Reason    string `json:"reason"`
}

func (el *EmployeeLeave) GetAllLeaves(email string) ([]*EmployeeLeave, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT el.id, 
		el.leave_type_id, 
		lt.leave_type, 
		el.entitled, 
		el.taken, 
		el.credited_expired, 
		el.created_at, 
		el.updated_at, 
		el.updated_by, 
		u.id AS employee_id
		FROM LEAVE_TYPE_CONFIG lt, 
		EMPLOYEE_LEAVE el,
		users u 
		WHERE el.leave_type_id  = lt.id 
		AND u.email = $1;
	`

	rows, err := db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employeeLeaves []*EmployeeLeave

	for rows.Next() {
		var employeeLeave EmployeeLeave
		err := rows.Scan(
			&employeeLeave.ID,
			&employeeLeave.LeaveTypeID,
			&employeeLeave.LeaveType,
			&employeeLeave.Entitled,
			&employeeLeave.Taken,
			&employeeLeave.CreditExpired,
			&employeeLeave.CreatedAt,
			&employeeLeave.UpdatedAt,
			&employeeLeave.CreatedBy,
			&employeeLeave.UpdatedBy,
			&employeeLeave.EmployeeID,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		employeeLeaves = append(employeeLeaves, &employeeLeave)
	}

	return employeeLeaves, nil
}

func (la *LeaveApplication) GetAllApplications(email string) ([]*LeaveApplication, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT la.id, 
		la.leave_type_id, 
		u.id AS employee_id,
		lt.leave_type, 
		la.apply_reason,
		la.approved,
		la.cancelled,
		la.reject_reason,
		la.cancel_reason,
		la.created_at,
		la.approved_at,
		la.created_by,
		la.approved_by,
		ld.leave_date,
		lf.filename,
		FROM LEAVE_APPLICATION la, 
		LEAVE_TYPE_CONFIG lt,
		LEAVE_DETAILS ld,
		LEAVE_ATTACHMENTS lf,
		users u
		WHERE la.leave_type_id = lt.id 
		AND u.email = $1
		AND ld.leave_application_id = la.id
		AND lf.leave_application_id = la.id;
	`

	rows, err := db.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaveApplications []*LeaveApplication

	for rows.Next() {
		var leaveApplication LeaveApplication
		err := rows.Scan(
			&leaveApplication.ID,
			&leaveApplication.LeaveTypeID,
			&leaveApplication.EmployeeID,
			&leaveApplication.LeaveType,
			&leaveApplication.ApplyReason,
			&leaveApplication.Approved,
			&leaveApplication.Cancelled,
			&leaveApplication.RejectReason,
			&leaveApplication.CancelReason,
			&leaveApplication.CreatedAt,
			&leaveApplication.ApprovedAt,
			&leaveApplication.CreatedBy,
			&leaveApplication.ApprovedBy,
			&leaveApplication.LeaveDate,
			&leaveApplication.Filename,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		leaveApplications = append(leaveApplications, &leaveApplication)
	}

	return leaveApplications, nil
}

func (l *LeaveApplicationTest) ApplyLeave(LeaveApplicationTest LeaveApplicationTest) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	Email     string `json:"email"`
	LeaveType int    `json: leave_type`
	LeaveDate string `json:"leave_date"`
	HalfDay   bool   `json:"half_day"`
	Reason    string `json:"reason"`

	stmt := `insert into LEAVES_APPLICATION (leave_type_id, employee_id, apply_reason, leave_status_id) VALUES($1, 3,)`
	err := db.QueryRowContext(ctx, stmt, 
}
