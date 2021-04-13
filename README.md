
# simple-golang-demo

This is the DEMO of Building the REST APIs using golang.

## Prerequisites
* Go 1.15+ [installed](https://github.com/golang/go)


## Steps for building the project executables

1.0 Clone the project repository

    `git clone https://github.com/githubckgoh1439/simple-golang-demo`


2.0 Change to the project directory.

    `cd  https://github.com/githubckgoh1439/simple-golang-demo/`


3.1 Get all the dependecies and build the binary of project

    `go mod init`
    `go mod vendor`
    `go build main.go`

3.2 Execute the binary

    `go run main.go`

4.1 Sample of View Single Post

`http://localhost:8000/api/posts/3`

```
{
	"postId":1,
	"id":3,
	"name":"odio adipisci rerum aut animi",
	"email":"Nikita@garfield.biz",
	"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
}

```


4.2 Sample of View All Posts

`http://localhost:8000/api/posts`

```
[
	{
		"postId":1,
		"id":1,
		"name":"id labore ex et quam laborum",
		"email":"Eliseo@gardner.biz",
		"body":"laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
	},
	{
		"postId":1,
		"id":2,
		"name":"quo vero reiciendis velit similique earum",
		"email":"Jayne_Kuhic@sydney.com",
		"body":"est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et"
	},
	{
		"postId":1,
		"id":3,
		"name":"odio adipisci rerum aut animi",
		"email":"Nikita@garfield.biz",
		"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
	},
	{
		"postId":1,
		"id":4,
		"name":"alias odio sit",
		"email":"Lew@alysha.tv",
		"body":"non et atque\noccaecati deserunt quas accusantium unde odit nobis qui voluptatem\nquia voluptas consequuntur itaque dolor\net qui rerum deleniti ut occaecati"
	},
	{
		"postId":1,
		"id":5,
		"name":"vero eaque aliquid doloribus et culpa",
		"email":"Hayden@althea.biz",
		"body":"harum non quasi et ratione\ntempore iure ex voluptates in ratione\nharum architecto fugit inventore cupiditate\nvoluptates magni quo et"
	},
	{
		"postId":2,
		"id":6,
		"name":"et fugit eligendi deleniti quidem qui sint nihil autem",
		"email":"Presley.Mueller@myrl.com",
		"body":"doloribus at sed quis culpa deserunt consectetur qui praesentium\naccusamus fugiat dicta\nvoluptatem rerum ut voluptate autem\nvoluptatem repellendus aspernatur dolorem in"
	},
	{
		"postId":2,
		"id":10,
		"name":"eaque et deleniti atque tenetur ut quo ut",
		"email":"Carmen_Keeling@caroline.name",
		"body":"voluptate iusto quis nobis reprehenderit ipsum amet nulla\nquia quas dolores velit et non\naut quia necessitatibus\nnostrum quaerat nulla et accusamus nisi facilis"
	},
	{
		"postId":3,
		"id":14,
		"name":"et officiis id praesentium hic aut ipsa dolorem repudiandae",
		"email":"Nathan@solon.io",
		"body":"vel quae voluptas qui exercitationem\nvoluptatibus unde sed\nminima et qui ipsam aspernatur\nexpedita magnam laudant{
	"PostID":1,
	"ID":3,
	"Name":"odio adipisci rerum aut animi",
	"Email":"Nikita@garfield.biz",
	"Body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
}
ium et et quaerat ut qui dolorum"
	}
]

```


4.3 Sample of View Group Posts

`http://localhost:8000/api/postgroup/1`

```
{
	"postId":1,
	"totalNumberComments":5,
	"post":[
		{
			"id":1,
			"name":"id labore ex et quam laborum",
			"email":"Eliseo@gardner.biz",
			"body":"laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
		},
		{
			"id":2,
			"name":"quo vero reiciendis velit similique earum",
			"email":"Jayne_Kuhic@sydney.com",
			"body":"est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et"
		},
		{
			"id":3,
			"name":"odio adipisci rerum aut animi",
			"email":"Nikita@garfield.biz",
			"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
		},
		{
			"id":4,
			"name":"alias odio sit",
			"email":"Lew@alysha.tv",
			"body":"non et atque\noccaecati deserunt quas accusantium unde odit nobis qui voluptatem\nquia voluptas consequuntur itaque dolor\net qui rerum deleniti ut occaecati"
		},
		{
			"id":5,
			"name":"vero eaque aliquid doloribus et culpa",
			"email":"Hayden@althea.biz",
			"body":"harum non quasi et ratione\ntempore iure ex voluptates in ratione\nharum architecto fugit inventore cupiditate\nvoluptates magni quo et"
		}
	]
}

```

5.1 Sample of Search API : by ID 

`http://localhost:8000/api/comments?post_id=3`

```
{
	"postId":1,
	"id":3,
	"name":"odio adipisci rerum aut animi",
	"email":"Nikita@garfield.biz",
	"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
}
```

5.2 Sample of Search API : by email 

`http://localhost:8000/api/comments?email=Nikita@garfield.biz`


```
{
	"postId":1,
	"id":3,
	"name":"odio adipisci rerum aut animi",
	"email":"Nikita@garfield.biz",
	"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
}

```

5.3 Sample of Search API : by email, name 

`http://localhost:8000/api/comments?email=Nikita@garfield.biz&name=odio adipisci rerum aut animi`

```
{
	"postId":1,
	"id":3,
	"name":"odio adipisci rerum aut animi",
	"email":"Nikita@garfield.biz",
	"body":"quia molestiae reprehenderit quasi aspernatur\naut expedita occaecati aliquam eveniet laudantium\nomnis quibusdam delectus saepe quia accusamus maiores nam est\ncum et ducimus et vero voluptates excepturi deleniti ratione"
}
```
