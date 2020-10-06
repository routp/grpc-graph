# 1. Introduction

gRPC service for undirected non-weighted graph. 
- Service creates a graph from the specified payload and returns an id for the graph
- Finds the shortest path for the given graph id
- Deletes the graph for the given graph id

Uses in-memory singleton map for persistence.
 
# 2. Test
Unit, Functional and integration tests are included.

```shell script
$ go test ./... -v
=== RUN   TestIntegration
graphID: d92e21ee-003a-43a0-8e5a-558112367dcb
A -> B C 
B -> A D 
C -> A 
D -> B 

time=2020-08-10T21:25:05.875344Z level=info caller=service.go:34 message="Graph created with id d92e21ee-003a-43a0-8e5a-558112367dcb"
    TestIntegration: integration_test.go:67: GraphId:  d92e21ee-003a-43a0-8e5a-558112367dcb
graphID: d92e21ee-003a-43a0-8e5a-558112367dcb
A -> B C 
B -> A D 
C -> A 
D -> B 

time=2020-08-10T21:25:05.875655Z level=info caller=service.go:57 message="Shortest Path found for graph d92e21ee-003a-43a0-8e5a-558112367dcb from D to A: D -> B -> A"
    TestIntegration: integration_test.go:74: Shortest Path:  D -> B -> A
time=2020-08-10T21:25:05.875921Z level=error caller=service.go:47 error="Shortest Path finding failed. error: Graph id a123 not found"
    TestIntegration: integration_test.go:80: rpc error: code = Unknown desc = Graph id a123 not found
time=2020-08-10T21:25:05.876775Z level=info caller=service.go:70 message="Graph with id d92e21ee-003a-43a0-8e5a-558112367dcb deleted"
time=2020-08-10T21:25:05.876987Z level=error caller=service.go:73 error="Delete graph failed. error: Graph id a123 not found"
    TestIntegration: integration_test.go:91: rpc error: code = Unknown desc = Graph id a123 not found
--- PASS: TestIntegration (0.00s)
PASS
ok  	graph	1.076s
?   	graph/client	[no test files]
=== RUN   TestConstructGraph
graphID: 9de3f4dc-7473-48f6-bf98-ab938cdf24d8
B -> A C D 
C -> B E F 
A -> B 
D -> B 
E -> C G 
F -> C H I 
G -> E 
H -> F 
I -> F 

--- PASS: TestConstructGraph (0.00s)
PASS
ok  	graph/collections	1.100s
?   	graph/gerror	[no test files]
?   	graph/logger	[no test files]
?   	graph/pb	[no test files]
?   	graph/service	[no test files]
?   	graph/service/model	[no test files]
=== RUN   TestStore
graphID: b2d5cd0a-b405-437d-a6b9-882c9edf1c29
Node1 -> Node2 Node3 Node5 
Node2 -> Node1 Node3 Node2 
Node3 -> Node2 Node1 Node4 
Node4 -> Node3 Node5 
Node5 -> Node4 Node1 

graphID: b2d5cd0a-b405-437d-a6b9-882c9edf1c29
Node4 -> Node3 Node5 
Node5 -> Node4 Node1 
Node1 -> Node2 Node3 Node5 
Node2 -> Node1 Node3 Node2 
Node3 -> Node2 Node1 Node4 

--- PASS: TestStore (0.00s)
PASS
ok  	graph/store	0.615s
=== RUN   TestShortestPath
graphID: f9cbb985-fce7-4251-8a0d-8cfec61a2964
B -> A D 
C -> A E F 
F -> C H I 
H -> F 
I -> F 
A -> B C 
D -> B 
E -> C G 
G -> E 

    TestShortestPath: shortest_path_test.go:26: D -> B -> A -> C -> F -> I
--- PASS: TestShortestPath (0.00s)
PASS
ok  	graph/traversal	0.780s
$ 
```

# 3. Build
From the project's home directory execute `mvn clean install` to generate executable jar.
### 3.1 Build Server
```shell script
$ cd ../graph
$ go build .
```
### 3.2 Build Client
```shell script
$ cd ../graph/client
$ go build .
```

# 4. Start Server
```shell script
$ cd ../graph
$ ./graph --help
  Usage of ./graph:
    -listen string
      	server listener (default ":9007")
    -logLevel string
      	log level (default "Info")
$ ./graph 
  time=2020-08-10T21:28:14.912583Z level=info caller=main.go:47 message="Starting server with Handler gRCP and port :9007."
  time=2020-08-10T21:28:14.912878Z level=info caller=main.go:56 message="Server started"
```

# 5. Client

### 5.1 Command Execution
```shell script
$ cd ../graph/client
```

#### 5.1.1 Create a graph
```shell script
$ ./client create -graph=A-B,B-C,C-D,D-E,E-H,H-A
  2020/08/10 14:37:15 GraphId: 7974f605-d83a-4857-911a-77724e7eef63
$ ./client create -graph=1-2,1-3,2-4,3-5,4-6,4-1,5-2
  2020/08/10 14:46:03 GraphId: 9703ae7b-80c8-4dac-bb8a-140da0a20ddf
$ 
```

#### 5.1.2 Find shortest path
```shell script
$ ./client find -id=7974f605-d83a-4857-911a-77724e7eef63 -src=B -dest=E
  2020/08/10 14:43:38 Shortest Path: B -> A -> H -> E
$ ./client find -id=9703ae7b-80c8-4dac-bb8a-140da0a20ddf -src=5 -dest=6
  2020/08/10 14:48:59 Shortest Path: 5 -> 2 -> 4 -> 6
$
```

#### 5.1.3 Delete a graph
```shell script
$ ./client delete -id=7974f605-d83a-4857-911a-77724e7eef63
  2020/08/10 14:44:26 Message: Graph with id 7974f605-d83a-4857-911a-77724e7eef63 deleted
$
```

### 5.2 Command Usage

#### 5.2.1 Create a graph
```shell script
$ ./client create --help
  Usage of create:
    -address string
      	server listener address (default ":9007")
    -graph string
      	graph edges example: A-B,B-C or 'null' for null graph
$
```
#### 5.2.2 Find shortest path
```shell script
$ ./client find --help
  Usage of find:
    -address string
      	server listener address (default ":9007")
    -dest string
      	destination node
    -id string
      	graph id
    -src string
      	source node
$ 
```

#### 5.2.3  Delete a graph
```shell script
$ ./client delete --help
  Usage of delete:
    -address string
      	server listener address (default ":9007")
    -id string
      	graph id
```