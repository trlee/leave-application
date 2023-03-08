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
	ID                     int     `json:"ID"`
	LeaveType              string  `json:"leaveType"`
	Unpaid                 bool    `json:"unpaid"`
	Limit                  int     `json:"limit"`
	EntitlementCalculation float32 `json:"entitlementCalculation"`
	Gender                 int     `json:"gender"`
	AttachmentMandatory    bool    `json:"attachmentMandatory"`
	EncashmentLeave        bool    `json:"encashmentLeave"`
	SoftDelete             bool    `json:"softDelete"`
}

type EmployeeLeave struct {
	ID            int     `json:"ID"`
	LeaveTypeID   int     `json:"leaveTypeID"`
	LeaveType     string  `json:"leaveType"`
	Entitled      float32 `json:"entitled"`
	Taken         float32 `json:"taken"`
	CreditExpired float32 `json:"creditExpired"`
	EmployeeID    int     `json:"employeeID"`
}

type LeaveApplication struct {
	ID           int       `json:"ID,string,omitempty"`
	LeaveTypeID  int       `json:"leaveTypeID,string,omitempty"`
	EmployeeID   int       `json:"employeeID,string,omitempty"`
	LeaveType    string    `json:"leaveType"`
	ApplyReason  string    `json:"applyReason"`
	Approved     bool      `json:"approved,string"`
	Cancelled    bool      `json:"cancelled,string"`
	RejectReason string    `json:"rejectReason"`
	CancelReason string    `json:"cancelReason"`
	ApprovedAt   time.Time `json:"approvedAt,string"`
	ApprovedBy   string    `json:"approvedBy"`
	LeaveDate    time.Time `json:"leaveDate,string"`
	Filename     string    `json:"attachments"`
	SoftDelete   bool      `json:"softDelete,string"`
}

type LeaveApplicationTest struct {
	Email     string `json:"email"`
	LeaveType int    `json:"leaveType"`
	LeaveDate string `json:"leaveDate"`
	HalfDay   bool   `json:"halfDay"`
	Reason    string `json:"reason"`
}

func (el *EmployeeLeave) GetAllLeaves(email string) ([]*EmployeeLeave, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT el.id, 
		el.leave_type_id, 
		lt.name, 
		el.entitled, 
		el.taken, 
		el.credited_expired, 
		u.id AS employee_id,
		FROM public."LEAVE_TYPE_CONFIG" lt, 
		public."EMPLOYEE_LEAVE" el,
		public."users" u 
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
		lt.name, 
		la.apply_reason,
		la.approved,
		la.cancelled,
		la.reject_reason,
		la.cancel_reason,
		la.approved_at,
		la.approved_by,
		ld.leave_date,
		lf.filename,
		la.soft_delete
		FROM public."LEAVE_APPLICATION" la, 
		public."LEAVE_TYPE_CONFIG" lt,
		public."LEAVE_DETAILS" ld,
		public."LEAVE_ATTACHMENTS" lf,
		public."users" u
		WHERE la.leave_type_id = lt.id 
		AND u.email = $1
		AND ld.leave_application_id = la.id
		AND lf.leave_application_id = la.id
		AND la.soft_delete = false;
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
			&leaveApplication.ApprovedAt,
			&leaveApplication.ApprovedBy,
			&leaveApplication.LeaveDate,
			&leaveApplication.Filename,
			&leaveApplication.SoftDelete,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
		leaveApplications = append(leaveApplications, &leaveApplication)
	}

	log.Println(leaveApplications[0])

	return leaveApplications, nil
}

