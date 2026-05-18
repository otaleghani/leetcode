-- Condition 1: Employees with multiple departments (get the 'Y' flag)
SELECT employee_id, department_id
FROM Employee
WHERE primary_flag = 'Y'

UNION

-- Condition 2: Employees with only one department
SELECT employee_id, department_id
FROM Employee
GROUP BY employee_id
HAVING COUNT(department_id) = 1;
