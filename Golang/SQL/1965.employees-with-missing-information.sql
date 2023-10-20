select e.employee_id from Employees e left join Salaries s on e.employee_id = s.employee_id where where s.salary is null
union
select e.employee_id from Employees e right join Salaries s on employee_id = s.employee_id where e.name is null
order by employee_id;