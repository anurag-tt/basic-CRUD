package postgre

var (
	queryGetEmployeeById = `
		SELECT 
			* 
		FROM 
			employee_info 
		WHERE 
			id=$1
	`

	queryGetAllEmployees = `
	SELECT 
		* 
	FROM 
		employee_info 
	`
	queryCreateEmployee = `
	INSERT INTO 
		employee_info 
		(
			name, 
			department, 
			manager_name, 
			joining_date, 
			updated_at
		) 
	VALUES 
		($1, $2, $3, $4, NOW()) 
	RETURNING 
		id
	`

	queryUpdateEmployeeById = `
		UPDATE 
			employee_info 
		SET 
			name=$2, 
			department=$3, 
			manager_name=$4, 
			joining_date=$5,
			updated_at=NOW() 
		WHERE 
			id=$1
	`

	queryDeleteEmployeeById = `
		DELETE FROM 
			employee_info 
		WHERE 
			id=$1
	`
)
