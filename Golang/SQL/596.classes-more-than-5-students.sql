# Use HAVING to filter after GROUP BY
select class from Courses group by class having count(student) >= 5 