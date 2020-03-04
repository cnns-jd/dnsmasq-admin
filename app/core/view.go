package core

const(
	LIST_VIEW = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>DNS项目列表</title>
</head>
<body>
<h1>DNS项目列表</h1>
<div>
    <ul>
		{{context}}
    </ul>
	<a href="/add">添加项目</a>
</div>
</body>
<style>
li { list-style: none;}
a { color: blue; }
</style>
</html>
	`

	ITEM_VIEW = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>{{item}}</title>
</head>
<body>
<h1>{{item}}</h1>
<form action="/item/edit/{{item}}" method="post">
<div>
	<textarea name="context" rows="30" cols="300">{{context}}</textarea>
</div>
<button>保存</button>
</form>
</body>
</html>
`

	ADD_VIEW = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>添加项目</title>
</head>
<body>
<h1>添加项目</h1>
<form action="/item/add" method="post">
<div>
	<input type="text" name="name" />
	<textarea name="context" rows="30" cols="300"></textarea>
</div>
<button>保存</button>
</form>
</body>
</html>
`
)
