SELECT x, y, z
CASE
    WHEN z > 0 AND x + y > z AND x > 0 AND y + z > x AND y > 0 AND x + z > y THEN 'Yes'
    ELSE 'No'
END AS 'triangle'
FROM Triangle