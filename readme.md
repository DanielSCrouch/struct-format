# Golang Struct Format

FormattedList - Returns a formatted string representation of a list of structs with selected fields represented as columns. Field names and fieldPaths must align. Field paths specify the nested fields of a struct in the format a.b.c

```bash
ID       Name    Status          Node Port   Cluster           
───────────────────────────────────────────────────────────────
foo      baaar   AVAILABLE       30002       baaar-cluster     
fooooo   bar     BROKEN          30005       baaaaar-cluster   
fooo     bar     NOT_CONNECTED   30004       bar-cluster   
```

```go
item1 := &Cell{
  Identity: &Identifier{
    Guid: "foo", Alias: "baaar"
  }, 
  CellStatus: CellStatus_AVAILABLE, 
  NodePort: 30002, 
  clusterIdentifier: &Identifier{
    Guid: "foo-cluster", 
    Alias: "baaar-cluster"
  }
}

item2 := &Cell{
  Identity: &Identifier{
    Guid: "fooooo", 
    Alias: "bar"
  }, 
  CellStatus: CellStatus_BROKEN, 
  NodePort: 30005, 
  clusterIdentifier: &Identifier{
    Guid: "foooooo-cluster", 
    Alias: "baaaaar-cluster"
  }
}
item3 := &Cell{
  Identity: &Identifier{
    Guid: "fooo", 
    Alias: "bar"
  }, 
  CellStatus: CellStatus_NOT_CONNECTED, 
  NodePort: 30004, 
  clusterIdentifier: &Identifier{
    Guid: "fo-cluster", 
    Alias: "bar-cluster"
    }
  }

itemList := []*Cell{item1, item2, item3}
fieldNames := []string{"ID", "Name", "Status", "Node Port", "cluster"}
fieldPaths := []string{"Identity.Guid", "Identity.Alias", "CellStatus", "NodePort", "clusterIdentifier.Alias"}

output, err := format.FormattedList(itemList, fieldNames, fieldPaths)
```
