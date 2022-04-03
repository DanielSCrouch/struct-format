# Struct Format

FormattedList - Returns a formatted string representation of a list of structs with selected fields represented as columns. Field names and fieldPaths must align. Field paths specify the nested fields of a struct in the format a.b.c

<img src="assets/output.png" alt="drawing" width="500"/>

```go
item1 := &datatypes.Cell{
  Identity: &datatypes.Identifier{
    Guid: "foo", Alias: "baaar"
  }, 
  CellStatus: datatypes.CellStatus_AVAILABLE, 
  NodePort: 30002, 
  FactoryIdentifier: &datatypes.Identifier{
    Guid: "foo-factory", 
    Alias: "baaar-factory"
  }
}

item2 := &datatypes.Cell{
  Identity: &datatypes.Identifier{
    Guid: "fooooo", 
    Alias: "bar"
  }, 
  CellStatus: datatypes.CellStatus_BROKEN, 
  NodePort: 30005, 
  FactoryIdentifier: &datatypes.Identifier{
    Guid: "foooooo-factory", 
    Alias: "baaaaar-factory"
  }
}
item3 := &datatypes.Cell{
  Identity: &datatypes.Identifier{
    Guid: "fooo", 
    Alias: "bar"
  }, 
  CellStatus: datatypes.CellStatus_NOT_CONNECTED, 
  NodePort: 30004, 
  FactoryIdentifier: &datatypes.Identifier{
    Guid: "fo-factory", 
    Alias: "bar-factory"
    }
  }

itemList := []*datatypes.Cell{item1, item2, item3}
fieldNames := []string{"ID", "Name", "Status", "Node Port", "Factory"}
fieldPaths := []string{"Identity.Guid", "Identity.Alias", "CellStatus", "NodePort", "FactoryIdentifier.Alias"}
```