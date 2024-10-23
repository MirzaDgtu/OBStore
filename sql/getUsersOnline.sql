SELECT COUNT(u.id) as userCount
 FROM ordersbuild.users u
 WHERE u.loggedin = 1;