func (l *LeaveType) GetAllLeaveTypes() ([]*LeaveType, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * from public."LEAVE_TYPE_CONFIG" WHERE soft_delete = false order by id;`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var leaveTypes []*LeaveType

	for rows.Next() {
		var leaveType LeaveType
		err := rows.Scan(
			&leaveType.ID,
			&leaveType.LeaveType,
			&leaveType.Unpaid,
			&leaveType.Limit,
			&leaveType.EntitlementCalculation,
			&leaveType.Gender,
			&leaveType.AttachmentMandatory,
			&leaveType.EncashmentLeave,
			&leaveType.SoftDelete,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		leaveTypes = append(leaveTypes, &leaveType)
	}

	return leaveTypes, nil
}

func (l *LeaveApplicationTest) ApplyLeave(LeaveApplicationTest LeaveApplicationTest) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	// var newID2 int
	// var newID3 int
	stmt := `insert into public."LEAVE_APPLICATION" (leave_type_id, employee_id, apply_reason, leave_status_id, reject_reason, cancel_reason, approved_at, approved_by) VALUES($1, 3, $2, $3, '', '', '1990-01-01', 1) returning id;`
	err := db.QueryRowContext(ctx, stmt, LeaveApplicationTest.LeaveType, LeaveApplicationTest.Reason, LeaveApplicationTest.LeaveType).Scan(&newID)
	if err != nil {
		log.Println("SQL error 1: ", err)
		return false, nil
	} else {
		if InsertLeaveDetails(newID, LeaveApplicationTest.LeaveDate, LeaveApplicationTest.HalfDay) {
			return true, nil
		} else {
			return false, nil
		}
		// stmt2 := `insert into public."LEAVE_DETAILS" (leave_application_id ,leave_date, half_day) VALUES($1, $2, $3) returning id;`
		// err2 := db.QueryRowContext(ctx, stmt2, newID, LeaveApplicationTest.LeaveDate, LeaveApplicationTest.HalfDay).Scan(&newID2)
		// if err2 != nil {
		// 	log.Println("SQL error 2: ", err)
		// 	return false, nil
		// }

		// stmt3 := `insert into public."LEAVE_ATTACHMENTS" (filename, leave_application_id) VALUES('', $1) returning id;`
		// err3 := db.QueryRowContext(ctx, stmt3, newID).Scan(&newID3)
		// if err3 != nil {
		// 	log.Println("SQL error 3: ", err)
		// 	return false, nil
		// } else {
		// 	return true, nil
		// }
	}
}

func InsertLeaveDetails(ID int, LeaveDate string, HalfDay bool) bool {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `insert into public."LEAVE_DETAILS" (leave_application_id ,leave_date, half_day) VALUES($1, $2, $3) returning id;`
	err := db.QueryRowContext(ctx, stmt, ID, LeaveDate, HalfDay).Scan(&newID)
	if err != nil {
		log.Println("SQL error 2: ", err)
		return false
	}
	return true
}

func (l *LeaveApplication) CancelLeaveApplication(id int, reason string) (bool, error) {
	log.Println("In Cancel...")
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_APPLICATION" SET cancel_reason = $1, cancelled = true WHERE id = $2 returning id;`
	err := db.QueryRowContext(ctx, stmt, reason, id).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		return true, nil
	}
}

func (l *LeaveApplication) ApproveLeaveApplication(id int, reason string) (bool, error) {
	log.Println("In Approve...")
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_APPLICATION" SET reject_reason = $1, approved = true WHERE id = $2 returning id;`
	err := db.QueryRowContext(ctx, stmt, reason, id).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		return true, nil
	}
}

func (l *LeaveApplication) RejectLeaveApplication(id int, reason string) (bool, error) {
	log.Println("In Reject...")
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_APPLICATION" SET reject_reason = $1, approved = false WHERE id = $2 returning id;`
	err := db.QueryRowContext(ctx, stmt, reason, id).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		return true, nil
	}
}

func (l *LeaveApplication) DeleteLeaveApplication(ID any) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_APPLICATION" SET soft_delete = true WHERE id = $1 returning id;`
	err := db.QueryRowContext(ctx, stmt, ID).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		log.Print("Updated successfully: ", newID)
		return true, nil
	}
}

func (l *LeaveType) AddLeave(LeaveType LeaveType) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `insert into public."LEAVE_TYPE_CONFIG" (name, unpaid, max_limit, entitlement_calculation, gender, attachment_mandatory, encashment_leave) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id;`
	err := db.QueryRowContext(ctx, stmt, LeaveType.LeaveType, LeaveType.Unpaid, LeaveType.Limit, LeaveType.EntitlementCalculation, LeaveType.Gender, LeaveType.AttachmentMandatory, LeaveType.EncashmentLeave).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		log.Print("Inserted successfully: ", newID)
		return true, nil
	}
}

func (l *LeaveType) UpdateLeave(LeaveType LeaveType) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_TYPE_CONFIG" SET name = $2, unpaid = $3, max_limit = $4, entitlement_calculation = $5, gender = $6, attachment_mandatory = $7, encashment_leave = $8 WHERE id = $1 returning id;`
	err := db.QueryRowContext(ctx, stmt, LeaveType.ID, LeaveType.LeaveType, LeaveType.Unpaid, LeaveType.Limit, LeaveType.EntitlementCalculation, LeaveType.Gender, LeaveType.AttachmentMandatory, LeaveType.EncashmentLeave).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		log.Print("Updated successfully: ", newID)
		return true, nil
	}
}

func (l *LeaveType) DeleteLeave(ID any) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	var newID int
	stmt := `update public."LEAVE_TYPE_CONFIG" SET soft_delete = true WHERE id = $1 returning id;`
	err := db.QueryRowContext(ctx, stmt, ID).Scan(&newID)
	if err != nil {
		log.Println("SQL error: ", err)
		return false, nil
	} else {
		log.Print("Updated successfully: ", newID)
		return true, nil
	}
}
