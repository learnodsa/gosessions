package pages

const Login string = `
<html>
<head><title>Sessions Example</title></head>
<body>
	<form action="login" method="post">
		<input name="username" />
		<input name="password" type="password" />
		<input type="Submit" />
	</form>
</body>
</html>`

const Home string = `
<html>
<head><title>Sessions Example</title></head>
<body>
	<form action="logout" method="post">
		<input type="Submit" value="Logout"/>
	</form>
</body>
</html>`
