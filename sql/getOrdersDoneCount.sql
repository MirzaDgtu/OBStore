SELECT count(o.id) as countOrders
FROM ordersbuild.orders o
WHERE o.orderdate = "2024-10-24"
	  and o.done = 0