SELECT od.articul, 
	   p.strikecode,
       p.fasovka,
       p.tiptovr,
       p.unit,
       p.nameartic,
       od.qty,
       round(cast((od.qty/p.fasovka) AS decimal(10,2)), 0) as upk,
       round(abs(round(cast((od.qty/p.fasovka) AS decimal(10,2)), 0) - 
		   cast((od.qty/p.fasovka) AS decimal(10,2))), 1) * 10 as unt
      
       
FROM ordersbuild.orderdetails od
	LEFT JOIN ordersbuild.products p on p.article = od.articul
WHERE OrderId = 7
